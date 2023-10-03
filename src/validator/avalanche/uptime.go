package avalanche

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type ValidatorAvaResult struct {
	Result *struct {
		ValidatorsAva []ValidatorAvaUptime `json:"validators"`
	} `json:"result"`
	Error *struct {
		Code int `json:"code"`
	} `json:"error"`
}

type ValidatorAvaUptime struct {
	Uptime string `json:"uptime"`
}

func GetValidatorUptime(nodeID string) (float32, error) {
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
		return 0, errors.New("Avalanche validator: " + err.Error())
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return 0, errors.New("Avalanche validator: " + err.Error())
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var response ValidatorAvaResult
	err = json.Unmarshal(body, &response)
	if err != nil {
		return 0, errors.New("Avalanche validator: " + err.Error())
	}

	if response.Error != nil {
		return 0, errors.New("Avalanche validator: NODE-ID is not correct")
	}

	result, err := strconv.ParseFloat(response.Result.ValidatorsAva[len(response.Result.ValidatorsAva)-1].Uptime, 32)
	if err != nil {
		return 0, errors.New("Avalanche validator: " + err.Error())
	}

	return float32(result), nil
}
