package ethereum

import (
	"context"
	"math/big"
	"time"
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

	s.validator.Rewards = reward
	return s.validator.Rewards, nil
}

func (s *Service) GetValidatorUptime(ctx context.Context) (float32, error) {

	return s.validator.Uptime, nil
}

func (s *Service) IsValidatorHealthy(ctx context.Context) (bool, error) {
	status, err := GetValidatorStatusAndSlashes()
	if err != nil {
		log.Errorf("Can not get validator's healthy")
	}
	if status.Data.Status == "active_online" {
		s.validator.IsHealty = true
	} else {
		s.validator.IsHealty = false
	}

	return s.validator.IsHealty, nil
}

func (s *Service) IsValidatorSlashed(ctx context.Context) (bool, error) {
	status, err := GetValidatorStatusAndSlashes()
	if err != nil {
		log.Errorf("Can not get validator's slash info")
	}

	s.validator.IsSlashed = status.Data.Slashed

	return s.validator.IsSlashed, nil
}

func (s *Service) GetMissedBlocksOfValidator(ctx context.Context) (int, error) {
	t := time.Now() /// "2023-09-17T12:00:23Z"
	iterator := 0
	var totalMissedBlocks = 0

	for {
		result, err := GetValidatorMissedBlocks(iterator+500, iterator)
		if err != nil {
			return 0, err
		}

		iterator += 500
		for _, value := range result.Data {
			totalMissedBlocks += value.MissedBlocks
		}

		rewardTime, err := parseStringToTime(result.Data[0].DayEnd)
		if err != nil {
			return 0, err
		}

		if t.Day() == rewardTime.Day() && t.Month() == rewardTime.Month() && t.Year() == rewardTime.Year() {
			break
		}
	}

	return totalMissedBlocks, nil
}
