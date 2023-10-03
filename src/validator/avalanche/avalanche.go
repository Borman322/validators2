package avalanche

import (
	"context"
	"math/big"
	"strconv"
	"validators2/src/config"
)

type Service struct {
	config    *config.Config
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
	s.validator.Platform = "Avalanche"
	return s.validator.Platform, nil
}

func (s *Service) GetValidatorReward(ctx context.Context) (float64, error) {
	reward, err := GetValidatorReward(s.config.NodeId)
	if err != nil {
		return 0, err
	}

	s.validator.Rewards, err = strconv.ParseFloat(reward, 64)
	if err != nil {
		return 0, err
	}

	return s.validator.Rewards, err
}

func (s *Service) GetValidatorUptime(ctx context.Context) (float32, error) {
	result, err := GetValidatorUptime(s.config.NodeId)
	if err != nil {
		return 0, err
	}

	s.validator.Uptime = result
	return s.validator.Uptime, nil
}

func (s *Service) IsValidatorHealthy(ctx context.Context) (bool, error) {
	result, err := IsValidatorHealthy(ctx, s.config.NodeId)
	if err != nil {
		return false, err
	}
	s.validator.IsHealty = result
	return s.validator.IsHealty, nil

}

func (s *Service) IsValidatorSlashed(ctx context.Context) (bool, error) {
	return s.validator.IsSlashed, nil
}
