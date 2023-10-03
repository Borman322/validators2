package fantom

import (
	"fmt"
	validator "validators2/src/contract/fantom"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func IsValidatorSlashed(address string) (bool, error) {

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
		fmt.Println(err)
	}
	slash, err := pohuiContract.IsSlashed(&bind.CallOpts{}, validatorID)
	if err != nil {
		fmt.Println(err)
	}
	return slash, nil
}
