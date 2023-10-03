package fantom

import (
	"errors"
	"fmt"
	"math/big"
	validator "validators2/src/contract/fantom"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetValidatorUptime(address string, allEpochs int64) (float32, error) {
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

	currentEpoch, err := pohuiContract.CurrentEpoch(&bind.CallOpts{})
	if err != nil {
		fmt.Println(err)
	}

	currentEpochInt := currentEpoch.Int64()

	finalEpoch := currentEpochInt - allEpochs

	var workUptime int64
	var offline int64

	for i := currentEpochInt; i >= finalEpoch; i-- {
		epochBigInt := new(big.Int).SetInt64(i)
		accumulatedUptime, err := pohuiContract.GetEpochAccumulatedUptime(nil, epochBigInt, validatorID)
		if err != nil {
			return 0, errors.New(err.Error())
		}
		workUptime += accumulatedUptime.Int64()

		offlineTime, err := pohuiContract.GetEpochOfflineTime(nil, epochBigInt, validatorID)
		if err != nil {
			return 0, errors.New(err.Error())
		}

		offline += offlineTime.Int64()
	}

	allTime := workUptime + offline
	uptime := (float64(workUptime) * 100) / float64(allTime)
	return float32(uptime), nil
}
