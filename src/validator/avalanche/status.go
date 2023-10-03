package avalanche

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type ValidatorAva struct {
	Validators []ValidatorAvaStatus `json:"validators"`
}

type ValidatorAvaStatus struct {
	ValidationStatus string `json:"validationStatus"`
}

func IsValidatorHealthy(ctx context.Context, nodeID string) (bool, error) {
	url := fmt.Sprintf("https://glacier-api.avax.network/v1/networks/mainnet/validators/%s?pageSize=10&sortOrder=desc&validationStatus=active", nodeID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return false, errors.New("Avalanche validator: " + err.Error())
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, errors.New("Avalanche validator: " + err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, errors.New("Avalanche validator: " + err.Error())
	}

	var response ValidatorAva
	err = json.Unmarshal(body, &response)
	if err != nil {
		return false, errors.New("Avalanche validator: " + err.Error())
	}

	if len(response.Validators) == 0 || response.Validators == nil {
		return false, errors.New("Avalanche validator: Couldn't get validator's status")
	}

	validationStatus := response.Validators[len(response.Validators)-1].ValidationStatus

	return validationStatus == "active", nil
}
