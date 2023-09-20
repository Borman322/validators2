package fantom

import (
	"fmt"
	validator "validators2/src/contract/fantom"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetValidatorReward() (string, error) {
	const api = "https://rpc.ftm.tools/"

	client, err := ethclient.Dial(api)
	if err != nil {
		fmt.Println(err)
	}

	contractAddress := common.HexToAddress("0xFC00FACE00000000000000000000000000000000")
	validatorAddress := common.HexToAddress("0x0AA7Aa665276A96acD25329354FeEa8F955CAf2b")

	pohuiContract, err := validator.NewPohui(contractAddress, client)
	if err != nil {
		fmt.Println(err)
	}

	validatorID, err := pohuiContract.GetValidatorID(&bind.CallOpts{}, validatorAddress)
	if err != nil {
		fmt.Println(err)
	}

	pendingRewards, err := pohuiContract.PendingRewards(&bind.CallOpts{}, validatorAddress, validatorID)
	if err != nil {
		fmt.Println(err)
	}

	var value = pendingRewards.String()
	insertIndex := len(value) - 18
	result := value[:insertIndex] + "." + value[insertIndex:insertIndex+4]

	return result, nil
}
