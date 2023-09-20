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
		log.Error(err)
	}

	bigFloat := new(big.Float).SetInt(result)

	// Convert the big.Float to a float32
	floatValue, _ := bigFloat.Float32()

	s.validator.Uptime = floatValue

	return s.validator.Uptime, nil
}

func (s *Service) IsValidatorHealthy(ctx context.Context) (bool, error) {

	return s.validator.IsHealty, nil
}

func (s *Service) IsValidatorSlashed(ctx context.Context) (bool, error) {

	return s.validator.IsSlashed, nil
}
