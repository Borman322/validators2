package config

import (
	"strings"

	"github.com/namsral/flag"
)

type Config struct {
	Chain                  string
	Endpoint               string
	StakingContractAddress string
	ValidatorAddress       string
	ProposerAddress        string
	NodeId                 string
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

	flag.StringVar(&config.Chain, "chain", "evmos", "type of chain")                                                                          // bsc
	flag.StringVar(&config.Endpoint, "endpoint", "http://127.0.0.1:1317", "url")                                                              // https://rpc.ankr.com/bsc
	flag.StringVar(&config.StakingContractAddress, "contract-address", "0x0000000000000000000000000000000000001001", "contract address")      // 0x0000000000000000000000000000000000001000
	flag.StringVar(&config.ValidatorAddress, "validator-address", "evmosvaloper125fkz3mq6qxxpkmphdl3ep92t0d3y969xmt8hz", "validator address") // bva1xnudjls7x4p48qrk0j247htt7rl2k2dzp3mr3j
	flag.StringVar(&config.ProposerAddress, "proposer-address", "639B7F45CCB4417C6536421A5138E711C3CFD3A5", "proposer address")               // 9F8cCdaFCc39F3c7D6EBf637c9151673CBc36b88
	flag.StringVar(&config.NodeId, "node-id", "31", "avalanche node id")                                                                      // 22
	flag.Parse()

	return config
}
