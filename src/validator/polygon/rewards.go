package polygon

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"math/big"
	"net/http"
)

type Response struct {
	Status  string `json:"status"`
	Success bool   `json:"success"`
	Result  struct {
		Id                        int         `json:"id"`
		Name                      string      `json:"name"`
		ClaimedReward             json.Number `json:"claimedReward"`
		ValidatorUnclaimedRewards json.Number `json:"validatorUnclaimedRewards"`
	} `json:"result"`
}

type ValidatorReward struct {
	Platform string
	Slashes  string
	Reward   string
}

func GetValidatorReward() (string, error) {
	const api = "https://staking-api.polygon.technology/api/v2/validators/31"

	client := http.DefaultClient

	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		return "", errors.New("MATIC validator: " + err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("MATIC validator: " + err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("BSC validator: " + err.Error())
	}

	var formattedData bytes.Buffer
	err = json.Indent(&formattedData, body, "", "\t")
	if err != nil {
		return "", errors.New("MATIC validator: " + err.Error())
	}

	var data Response
	err = json.Unmarshal(formattedData.Bytes(), &data)
	if err != nil {
		return "", errors.New("MATIC validator: Can not unmarshal JSON. " + err.Error())
	}

	totalRewards := new(big.Int)
	claimedReward := new(big.Int)
	validatorUnclaimedRewards := new(big.Int)

	claimedReward.SetString(data.Result.ClaimedReward.String(), 10)
	validatorUnclaimedRewards.SetString(data.Result.ValidatorUnclaimedRewards.String(), 10)

	totalRewards.Add(totalRewards, claimedReward)
	totalRewards.Add(totalRewards, validatorUnclaimedRewards)

	divisor := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
	maticValue := new(big.Float).Quo(new(big.Float).SetInt(totalRewards), divisor)
	maticStr := maticValue.Text('f', 4)

	return maticStr, nil
}
