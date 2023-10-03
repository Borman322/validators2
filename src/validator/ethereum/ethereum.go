package ethereum

import (
	"context"
	"fmt"
	"math/big"
	"time"
	"validators2/src/config"
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
	s.validator.Platform = "Ethereum"
	return s.validator.Platform, nil
}

func (s *Service) GetValidatorReward(ctx context.Context) (float64, error) {
	reward, err := GetValidatorReward(s.config.ValidatorIndex)
	fmt.Println(s.config.ValidatorIndex)
	if err != nil {
		return 0, err
	}

	s.validator.Rewards = reward
	return s.validator.Rewards, nil
}

func (s *Service) GetValidatorUptime(ctx context.Context) (float32, error) {
	uptime, err := GetValidatorUptime(s.config.ValidatorIndex)
	if err != nil {
		return 0, err
	}
	s.validator.Uptime = uptime
	return s.validator.Uptime, nil
}

func (s *Service) IsValidatorHealthy(ctx context.Context) (bool, error) {
	status, err := GetValidatorStatusAndSlashes(s.config.ValidatorIndex)
	if err != nil {
		return false, err
	}
	if status.Data.Status == "active_online" {
		s.validator.IsHealty = true
	} else {
		s.validator.IsHealty = false
	}

	return s.validator.IsHealty, nil
}

func (s *Service) IsValidatorSlashed(ctx context.Context) (bool, error) {
	status, err := GetValidatorStatusAndSlashes(s.config.ValidatorIndex)
	if err != nil {
		return false, err
	}

	s.validator.IsSlashed = status.Data.Slashed

	return s.validator.IsSlashed, nil
}

func (s *Service) GetMissedBlocksOfValidator(ctx context.Context) (int, error) {
	t := time.Now() /// "2023-09-17T12:00:23Z"
	iterator := 0
	var totalMissedBlocks = 0

	for {
		result, err := GetValidatorMissedBlocks(s.config.ValidatorIndex, iterator+500, iterator)
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
