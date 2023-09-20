package ethereum

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type ValidatorResponse struct {
	Status string `json:"status"`
	Data   struct {
		Slashed bool   `json:"slashed"`
		Status  string `json:"status"`
	}
}

func GetValidatorStatusAndSlashes() (*ValidatorResponse, error) {
	urlCheckStatus := "https://beaconcha.in/api/v1/validator"

	method := "POST"

	payload := []byte(`{
		"indicesOrPubkey": "2000"
	  }`)

	clientStatus := &http.Client{}
	req, err := http.NewRequest(method, urlCheckStatus, bytes.NewBuffer(payload))
	if err != nil {
		return nil, errors.New("ETH validator: Bad request status " + err.Error())
	}
	//req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := clientStatus.Do(req)
	if err != nil {
		return nil, errors.New("ETH validator: Bad response status " + err.Error())
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var response ValidatorResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.New("ETH validator: Can not unmarshal JSON " + err.Error())
	}

	return &response, nil

}
