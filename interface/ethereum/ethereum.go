package ethereum

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	validators2 "validators2/interface"
)

type ValidatorResponse struct {
	Status string `json:"status"`
	Data   struct {
		Slashed bool   `json:"slashed"`
		Status  string `json:"status"`
	}
}

type ValidatorStats struct {
	Status string `json:"status"`
	Data   []struct {
		Day          int `json:"day"`
		MissedBlocks int `json:"missed_blocks"`
	} `json:"data"`
}

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

type ValidatorReward struct {
	Platform string
	Slashes  string
	Reward   string
}

func Ethereum() *ValidatorReward {
	var eth = ValidatorReward{}
	return &eth
}

func (v ValidatorReward) GetValidatorReward() *validators2.ValidatorReward {
	urlCheckStatus := "https://beaconcha.in/api/v1/validator"
	urls := []string{"https://beaconcha.in/api/v1/validator/2000/execution/performance", "https://beaconcha.in/api/v1/validator/2000/performance"}

	method := "POST"

	payload := []byte(`{
		"indicesOrPubkey": "2000"
	  }`)

	clientStatus := &http.Client{}
	req, err := http.NewRequest(method, urlCheckStatus, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := clientStatus.Do(req)
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

	if !response.Data.Slashed && response.Data.Status == "active_online" {
		var rewards []string

		const urlStats = "https://beaconcha.in/api/v1/validator/stats/2000?end_day=750&start_day=700"

		reqStats, err := http.NewRequest("GET", urlStats, nil)
		if err != nil {
			log.Println("Error creating request:", err)
		}

		req.Header.Add("accept", "application/json")

		clientStats := &http.Client{}
		respStats, err := clientStats.Do(reqStats)
		if err != nil {
			log.Println("Error sending request:", err)
		}
		defer respStats.Body.Close()

		body, err := io.ReadAll(respStats.Body)
		if err != nil {
			log.Println("Error reading response:", err)
		}

		var responseStats ValidatorStats
		err = json.Unmarshal(body, &responseStats)
		if err != nil {
			fmt.Println("Error:", err)
		}

		var missedBlocks = 0

		for _, value := range responseStats.Data {
			missedBlocks += value.MissedBlocks
		}

		for index, url := range urls {
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				log.Println("Error creating request:", err)
			}

			req.Header.Add("accept", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Println("Error sending request:", err)
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Println("Error reading response:", err)
			}

			var response interface{}

			if index == 0 {
				var execution ExecutionReward
				err = json.Unmarshal(body, &execution)
				if err != nil {
					fmt.Println("Error:", err)
				}
				response = execution
			} else {
				var consensus ConsensusReward
				err = json.Unmarshal(body, &consensus)
				if err != nil {
					fmt.Println("Error:", err)
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
				fmt.Println("Error parsing number:", err)
			}
			total.Add(total, num)
		}

		ethStr := total.Text('f', 4) + " ETH"

		return &validators2.ValidatorReward{
			Platform: "Ethereum",
			Slashes:  string(rune(missedBlocks)),
			Reward:   ethStr,
		}
	} else {
		return &validators2.ValidatorReward{}
	}
}
