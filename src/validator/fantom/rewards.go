package fantom

import (
	"fmt"
	"math/big"
	validator "validators2/src/contract/fantom"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetValidatorReward(address string) (string, error) {
	const api = "https://rpc.ftm.tools/"

	client, err := ethclient.Dial(api)
	if err != nil {
		fmt.Println(err)
	}

	contractAddress := common.HexToAddress("0xFC00FACE00000000000000000000000000000000")
	validatorAddress := common.HexToAddress(address)

	pohuiContract, err := validator.NewPohui(contractAddress, client)
	if err != nil {
		fmt.Println(err)
	}

	validatorID, err := pohuiContract.GetValidatorID(&bind.CallOpts{}, validatorAddress)
	if err != nil {
		return "", fmt.Errorf("FTM validator: address is not correct")
	}

	if validatorID.Cmp(big.NewInt(0)) == 0 {
		return "", fmt.Errorf("FTM validator: address is not correct")
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
