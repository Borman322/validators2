package bsc

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Response struct {
	Total         int `json:"total"`
	RewardDetails []struct {
		ID         int     `json:"id"`
		ChainId    string  `json:"chainId"`
		Validator  string  `json:"validator"`
		ValName    string  `json:"valName"`
		Delegator  string  `json:"delegator"`
		Reward     float64 `json:"reward"`
		Height     uint64  `json:"height"`
		RewardTime string  `json:"rewardTime"`
	} `json:"rewardDetails"`
}

func getBinanceExplorerJSON(limit int, offset int) (*Response, error) {
	url := fmt.Sprintf("https://explorer.bnbchain.org/v1/staking/chains/bsc/delegators/bnb1xnudjls7x4p48qrk0j247htt7rl2k2dzpd6n0k/rewards?limit=%d&offset=%d", limit, offset)

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("BSC validator: " + err.Error())
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("BSC validator: " + err.Error())
	}

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("BSC validator: Can not unmarshal JSON. " + err.Error())
	}
	return &result, nil
}

func isRewardInfoFresh(rewards float64, rewardTimeString string) bool {

	if len(rewardTimeString) == 0 {
		return false
	}

	rewardDate, err := parseStringToTime(rewardTimeString) // "2023-02-15T00:00:00.000+00:00"
	if err != nil {
		return false
	}

	t := time.Now()

	// Check if the rewards are missing or the rewards need to be renewed
	if rewards == 0 || t.Day() != rewardDate.Day() || t.Month() != rewardDate.Month() {
		return false
	}
	return true
}
