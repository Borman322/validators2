package ethereum

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type ValidatorStats struct {
	Status string `json:"status"`
	Data   []struct {
		Day          int    `json:"day"`
		DayEnd       string `json:"day_end"` /// "2023-05-21T12:00:23Z"
		MissedBlocks int    `json:"missed_blocks"`
	} `json:"data"`
}

func GetValidatorMissedBlocks(index string, endDay int, startDay int) (*ValidatorStats, error) {
	urlStats := fmt.Sprintf("https://beaconcha.in/api/v1/validator/stats/%s?end_day=%d&start_day=%d", index, endDay, startDay)

	reqStats, err := http.NewRequest("GET", urlStats, nil)
	if err != nil {
		return nil, errors.New("ETH validator: Bad request status " + err.Error())
	}

	clientStats := &http.Client{}
	respStats, err := clientStats.Do(reqStats)
	if err != nil {
		return nil, errors.New("ETH validator: Bad response status " + err.Error())
	}
	defer respStats.Body.Close()

	body, err := io.ReadAll(respStats.Body)
	if err != nil {
		return nil, errors.New("ETH validator: " + err.Error())
	}

	var responseStats ValidatorStats
	err = json.Unmarshal(body, &responseStats)
	if err != nil {
		return nil, errors.New("ETH validator: Can not unmarshal JSON. " + err.Error())
	}

	return &responseStats, nil
}

func parseStringToTime(rewardTime string) (time.Time, error) {

	date, err := time.Parse("2006-01-02T15:04:05Z", rewardTime)
	if err != nil {
		log.Error(err)
	}
	return date, nil

}
