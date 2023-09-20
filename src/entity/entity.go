package entity

type InputValidator struct {
	Blockchain      string   `json:"blockchain"`
	ChainURL        string   `json:"chain-url"`
	StakingContract string   `json:"staking-contract"`
	Validators      []string `json:"validators"`
}
