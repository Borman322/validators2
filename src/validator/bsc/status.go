package bsc

import (
	"errors"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Status struct {
	Active bool `json:"active"`
}

func parseStatusFromBSC() (string, error) {
	url := "https://www.bnbchain.org/en/staking/validator/bva1xnudjls7x4p48qrk0j247htt7rl2k2dzp3mr3j"

	response, err := http.Get(url)
	if err != nil {
		return "", errors.New("BSC validator: Bad request status " + err.Error())
	}

	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return "", errors.New("BSC validator: Can not read HTML " + err.Error())
	}

	statusHTMLElement := doc.Find(".styled__Status-sc-1wu7y9t-2").First()
	status := statusHTMLElement.Text()

	return status, nil
}
