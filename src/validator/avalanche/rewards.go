package avalanche

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/log"
)

type ValidatorInfo struct {
	TxID                   string      `json:"txID"`
	StartTime              string      `json:"startTime"`
	EndTime                string      `json:"endTime"`
	Weight                 string      `json:"weight"`
	NodeID                 string      `json:"nodeID"`
	StakeAmount            string      `json:"stakeAmount"`
	PotentialReward        string      `json:"potentialReward"`
	AccruedDelegateeReward string      `json:"accruedDelegateeReward"`
	Delegator              []Delegator `json:"delegators"`
	Uptime                 string      `json:"uptime"`
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
	Result  *struct {
		Validators []ValidatorInfo `json:"validators"`
	} `json:"result"`
	Error *struct {
		Code int `json:"code"`
	} `json:"error"`
	ID int `json:"id"`
}

type ValidatorReward struct {
	Platform string
	Slashes  string
	Reward   string
}

func GetValidatorReward(nodeID string) (string, error) {
	const url = "https://api.avax.network/ext/bc/P"
	method := "POST"

	strPayload := fmt.Sprintf(`{
        "jsonrpc": "2.0",
        "method": "platform.getCurrentValidators",
        "params": {
            "nodeIDs": ["%s"]
        },
        "id": 1
    }`, nodeID)
	payload := []byte(strPayload)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return "", errors.New("Avalanche validator: " + err.Error())
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("Avalanche validator: " + err.Error())
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var response ValidatorResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", errors.New("Avalanche validator: " + err.Error())
	}

	if response.Error != nil {
		return "", errors.New("Avalanche validator: NODE-ID is not correct")
	}

	var validatorReward, delegatorsReward, totalReward big.Int

	for _, validator := range response.Result.Validators {
		validatorReward.SetString(validator.PotentialReward, 10)

		validatorCommission := new(big.Int)
		validatorCommission.SetString(validator.AccruedDelegateeReward, 10)
		validatorCommission.Mul(validatorCommission, big.NewInt(2))
		validatorCommission.Div(validatorCommission, big.NewInt(100))
		validatorReward.Sub(&validatorReward, validatorCommission)

		for _, delegator := range validator.Delegator {
			delegatorReward := new(big.Int)
			delegatorReward.SetString(delegator.PotentialReward, 10)

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
	log.Info("uptime ", response.Result.Validators[len(response.Result.Validators)-1].Uptime)

	reward := totalReward.String()
	insertIndex := len(reward) - 9
	result := reward[:insertIndex] + "." + reward[insertIndex:insertIndex+4]
	return result, err
}
