package avalanche

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	validators2 "validators2/interface"
)

type Validator struct {
	TxID                   string      `json:"txID"`
	StartTime              string      `json:"startTime"`
	EndTime                string      `json:"endTime"`
	Weight                 string      `json:"weight"`
	NodeID                 string      `json:"nodeID"`
	StakeAmount            string      `json:"stakeAmount"`
	PotentialReward        string      `json:"potentialReward"`
	AccruedDelegateeReward string      `json:"accruedDelegateeReward"`
	Delegator              []Delegator `json:"delegators"`
}

type Delegator struct {
	TxID            string `json:"txID"`
	StartTime       string `json:"startTime"`
	EndTime         string `json:"endTime"`
	Weight          string `json:"weight"`
	NodeID          string `json:"nodeID"`
	StakeAmount     string `json:"stakeAmount"`
	PotentialReward string `json:"potentialReward"`
	Commission      string `json:"commission"`
}

type ValidatorResponse struct {
	JSONRPC string `json:"jsonrpc"`
	Result  struct {
		Validators []Validator `json:"validators"`
	} `json:"result"`
	ID int `json:"id"`
}

type ValidatorReward struct {
	Platform string
	Slashes  string
	Reward   string
}

func Avalanche() *ValidatorReward {
	var avax = ValidatorReward{}
	return &avax
}

func (v ValidatorReward) GetValidatorReward() *validators2.ValidatorReward {
	const url = "https://api.avax.network/ext/bc/P"
	method := "POST"

	payload := []byte(`{
        "jsonrpc": "2.0",
        "method": "platform.getCurrentValidators",
        "params": {
            "nodeIDs": ["NodeID-NcZtrWEjPY7XDT5PHgZbwXLCW3LGBjxui"]
        },
        "id": 1
    }`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var response ValidatorResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error:", err)
	}

	var validatorReward, delegatorsReward, totalReward big.Int

	for _, validator := range response.Result.Validators {
		validatorReward.SetString(validator.PotentialReward, 10)

		// Учтем комиссию делегатов для валидатора (0.02%)
		validatorCommission := new(big.Int)
		validatorCommission.SetString(validator.AccruedDelegateeReward, 10)
		validatorCommission.Mul(validatorCommission, big.NewInt(2))
		validatorCommission.Div(validatorCommission, big.NewInt(100))
		validatorReward.Sub(&validatorReward, validatorCommission)

		for _, delegator := range validator.Delegator {
			delegatorReward := new(big.Int)
			delegatorReward.SetString(delegator.PotentialReward, 10)

			// Учтите комиссию делегата при добавлении к общей сумме (0.02%)
			delegatorCommission := new(big.Int)
			delegatorCommission.SetString(delegator.Commission, 10)
			delegatorCommission.Mul(delegatorCommission, big.NewInt(2))
			delegatorCommission.Div(delegatorCommission, big.NewInt(100))
			delegatorReward.Sub(delegatorReward, delegatorCommission)

			delegatorsReward.Add(&delegatorsReward, delegatorReward)
		}
	}

	delegatorsReward.Mul(&delegatorsReward, big.NewInt(2))
	delegatorsReward.Div(&delegatorsReward, big.NewInt(100))

	totalReward.Add(&validatorReward, &delegatorsReward)
	reward := totalReward.String()

	insertIndex := len(reward) - 9

	result := reward[:insertIndex] + "." + reward[insertIndex:insertIndex+4] + " AVAX"

	return &validators2.ValidatorReward{
		Platform: "Avalanche",
		Slashes:  "",
		Reward:   result,
	}
}
