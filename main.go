package main

import (
	"flag"
	"fmt"
	"os"
	avax "validators2/interface/avalanche"

	//bsc "validators2/interface/bsc"
	validators2 "validators2/interface"
	ethereum "validators2/interface/ethereum"
	fantom "validators2/interface/fantom"
	polygon "validators2/interface/polygon"
)

func main() {
	name := flag.String("value", "", "Value name")
	flag.Parse()

	cryptos := []validators2.ValidatorReward{}

	switch {
	case *name == "ethereum":
		cryptos = append(cryptos, *ethereum.Ethereum().GetValidatorReward())
	case *name == "fantom":
		cryptos = append(cryptos, *fantom.Fantom().GetValidatorReward())
	case *name == "avalanche":
		cryptos = append(cryptos, *avax.Avalanche().GetValidatorReward())
	// case *name == "binance":
	// 	cryptos = append(cryptos, *bsc.Binance().GetValidatorReward())
	case *name == "polygon":
		cryptos = append(cryptos, *polygon.Polygon().GetValidatorReward())
	case *name == "all":
		cryptos = append(cryptos,
			*ethereum.Ethereum().GetValidatorReward(),
			*fantom.Fantom().GetValidatorReward(),
			*avax.Avalanche().GetValidatorReward(),
			//*bsc.Binance().GetValidatorReward(api),
			*polygon.Polygon().GetValidatorReward())
	default:
		fmt.Println("Вы ввели неверный флаг")
		os.Exit(0)
	}

	for _, crypto := range cryptos {
		fmt.Printf("Platform: %s \n", crypto.Platform)
		fmt.Printf("Validator's slashes: %s\n", crypto.Slashes)
		fmt.Printf("Validator's reward: %s\n\n", crypto.Reward)
	}
}
