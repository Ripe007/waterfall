// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// CampaignContinuousCyclesABI is the input ABI used to generate the binding from.
const CampaignContinuousCyclesABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_target\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_joinwindow\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"apy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"}],\"internalType\":\"structCampaignContinuousCycles.TrancheParams[]\",\"name\":\"_tranches\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"trancheID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"optionBalance\",\"type\":\"uint256\"}],\"name\":\"ApplyForExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"trancheID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Exit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"trancheID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"optionBalance\",\"type\":\"uint256\"}],\"name\":\"Join\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"}],\"name\":\"NextCycle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"}],\"name\":\"Terminate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawFee\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AlpacaBUSD\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AlpacaFarm\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AlpacaToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BUSD\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CreamFarm\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PancakeRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PercentageParamScale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PercentageScale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VenusFarm\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"investment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"joinCycle\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"actualEndAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"actualStartAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"trancheID\",\"type\":\"uint256\"}],\"name\":\"applyForExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bustTrancheSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bustTranches\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"changeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currency\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cycle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cycleLifetime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"duration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exchangeRateCheckpoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"trancheID\",\"type\":\"uint256\"}],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"trancheID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"join\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"joinWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"launched\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextCycle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextCycleAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatorBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"setAlpacaBUSD\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"setBUSD\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cream\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"s1\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"venus\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"s2\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"alpaca\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"s3\",\"type\":\"uint256\"}],\"name\":\"setFarms\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"setPancakeRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"trancheID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setTranchesFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"target\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"terminate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"terminated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"trancheExitBuf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"trancheFeeInf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"trancheJoinBuf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tranches\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"target\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capital\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"apy\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tranchesTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capital\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CampaignContinuousCycles is an auto generated Go binding around an Ethereum contract.
type CampaignContinuousCycles struct {
	CampaignContinuousCyclesCaller     // Read-only binding to the contract
	CampaignContinuousCyclesTransactor // Write-only binding to the contract
	CampaignContinuousCyclesFilterer   // Log filterer for contract events
}

// CampaignContinuousCyclesCaller is an auto generated read-only Go binding around an Ethereum contract.
type CampaignContinuousCyclesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CampaignContinuousCyclesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CampaignContinuousCyclesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CampaignContinuousCyclesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CampaignContinuousCyclesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CampaignContinuousCyclesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CampaignContinuousCyclesSession struct {
	Contract     *CampaignContinuousCycles // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CampaignContinuousCyclesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CampaignContinuousCyclesCallerSession struct {
	Contract *CampaignContinuousCyclesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// CampaignContinuousCyclesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CampaignContinuousCyclesTransactorSession struct {
	Contract     *CampaignContinuousCyclesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// CampaignContinuousCyclesRaw is an auto generated low-level Go binding around an Ethereum contract.
type CampaignContinuousCyclesRaw struct {
	Contract *CampaignContinuousCycles // Generic contract binding to access the raw methods on
}

// CampaignContinuousCyclesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CampaignContinuousCyclesCallerRaw struct {
	Contract *CampaignContinuousCyclesCaller // Generic read-only contract binding to access the raw methods on
}

// CampaignContinuousCyclesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CampaignContinuousCyclesTransactorRaw struct {
	Contract *CampaignContinuousCyclesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCampaignContinuousCycles creates a new instance of CampaignContinuousCycles, bound to a specific deployed contract.
func NewCampaignContinuousCycles(address common.Address, backend bind.ContractBackend) (*CampaignContinuousCycles, error) {
	contract, err := bindCampaignContinuousCycles(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CampaignContinuousCycles{CampaignContinuousCyclesCaller: CampaignContinuousCyclesCaller{contract: contract}, CampaignContinuousCyclesTransactor: CampaignContinuousCyclesTransactor{contract: contract}, CampaignContinuousCyclesFilterer: CampaignContinuousCyclesFilterer{contract: contract}}, nil
}

// NewCampaignContinuousCyclesCaller creates a new read-only instance of CampaignContinuousCycles, bound to a specific deployed contract.
func NewCampaignContinuousCyclesCaller(address common.Address, caller bind.ContractCaller) (*CampaignContinuousCyclesCaller, error) {
	contract, err := bindCampaignContinuousCycles(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CampaignContinuousCyclesCaller{contract: contract}, nil
}

// NewCampaignContinuousCyclesTransactor creates a new write-only instance of CampaignContinuousCycles, bound to a specific deployed contract.
func NewCampaignContinuousCyclesTransactor(address common.Address, transactor bind.ContractTransactor) (*CampaignContinuousCyclesTransactor, error) {
	contract, err := bindCampaignContinuousCycles(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CampaignContinuousCyclesTransactor{contract: contract}, nil
}

// NewCampaignContinuousCyclesFilterer creates a new log filterer instance of CampaignContinuousCycles, bound to a specific deployed contract.
func NewCampaignContinuousCyclesFilterer(address common.Address, filterer bind.ContractFilterer) (*CampaignContinuousCyclesFilterer, error) {
	contract, err := bindCampaignContinuousCycles(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CampaignContinuousCyclesFilterer{contract: contract}, nil
}

// bindCampaignContinuousCycles binds a generic wrapper to an already deployed contract.
func bindCampaignContinuousCycles(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CampaignContinuousCyclesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CampaignContinuousCycles *CampaignContinuousCyclesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CampaignContinuousCycles.Contract.CampaignContinuousCyclesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CampaignContinuousCycles *CampaignContinuousCyclesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.CampaignContinuousCyclesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CampaignContinuousCycles *CampaignContinuousCyclesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.CampaignContinuousCyclesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CampaignContinuousCycles.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.contract.Transact(opts, method, params...)
}

// AlpacaBUSD is a free data retrieval call binding the contract method 0x9c87b197.
//
// Solidity: function AlpacaBUSD() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) AlpacaBUSD(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "AlpacaBUSD")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AlpacaBUSD is a free data retrieval call binding the contract method 0x9c87b197.
//
// Solidity: function AlpacaBUSD() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) AlpacaBUSD() (common.Address, error) {
	return _CampaignContinuousCycles.Contract.AlpacaBUSD(&_CampaignContinuousCycles.CallOpts)
}

// AlpacaBUSD is a free data retrieval call binding the contract method 0x9c87b197.
//
// Solidity: function AlpacaBUSD() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) AlpacaBUSD() (common.Address, error) {
	return _CampaignContinuousCycles.Contract.AlpacaBUSD(&_CampaignContinuousCycles.CallOpts)
}

// AlpacaFarm is a free data retrieval call binding the contract method 0xabe07bc7.
//
// Solidity: function AlpacaFarm() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) AlpacaFarm(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "AlpacaFarm")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AlpacaFarm is a free data retrieval call binding the contract method 0xabe07bc7.
//
// Solidity: function AlpacaFarm() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) AlpacaFarm() (common.Address, error) {
	return _CampaignContinuousCycles.Contract.AlpacaFarm(&_CampaignContinuousCycles.CallOpts)
}

// AlpacaFarm is a free data retrieval call binding the contract method 0xabe07bc7.
//
// Solidity: function AlpacaFarm() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) AlpacaFarm() (common.Address, error) {
	return _CampaignContinuousCycles.Contract.AlpacaFarm(&_CampaignContinuousCycles.CallOpts)
}

// AlpacaToken is a free data retrieval call binding the contract method 0xcf8db46b.
//
// Solidity: function AlpacaToken() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) AlpacaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "AlpacaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AlpacaToken is a free data retrieval call binding the contract method 0xcf8db46b.
//
// Solidity: function AlpacaToken() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) AlpacaToken() (common.Address, error) {
	return _CampaignContinuousCycles.Contract.AlpacaToken(&_CampaignContinuousCycles.CallOpts)
}

// AlpacaToken is a free data retrieval call binding the contract method 0xcf8db46b.
//
// Solidity: function AlpacaToken() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) AlpacaToken() (common.Address, error) {
	return _CampaignContinuousCycles.Contract.AlpacaToken(&_CampaignContinuousCycles.CallOpts)
}

// BUSD is a free data retrieval call binding the contract method 0x484f4ea9.
//
// Solidity: function BUSD() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) BUSD(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "BUSD")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BUSD is a free data retrieval call binding the contract method 0x484f4ea9.
//
// Solidity: function BUSD() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) BUSD() (common.Address, error) {
	return _CampaignContinuousCycles.Contract.BUSD(&_CampaignContinuousCycles.CallOpts)
}

// BUSD is a free data retrieval call binding the contract method 0x484f4ea9.
//
// Solidity: function BUSD() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) BUSD() (common.Address, error) {
	return _CampaignContinuousCycles.Contract.BUSD(&_CampaignContinuousCycles.CallOpts)
}

// CreamFarm is a free data retrieval call binding the contract method 0x0d32a551.
//
// Solidity: function CreamFarm() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) CreamFarm(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "CreamFarm")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreamFarm is a free data retrieval call binding the contract method 0x0d32a551.
//
// Solidity: function CreamFarm() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) CreamFarm() (common.Address, error) {
	return _CampaignContinuousCycles.Contract.CreamFarm(&_CampaignContinuousCycles.CallOpts)
}

// CreamFarm is a free data retrieval call binding the contract method 0x0d32a551.
//
// Solidity: function CreamFarm() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) CreamFarm() (common.Address, error) {
	return _CampaignContinuousCycles.Contract.CreamFarm(&_CampaignContinuousCycles.CallOpts)
}

// PancakeRouter is a free data retrieval call binding the contract method 0xeda0228f.
//
// Solidity: function PancakeRouter() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) PancakeRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "PancakeRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PancakeRouter is a free data retrieval call binding the contract method 0xeda0228f.
//
// Solidity: function PancakeRouter() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) PancakeRouter() (common.Address, error) {
	return _CampaignContinuousCycles.Contract.PancakeRouter(&_CampaignContinuousCycles.CallOpts)
}

// PancakeRouter is a free data retrieval call binding the contract method 0xeda0228f.
//
// Solidity: function PancakeRouter() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) PancakeRouter() (common.Address, error) {
	return _CampaignContinuousCycles.Contract.PancakeRouter(&_CampaignContinuousCycles.CallOpts)
}

// PercentageParamScale is a free data retrieval call binding the contract method 0xb94db6af.
//
// Solidity: function PercentageParamScale() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) PercentageParamScale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "PercentageParamScale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PercentageParamScale is a free data retrieval call binding the contract method 0xb94db6af.
//
// Solidity: function PercentageParamScale() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) PercentageParamScale() (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.PercentageParamScale(&_CampaignContinuousCycles.CallOpts)
}

// PercentageParamScale is a free data retrieval call binding the contract method 0xb94db6af.
//
// Solidity: function PercentageParamScale() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) PercentageParamScale() (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.PercentageParamScale(&_CampaignContinuousCycles.CallOpts)
}

// PercentageScale is a free data retrieval call binding the contract method 0x5cd90e58.
//
// Solidity: function PercentageScale() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) PercentageScale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "PercentageScale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PercentageScale is a free data retrieval call binding the contract method 0x5cd90e58.
//
// Solidity: function PercentageScale() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) PercentageScale() (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.PercentageScale(&_CampaignContinuousCycles.CallOpts)
}

// PercentageScale is a free data retrieval call binding the contract method 0x5cd90e58.
//
// Solidity: function PercentageScale() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) PercentageScale() (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.PercentageScale(&_CampaignContinuousCycles.CallOpts)
}

// VenusFarm is a free data retrieval call binding the contract method 0xe06b48ee.
//
// Solidity: function VenusFarm() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) VenusFarm(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "VenusFarm")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VenusFarm is a free data retrieval call binding the contract method 0xe06b48ee.
//
// Solidity: function VenusFarm() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) VenusFarm() (common.Address, error) {
	return _CampaignContinuousCycles.Contract.VenusFarm(&_CampaignContinuousCycles.CallOpts)
}

// VenusFarm is a free data retrieval call binding the contract method 0xe06b48ee.
//
// Solidity: function VenusFarm() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) VenusFarm() (common.Address, error) {
	return _CampaignContinuousCycles.Contract.VenusFarm(&_CampaignContinuousCycles.CallOpts)
}

// Accounts is a free data retrieval call binding the contract method 0x87524581.
//
// Solidity: function accounts(address , uint256 ) view returns(uint256 investment, uint256 joinCycle)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) Accounts(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Investment *big.Int
	JoinCycle  *big.Int
}, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "accounts", arg0, arg1)

	outstruct := new(struct {
		Investment *big.Int
		JoinCycle  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Investment = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.JoinCycle = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Accounts is a free data retrieval call binding the contract method 0x87524581.
//
// Solidity: function accounts(address , uint256 ) view returns(uint256 investment, uint256 joinCycle)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) Accounts(arg0 common.Address, arg1 *big.Int) (struct {
	Investment *big.Int
	JoinCycle  *big.Int
}, error) {
	return _CampaignContinuousCycles.Contract.Accounts(&_CampaignContinuousCycles.CallOpts, arg0, arg1)
}

// Accounts is a free data retrieval call binding the contract method 0x87524581.
//
// Solidity: function accounts(address , uint256 ) view returns(uint256 investment, uint256 joinCycle)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) Accounts(arg0 common.Address, arg1 *big.Int) (struct {
	Investment *big.Int
	JoinCycle  *big.Int
}, error) {
	return _CampaignContinuousCycles.Contract.Accounts(&_CampaignContinuousCycles.CallOpts, arg0, arg1)
}

// ActualEndAt is a free data retrieval call binding the contract method 0x75691a83.
//
// Solidity: function actualEndAt() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) ActualEndAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "actualEndAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActualEndAt is a free data retrieval call binding the contract method 0x75691a83.
//
// Solidity: function actualEndAt() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) ActualEndAt() (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.ActualEndAt(&_CampaignContinuousCycles.CallOpts)
}

// ActualEndAt is a free data retrieval call binding the contract method 0x75691a83.
//
// Solidity: function actualEndAt() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) ActualEndAt() (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.ActualEndAt(&_CampaignContinuousCycles.CallOpts)
}

// ActualStartAt is a free data retrieval call binding the contract method 0x27e79c44.
//
// Solidity: function actualStartAt() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) ActualStartAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "actualStartAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActualStartAt is a free data retrieval call binding the contract method 0x27e79c44.
//
// Solidity: function actualStartAt() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) ActualStartAt() (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.ActualStartAt(&_CampaignContinuousCycles.CallOpts)
}

// ActualStartAt is a free data retrieval call binding the contract method 0x27e79c44.
//
// Solidity: function actualStartAt() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) ActualStartAt() (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.ActualStartAt(&_CampaignContinuousCycles.CallOpts)
}

// BustTrancheSize is a free data retrieval call binding the contract method 0x221006eb.
//
// Solidity: function bustTrancheSize() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) BustTrancheSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "bustTrancheSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BustTrancheSize is a free data retrieval call binding the contract method 0x221006eb.
//
// Solidity: function bustTrancheSize() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) BustTrancheSize() (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.BustTrancheSize(&_CampaignContinuousCycles.CallOpts)
}

// BustTrancheSize is a free data retrieval call binding the contract method 0x221006eb.
//
// Solidity: function bustTrancheSize() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) BustTrancheSize() (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.BustTrancheSize(&_CampaignContinuousCycles.CallOpts)
}

// BustTranches is a free data retrieval call binding the contract method 0xc4914a3c.
//
// Solidity: function bustTranches(uint256 ) view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) BustTranches(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "bustTranches", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BustTranches is a free data retrieval call binding the contract method 0xc4914a3c.
//
// Solidity: function bustTranches(uint256 ) view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) BustTranches(arg0 *big.Int) (common.Address, error) {
	return _CampaignContinuousCycles.Contract.BustTranches(&_CampaignContinuousCycles.CallOpts, arg0)
}

// BustTranches is a free data retrieval call binding the contract method 0xc4914a3c.
//
// Solidity: function bustTranches(uint256 ) view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) BustTranches(arg0 *big.Int) (common.Address, error) {
	return _CampaignContinuousCycles.Contract.BustTranches(&_CampaignContinuousCycles.CallOpts, arg0)
}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) Currency(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "currency")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) Currency() (common.Address, error) {
	return _CampaignContinuousCycles.Contract.Currency(&_CampaignContinuousCycles.CallOpts)
}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) Currency() (common.Address, error) {
	return _CampaignContinuousCycles.Contract.Currency(&_CampaignContinuousCycles.CallOpts)
}

// Cycle is a free data retrieval call binding the contract method 0x6190c9d5.
//
// Solidity: function cycle() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) Cycle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "cycle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cycle is a free data retrieval call binding the contract method 0x6190c9d5.
//
// Solidity: function cycle() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) Cycle() (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.Cycle(&_CampaignContinuousCycles.CallOpts)
}

// Cycle is a free data retrieval call binding the contract method 0x6190c9d5.
//
// Solidity: function cycle() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) Cycle() (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.Cycle(&_CampaignContinuousCycles.CallOpts)
}

// CycleLifetime is a free data retrieval call binding the contract method 0x1ab74502.
//
// Solidity: function cycleLifetime() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) CycleLifetime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "cycleLifetime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CycleLifetime is a free data retrieval call binding the contract method 0x1ab74502.
//
// Solidity: function cycleLifetime() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) CycleLifetime() (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.CycleLifetime(&_CampaignContinuousCycles.CallOpts)
}

// CycleLifetime is a free data retrieval call binding the contract method 0x1ab74502.
//
// Solidity: function cycleLifetime() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) CycleLifetime() (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.CycleLifetime(&_CampaignContinuousCycles.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) Duration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "duration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) Duration() (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.Duration(&_CampaignContinuousCycles.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) Duration() (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.Duration(&_CampaignContinuousCycles.CallOpts)
}

// ExchangeRateCheckpoint is a free data retrieval call binding the contract method 0x30c654a5.
//
// Solidity: function exchangeRateCheckpoint(uint256 , uint256 ) view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) ExchangeRateCheckpoint(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "exchangeRateCheckpoint", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRateCheckpoint is a free data retrieval call binding the contract method 0x30c654a5.
//
// Solidity: function exchangeRateCheckpoint(uint256 , uint256 ) view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) ExchangeRateCheckpoint(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.ExchangeRateCheckpoint(&_CampaignContinuousCycles.CallOpts, arg0, arg1)
}

// ExchangeRateCheckpoint is a free data retrieval call binding the contract method 0x30c654a5.
//
// Solidity: function exchangeRateCheckpoint(uint256 , uint256 ) view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) ExchangeRateCheckpoint(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.ExchangeRateCheckpoint(&_CampaignContinuousCycles.CallOpts, arg0, arg1)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) Id(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "id")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) Id() (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.Id(&_CampaignContinuousCycles.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) Id() (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.Id(&_CampaignContinuousCycles.CallOpts)
}

// JoinWindow is a free data retrieval call binding the contract method 0x96de63e9.
//
// Solidity: function joinWindow() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) JoinWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "joinWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// JoinWindow is a free data retrieval call binding the contract method 0x96de63e9.
//
// Solidity: function joinWindow() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) JoinWindow() (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.JoinWindow(&_CampaignContinuousCycles.CallOpts)
}

// JoinWindow is a free data retrieval call binding the contract method 0x96de63e9.
//
// Solidity: function joinWindow() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) JoinWindow() (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.JoinWindow(&_CampaignContinuousCycles.CallOpts)
}

// Launched is a free data retrieval call binding the contract method 0x8091f3bf.
//
// Solidity: function launched() view returns(bool)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) Launched(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "launched")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Launched is a free data retrieval call binding the contract method 0x8091f3bf.
//
// Solidity: function launched() view returns(bool)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) Launched() (bool, error) {
	return _CampaignContinuousCycles.Contract.Launched(&_CampaignContinuousCycles.CallOpts)
}

// Launched is a free data retrieval call binding the contract method 0x8091f3bf.
//
// Solidity: function launched() view returns(bool)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) Launched() (bool, error) {
	return _CampaignContinuousCycles.Contract.Launched(&_CampaignContinuousCycles.CallOpts)
}

// NextCycleAt is a free data retrieval call binding the contract method 0xd725ccab.
//
// Solidity: function nextCycleAt() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) NextCycleAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "nextCycleAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextCycleAt is a free data retrieval call binding the contract method 0xd725ccab.
//
// Solidity: function nextCycleAt() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) NextCycleAt() (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.NextCycleAt(&_CampaignContinuousCycles.CallOpts)
}

// NextCycleAt is a free data retrieval call binding the contract method 0xd725ccab.
//
// Solidity: function nextCycleAt() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) NextCycleAt() (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.NextCycleAt(&_CampaignContinuousCycles.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) Operator() (common.Address, error) {
	return _CampaignContinuousCycles.Contract.Operator(&_CampaignContinuousCycles.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) Operator() (common.Address, error) {
	return _CampaignContinuousCycles.Contract.Operator(&_CampaignContinuousCycles.CallOpts)
}

// OperatorBalance is a free data retrieval call binding the contract method 0xfcdd8aac.
//
// Solidity: function operatorBalance() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) OperatorBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "operatorBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorBalance is a free data retrieval call binding the contract method 0xfcdd8aac.
//
// Solidity: function operatorBalance() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) OperatorBalance() (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.OperatorBalance(&_CampaignContinuousCycles.CallOpts)
}

// OperatorBalance is a free data retrieval call binding the contract method 0xfcdd8aac.
//
// Solidity: function operatorBalance() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) OperatorBalance() (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.OperatorBalance(&_CampaignContinuousCycles.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) Owner() (common.Address, error) {
	return _CampaignContinuousCycles.Contract.Owner(&_CampaignContinuousCycles.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) Owner() (common.Address, error) {
	return _CampaignContinuousCycles.Contract.Owner(&_CampaignContinuousCycles.CallOpts)
}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) Target(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "target")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) Target() (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.Target(&_CampaignContinuousCycles.CallOpts)
}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) Target() (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.Target(&_CampaignContinuousCycles.CallOpts)
}

// Terminated is a free data retrieval call binding the contract method 0x194307bf.
//
// Solidity: function terminated() view returns(bool)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) Terminated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "terminated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Terminated is a free data retrieval call binding the contract method 0x194307bf.
//
// Solidity: function terminated() view returns(bool)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) Terminated() (bool, error) {
	return _CampaignContinuousCycles.Contract.Terminated(&_CampaignContinuousCycles.CallOpts)
}

// Terminated is a free data retrieval call binding the contract method 0x194307bf.
//
// Solidity: function terminated() view returns(bool)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) Terminated() (bool, error) {
	return _CampaignContinuousCycles.Contract.Terminated(&_CampaignContinuousCycles.CallOpts)
}

// TrancheExitBuf is a free data retrieval call binding the contract method 0x6d2916de.
//
// Solidity: function trancheExitBuf(uint256 ) view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) TrancheExitBuf(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "trancheExitBuf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TrancheExitBuf is a free data retrieval call binding the contract method 0x6d2916de.
//
// Solidity: function trancheExitBuf(uint256 ) view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) TrancheExitBuf(arg0 *big.Int) (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.TrancheExitBuf(&_CampaignContinuousCycles.CallOpts, arg0)
}

// TrancheExitBuf is a free data retrieval call binding the contract method 0x6d2916de.
//
// Solidity: function trancheExitBuf(uint256 ) view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) TrancheExitBuf(arg0 *big.Int) (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.TrancheExitBuf(&_CampaignContinuousCycles.CallOpts, arg0)
}

// TrancheFeeInf is a free data retrieval call binding the contract method 0x6ae1198e.
//
// Solidity: function trancheFeeInf(uint256 ) view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) TrancheFeeInf(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "trancheFeeInf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TrancheFeeInf is a free data retrieval call binding the contract method 0x6ae1198e.
//
// Solidity: function trancheFeeInf(uint256 ) view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) TrancheFeeInf(arg0 *big.Int) (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.TrancheFeeInf(&_CampaignContinuousCycles.CallOpts, arg0)
}

// TrancheFeeInf is a free data retrieval call binding the contract method 0x6ae1198e.
//
// Solidity: function trancheFeeInf(uint256 ) view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) TrancheFeeInf(arg0 *big.Int) (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.TrancheFeeInf(&_CampaignContinuousCycles.CallOpts, arg0)
}

// TrancheJoinBuf is a free data retrieval call binding the contract method 0xd503f715.
//
// Solidity: function trancheJoinBuf(uint256 ) view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) TrancheJoinBuf(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "trancheJoinBuf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TrancheJoinBuf is a free data retrieval call binding the contract method 0xd503f715.
//
// Solidity: function trancheJoinBuf(uint256 ) view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) TrancheJoinBuf(arg0 *big.Int) (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.TrancheJoinBuf(&_CampaignContinuousCycles.CallOpts, arg0)
}

// TrancheJoinBuf is a free data retrieval call binding the contract method 0xd503f715.
//
// Solidity: function trancheJoinBuf(uint256 ) view returns(uint256)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) TrancheJoinBuf(arg0 *big.Int) (*big.Int, error) {
	return _CampaignContinuousCycles.Contract.TrancheJoinBuf(&_CampaignContinuousCycles.CallOpts, arg0)
}

// Tranches is a free data retrieval call binding the contract method 0x26c25962.
//
// Solidity: function tranches(uint256 ) view returns(uint256 target, uint256 principal, uint256 capital, uint256 percentage, uint256 apy, address token)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) Tranches(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Target     *big.Int
	Principal  *big.Int
	Capital    *big.Int
	Percentage *big.Int
	Apy        *big.Int
	Token      common.Address
}, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "tranches", arg0)

	outstruct := new(struct {
		Target     *big.Int
		Principal  *big.Int
		Capital    *big.Int
		Percentage *big.Int
		Apy        *big.Int
		Token      common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Target = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Principal = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Capital = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Percentage = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Apy = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Token = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Tranches is a free data retrieval call binding the contract method 0x26c25962.
//
// Solidity: function tranches(uint256 ) view returns(uint256 target, uint256 principal, uint256 capital, uint256 percentage, uint256 apy, address token)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) Tranches(arg0 *big.Int) (struct {
	Target     *big.Int
	Principal  *big.Int
	Capital    *big.Int
	Percentage *big.Int
	Apy        *big.Int
	Token      common.Address
}, error) {
	return _CampaignContinuousCycles.Contract.Tranches(&_CampaignContinuousCycles.CallOpts, arg0)
}

// Tranches is a free data retrieval call binding the contract method 0x26c25962.
//
// Solidity: function tranches(uint256 ) view returns(uint256 target, uint256 principal, uint256 capital, uint256 percentage, uint256 apy, address token)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) Tranches(arg0 *big.Int) (struct {
	Target     *big.Int
	Principal  *big.Int
	Capital    *big.Int
	Percentage *big.Int
	Apy        *big.Int
	Token      common.Address
}, error) {
	return _CampaignContinuousCycles.Contract.Tranches(&_CampaignContinuousCycles.CallOpts, arg0)
}

// TranchesTotal is a free data retrieval call binding the contract method 0xbd7bc46d.
//
// Solidity: function tranchesTotal() view returns(uint256 principal, uint256 capital)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCaller) TranchesTotal(opts *bind.CallOpts) (struct {
	Principal *big.Int
	Capital   *big.Int
}, error) {
	var out []interface{}
	err := _CampaignContinuousCycles.contract.Call(opts, &out, "tranchesTotal")

	outstruct := new(struct {
		Principal *big.Int
		Capital   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Principal = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Capital = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TranchesTotal is a free data retrieval call binding the contract method 0xbd7bc46d.
//
// Solidity: function tranchesTotal() view returns(uint256 principal, uint256 capital)
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) TranchesTotal() (struct {
	Principal *big.Int
	Capital   *big.Int
}, error) {
	return _CampaignContinuousCycles.Contract.TranchesTotal(&_CampaignContinuousCycles.CallOpts)
}

// TranchesTotal is a free data retrieval call binding the contract method 0xbd7bc46d.
//
// Solidity: function tranchesTotal() view returns(uint256 principal, uint256 capital)
func (_CampaignContinuousCycles *CampaignContinuousCyclesCallerSession) TranchesTotal() (struct {
	Principal *big.Int
	Capital   *big.Int
}, error) {
	return _CampaignContinuousCycles.Contract.TranchesTotal(&_CampaignContinuousCycles.CallOpts)
}

// ApplyForExit is a paid mutator transaction binding the contract method 0xe9b3059b.
//
// Solidity: function applyForExit(uint256 trancheID) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactor) ApplyForExit(opts *bind.TransactOpts, trancheID *big.Int) (*types.Transaction, error) {
	return _CampaignContinuousCycles.contract.Transact(opts, "applyForExit", trancheID)
}

// ApplyForExit is a paid mutator transaction binding the contract method 0xe9b3059b.
//
// Solidity: function applyForExit(uint256 trancheID) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) ApplyForExit(trancheID *big.Int) (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.ApplyForExit(&_CampaignContinuousCycles.TransactOpts, trancheID)
}

// ApplyForExit is a paid mutator transaction binding the contract method 0xe9b3059b.
//
// Solidity: function applyForExit(uint256 trancheID) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactorSession) ApplyForExit(trancheID *big.Int) (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.ApplyForExit(&_CampaignContinuousCycles.TransactOpts, trancheID)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address o) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactor) ChangeOperator(opts *bind.TransactOpts, o common.Address) (*types.Transaction, error) {
	return _CampaignContinuousCycles.contract.Transact(opts, "changeOperator", o)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address o) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) ChangeOperator(o common.Address) (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.ChangeOperator(&_CampaignContinuousCycles.TransactOpts, o)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address o) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactorSession) ChangeOperator(o common.Address) (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.ChangeOperator(&_CampaignContinuousCycles.TransactOpts, o)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(uint256 trancheID) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactor) Exit(opts *bind.TransactOpts, trancheID *big.Int) (*types.Transaction, error) {
	return _CampaignContinuousCycles.contract.Transact(opts, "exit", trancheID)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(uint256 trancheID) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) Exit(trancheID *big.Int) (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.Exit(&_CampaignContinuousCycles.TransactOpts, trancheID)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(uint256 trancheID) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactorSession) Exit(trancheID *big.Int) (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.Exit(&_CampaignContinuousCycles.TransactOpts, trancheID)
}

// Join is a paid mutator transaction binding the contract method 0x79e66b46.
//
// Solidity: function join(uint256 trancheID, uint256 amount) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactor) Join(opts *bind.TransactOpts, trancheID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CampaignContinuousCycles.contract.Transact(opts, "join", trancheID, amount)
}

// Join is a paid mutator transaction binding the contract method 0x79e66b46.
//
// Solidity: function join(uint256 trancheID, uint256 amount) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) Join(trancheID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.Join(&_CampaignContinuousCycles.TransactOpts, trancheID, amount)
}

// Join is a paid mutator transaction binding the contract method 0x79e66b46.
//
// Solidity: function join(uint256 trancheID, uint256 amount) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactorSession) Join(trancheID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.Join(&_CampaignContinuousCycles.TransactOpts, trancheID, amount)
}

// NextCycle is a paid mutator transaction binding the contract method 0x6b6923b8.
//
// Solidity: function nextCycle() returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactor) NextCycle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CampaignContinuousCycles.contract.Transact(opts, "nextCycle")
}

// NextCycle is a paid mutator transaction binding the contract method 0x6b6923b8.
//
// Solidity: function nextCycle() returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) NextCycle() (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.NextCycle(&_CampaignContinuousCycles.TransactOpts)
}

// NextCycle is a paid mutator transaction binding the contract method 0x6b6923b8.
//
// Solidity: function nextCycle() returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactorSession) NextCycle() (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.NextCycle(&_CampaignContinuousCycles.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CampaignContinuousCycles.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) RenounceOwnership() (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.RenounceOwnership(&_CampaignContinuousCycles.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.RenounceOwnership(&_CampaignContinuousCycles.TransactOpts)
}

// SetAlpacaBUSD is a paid mutator transaction binding the contract method 0xd15ed84d.
//
// Solidity: function setAlpacaBUSD(address a) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactor) SetAlpacaBUSD(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _CampaignContinuousCycles.contract.Transact(opts, "setAlpacaBUSD", a)
}

// SetAlpacaBUSD is a paid mutator transaction binding the contract method 0xd15ed84d.
//
// Solidity: function setAlpacaBUSD(address a) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) SetAlpacaBUSD(a common.Address) (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.SetAlpacaBUSD(&_CampaignContinuousCycles.TransactOpts, a)
}

// SetAlpacaBUSD is a paid mutator transaction binding the contract method 0xd15ed84d.
//
// Solidity: function setAlpacaBUSD(address a) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactorSession) SetAlpacaBUSD(a common.Address) (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.SetAlpacaBUSD(&_CampaignContinuousCycles.TransactOpts, a)
}

// SetBUSD is a paid mutator transaction binding the contract method 0x9805f01b.
//
// Solidity: function setBUSD(address a) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactor) SetBUSD(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _CampaignContinuousCycles.contract.Transact(opts, "setBUSD", a)
}

// SetBUSD is a paid mutator transaction binding the contract method 0x9805f01b.
//
// Solidity: function setBUSD(address a) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) SetBUSD(a common.Address) (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.SetBUSD(&_CampaignContinuousCycles.TransactOpts, a)
}

// SetBUSD is a paid mutator transaction binding the contract method 0x9805f01b.
//
// Solidity: function setBUSD(address a) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactorSession) SetBUSD(a common.Address) (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.SetBUSD(&_CampaignContinuousCycles.TransactOpts, a)
}

// SetFarms is a paid mutator transaction binding the contract method 0xa557a9fe.
//
// Solidity: function setFarms(address cream, uint256 s1, address venus, uint256 s2, address alpaca, uint256 s3) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactor) SetFarms(opts *bind.TransactOpts, cream common.Address, s1 *big.Int, venus common.Address, s2 *big.Int, alpaca common.Address, s3 *big.Int) (*types.Transaction, error) {
	return _CampaignContinuousCycles.contract.Transact(opts, "setFarms", cream, s1, venus, s2, alpaca, s3)
}

// SetFarms is a paid mutator transaction binding the contract method 0xa557a9fe.
//
// Solidity: function setFarms(address cream, uint256 s1, address venus, uint256 s2, address alpaca, uint256 s3) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) SetFarms(cream common.Address, s1 *big.Int, venus common.Address, s2 *big.Int, alpaca common.Address, s3 *big.Int) (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.SetFarms(&_CampaignContinuousCycles.TransactOpts, cream, s1, venus, s2, alpaca, s3)
}

// SetFarms is a paid mutator transaction binding the contract method 0xa557a9fe.
//
// Solidity: function setFarms(address cream, uint256 s1, address venus, uint256 s2, address alpaca, uint256 s3) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactorSession) SetFarms(cream common.Address, s1 *big.Int, venus common.Address, s2 *big.Int, alpaca common.Address, s3 *big.Int) (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.SetFarms(&_CampaignContinuousCycles.TransactOpts, cream, s1, venus, s2, alpaca, s3)
}

// SetPancakeRouter is a paid mutator transaction binding the contract method 0xf8316c90.
//
// Solidity: function setPancakeRouter(address a) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactor) SetPancakeRouter(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _CampaignContinuousCycles.contract.Transact(opts, "setPancakeRouter", a)
}

// SetPancakeRouter is a paid mutator transaction binding the contract method 0xf8316c90.
//
// Solidity: function setPancakeRouter(address a) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) SetPancakeRouter(a common.Address) (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.SetPancakeRouter(&_CampaignContinuousCycles.TransactOpts, a)
}

// SetPancakeRouter is a paid mutator transaction binding the contract method 0xf8316c90.
//
// Solidity: function setPancakeRouter(address a) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactorSession) SetPancakeRouter(a common.Address) (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.SetPancakeRouter(&_CampaignContinuousCycles.TransactOpts, a)
}

// SetTranchesFee is a paid mutator transaction binding the contract method 0x1516ef81.
//
// Solidity: function setTranchesFee(uint256 trancheID, uint256 fee) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactor) SetTranchesFee(opts *bind.TransactOpts, trancheID *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _CampaignContinuousCycles.contract.Transact(opts, "setTranchesFee", trancheID, fee)
}

// SetTranchesFee is a paid mutator transaction binding the contract method 0x1516ef81.
//
// Solidity: function setTranchesFee(uint256 trancheID, uint256 fee) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) SetTranchesFee(trancheID *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.SetTranchesFee(&_CampaignContinuousCycles.TransactOpts, trancheID, fee)
}

// SetTranchesFee is a paid mutator transaction binding the contract method 0x1516ef81.
//
// Solidity: function setTranchesFee(uint256 trancheID, uint256 fee) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactorSession) SetTranchesFee(trancheID *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.SetTranchesFee(&_CampaignContinuousCycles.TransactOpts, trancheID, fee)
}

// Terminate is a paid mutator transaction binding the contract method 0x0c08bf88.
//
// Solidity: function terminate() returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactor) Terminate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CampaignContinuousCycles.contract.Transact(opts, "terminate")
}

// Terminate is a paid mutator transaction binding the contract method 0x0c08bf88.
//
// Solidity: function terminate() returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) Terminate() (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.Terminate(&_CampaignContinuousCycles.TransactOpts)
}

// Terminate is a paid mutator transaction binding the contract method 0x0c08bf88.
//
// Solidity: function terminate() returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactorSession) Terminate() (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.Terminate(&_CampaignContinuousCycles.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CampaignContinuousCycles.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.TransferOwnership(&_CampaignContinuousCycles.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.TransferOwnership(&_CampaignContinuousCycles.TransactOpts, newOwner)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xbe357616.
//
// Solidity: function withdrawFee(uint256 amount) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactor) WithdrawFee(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CampaignContinuousCycles.contract.Transact(opts, "withdrawFee", amount)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xbe357616.
//
// Solidity: function withdrawFee(uint256 amount) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesSession) WithdrawFee(amount *big.Int) (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.WithdrawFee(&_CampaignContinuousCycles.TransactOpts, amount)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xbe357616.
//
// Solidity: function withdrawFee(uint256 amount) returns()
func (_CampaignContinuousCycles *CampaignContinuousCyclesTransactorSession) WithdrawFee(amount *big.Int) (*types.Transaction, error) {
	return _CampaignContinuousCycles.Contract.WithdrawFee(&_CampaignContinuousCycles.TransactOpts, amount)
}

// CampaignContinuousCyclesApplyForExitIterator is returned from FilterApplyForExit and is used to iterate over the raw logs and unpacked data for ApplyForExit events raised by the CampaignContinuousCycles contract.
type CampaignContinuousCyclesApplyForExitIterator struct {
	Event *CampaignContinuousCyclesApplyForExit // Event containing the contract specifics and raw log

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
func (it *CampaignContinuousCyclesApplyForExitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CampaignContinuousCyclesApplyForExit)
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
		it.Event = new(CampaignContinuousCyclesApplyForExit)
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
func (it *CampaignContinuousCyclesApplyForExitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CampaignContinuousCyclesApplyForExitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CampaignContinuousCyclesApplyForExit represents a ApplyForExit event raised by the CampaignContinuousCycles contract.
type CampaignContinuousCyclesApplyForExit struct {
	User          common.Address
	TrancheID     *big.Int
	Cycle         *big.Int
	OptionBalance *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterApplyForExit is a free log retrieval operation binding the contract event 0xfe4cedd15ba5c2324e927b3143ca076f1354f4c43c4bd90f80017f30816e82dd.
//
// Solidity: event ApplyForExit(address indexed user, uint256 indexed trancheID, uint256 cycle, uint256 optionBalance)
func (_CampaignContinuousCycles *CampaignContinuousCyclesFilterer) FilterApplyForExit(opts *bind.FilterOpts, user []common.Address, trancheID []*big.Int) (*CampaignContinuousCyclesApplyForExitIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var trancheIDRule []interface{}
	for _, trancheIDItem := range trancheID {
		trancheIDRule = append(trancheIDRule, trancheIDItem)
	}

	logs, sub, err := _CampaignContinuousCycles.contract.FilterLogs(opts, "ApplyForExit", userRule, trancheIDRule)
	if err != nil {
		return nil, err
	}
	return &CampaignContinuousCyclesApplyForExitIterator{contract: _CampaignContinuousCycles.contract, event: "ApplyForExit", logs: logs, sub: sub}, nil
}

// WatchApplyForExit is a free log subscription operation binding the contract event 0xfe4cedd15ba5c2324e927b3143ca076f1354f4c43c4bd90f80017f30816e82dd.
//
// Solidity: event ApplyForExit(address indexed user, uint256 indexed trancheID, uint256 cycle, uint256 optionBalance)
func (_CampaignContinuousCycles *CampaignContinuousCyclesFilterer) WatchApplyForExit(opts *bind.WatchOpts, sink chan<- *CampaignContinuousCyclesApplyForExit, user []common.Address, trancheID []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var trancheIDRule []interface{}
	for _, trancheIDItem := range trancheID {
		trancheIDRule = append(trancheIDRule, trancheIDItem)
	}

	logs, sub, err := _CampaignContinuousCycles.contract.WatchLogs(opts, "ApplyForExit", userRule, trancheIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CampaignContinuousCyclesApplyForExit)
				if err := _CampaignContinuousCycles.contract.UnpackLog(event, "ApplyForExit", log); err != nil {
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

// ParseApplyForExit is a log parse operation binding the contract event 0xfe4cedd15ba5c2324e927b3143ca076f1354f4c43c4bd90f80017f30816e82dd.
//
// Solidity: event ApplyForExit(address indexed user, uint256 indexed trancheID, uint256 cycle, uint256 optionBalance)
func (_CampaignContinuousCycles *CampaignContinuousCyclesFilterer) ParseApplyForExit(log types.Log) (*CampaignContinuousCyclesApplyForExit, error) {
	event := new(CampaignContinuousCyclesApplyForExit)
	if err := _CampaignContinuousCycles.contract.UnpackLog(event, "ApplyForExit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CampaignContinuousCyclesExitIterator is returned from FilterExit and is used to iterate over the raw logs and unpacked data for Exit events raised by the CampaignContinuousCycles contract.
type CampaignContinuousCyclesExitIterator struct {
	Event *CampaignContinuousCyclesExit // Event containing the contract specifics and raw log

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
func (it *CampaignContinuousCyclesExitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CampaignContinuousCyclesExit)
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
		it.Event = new(CampaignContinuousCyclesExit)
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
func (it *CampaignContinuousCyclesExitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CampaignContinuousCyclesExitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CampaignContinuousCyclesExit represents a Exit event raised by the CampaignContinuousCycles contract.
type CampaignContinuousCyclesExit struct {
	User      common.Address
	TrancheID *big.Int
	Cycle     *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExit is a free log retrieval operation binding the contract event 0x275029c7b988945c03ac5499c0d532fce79ce36efab42e1b3f180a62001cad2c.
//
// Solidity: event Exit(address indexed user, uint256 indexed trancheID, uint256 cycle, uint256 amount)
func (_CampaignContinuousCycles *CampaignContinuousCyclesFilterer) FilterExit(opts *bind.FilterOpts, user []common.Address, trancheID []*big.Int) (*CampaignContinuousCyclesExitIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var trancheIDRule []interface{}
	for _, trancheIDItem := range trancheID {
		trancheIDRule = append(trancheIDRule, trancheIDItem)
	}

	logs, sub, err := _CampaignContinuousCycles.contract.FilterLogs(opts, "Exit", userRule, trancheIDRule)
	if err != nil {
		return nil, err
	}
	return &CampaignContinuousCyclesExitIterator{contract: _CampaignContinuousCycles.contract, event: "Exit", logs: logs, sub: sub}, nil
}

// WatchExit is a free log subscription operation binding the contract event 0x275029c7b988945c03ac5499c0d532fce79ce36efab42e1b3f180a62001cad2c.
//
// Solidity: event Exit(address indexed user, uint256 indexed trancheID, uint256 cycle, uint256 amount)
func (_CampaignContinuousCycles *CampaignContinuousCyclesFilterer) WatchExit(opts *bind.WatchOpts, sink chan<- *CampaignContinuousCyclesExit, user []common.Address, trancheID []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var trancheIDRule []interface{}
	for _, trancheIDItem := range trancheID {
		trancheIDRule = append(trancheIDRule, trancheIDItem)
	}

	logs, sub, err := _CampaignContinuousCycles.contract.WatchLogs(opts, "Exit", userRule, trancheIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CampaignContinuousCyclesExit)
				if err := _CampaignContinuousCycles.contract.UnpackLog(event, "Exit", log); err != nil {
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

// ParseExit is a log parse operation binding the contract event 0x275029c7b988945c03ac5499c0d532fce79ce36efab42e1b3f180a62001cad2c.
//
// Solidity: event Exit(address indexed user, uint256 indexed trancheID, uint256 cycle, uint256 amount)
func (_CampaignContinuousCycles *CampaignContinuousCyclesFilterer) ParseExit(log types.Log) (*CampaignContinuousCyclesExit, error) {
	event := new(CampaignContinuousCyclesExit)
	if err := _CampaignContinuousCycles.contract.UnpackLog(event, "Exit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CampaignContinuousCyclesJoinIterator is returned from FilterJoin and is used to iterate over the raw logs and unpacked data for Join events raised by the CampaignContinuousCycles contract.
type CampaignContinuousCyclesJoinIterator struct {
	Event *CampaignContinuousCyclesJoin // Event containing the contract specifics and raw log

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
func (it *CampaignContinuousCyclesJoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CampaignContinuousCyclesJoin)
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
		it.Event = new(CampaignContinuousCyclesJoin)
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
func (it *CampaignContinuousCyclesJoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CampaignContinuousCyclesJoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CampaignContinuousCyclesJoin represents a Join event raised by the CampaignContinuousCycles contract.
type CampaignContinuousCyclesJoin struct {
	User          common.Address
	TrancheID     *big.Int
	Cycle         *big.Int
	Amount        *big.Int
	OptionBalance *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterJoin is a free log retrieval operation binding the contract event 0x6e45e6f0d59ac49faf62252114f3e63d57b36890e974b486f8b14ee741c07f7f.
//
// Solidity: event Join(address indexed user, uint256 indexed trancheID, uint256 cycle, uint256 amount, uint256 optionBalance)
func (_CampaignContinuousCycles *CampaignContinuousCyclesFilterer) FilterJoin(opts *bind.FilterOpts, user []common.Address, trancheID []*big.Int) (*CampaignContinuousCyclesJoinIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var trancheIDRule []interface{}
	for _, trancheIDItem := range trancheID {
		trancheIDRule = append(trancheIDRule, trancheIDItem)
	}

	logs, sub, err := _CampaignContinuousCycles.contract.FilterLogs(opts, "Join", userRule, trancheIDRule)
	if err != nil {
		return nil, err
	}
	return &CampaignContinuousCyclesJoinIterator{contract: _CampaignContinuousCycles.contract, event: "Join", logs: logs, sub: sub}, nil
}

// WatchJoin is a free log subscription operation binding the contract event 0x6e45e6f0d59ac49faf62252114f3e63d57b36890e974b486f8b14ee741c07f7f.
//
// Solidity: event Join(address indexed user, uint256 indexed trancheID, uint256 cycle, uint256 amount, uint256 optionBalance)
func (_CampaignContinuousCycles *CampaignContinuousCyclesFilterer) WatchJoin(opts *bind.WatchOpts, sink chan<- *CampaignContinuousCyclesJoin, user []common.Address, trancheID []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var trancheIDRule []interface{}
	for _, trancheIDItem := range trancheID {
		trancheIDRule = append(trancheIDRule, trancheIDItem)
	}

	logs, sub, err := _CampaignContinuousCycles.contract.WatchLogs(opts, "Join", userRule, trancheIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CampaignContinuousCyclesJoin)
				if err := _CampaignContinuousCycles.contract.UnpackLog(event, "Join", log); err != nil {
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

// ParseJoin is a log parse operation binding the contract event 0x6e45e6f0d59ac49faf62252114f3e63d57b36890e974b486f8b14ee741c07f7f.
//
// Solidity: event Join(address indexed user, uint256 indexed trancheID, uint256 cycle, uint256 amount, uint256 optionBalance)
func (_CampaignContinuousCycles *CampaignContinuousCyclesFilterer) ParseJoin(log types.Log) (*CampaignContinuousCyclesJoin, error) {
	event := new(CampaignContinuousCyclesJoin)
	if err := _CampaignContinuousCycles.contract.UnpackLog(event, "Join", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CampaignContinuousCyclesNextCycleIterator is returned from FilterNextCycle and is used to iterate over the raw logs and unpacked data for NextCycle events raised by the CampaignContinuousCycles contract.
type CampaignContinuousCyclesNextCycleIterator struct {
	Event *CampaignContinuousCyclesNextCycle // Event containing the contract specifics and raw log

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
func (it *CampaignContinuousCyclesNextCycleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CampaignContinuousCyclesNextCycle)
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
		it.Event = new(CampaignContinuousCyclesNextCycle)
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
func (it *CampaignContinuousCyclesNextCycleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CampaignContinuousCyclesNextCycleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CampaignContinuousCyclesNextCycle represents a NextCycle event raised by the CampaignContinuousCycles contract.
type CampaignContinuousCyclesNextCycle struct {
	Cycle *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNextCycle is a free log retrieval operation binding the contract event 0x2a39a1902551eacb05edd2a307cdbcce29c91601e0595e7fdacb21c9831d3e83.
//
// Solidity: event NextCycle(uint256 cycle)
func (_CampaignContinuousCycles *CampaignContinuousCyclesFilterer) FilterNextCycle(opts *bind.FilterOpts) (*CampaignContinuousCyclesNextCycleIterator, error) {

	logs, sub, err := _CampaignContinuousCycles.contract.FilterLogs(opts, "NextCycle")
	if err != nil {
		return nil, err
	}
	return &CampaignContinuousCyclesNextCycleIterator{contract: _CampaignContinuousCycles.contract, event: "NextCycle", logs: logs, sub: sub}, nil
}

// WatchNextCycle is a free log subscription operation binding the contract event 0x2a39a1902551eacb05edd2a307cdbcce29c91601e0595e7fdacb21c9831d3e83.
//
// Solidity: event NextCycle(uint256 cycle)
func (_CampaignContinuousCycles *CampaignContinuousCyclesFilterer) WatchNextCycle(opts *bind.WatchOpts, sink chan<- *CampaignContinuousCyclesNextCycle) (event.Subscription, error) {

	logs, sub, err := _CampaignContinuousCycles.contract.WatchLogs(opts, "NextCycle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CampaignContinuousCyclesNextCycle)
				if err := _CampaignContinuousCycles.contract.UnpackLog(event, "NextCycle", log); err != nil {
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

// ParseNextCycle is a log parse operation binding the contract event 0x2a39a1902551eacb05edd2a307cdbcce29c91601e0595e7fdacb21c9831d3e83.
//
// Solidity: event NextCycle(uint256 cycle)
func (_CampaignContinuousCycles *CampaignContinuousCyclesFilterer) ParseNextCycle(log types.Log) (*CampaignContinuousCyclesNextCycle, error) {
	event := new(CampaignContinuousCyclesNextCycle)
	if err := _CampaignContinuousCycles.contract.UnpackLog(event, "NextCycle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CampaignContinuousCyclesOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CampaignContinuousCycles contract.
type CampaignContinuousCyclesOwnershipTransferredIterator struct {
	Event *CampaignContinuousCyclesOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CampaignContinuousCyclesOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CampaignContinuousCyclesOwnershipTransferred)
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
		it.Event = new(CampaignContinuousCyclesOwnershipTransferred)
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
func (it *CampaignContinuousCyclesOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CampaignContinuousCyclesOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CampaignContinuousCyclesOwnershipTransferred represents a OwnershipTransferred event raised by the CampaignContinuousCycles contract.
type CampaignContinuousCyclesOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CampaignContinuousCycles *CampaignContinuousCyclesFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CampaignContinuousCyclesOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CampaignContinuousCycles.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CampaignContinuousCyclesOwnershipTransferredIterator{contract: _CampaignContinuousCycles.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CampaignContinuousCycles *CampaignContinuousCyclesFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CampaignContinuousCyclesOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CampaignContinuousCycles.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CampaignContinuousCyclesOwnershipTransferred)
				if err := _CampaignContinuousCycles.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CampaignContinuousCycles *CampaignContinuousCyclesFilterer) ParseOwnershipTransferred(log types.Log) (*CampaignContinuousCyclesOwnershipTransferred, error) {
	event := new(CampaignContinuousCyclesOwnershipTransferred)
	if err := _CampaignContinuousCycles.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CampaignContinuousCyclesTerminateIterator is returned from FilterTerminate and is used to iterate over the raw logs and unpacked data for Terminate events raised by the CampaignContinuousCycles contract.
type CampaignContinuousCyclesTerminateIterator struct {
	Event *CampaignContinuousCyclesTerminate // Event containing the contract specifics and raw log

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
func (it *CampaignContinuousCyclesTerminateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CampaignContinuousCyclesTerminate)
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
		it.Event = new(CampaignContinuousCyclesTerminate)
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
func (it *CampaignContinuousCyclesTerminateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CampaignContinuousCyclesTerminateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CampaignContinuousCyclesTerminate represents a Terminate event raised by the CampaignContinuousCycles contract.
type CampaignContinuousCyclesTerminate struct {
	Cycle *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTerminate is a free log retrieval operation binding the contract event 0xd681175168470800567b22d50d831df189686adc5b155827823a5ada6a97a4fe.
//
// Solidity: event Terminate(uint256 cycle)
func (_CampaignContinuousCycles *CampaignContinuousCyclesFilterer) FilterTerminate(opts *bind.FilterOpts) (*CampaignContinuousCyclesTerminateIterator, error) {

	logs, sub, err := _CampaignContinuousCycles.contract.FilterLogs(opts, "Terminate")
	if err != nil {
		return nil, err
	}
	return &CampaignContinuousCyclesTerminateIterator{contract: _CampaignContinuousCycles.contract, event: "Terminate", logs: logs, sub: sub}, nil
}

// WatchTerminate is a free log subscription operation binding the contract event 0xd681175168470800567b22d50d831df189686adc5b155827823a5ada6a97a4fe.
//
// Solidity: event Terminate(uint256 cycle)
func (_CampaignContinuousCycles *CampaignContinuousCyclesFilterer) WatchTerminate(opts *bind.WatchOpts, sink chan<- *CampaignContinuousCyclesTerminate) (event.Subscription, error) {

	logs, sub, err := _CampaignContinuousCycles.contract.WatchLogs(opts, "Terminate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CampaignContinuousCyclesTerminate)
				if err := _CampaignContinuousCycles.contract.UnpackLog(event, "Terminate", log); err != nil {
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

// ParseTerminate is a log parse operation binding the contract event 0xd681175168470800567b22d50d831df189686adc5b155827823a5ada6a97a4fe.
//
// Solidity: event Terminate(uint256 cycle)
func (_CampaignContinuousCycles *CampaignContinuousCyclesFilterer) ParseTerminate(log types.Log) (*CampaignContinuousCyclesTerminate, error) {
	event := new(CampaignContinuousCyclesTerminate)
	if err := _CampaignContinuousCycles.contract.UnpackLog(event, "Terminate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CampaignContinuousCyclesWithdrawFeeIterator is returned from FilterWithdrawFee and is used to iterate over the raw logs and unpacked data for WithdrawFee events raised by the CampaignContinuousCycles contract.
type CampaignContinuousCyclesWithdrawFeeIterator struct {
	Event *CampaignContinuousCyclesWithdrawFee // Event containing the contract specifics and raw log

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
func (it *CampaignContinuousCyclesWithdrawFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CampaignContinuousCyclesWithdrawFee)
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
		it.Event = new(CampaignContinuousCyclesWithdrawFee)
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
func (it *CampaignContinuousCyclesWithdrawFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CampaignContinuousCyclesWithdrawFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CampaignContinuousCyclesWithdrawFee represents a WithdrawFee event raised by the CampaignContinuousCycles contract.
type CampaignContinuousCyclesWithdrawFee struct {
	Operator common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawFee is a free log retrieval operation binding the contract event 0x66bf9186b00db666fc37aaffbb95a050c66e599e000c785c1dff0467d868f1b1.
//
// Solidity: event WithdrawFee(address operator, uint256 amount)
func (_CampaignContinuousCycles *CampaignContinuousCyclesFilterer) FilterWithdrawFee(opts *bind.FilterOpts) (*CampaignContinuousCyclesWithdrawFeeIterator, error) {

	logs, sub, err := _CampaignContinuousCycles.contract.FilterLogs(opts, "WithdrawFee")
	if err != nil {
		return nil, err
	}
	return &CampaignContinuousCyclesWithdrawFeeIterator{contract: _CampaignContinuousCycles.contract, event: "WithdrawFee", logs: logs, sub: sub}, nil
}

// WatchWithdrawFee is a free log subscription operation binding the contract event 0x66bf9186b00db666fc37aaffbb95a050c66e599e000c785c1dff0467d868f1b1.
//
// Solidity: event WithdrawFee(address operator, uint256 amount)
func (_CampaignContinuousCycles *CampaignContinuousCyclesFilterer) WatchWithdrawFee(opts *bind.WatchOpts, sink chan<- *CampaignContinuousCyclesWithdrawFee) (event.Subscription, error) {

	logs, sub, err := _CampaignContinuousCycles.contract.WatchLogs(opts, "WithdrawFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CampaignContinuousCyclesWithdrawFee)
				if err := _CampaignContinuousCycles.contract.UnpackLog(event, "WithdrawFee", log); err != nil {
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

// ParseWithdrawFee is a log parse operation binding the contract event 0x66bf9186b00db666fc37aaffbb95a050c66e599e000c785c1dff0467d868f1b1.
//
// Solidity: event WithdrawFee(address operator, uint256 amount)
func (_CampaignContinuousCycles *CampaignContinuousCyclesFilterer) ParseWithdrawFee(log types.Log) (*CampaignContinuousCyclesWithdrawFee, error) {
	event := new(CampaignContinuousCyclesWithdrawFee)
	if err := _CampaignContinuousCycles.contract.UnpackLog(event, "WithdrawFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
