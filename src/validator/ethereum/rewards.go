package ethereum

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strconv"
)

type ExecutionReward struct {
	Data []struct {
		PerformanceTotal big.Int `json:"performanceTotal"`
	} `json:"data"`
}

type ConsensusReward struct {
	Data []struct {
		PerformanceTotal big.Int `json:"performancetotal"`
	} `json:"data"`
}

func GetValidatorReward(index string) (float64, error) {

	urls := []string{fmt.Sprintf("https://beaconcha.in/api/v1/validator/%s/execution/performance", index), fmt.Sprintf("https://beaconcha.in/api/v1/validator/%s/performance", index)}

	var rewards []string

	for index, url := range urls {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return 0, errors.New("ETH validator: Bad request status " + err.Error())
		}

		req.Header.Add("accept", "application/json")

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

		var response interface{}

		if index == 0 {
			var execution ExecutionReward
			err = json.Unmarshal(body, &execution)
			if err != nil {
				return 0, errors.New("ETH validator: Can not unmarshal JSON. " + err.Error())
			}
			if len(execution.Data) == 0 {
				return 0, fmt.Errorf("ETH validator: Index is not correct")
			}
			response = execution
		} else {
			var consensus ConsensusReward
			err = json.Unmarshal(body, &consensus)
			if err != nil {
				return 0, errors.New("ETH validator: Can not unmarshal JSON. " + err.Error())
			}
			if len(consensus.Data) == 0 {
				return 0, fmt.Errorf("ETH validator: Index is not correct")
			}
			response = consensus
		}

		switch v := response.(type) {
		case ExecutionReward:
			for _, data := range v.Data {
				var value = data.PerformanceTotal.String()
				insertIndex := len(value) - 18
				result := value[:insertIndex] + "." + value[insertIndex:]
				rewards = append(rewards, result)
			}
		case ConsensusReward:
			for _, data := range v.Data {
				var value = data.PerformanceTotal.String()
				insertIndex := len(value) - 9
				result := value[:insertIndex] + "." + value[insertIndex:]
				rewards = append(rewards, result)
			}
		}

	}
	total := new(big.Float)

	for _, strNum := range rewards {
		num, _, err := big.ParseFloat(strNum, 10, 53, big.ToNearestEven)
		if err != nil {
			return 0, errors.New("ETH validator: Can not parse " + err.Error())
		}
		total.Add(total, num)
	}

	ethStr := total.Text('f', 4)

	eth, err := strconv.ParseFloat(ethStr, 64)
	if err != nil {
		return 0, errors.New("ETH validator: Bad parse " + err.Error())
	}

	return eth, nil
}
