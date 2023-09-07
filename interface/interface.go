package validators2

type ValidatorsRewardsInterface interface {
	GetValidatorReward() *ValidatorReward
}

type ValidatorReward struct {
	Platform string
	Slashes  string
	Reward   string
}
