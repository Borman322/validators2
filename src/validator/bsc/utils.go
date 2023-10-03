package bsc

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
)

func parseStringToTime(rewardTime string) (time.Time, error) {

	date, err := time.Parse("2006-01-02T00:00:00.000+00:00", rewardTime)
	if err != nil {
		log.Error(err)
	}
	return date, nil

}

func parseAddresses(address string) ([]string, error) {
	url := fmt.Sprintf("https://www.bnbchain.org/en/staking/validator/%s", address)

	response, err := http.Get(url)
	if err != nil {
		return nil, errors.New("BSC validator: Bad request status " + err.Error())
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, errors.New("BSC validator: Can not read HTML " + err.Error())
	}

	var listAddress = []string{}

	doc.Find(".styled__Address-sc-1wu7y9t-5").Each(func(index int, element *goquery.Selection) {
		if index == 1 || index == 2 {
			listAddress = append(listAddress, element.Text())
		}
	})

	return listAddress, nil
}
