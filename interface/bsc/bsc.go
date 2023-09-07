package bsc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	validators2 "validators2/interface"
)

type ValidatorReward struct {
	Platform string
	Slashes  string
	Reward   string
}

func Binance() *ValidatorReward {
	var bnb = ValidatorReward{}
	return &bnb
}

func (v ValidatorReward) GetValidatorReward() *validators2.ValidatorReward {
	const api = "https://api.binance.org/v1/staking/chains/bsc/validators/bva1xnudjls7x4p48qrk0j247htt7rl2k2dzp3mr3j/rewards?endTime=1689857489&startTime=1689252689"

	response, err := http.Get(api)
	if err != nil {
		fmt.Println("Ошибка при выполнении GET-запроса:", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
	}

	type Reward struct {
		ChainID     string  `json:"chainId"`
		Validator   string  `json:"validator"`
		ValTokens   float64 `json:"valTokens"`
		TotalReward float64 `json:"totalReward"`
		Commission  float64 `json:"commission"`
		Height      uint64  `json:"height"`
		RewardTime  string  `json:"rewardTime"`
	}

	var rewards []Reward
	err = json.Unmarshal(body, &rewards)
	if err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
	}

	totalReward := 0.0

	for _, reward := range rewards {
		totalReward += reward.TotalReward
	}
	fmt.Println("BSC:")
	fmt.Println("Total Rewards 7d: ", totalReward)
	reward := strconv.FormatFloat(totalReward, 'f', 10, 64)

	return &validators2.ValidatorReward{
		Platform: "Binance Smart Chain",
		Slashes:  "",
		Reward:   reward,
	}
}
