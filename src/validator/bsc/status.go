package bsc

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Status struct {
	Active bool `json:"active"`
}

func parseStatusFromBSC(address string) (string, error) {
	url := fmt.Sprintf("https://www.bnbchain.org/en/staking/validator/%s", address)

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
