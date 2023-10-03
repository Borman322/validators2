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
	Platform   string
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

func (s *Service) GetValidatorPlatform(ctx context.Context) (string, error) {
	s.validator.Platform = "Polygon"
	return s.validator.Platform, nil
}

func (s *Service) GetValidatorReward(ctx context.Context) (float64, error) {
	reward, err := GetValidatorReward(s.config.ValidatorIndex)
	if err != nil {
		return 0, err
	}

	s.validator.Rewards, err = strconv.ParseFloat(reward, 64)
	if err != nil {
		log.Errorf("Can not parse the value: %s", err)
	}
	return s.validator.Rewards, nil
}

func (s *Service) GetValidatorUptime(ctx context.Context) (float32, error) {
	result, err := GetValidatorUptime(s.config.ValidatorIndex)
	if err != nil {
		return 0, err
	}

	s.validator.Uptime = result

	return s.validator.Uptime, nil
}

func (s *Service) IsValidatorHealthy(ctx context.Context) (bool, error) {
	status, err := GetValidatorStatus(s.config.ValidatorIndex)
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
	result, err := GetValidatorMissedBlocks(s.config.ValidatorIndex)
	if err != nil {
		return 0, err
	}
	return result, nil
}
