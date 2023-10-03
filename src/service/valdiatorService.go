package service

import (
	"context"
	"encoding/json"
	"os"
	"validators2/src/config"
	"validators2/src/validator"
	"validators2/src/validator/avalanche"
	"validators2/src/validator/bsc"
	"validators2/src/validator/ethereum"
	"validators2/src/validator/fantom"
	"validators2/src/validator/polygon"

	log "github.com/sirupsen/logrus"
)

type ValidatorService struct {
	config *config.Config
}

func NewValService(
	config *config.Config,

) *ValidatorService {
	return &ValidatorService{
		config: config,
	}
}

func (s *ValidatorService) Start(ctx context.Context) {

	s.monitorValidator(s.config, ctx)
}

func (s *ValidatorService) monitorValidator(config *config.Config, ctx context.Context) {
	log.Info("Initialized validator")
	validator, err := s.CreateValidator(config)
	log.Info("Created validator")
	if err == nil {
		platform, err := validator.GetValidatorPlatform(ctx)
		if err != nil {
			log.Error(err)
		}

		reward, err := validator.GetValidatorReward(ctx)
		if err != nil {
			log.Error(err)
			os.Exit(0)
		}

		uptime, err := validator.GetValidatorUptime(ctx)
		if err != nil {
			log.Error(err)
			os.Exit(0)
		}

		isHealthy, err := validator.IsValidatorHealthy(ctx)
		if err != nil {
			log.Error(err)
			os.Exit(0)
		}

		isSlashed, err := validator.IsValidatorSlashed(ctx)
		if err != nil {
			log.Error(err)
			os.Exit(0)
		}

		data := map[string]interface{}{
			"platform":  platform,
			"reward":    reward,
			"uptime":    uptime,
			"isHealthy": isHealthy,
			"isSlashed": isSlashed,
		}
		// encoder := json.NewEncoder(os.Stdout)
		// encoder.Encode(data)

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Error(err)
		}
		output := log.New()
		output.SetFormatter(&log.TextFormatter{
			DisableQuote: true,
		})
		output.WithField("output", string(jsonData)).Info()

	} else {
		log.Error("Couldn't get validator info")
	}

}

func (s *ValidatorService) ParseBoolToInt(value bool) uint64 {
	if value {
		return 1
	} else {
		return 0
	}
}

func (s *ValidatorService) CreateValidator(config *config.Config) (validator.Validator, error) {

	var validator validator.Validator
	var err error

	switch config.Chain {
	case "bsc":
		validator, err = bsc.NewService(config)
	case "avax":
		validator, err = avalanche.NewService(config)
	case "eth":
		validator, err = ethereum.NewService(config)
	case "ftm":
		validator, err = fantom.NewService(config)
	case "pol":
		validator, err = polygon.NewService(config)
	default:
		log.Info("Incorrect input data")
		os.Exit(0)
	}

	if err != nil {
		log.Error(err)
		os.Exit(0)
	}
	return validator, err
}
