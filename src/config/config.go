package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/namsral/flag"
)

type Config struct {
	Chain                  string
	Endpoint               string
	StakingContractAddress string
	ValidatorAddress       string
	OperatorAddress        string
	NodeId                 string
	ValidatorIndex         string
}

type arrayFlags []string

func (i *arrayFlags) String() string {
	return ""
}

func (i *arrayFlags) Set(value string) error {
	list := strings.Split(value, ",")
	for _, v := range list {
		*i = append(*i, v)
	}

	return nil
}

func NewConfig() *Config {
	config := &Config{}

	flag.StringVar(&config.Chain, "chain", "eth", "type of chain")                                                                       // bsc
	flag.StringVar(&config.Endpoint, "endpoint", "https://rpc.ankr.com/bsc", "url")                                                      // https://rpc.ankr.com/bsc
	flag.StringVar(&config.StakingContractAddress, "contract-address", "0x0000000000000000000000000000000000001001", "contract address") // 0x0000000000000000000000000000000000001000
	flag.StringVar(&config.ValidatorAddress, "validator-address", "0x0AA7Aa665276A96acD25329354FeEa8F955CAf2b", "validator address")     // bva1xnudjls7x4p48qrk0j247htt7rl2k2dzp3mr3j
	flag.StringVar(&config.OperatorAddress, "operator-address", "bva1xnudjls7x4p48qrk0j247htt7rl2k2dzp3mr3j", "operator address")        // 9F8cCdaFCc39F3c7D6EBf637c9151673CBc36b88
	flag.StringVar(&config.NodeId, "node-id", "31", "avalanche node id")                                                                 // 22
	flag.StringVar(&config.ValidatorIndex, "validator-index", "2000", "ethereum and polygon validators index")                           // 2000
	help := flag.Bool("help", false, "Команды запуска программы с флагов")
	flag.Parse()

	if *help {
		outputHelp()
		os.Exit(0)
	}

	return config
}

func outputHelp() {
	fmt.Println("COMMANDS:")
	fmt.Println("\n   --chain    (default: eth)")
	fmt.Println(`        type of chain from this list: "eth", "bsc", "avax", "pol", "ftm"`)
	fmt.Println("\n   --endpoint")
	fmt.Println(`        url needs to get info from smart contract (need only for BSC)`)
	fmt.Println("\n   --validator-address")
	fmt.Println(`        validator's address is required flag for fantom chain`)
	fmt.Println("\n   --node-id")
	fmt.Println("        node id is required flag for avalanche chain")
	fmt.Println("\n   --validator-index")
	fmt.Println("        validator index is required flag for ethereum and polygon chains")
	fmt.Println("\n   --operator-address")
	fmt.Println(`        operator's address is required flag for BSC`)
}
