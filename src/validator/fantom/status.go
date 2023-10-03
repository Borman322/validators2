package fantom

import (
	"errors"
	validator "validators2/src/contract/fantom"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func IsValidatorHealthy(address string) (bool, error) {
	const api = "https://rpc.ftm.tools/"

	client, err := ethclient.Dial(api)
	if err != nil {
		return false, errors.New("FTM validator: " + err.Error())
	}
	contractAddress := common.HexToAddress("0xFC00FACE00000000000000000000000000000000")
	validatorAddress := common.HexToAddress(address)

	pohuiContract, err := validator.NewPohui(contractAddress, client)
	if err != nil {
		return false, errors.New("FTM validator: " + err.Error())
	}

	validatorID, err := pohuiContract.GetValidatorID(&bind.CallOpts{}, validatorAddress)
	if err != nil {
		return false, errors.New("FTM validator: " + err.Error())
	}
	status, err := pohuiContract.GetValidator(&bind.CallOpts{}, validatorID)
	if err != nil {
		return false, errors.New("FTM validator: " + err.Error())
	}
	var result bool
	value := status.Status.Int64()
	if value == 0 {
		result = true
	} else {
		result = false
	}
	return result, nil
}
