package fantom

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
	s.validator.Platform = "Fantom"
	return s.validator.Platform, nil
}

func (s *Service) GetValidatorReward(ctx context.Context) (float64, error) {
	reward, err := GetValidatorReward(s.config.ValidatorAddress)
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
	result, err := GetValidatorUptime(s.config.ValidatorAddress, 10)
	if err != nil {
		log.Error(err)
	}

	s.validator.Uptime = result

	return s.validator.Uptime, nil
}

func (s *Service) IsValidatorHealthy(ctx context.Context) (bool, error) {
	status, err := IsValidatorHealthy(s.config.ValidatorAddress)
	if err != nil {
		log.Error(err)
	}
	s.validator.IsHealty = status
	return s.validator.IsHealty, nil
}

func (s *Service) IsValidatorSlashed(ctx context.Context) (bool, error) {
	slash, err := IsValidatorSlashed(s.config.ValidatorAddress)
	if err != nil {
		log.Error(err)
	}
	s.validator.IsSlashed = slash
	return s.validator.IsSlashed, nil
}
