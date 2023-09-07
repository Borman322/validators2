package fantom

import (
	"fmt"
	"math/big"
	validators2 "validators2/interface"
	"validators2/interface/fantom/validator"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ValidatorReward struct {
	Platform string
	Slashes  string
	Reward   string
}

func Fantom() *ValidatorReward {
	var ftm = ValidatorReward{}
	return &ftm
}

func (v ValidatorReward) GetValidatorReward() *validators2.ValidatorReward {
	const api = "https://rpc.ftm.tools/"

	// Подключение к узлу Ethereum
	client, err := ethclient.Dial(api)
	if err != nil {
		fmt.Println(err)
	}

	// Адрес контракта
	contractAddress := common.HexToAddress("0xFC00FACE00000000000000000000000000000000") // Замените на фактический адрес контракта
	validatorAddress := common.HexToAddress("0x0AA7Aa665276A96acD25329354FeEa8F955CAf2b")

	// Создайте объект вашего контракта, используя загруженный ABI и клиента Ethereum
	pohuiContract, err := validator.NewPohui(contractAddress, client)
	if err != nil {
		fmt.Println(err)
	}

	validatorID, err := pohuiContract.GetValidatorID(&bind.CallOpts{}, validatorAddress)
	if err != nil {
		fmt.Println(err)
	}

	validatorInfo, err := pohuiContract.GetValidator(&bind.CallOpts{}, validatorID)
	if err != nil {
		fmt.Println(err)
	}

	var zero = big.NewInt(0)

	if validatorInfo.Status.Cmp(zero) == 0 && validatorInfo.DeactivatedTime.Cmp(zero) == 0 {
		// Выполнение вызова метода TotalSupply
		pendingRewards, err := pohuiContract.PendingRewards(&bind.CallOpts{}, validatorAddress, validatorID)
		if err != nil {
			fmt.Println(err)
		}

		var value = pendingRewards.String()
		insertIndex := len(value) - 18
		result := value[:insertIndex] + "." + value[insertIndex:insertIndex+4] + " FTM"

		return &validators2.ValidatorReward{
			Platform: "Fantom",
			Slashes:  "",
			Reward:   result,
		}
	} else {
		return &validators2.ValidatorReward{
			Platform: "Fantom",
			Slashes:  validatorInfo.DeactivatedTime.String(),
			Reward:   "Null",
		}
	}
}
