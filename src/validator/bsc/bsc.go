package bsc

import (
	"bytes"
	"context"
	"errors"
	"io"
	"math/big"
	"net/http"
	"validators2/src/config"
	"validators2/src/constants"
	"validators2/src/utils"

	"github.com/buger/jsonparser"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	contract "validators2/src/contract/bsc"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	ethClient *ethclient.Client
	config    *config.Config
	contract  *contract.BscValidator
	validator Validator
}

type Validator struct {
	Rewards    float64
	RewardTime string
	Uptime     float32
	IsHealty   bool
	IsSlashed  bool
}

type Indicators struct {
	Height *big.Int
	Count  *big.Int
	Exist  bool
}

func NewService(
	config *config.Config,
) (*Service, error) {
	var service = Service{}
	service.config = config
	err := service.Start()
	return &service, err
}

func (s *Service) Start() error {
	err := s.connectEthereum()
	if err != nil {
		log.Errorf("unable to start contract service: %s", err)
		return err
	}
	contractAddress := common.HexToAddress(s.config.StakingContractAddress)
	instance, err := contract.NewBscValidator(contractAddress, s.ethClient)
	if err != nil {
		log.Error(err)
		return err
	}
	s.contract = instance
	return nil
}

func (s *Service) dialEthClientOrFatal(url string) (*ethclient.Client, *big.Int, error) {
	dial, err := ethclient.Dial(url)
	if err != nil {
		log.Errorf("Unable to dial eth1 client for url (%s): %s", url, err)
		return nil, big.NewInt(0), err
	}
	chainID, err := dial.ChainID(context.Background())
	if err != nil {
		log.Errorf("unable to get chain id: %s", err)
		return nil, big.NewInt(0), err
	}
	return dial, chainID, nil
}

func (s *Service) connectEthereum() error {
	client, _, err := s.dialEthClientOrFatal(s.config.Endpoint)
	s.ethClient = client
	return err
}

func (s *Service) GetValidatorReward(ctx context.Context) (float64, error) {

	if isRewardInfoFresh(s.validator.Rewards, s.validator.RewardTime) {
		return s.validator.Rewards, nil
	}

	result, err := getBinanceExplorerJSON(100, 0)
	if err != nil {
		return 0, errors.New("BSC validator: " + err.Error())
	}

	var totalRewards float64
	for _, reward := range result.RewardDetails {
		totalRewards += reward.Reward
	}

	iterations := result.Total / 100
	for i := 1; i <= iterations; i++ {
		result, err := getBinanceExplorerJSON(100, i*100)
		if err != nil {
			return 0, errors.New("BSC validator: " + err.Error())
		}

		for _, reward := range result.RewardDetails {
			totalRewards += reward.Reward
		}
	}
	s.validator.Rewards = totalRewards
	return s.validator.Rewards, nil
}

func (s *Service) GetValidatorUptime(ctx context.Context) (float32, error) {
	return s.validator.Uptime, nil
}

func (s *Service) IsValidatorHealthy(ctx context.Context) (bool, error) {
	status, err := parseStatusFromBSC()
	if err != nil {
		return false, errors.New("BSC validator: " + err.Error())
	}

	if status == "Active" {
		s.validator.IsHealty = true
	} else {
		s.validator.IsHealty = false
	}

	return s.validator.IsHealty, nil
}

func (s *Service) IsValidatorSlashed(ctx context.Context) (bool, error) {
	validatorHexAddress := common.HexToAddress(s.config.ValidatorAddress)

	var isSlashed Indicators
	isSlashed, err := s.contract.Indicators(nil, validatorHexAddress)
	if err != nil {
		return false, errors.New("BSC validator: " + err.Error())
	}

	if isSlashed.Exist {
		s.validator.IsSlashed = isSlashed.Exist
	} else {
		s.validator.IsSlashed = isSlashed.Exist
	}

	return s.validator.IsSlashed, nil
}

func (s *Service) GetMissingBlocksCount(ctx context.Context) (uint64, error) {
	_, count, err := s.GetMissedBlocksOfValidator(ctx)
	if err != nil {
		return 0, err
	}
	return count.Uint64(), err
}

func (s *Service) IsBlocksSigning(ctx context.Context) (bool, error) {
	_, count, err := s.GetMissedBlocksOfValidator(ctx)
	if err != nil {
		return false, err
	}

	if count.Int64() < 50 {
		return true, err
	}
	return false, err
}

func (s *Service) IsSynced(ctx context.Context) (bool, error) {
	currentBlock, err := s.ethClient.HeaderByNumber(context.Background(), nil)
	lastBlock, err := utils.GetEthBLockNumber(constants.OfficialEndpoints[s.config.Chain])

	if err != nil {
		log.Error("Error getting sync status: ", err)
		return false, err
	}
	log.Info("Current Block: ", s.config.Chain, " ", currentBlock.Number.Uint64())
	log.Info("Highest Block: ", s.config.Chain, " ", lastBlock)

	difference := lastBlock - currentBlock.Number.Uint64()

	if difference > constants.BlockDifference {
		return false, err
	}
	return true, err
}

func (s *Service) GetHeight(ctx context.Context) (uint64, error) {
	header, err := s.ethClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return big.NewInt(0).Uint64(), err
	}

	return header.Number.Uint64(), nil
}

func (s *Service) GetOfficialHeight(endpoint string, ctx context.Context) (uint64, error) {
	return utils.GetEthBLockNumber(endpoint)
}

func (s *Service) GetMissedBlocksOfValidator(ctx context.Context) (*big.Int, *big.Int, error) {
	opts := bind.CallOpts{
		Context: ctx,
	}
	address := common.HexToAddress(s.config.ValidatorAddress)
	log.Info("common.HexToAddress ", address)
	height, count, err := s.contract.GetSlashIndicator(&opts, address)
	if err != nil {
		return big.NewInt(0), big.NewInt(0), err
	}
	log.Info("Height of last missed block bsc: ", height)
	log.Info("Amount of missed blocks bsc: ", count)

	return height, count, nil
}

func (s *Service) GetEarnedTokens(ctx context.Context) (float64, error) {
	if isRewardInfoFresh(s.validator.Rewards, s.validator.RewardTime) {
		return s.validator.Rewards, nil
	}

	result, err := getBinanceExplorerJSON(100, 0)
	if err != nil {
		return 0, errors.New("BSC validator: " + err.Error())
	}

	var totalRewards float64
	for _, reward := range result.RewardDetails {
		totalRewards += reward.Reward
	}

	iterations := result.Total / 100
	for i := 1; i <= iterations; i++ {
		result, err := getBinanceExplorerJSON(100, i*100)
		if err != nil {
			return 0, errors.New("BSC validator: " + err.Error())
		}

		for _, reward := range result.RewardDetails {
			totalRewards += reward.Reward
		}
	}
	s.validator.Rewards = totalRewards
	return s.validator.Rewards, nil
}

func (s *Service) GetVersion(ctx context.Context) (string, error) {
	var jsonStr = []byte(`{"id":1, "jsonrpc": "2.0", "method": "web3_clientVersion", "params": []}`)
	req, err := http.NewRequest("POST", s.config.Endpoint, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	httpClient := &http.Client{
		CheckRedirect: http.DefaultClient.CheckRedirect,
		Timeout:       http.DefaultClient.Timeout,
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "Error getting version", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Error getting version", err
	}

	pathJSON := "result"
	version, err := jsonparser.GetString(body, pathJSON)
	if err != nil {
		return "Error getting version", err
	}
	return version, nil
}
