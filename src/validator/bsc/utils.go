package bsc

import (
	"time"

	log "github.com/sirupsen/logrus"
)

func parseStringToTime(rewardTime string) (time.Time, error) {

	date, err := time.Parse("2006-01-02T00:00:00.000+00:00", rewardTime)
	if err != nil {
		log.Error(err)
	}
	return date, nil

}
