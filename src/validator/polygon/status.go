package polygon

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type ResponseStatus struct {
	Status  string `json:"status"`
	Success bool   `json:"success"`
	Result  struct {
		Id           int    `json:"id"`
		Name         string `json:"name"`
		CurrentState string `json:"currentState"`
	} `json:"result"`
}

func GetValidatorStatus() (bool, error) {
	const api = "https://staking-api.polygon.technology/api/v2/validators/31"

	client := http.DefaultClient

	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		return false, errors.New("MATIC validator: " + err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		return false, errors.New("MATIC validator: " + err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, errors.New("BSC validator: " + err.Error())
	}

	var formattedData bytes.Buffer
	err = json.Indent(&formattedData, body, "", "\t")
	if err != nil {
		return false, errors.New("MATIC validator: " + err.Error())
	}

	var data ResponseStatus
	err = json.Unmarshal(formattedData.Bytes(), &data)
	if err != nil {
		return false, errors.New("MATIC validator: Can not unmarshal JSON. " + err.Error())
	}

	if data.Result.CurrentState == "HEALTHY" {
		return true, nil
	} else {
		return false, nil
	}
}
