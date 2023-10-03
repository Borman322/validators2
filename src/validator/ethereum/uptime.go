package ethereum

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type ValidatorUptime struct {
	Status string `json:"status"`
	Data   []struct {
		ValidatorIndex        int     `json:"validatorindex"`
		AttestationEfficiency float32 `json:"attestation_efficiency"`
	} `json:"data"`
}

func GetValidatorUptime(index string) (float32, error) {
	url := fmt.Sprintf("https://beaconcha.in/api/v1/validator/%s/attestationefficiency", index)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, errors.New("ETH validator: Bad request status " + err.Error())
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, errors.New("ETH validator: Bad response status " + err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, errors.New("ETH validator: " + err.Error())
	}

	var response ValidatorUptime
	err = json.Unmarshal(body, &response)
	if err != nil {
		return 0, errors.New("ETH validator: Can not unmarshal JSON " + err.Error())
	}

	if len(response.Data) == 0 {
		return 0, fmt.Errorf("ETH validator: Index is not correct")
	}
	value := (2 - response.Data[0].AttestationEfficiency) * 100

	uptime, err := strconv.ParseFloat(fmt.Sprintf("%.0f", value), 32)
	if err != nil {
		return 0, nil
	}
	return float32(uptime), nil
}
