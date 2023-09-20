package validator

import "context"

type Validator interface {
	GetValidatorReward(ctx context.Context) (float64, error)
	GetValidatorUptime(ctx context.Context) (float32, error)
	IsValidatorHealthy(ctx context.Context) (bool, error)
	IsValidatorSlashed(ctx context.Context) (bool, error)
}
