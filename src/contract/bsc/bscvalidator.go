// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bsc

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BscValidatorABI is the input ABI used to generate the binding from.
const BscValidatorABI = "[{\"type\":\"event\",\"name\":\"crashResponse\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"indicatorCleaned\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"knownResponse\",\"inputs\":[{\"type\":\"uint32\",\"name\":\"code\",\"internalType\":\"uint32\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"paramChange\",\"inputs\":[{\"type\":\"string\",\"name\":\"key\",\"internalType\":\"string\",\"indexed\":false},{\"type\":\"bytes\",\"name\":\"value\",\"internalType\":\"bytes\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"unKnownResponse\",\"inputs\":[{\"type\":\"uint32\",\"name\":\"code\",\"internalType\":\"uint32\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"validatorSlashed\",\"inputs\":[{\"type\":\"address\",\"name\":\"validator\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"BIND_CHANNELID\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"BSC_RELAYER_REWARD\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint32\",\"name\":\"\",\"internalType\":\"uint32\"}],\"name\":\"CODE_OK\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"CROSS_CHAIN_CONTRACT_ADDR\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"DECREASE_RATE\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint32\",\"name\":\"\",\"internalType\":\"uint32\"}],\"name\":\"ERROR_FAIL_DECODE\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"FELONY_THRESHOLD\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"GOV_CHANNELID\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"GOV_HUB_ADDR\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"INCENTIVIZE_ADDR\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"LIGHT_CLIENT_ADDR\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"MISDEMEANOR_THRESHOLD\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"RELAYERHUB_CONTRACT_ADDR\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"SLASH_CHANNELID\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"SLASH_CONTRACT_ADDR\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"STAKING_CHANNELID\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"SYSTEM_REWARD_ADDR\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"TOKEN_HUB_ADDR\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"TOKEN_MANAGER_ADDR\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"TRANSFER_IN_CHANNELID\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"TRANSFER_OUT_CHANNELID\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"VALIDATOR_CONTRACT_ADDR\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"alreadyInit\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint16\",\"name\":\"\",\"internalType\":\"uint16\"}],\"name\":\"bscChainID\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"clean\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"felonyThreshold\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getSlashIndicator\",\"inputs\":[{\"type\":\"address\",\"name\":\"validator\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"handleAckPackage\",\"inputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"},{\"type\":\"bytes\",\"name\":\"msgBytes\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"handleFailAckPackage\",\"inputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"},{\"type\":\"bytes\",\"name\":\"\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bytes\",\"name\":\"\",\"internalType\":\"bytes\"}],\"name\":\"handleSynPackage\",\"inputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"},{\"type\":\"bytes\",\"name\":\"\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"height\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"count\",\"internalType\":\"uint256\"},{\"type\":\"bool\",\"name\":\"exist\",\"internalType\":\"bool\"}],\"name\":\"indicators\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"init\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"misdemeanorThreshold\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"previousHeight\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"slash\",\"inputs\":[{\"type\":\"address\",\"name\":\"validator\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"updateParam\",\"inputs\":[{\"type\":\"string\",\"name\":\"key\",\"internalType\":\"string\"},{\"type\":\"bytes\",\"name\":\"value\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"validators\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}]}]"

// BscValidator is an auto generated Go binding around an Ethereum contract.
type BscValidator struct {
	BscValidatorCaller     // Read-only binding to the contract
	BscValidatorTransactor // Write-only binding to the contract
	BscValidatorFilterer   // Log filterer for contract events
}

// BscValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type BscValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BscValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BscValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BscValidatorSession struct {
	Contract     *BscValidator     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BscValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BscValidatorCallerSession struct {
	Contract *BscValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BscValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BscValidatorTransactorSession struct {
	Contract     *BscValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BscValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type BscValidatorRaw struct {
	Contract *BscValidator // Generic contract binding to access the raw methods on
}

// BscValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BscValidatorCallerRaw struct {
	Contract *BscValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// BscValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BscValidatorTransactorRaw struct {
	Contract *BscValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBscValidator creates a new instance of BscValidator, bound to a specific deployed contract.
func NewBscValidator(address common.Address, backend bind.ContractBackend) (*BscValidator, error) {
	contract, err := bindBscValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BscValidator{BscValidatorCaller: BscValidatorCaller{contract: contract}, BscValidatorTransactor: BscValidatorTransactor{contract: contract}, BscValidatorFilterer: BscValidatorFilterer{contract: contract}}, nil
}

// NewBscValidatorCaller creates a new read-only instance of BscValidator, bound to a specific deployed contract.
func NewBscValidatorCaller(address common.Address, caller bind.ContractCaller) (*BscValidatorCaller, error) {
	contract, err := bindBscValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BscValidatorCaller{contract: contract}, nil
}

// NewBscValidatorTransactor creates a new write-only instance of BscValidator, bound to a specific deployed contract.
func NewBscValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*BscValidatorTransactor, error) {
	contract, err := bindBscValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BscValidatorTransactor{contract: contract}, nil
}

// NewBscValidatorFilterer creates a new log filterer instance of BscValidator, bound to a specific deployed contract.
func NewBscValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*BscValidatorFilterer, error) {
	contract, err := bindBscValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BscValidatorFilterer{contract: contract}, nil
}

// bindBscValidator binds a generic wrapper to an already deployed contract.
func bindBscValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BscValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BscValidator *BscValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BscValidator.Contract.BscValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BscValidator *BscValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscValidator.Contract.BscValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BscValidator *BscValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BscValidator.Contract.BscValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BscValidator *BscValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BscValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BscValidator *BscValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BscValidator *BscValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BscValidator.Contract.contract.Transact(opts, method, params...)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_BscValidator *BscValidatorCaller) BINDCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "BIND_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_BscValidator *BscValidatorSession) BINDCHANNELID() (uint8, error) {
	return _BscValidator.Contract.BINDCHANNELID(&_BscValidator.CallOpts)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_BscValidator *BscValidatorCallerSession) BINDCHANNELID() (uint8, error) {
	return _BscValidator.Contract.BINDCHANNELID(&_BscValidator.CallOpts)
}

// BSCRELAYERREWARD is a free data retrieval call binding the contract method 0x9bc8e4f2.
//
// Solidity: function BSC_RELAYER_REWARD() view returns(uint256)
func (_BscValidator *BscValidatorCaller) BSCRELAYERREWARD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "BSC_RELAYER_REWARD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BSCRELAYERREWARD is a free data retrieval call binding the contract method 0x9bc8e4f2.
//
// Solidity: function BSC_RELAYER_REWARD() view returns(uint256)
func (_BscValidator *BscValidatorSession) BSCRELAYERREWARD() (*big.Int, error) {
	return _BscValidator.Contract.BSCRELAYERREWARD(&_BscValidator.CallOpts)
}

// BSCRELAYERREWARD is a free data retrieval call binding the contract method 0x9bc8e4f2.
//
// Solidity: function BSC_RELAYER_REWARD() view returns(uint256)
func (_BscValidator *BscValidatorCallerSession) BSCRELAYERREWARD() (*big.Int, error) {
	return _BscValidator.Contract.BSCRELAYERREWARD(&_BscValidator.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_BscValidator *BscValidatorCaller) CODEOK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "CODE_OK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_BscValidator *BscValidatorSession) CODEOK() (uint32, error) {
	return _BscValidator.Contract.CODEOK(&_BscValidator.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_BscValidator *BscValidatorCallerSession) CODEOK() (uint32, error) {
	return _BscValidator.Contract.CODEOK(&_BscValidator.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_BscValidator *BscValidatorCaller) CROSSCHAINCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "CROSS_CHAIN_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_BscValidator *BscValidatorSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _BscValidator.Contract.CROSSCHAINCONTRACTADDR(&_BscValidator.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_BscValidator *BscValidatorCallerSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _BscValidator.Contract.CROSSCHAINCONTRACTADDR(&_BscValidator.CallOpts)
}

// DECREASERATE is a free data retrieval call binding the contract method 0xac0af629.
//
// Solidity: function DECREASE_RATE() view returns(uint256)
func (_BscValidator *BscValidatorCaller) DECREASERATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "DECREASE_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DECREASERATE is a free data retrieval call binding the contract method 0xac0af629.
//
// Solidity: function DECREASE_RATE() view returns(uint256)
func (_BscValidator *BscValidatorSession) DECREASERATE() (*big.Int, error) {
	return _BscValidator.Contract.DECREASERATE(&_BscValidator.CallOpts)
}

// DECREASERATE is a free data retrieval call binding the contract method 0xac0af629.
//
// Solidity: function DECREASE_RATE() view returns(uint256)
func (_BscValidator *BscValidatorCallerSession) DECREASERATE() (*big.Int, error) {
	return _BscValidator.Contract.DECREASERATE(&_BscValidator.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_BscValidator *BscValidatorCaller) ERRORFAILDECODE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "ERROR_FAIL_DECODE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_BscValidator *BscValidatorSession) ERRORFAILDECODE() (uint32, error) {
	return _BscValidator.Contract.ERRORFAILDECODE(&_BscValidator.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_BscValidator *BscValidatorCallerSession) ERRORFAILDECODE() (uint32, error) {
	return _BscValidator.Contract.ERRORFAILDECODE(&_BscValidator.CallOpts)
}

// FELONYTHRESHOLD is a free data retrieval call binding the contract method 0xc80d4b8f.
//
// Solidity: function FELONY_THRESHOLD() view returns(uint256)
func (_BscValidator *BscValidatorCaller) FELONYTHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "FELONY_THRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FELONYTHRESHOLD is a free data retrieval call binding the contract method 0xc80d4b8f.
//
// Solidity: function FELONY_THRESHOLD() view returns(uint256)
func (_BscValidator *BscValidatorSession) FELONYTHRESHOLD() (*big.Int, error) {
	return _BscValidator.Contract.FELONYTHRESHOLD(&_BscValidator.CallOpts)
}

// FELONYTHRESHOLD is a free data retrieval call binding the contract method 0xc80d4b8f.
//
// Solidity: function FELONY_THRESHOLD() view returns(uint256)
func (_BscValidator *BscValidatorCallerSession) FELONYTHRESHOLD() (*big.Int, error) {
	return _BscValidator.Contract.FELONYTHRESHOLD(&_BscValidator.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_BscValidator *BscValidatorCaller) GOVCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "GOV_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_BscValidator *BscValidatorSession) GOVCHANNELID() (uint8, error) {
	return _BscValidator.Contract.GOVCHANNELID(&_BscValidator.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_BscValidator *BscValidatorCallerSession) GOVCHANNELID() (uint8, error) {
	return _BscValidator.Contract.GOVCHANNELID(&_BscValidator.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_BscValidator *BscValidatorCaller) GOVHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "GOV_HUB_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_BscValidator *BscValidatorSession) GOVHUBADDR() (common.Address, error) {
	return _BscValidator.Contract.GOVHUBADDR(&_BscValidator.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_BscValidator *BscValidatorCallerSession) GOVHUBADDR() (common.Address, error) {
	return _BscValidator.Contract.GOVHUBADDR(&_BscValidator.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_BscValidator *BscValidatorCaller) INCENTIVIZEADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "INCENTIVIZE_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_BscValidator *BscValidatorSession) INCENTIVIZEADDR() (common.Address, error) {
	return _BscValidator.Contract.INCENTIVIZEADDR(&_BscValidator.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_BscValidator *BscValidatorCallerSession) INCENTIVIZEADDR() (common.Address, error) {
	return _BscValidator.Contract.INCENTIVIZEADDR(&_BscValidator.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_BscValidator *BscValidatorCaller) LIGHTCLIENTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "LIGHT_CLIENT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_BscValidator *BscValidatorSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _BscValidator.Contract.LIGHTCLIENTADDR(&_BscValidator.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_BscValidator *BscValidatorCallerSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _BscValidator.Contract.LIGHTCLIENTADDR(&_BscValidator.CallOpts)
}

// MISDEMEANORTHRESHOLD is a free data retrieval call binding the contract method 0x7912a65d.
//
// Solidity: function MISDEMEANOR_THRESHOLD() view returns(uint256)
func (_BscValidator *BscValidatorCaller) MISDEMEANORTHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "MISDEMEANOR_THRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MISDEMEANORTHRESHOLD is a free data retrieval call binding the contract method 0x7912a65d.
//
// Solidity: function MISDEMEANOR_THRESHOLD() view returns(uint256)
func (_BscValidator *BscValidatorSession) MISDEMEANORTHRESHOLD() (*big.Int, error) {
	return _BscValidator.Contract.MISDEMEANORTHRESHOLD(&_BscValidator.CallOpts)
}

// MISDEMEANORTHRESHOLD is a free data retrieval call binding the contract method 0x7912a65d.
//
// Solidity: function MISDEMEANOR_THRESHOLD() view returns(uint256)
func (_BscValidator *BscValidatorCallerSession) MISDEMEANORTHRESHOLD() (*big.Int, error) {
	return _BscValidator.Contract.MISDEMEANORTHRESHOLD(&_BscValidator.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_BscValidator *BscValidatorCaller) RELAYERHUBCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "RELAYERHUB_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_BscValidator *BscValidatorSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _BscValidator.Contract.RELAYERHUBCONTRACTADDR(&_BscValidator.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_BscValidator *BscValidatorCallerSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _BscValidator.Contract.RELAYERHUBCONTRACTADDR(&_BscValidator.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_BscValidator *BscValidatorCaller) SLASHCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "SLASH_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_BscValidator *BscValidatorSession) SLASHCHANNELID() (uint8, error) {
	return _BscValidator.Contract.SLASHCHANNELID(&_BscValidator.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_BscValidator *BscValidatorCallerSession) SLASHCHANNELID() (uint8, error) {
	return _BscValidator.Contract.SLASHCHANNELID(&_BscValidator.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_BscValidator *BscValidatorCaller) SLASHCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "SLASH_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_BscValidator *BscValidatorSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _BscValidator.Contract.SLASHCONTRACTADDR(&_BscValidator.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_BscValidator *BscValidatorCallerSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _BscValidator.Contract.SLASHCONTRACTADDR(&_BscValidator.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_BscValidator *BscValidatorCaller) STAKINGCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "STAKING_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_BscValidator *BscValidatorSession) STAKINGCHANNELID() (uint8, error) {
	return _BscValidator.Contract.STAKINGCHANNELID(&_BscValidator.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_BscValidator *BscValidatorCallerSession) STAKINGCHANNELID() (uint8, error) {
	return _BscValidator.Contract.STAKINGCHANNELID(&_BscValidator.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_BscValidator *BscValidatorCaller) SYSTEMREWARDADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "SYSTEM_REWARD_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_BscValidator *BscValidatorSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _BscValidator.Contract.SYSTEMREWARDADDR(&_BscValidator.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_BscValidator *BscValidatorCallerSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _BscValidator.Contract.SYSTEMREWARDADDR(&_BscValidator.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_BscValidator *BscValidatorCaller) TOKENHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "TOKEN_HUB_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_BscValidator *BscValidatorSession) TOKENHUBADDR() (common.Address, error) {
	return _BscValidator.Contract.TOKENHUBADDR(&_BscValidator.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_BscValidator *BscValidatorCallerSession) TOKENHUBADDR() (common.Address, error) {
	return _BscValidator.Contract.TOKENHUBADDR(&_BscValidator.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_BscValidator *BscValidatorCaller) TOKENMANAGERADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "TOKEN_MANAGER_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_BscValidator *BscValidatorSession) TOKENMANAGERADDR() (common.Address, error) {
	return _BscValidator.Contract.TOKENMANAGERADDR(&_BscValidator.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_BscValidator *BscValidatorCallerSession) TOKENMANAGERADDR() (common.Address, error) {
	return _BscValidator.Contract.TOKENMANAGERADDR(&_BscValidator.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_BscValidator *BscValidatorCaller) TRANSFERINCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "TRANSFER_IN_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_BscValidator *BscValidatorSession) TRANSFERINCHANNELID() (uint8, error) {
	return _BscValidator.Contract.TRANSFERINCHANNELID(&_BscValidator.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_BscValidator *BscValidatorCallerSession) TRANSFERINCHANNELID() (uint8, error) {
	return _BscValidator.Contract.TRANSFERINCHANNELID(&_BscValidator.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_BscValidator *BscValidatorCaller) TRANSFEROUTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "TRANSFER_OUT_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_BscValidator *BscValidatorSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _BscValidator.Contract.TRANSFEROUTCHANNELID(&_BscValidator.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_BscValidator *BscValidatorCallerSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _BscValidator.Contract.TRANSFEROUTCHANNELID(&_BscValidator.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_BscValidator *BscValidatorCaller) VALIDATORCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "VALIDATOR_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_BscValidator *BscValidatorSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _BscValidator.Contract.VALIDATORCONTRACTADDR(&_BscValidator.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_BscValidator *BscValidatorCallerSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _BscValidator.Contract.VALIDATORCONTRACTADDR(&_BscValidator.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_BscValidator *BscValidatorCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "alreadyInit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_BscValidator *BscValidatorSession) AlreadyInit() (bool, error) {
	return _BscValidator.Contract.AlreadyInit(&_BscValidator.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_BscValidator *BscValidatorCallerSession) AlreadyInit() (bool, error) {
	return _BscValidator.Contract.AlreadyInit(&_BscValidator.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_BscValidator *BscValidatorCaller) BscChainID(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "bscChainID")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_BscValidator *BscValidatorSession) BscChainID() (uint16, error) {
	return _BscValidator.Contract.BscChainID(&_BscValidator.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_BscValidator *BscValidatorCallerSession) BscChainID() (uint16, error) {
	return _BscValidator.Contract.BscChainID(&_BscValidator.CallOpts)
}

// FelonyThreshold is a free data retrieval call binding the contract method 0x389f4f71.
//
// Solidity: function felonyThreshold() view returns(uint256)
func (_BscValidator *BscValidatorCaller) FelonyThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "felonyThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FelonyThreshold is a free data retrieval call binding the contract method 0x389f4f71.
//
// Solidity: function felonyThreshold() view returns(uint256)
func (_BscValidator *BscValidatorSession) FelonyThreshold() (*big.Int, error) {
	return _BscValidator.Contract.FelonyThreshold(&_BscValidator.CallOpts)
}

// FelonyThreshold is a free data retrieval call binding the contract method 0x389f4f71.
//
// Solidity: function felonyThreshold() view returns(uint256)
func (_BscValidator *BscValidatorCallerSession) FelonyThreshold() (*big.Int, error) {
	return _BscValidator.Contract.FelonyThreshold(&_BscValidator.CallOpts)
}

// GetSlashIndicator is a free data retrieval call binding the contract method 0x37c8dab9.
//
// Solidity: function getSlashIndicator(address validator) view returns(uint256, uint256)
func (_BscValidator *BscValidatorCaller) GetSlashIndicator(opts *bind.CallOpts, validator common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "getSlashIndicator", validator)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetSlashIndicator is a free data retrieval call binding the contract method 0x37c8dab9.
//
// Solidity: function getSlashIndicator(address validator) view returns(uint256, uint256)
func (_BscValidator *BscValidatorSession) GetSlashIndicator(validator common.Address) (*big.Int, *big.Int, error) {
	return _BscValidator.Contract.GetSlashIndicator(&_BscValidator.CallOpts, validator)
}

// GetSlashIndicator is a free data retrieval call binding the contract method 0x37c8dab9.
//
// Solidity: function getSlashIndicator(address validator) view returns(uint256, uint256)
func (_BscValidator *BscValidatorCallerSession) GetSlashIndicator(validator common.Address) (*big.Int, *big.Int, error) {
	return _BscValidator.Contract.GetSlashIndicator(&_BscValidator.CallOpts, validator)
}

// Indicators is a free data retrieval call binding the contract method 0x23bac5a2.
//
// Solidity: function indicators(address ) view returns(uint256 height, uint256 count, bool exist)
func (_BscValidator *BscValidatorCaller) Indicators(opts *bind.CallOpts, arg0 common.Address) (struct {
	Height *big.Int
	Count  *big.Int
	Exist  bool
}, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "indicators", arg0)

	outstruct := new(struct {
		Height *big.Int
		Count  *big.Int
		Exist  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Height = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Count = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Exist = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Indicators is a free data retrieval call binding the contract method 0x23bac5a2.
//
// Solidity: function indicators(address ) view returns(uint256 height, uint256 count, bool exist)
func (_BscValidator *BscValidatorSession) Indicators(arg0 common.Address) (struct {
	Height *big.Int
	Count  *big.Int
	Exist  bool
}, error) {
	return _BscValidator.Contract.Indicators(&_BscValidator.CallOpts, arg0)
}

// Indicators is a free data retrieval call binding the contract method 0x23bac5a2.
//
// Solidity: function indicators(address ) view returns(uint256 height, uint256 count, bool exist)
func (_BscValidator *BscValidatorCallerSession) Indicators(arg0 common.Address) (struct {
	Height *big.Int
	Count  *big.Int
	Exist  bool
}, error) {
	return _BscValidator.Contract.Indicators(&_BscValidator.CallOpts, arg0)
}

// MisdemeanorThreshold is a free data retrieval call binding the contract method 0x567a372d.
//
// Solidity: function misdemeanorThreshold() view returns(uint256)
func (_BscValidator *BscValidatorCaller) MisdemeanorThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "misdemeanorThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MisdemeanorThreshold is a free data retrieval call binding the contract method 0x567a372d.
//
// Solidity: function misdemeanorThreshold() view returns(uint256)
func (_BscValidator *BscValidatorSession) MisdemeanorThreshold() (*big.Int, error) {
	return _BscValidator.Contract.MisdemeanorThreshold(&_BscValidator.CallOpts)
}

// MisdemeanorThreshold is a free data retrieval call binding the contract method 0x567a372d.
//
// Solidity: function misdemeanorThreshold() view returns(uint256)
func (_BscValidator *BscValidatorCallerSession) MisdemeanorThreshold() (*big.Int, error) {
	return _BscValidator.Contract.MisdemeanorThreshold(&_BscValidator.CallOpts)
}

// PreviousHeight is a free data retrieval call binding the contract method 0x62b72cf5.
//
// Solidity: function previousHeight() view returns(uint256)
func (_BscValidator *BscValidatorCaller) PreviousHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "previousHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviousHeight is a free data retrieval call binding the contract method 0x62b72cf5.
//
// Solidity: function previousHeight() view returns(uint256)
func (_BscValidator *BscValidatorSession) PreviousHeight() (*big.Int, error) {
	return _BscValidator.Contract.PreviousHeight(&_BscValidator.CallOpts)
}

// PreviousHeight is a free data retrieval call binding the contract method 0x62b72cf5.
//
// Solidity: function previousHeight() view returns(uint256)
func (_BscValidator *BscValidatorCallerSession) PreviousHeight() (*big.Int, error) {
	return _BscValidator.Contract.PreviousHeight(&_BscValidator.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address)
func (_BscValidator *BscValidatorCaller) Validators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BscValidator.contract.Call(opts, &out, "validators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address)
func (_BscValidator *BscValidatorSession) Validators(arg0 *big.Int) (common.Address, error) {
	return _BscValidator.Contract.Validators(&_BscValidator.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address)
func (_BscValidator *BscValidatorCallerSession) Validators(arg0 *big.Int) (common.Address, error) {
	return _BscValidator.Contract.Validators(&_BscValidator.CallOpts, arg0)
}

// Clean is a paid mutator transaction binding the contract method 0xfc4333cd.
//
// Solidity: function clean() returns()
func (_BscValidator *BscValidatorTransactor) Clean(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscValidator.contract.Transact(opts, "clean")
}

// Clean is a paid mutator transaction binding the contract method 0xfc4333cd.
//
// Solidity: function clean() returns()
func (_BscValidator *BscValidatorSession) Clean() (*types.Transaction, error) {
	return _BscValidator.Contract.Clean(&_BscValidator.TransactOpts)
}

// Clean is a paid mutator transaction binding the contract method 0xfc4333cd.
//
// Solidity: function clean() returns()
func (_BscValidator *BscValidatorTransactorSession) Clean() (*types.Transaction, error) {
	return _BscValidator.Contract.Clean(&_BscValidator.TransactOpts)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 , bytes msgBytes) returns()
func (_BscValidator *BscValidatorTransactor) HandleAckPackage(opts *bind.TransactOpts, arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BscValidator.contract.Transact(opts, "handleAckPackage", arg0, msgBytes)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 , bytes msgBytes) returns()
func (_BscValidator *BscValidatorSession) HandleAckPackage(arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BscValidator.Contract.HandleAckPackage(&_BscValidator.TransactOpts, arg0, msgBytes)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 , bytes msgBytes) returns()
func (_BscValidator *BscValidatorTransactorSession) HandleAckPackage(arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BscValidator.Contract.HandleAckPackage(&_BscValidator.TransactOpts, arg0, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 , bytes ) returns()
func (_BscValidator *BscValidatorTransactor) HandleFailAckPackage(opts *bind.TransactOpts, arg0 uint8, arg1 []byte) (*types.Transaction, error) {
	return _BscValidator.contract.Transact(opts, "handleFailAckPackage", arg0, arg1)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 , bytes ) returns()
func (_BscValidator *BscValidatorSession) HandleFailAckPackage(arg0 uint8, arg1 []byte) (*types.Transaction, error) {
	return _BscValidator.Contract.HandleFailAckPackage(&_BscValidator.TransactOpts, arg0, arg1)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 , bytes ) returns()
func (_BscValidator *BscValidatorTransactorSession) HandleFailAckPackage(arg0 uint8, arg1 []byte) (*types.Transaction, error) {
	return _BscValidator.Contract.HandleFailAckPackage(&_BscValidator.TransactOpts, arg0, arg1)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes ) returns(bytes)
func (_BscValidator *BscValidatorTransactor) HandleSynPackage(opts *bind.TransactOpts, arg0 uint8, arg1 []byte) (*types.Transaction, error) {
	return _BscValidator.contract.Transact(opts, "handleSynPackage", arg0, arg1)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes ) returns(bytes)
func (_BscValidator *BscValidatorSession) HandleSynPackage(arg0 uint8, arg1 []byte) (*types.Transaction, error) {
	return _BscValidator.Contract.HandleSynPackage(&_BscValidator.TransactOpts, arg0, arg1)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes ) returns(bytes)
func (_BscValidator *BscValidatorTransactorSession) HandleSynPackage(arg0 uint8, arg1 []byte) (*types.Transaction, error) {
	return _BscValidator.Contract.HandleSynPackage(&_BscValidator.TransactOpts, arg0, arg1)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_BscValidator *BscValidatorTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscValidator.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_BscValidator *BscValidatorSession) Init() (*types.Transaction, error) {
	return _BscValidator.Contract.Init(&_BscValidator.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_BscValidator *BscValidatorTransactorSession) Init() (*types.Transaction, error) {
	return _BscValidator.Contract.Init(&_BscValidator.TransactOpts)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address validator) returns()
func (_BscValidator *BscValidatorTransactor) Slash(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _BscValidator.contract.Transact(opts, "slash", validator)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address validator) returns()
func (_BscValidator *BscValidatorSession) Slash(validator common.Address) (*types.Transaction, error) {
	return _BscValidator.Contract.Slash(&_BscValidator.TransactOpts, validator)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address validator) returns()
func (_BscValidator *BscValidatorTransactorSession) Slash(validator common.Address) (*types.Transaction, error) {
	return _BscValidator.Contract.Slash(&_BscValidator.TransactOpts, validator)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_BscValidator *BscValidatorTransactor) UpdateParam(opts *bind.TransactOpts, key string, value []byte) (*types.Transaction, error) {
	return _BscValidator.contract.Transact(opts, "updateParam", key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_BscValidator *BscValidatorSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _BscValidator.Contract.UpdateParam(&_BscValidator.TransactOpts, key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_BscValidator *BscValidatorTransactorSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _BscValidator.Contract.UpdateParam(&_BscValidator.TransactOpts, key, value)
}

// BscValidatorCrashResponseIterator is returned from FilterCrashResponse and is used to iterate over the raw logs and unpacked data for CrashResponse events raised by the BscValidator contract.
type BscValidatorCrashResponseIterator struct {
	Event *BscValidatorCrashResponse // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BscValidatorCrashResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscValidatorCrashResponse)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BscValidatorCrashResponse)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BscValidatorCrashResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscValidatorCrashResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscValidatorCrashResponse represents a CrashResponse event raised by the BscValidator contract.
type BscValidatorCrashResponse struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCrashResponse is a free log retrieval operation binding the contract event 0x07db600eebe2ac176be8dcebad61858c245a4961bb32ca2aa3d159b09aa0810e.
//
// Solidity: event crashResponse()
func (_BscValidator *BscValidatorFilterer) FilterCrashResponse(opts *bind.FilterOpts) (*BscValidatorCrashResponseIterator, error) {

	logs, sub, err := _BscValidator.contract.FilterLogs(opts, "crashResponse")
	if err != nil {
		return nil, err
	}
	return &BscValidatorCrashResponseIterator{contract: _BscValidator.contract, event: "crashResponse", logs: logs, sub: sub}, nil
}

// WatchCrashResponse is a free log subscription operation binding the contract event 0x07db600eebe2ac176be8dcebad61858c245a4961bb32ca2aa3d159b09aa0810e.
//
// Solidity: event crashResponse()
func (_BscValidator *BscValidatorFilterer) WatchCrashResponse(opts *bind.WatchOpts, sink chan<- *BscValidatorCrashResponse) (event.Subscription, error) {

	logs, sub, err := _BscValidator.contract.WatchLogs(opts, "crashResponse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscValidatorCrashResponse)
				if err := _BscValidator.contract.UnpackLog(event, "crashResponse", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCrashResponse is a log parse operation binding the contract event 0x07db600eebe2ac176be8dcebad61858c245a4961bb32ca2aa3d159b09aa0810e.
//
// Solidity: event crashResponse()
func (_BscValidator *BscValidatorFilterer) ParseCrashResponse(log types.Log) (*BscValidatorCrashResponse, error) {
	event := new(BscValidatorCrashResponse)
	if err := _BscValidator.contract.UnpackLog(event, "crashResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscValidatorIndicatorCleanedIterator is returned from FilterIndicatorCleaned and is used to iterate over the raw logs and unpacked data for IndicatorCleaned events raised by the BscValidator contract.
type BscValidatorIndicatorCleanedIterator struct {
	Event *BscValidatorIndicatorCleaned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BscValidatorIndicatorCleanedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscValidatorIndicatorCleaned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BscValidatorIndicatorCleaned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BscValidatorIndicatorCleanedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscValidatorIndicatorCleanedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscValidatorIndicatorCleaned represents a IndicatorCleaned event raised by the BscValidator contract.
type BscValidatorIndicatorCleaned struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterIndicatorCleaned is a free log retrieval operation binding the contract event 0xcfdb3b6ccaeccbdc68be3c59c840e3b3c90f0a7c491f5fff1cf56cfda200dd9c.
//
// Solidity: event indicatorCleaned()
func (_BscValidator *BscValidatorFilterer) FilterIndicatorCleaned(opts *bind.FilterOpts) (*BscValidatorIndicatorCleanedIterator, error) {

	logs, sub, err := _BscValidator.contract.FilterLogs(opts, "indicatorCleaned")
	if err != nil {
		return nil, err
	}
	return &BscValidatorIndicatorCleanedIterator{contract: _BscValidator.contract, event: "indicatorCleaned", logs: logs, sub: sub}, nil
}

// WatchIndicatorCleaned is a free log subscription operation binding the contract event 0xcfdb3b6ccaeccbdc68be3c59c840e3b3c90f0a7c491f5fff1cf56cfda200dd9c.
//
// Solidity: event indicatorCleaned()
func (_BscValidator *BscValidatorFilterer) WatchIndicatorCleaned(opts *bind.WatchOpts, sink chan<- *BscValidatorIndicatorCleaned) (event.Subscription, error) {

	logs, sub, err := _BscValidator.contract.WatchLogs(opts, "indicatorCleaned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscValidatorIndicatorCleaned)
				if err := _BscValidator.contract.UnpackLog(event, "indicatorCleaned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIndicatorCleaned is a log parse operation binding the contract event 0xcfdb3b6ccaeccbdc68be3c59c840e3b3c90f0a7c491f5fff1cf56cfda200dd9c.
//
// Solidity: event indicatorCleaned()
func (_BscValidator *BscValidatorFilterer) ParseIndicatorCleaned(log types.Log) (*BscValidatorIndicatorCleaned, error) {
	event := new(BscValidatorIndicatorCleaned)
	if err := _BscValidator.contract.UnpackLog(event, "indicatorCleaned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscValidatorKnownResponseIterator is returned from FilterKnownResponse and is used to iterate over the raw logs and unpacked data for KnownResponse events raised by the BscValidator contract.
type BscValidatorKnownResponseIterator struct {
	Event *BscValidatorKnownResponse // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BscValidatorKnownResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscValidatorKnownResponse)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BscValidatorKnownResponse)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BscValidatorKnownResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscValidatorKnownResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscValidatorKnownResponse represents a KnownResponse event raised by the BscValidator contract.
type BscValidatorKnownResponse struct {
	Code uint32
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterKnownResponse is a free log retrieval operation binding the contract event 0x7f0956d47419b9525356e7111652b653b530ec6f5096dccc04589bc38e629967.
//
// Solidity: event knownResponse(uint32 code)
func (_BscValidator *BscValidatorFilterer) FilterKnownResponse(opts *bind.FilterOpts) (*BscValidatorKnownResponseIterator, error) {

	logs, sub, err := _BscValidator.contract.FilterLogs(opts, "knownResponse")
	if err != nil {
		return nil, err
	}
	return &BscValidatorKnownResponseIterator{contract: _BscValidator.contract, event: "knownResponse", logs: logs, sub: sub}, nil
}

// WatchKnownResponse is a free log subscription operation binding the contract event 0x7f0956d47419b9525356e7111652b653b530ec6f5096dccc04589bc38e629967.
//
// Solidity: event knownResponse(uint32 code)
func (_BscValidator *BscValidatorFilterer) WatchKnownResponse(opts *bind.WatchOpts, sink chan<- *BscValidatorKnownResponse) (event.Subscription, error) {

	logs, sub, err := _BscValidator.contract.WatchLogs(opts, "knownResponse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscValidatorKnownResponse)
				if err := _BscValidator.contract.UnpackLog(event, "knownResponse", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseKnownResponse is a log parse operation binding the contract event 0x7f0956d47419b9525356e7111652b653b530ec6f5096dccc04589bc38e629967.
//
// Solidity: event knownResponse(uint32 code)
func (_BscValidator *BscValidatorFilterer) ParseKnownResponse(log types.Log) (*BscValidatorKnownResponse, error) {
	event := new(BscValidatorKnownResponse)
	if err := _BscValidator.contract.UnpackLog(event, "knownResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscValidatorParamChangeIterator is returned from FilterParamChange and is used to iterate over the raw logs and unpacked data for ParamChange events raised by the BscValidator contract.
type BscValidatorParamChangeIterator struct {
	Event *BscValidatorParamChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BscValidatorParamChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscValidatorParamChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BscValidatorParamChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BscValidatorParamChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscValidatorParamChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscValidatorParamChange represents a ParamChange event raised by the BscValidator contract.
type BscValidatorParamChange struct {
	Key   string
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParamChange is a free log retrieval operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_BscValidator *BscValidatorFilterer) FilterParamChange(opts *bind.FilterOpts) (*BscValidatorParamChangeIterator, error) {

	logs, sub, err := _BscValidator.contract.FilterLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return &BscValidatorParamChangeIterator{contract: _BscValidator.contract, event: "paramChange", logs: logs, sub: sub}, nil
}

// WatchParamChange is a free log subscription operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_BscValidator *BscValidatorFilterer) WatchParamChange(opts *bind.WatchOpts, sink chan<- *BscValidatorParamChange) (event.Subscription, error) {

	logs, sub, err := _BscValidator.contract.WatchLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscValidatorParamChange)
				if err := _BscValidator.contract.UnpackLog(event, "paramChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseParamChange is a log parse operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_BscValidator *BscValidatorFilterer) ParseParamChange(log types.Log) (*BscValidatorParamChange, error) {
	event := new(BscValidatorParamChange)
	if err := _BscValidator.contract.UnpackLog(event, "paramChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscValidatorUnKnownResponseIterator is returned from FilterUnKnownResponse and is used to iterate over the raw logs and unpacked data for UnKnownResponse events raised by the BscValidator contract.
type BscValidatorUnKnownResponseIterator struct {
	Event *BscValidatorUnKnownResponse // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BscValidatorUnKnownResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscValidatorUnKnownResponse)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BscValidatorUnKnownResponse)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BscValidatorUnKnownResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscValidatorUnKnownResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscValidatorUnKnownResponse represents a UnKnownResponse event raised by the BscValidator contract.
type BscValidatorUnKnownResponse struct {
	Code uint32
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUnKnownResponse is a free log retrieval operation binding the contract event 0x7d45f62d17443dd4547bca8a8112c60e2385669318dc300ec61a5d2492f262e7.
//
// Solidity: event unKnownResponse(uint32 code)
func (_BscValidator *BscValidatorFilterer) FilterUnKnownResponse(opts *bind.FilterOpts) (*BscValidatorUnKnownResponseIterator, error) {

	logs, sub, err := _BscValidator.contract.FilterLogs(opts, "unKnownResponse")
	if err != nil {
		return nil, err
	}
	return &BscValidatorUnKnownResponseIterator{contract: _BscValidator.contract, event: "unKnownResponse", logs: logs, sub: sub}, nil
}

// WatchUnKnownResponse is a free log subscription operation binding the contract event 0x7d45f62d17443dd4547bca8a8112c60e2385669318dc300ec61a5d2492f262e7.
//
// Solidity: event unKnownResponse(uint32 code)
func (_BscValidator *BscValidatorFilterer) WatchUnKnownResponse(opts *bind.WatchOpts, sink chan<- *BscValidatorUnKnownResponse) (event.Subscription, error) {

	logs, sub, err := _BscValidator.contract.WatchLogs(opts, "unKnownResponse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscValidatorUnKnownResponse)
				if err := _BscValidator.contract.UnpackLog(event, "unKnownResponse", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnKnownResponse is a log parse operation binding the contract event 0x7d45f62d17443dd4547bca8a8112c60e2385669318dc300ec61a5d2492f262e7.
//
// Solidity: event unKnownResponse(uint32 code)
func (_BscValidator *BscValidatorFilterer) ParseUnKnownResponse(log types.Log) (*BscValidatorUnKnownResponse, error) {
	event := new(BscValidatorUnKnownResponse)
	if err := _BscValidator.contract.UnpackLog(event, "unKnownResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscValidatorValidatorSlashedIterator is returned from FilterValidatorSlashed and is used to iterate over the raw logs and unpacked data for ValidatorSlashed events raised by the BscValidator contract.
type BscValidatorValidatorSlashedIterator struct {
	Event *BscValidatorValidatorSlashed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BscValidatorValidatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscValidatorValidatorSlashed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BscValidatorValidatorSlashed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BscValidatorValidatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscValidatorValidatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscValidatorValidatorSlashed represents a ValidatorSlashed event raised by the BscValidator contract.
type BscValidatorValidatorSlashed struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorSlashed is a free log retrieval operation binding the contract event 0xddb6012116e51abf5436d956a4f0ebd927e92c576ff96d7918290c8782291e3e.
//
// Solidity: event validatorSlashed(address indexed validator)
func (_BscValidator *BscValidatorFilterer) FilterValidatorSlashed(opts *bind.FilterOpts, validator []common.Address) (*BscValidatorValidatorSlashedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BscValidator.contract.FilterLogs(opts, "validatorSlashed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &BscValidatorValidatorSlashedIterator{contract: _BscValidator.contract, event: "validatorSlashed", logs: logs, sub: sub}, nil
}

// WatchValidatorSlashed is a free log subscription operation binding the contract event 0xddb6012116e51abf5436d956a4f0ebd927e92c576ff96d7918290c8782291e3e.
//
// Solidity: event validatorSlashed(address indexed validator)
func (_BscValidator *BscValidatorFilterer) WatchValidatorSlashed(opts *bind.WatchOpts, sink chan<- *BscValidatorValidatorSlashed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BscValidator.contract.WatchLogs(opts, "validatorSlashed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscValidatorValidatorSlashed)
				if err := _BscValidator.contract.UnpackLog(event, "validatorSlashed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorSlashed is a log parse operation binding the contract event 0xddb6012116e51abf5436d956a4f0ebd927e92c576ff96d7918290c8782291e3e.
//
// Solidity: event validatorSlashed(address indexed validator)
func (_BscValidator *BscValidatorFilterer) ParseValidatorSlashed(log types.Log) (*BscValidatorValidatorSlashed, error) {
	event := new(BscValidatorValidatorSlashed)
	if err := _BscValidator.contract.UnpackLog(event, "validatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}