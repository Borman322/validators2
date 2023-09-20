package polygon

import (
	"context"
	"math/big"
	"strconv"
	"validators2/src/config"

	log "github.com/sirupsen/logrus"
)

type Service struct {
	config *config.Config

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

	return &service, nil
}

func (s *Service) GetValidatorReward(ctx context.Context) (float64, error) {
	reward, err := GetValidatorReward()
	if err != nil {
		log.Errorf("Can not get validator's reward: %s", err)
		return 0, err
	}

	s.validator.Rewards, err = strconv.ParseFloat(reward, 64)
	if err != nil {
		log.Errorf("Can not parse the value: %s", err)
	}
	return s.validator.Rewards, nil
}

func (s *Service) GetValidatorUptime(ctx context.Context) (float32, error) {
	result, err := GetValidatorUptime()
	if err != nil {
		return 0, err
	}

	s.validator.Uptime = result

	return s.validator.Uptime, nil
}

func (s *Service) IsValidatorHealthy(ctx context.Context) (bool, error) {
	status, err := GetValidatorStatus()
	if err != nil {
		log.Errorf("Can not get validator's healthy")
	}

	s.validator.IsHealty = status

	return s.validator.IsHealty, nil
}

func (s *Service) IsValidatorSlashed(ctx context.Context) (bool, error) {

	return s.validator.IsSlashed, nil
}

func (s *Service) GetMissingBlocksCount(ctx context.Context) (int, error) {
	count, err := s.GetMissedBlocksOfValidator(ctx)
	if err != nil {
		return 0, err
	}
	return count, err
}

func (s *Service) GetMissedBlocksOfValidator(ctx context.Context) (int, error) {
	result, err := GetValidatorMissedBlocks()
	if err != nil {
		return 0, err
	}
	return result, nil
}
