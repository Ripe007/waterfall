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

// IAlpacaLikeABI is the input ABI used to generate the binding from.
const IAlpacaLikeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"amount\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"pendingInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IAlpacaLike is an auto generated Go binding around an Ethereum contract.
type IAlpacaLike struct {
	IAlpacaLikeCaller     // Read-only binding to the contract
	IAlpacaLikeTransactor // Write-only binding to the contract
	IAlpacaLikeFilterer   // Log filterer for contract events
}

// IAlpacaLikeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAlpacaLikeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAlpacaLikeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAlpacaLikeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAlpacaLikeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAlpacaLikeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAlpacaLikeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAlpacaLikeSession struct {
	Contract     *IAlpacaLike      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAlpacaLikeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAlpacaLikeCallerSession struct {
	Contract *IAlpacaLikeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IAlpacaLikeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAlpacaLikeTransactorSession struct {
	Contract     *IAlpacaLikeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IAlpacaLikeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAlpacaLikeRaw struct {
	Contract *IAlpacaLike // Generic contract binding to access the raw methods on
}

// IAlpacaLikeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAlpacaLikeCallerRaw struct {
	Contract *IAlpacaLikeCaller // Generic read-only contract binding to access the raw methods on
}

// IAlpacaLikeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAlpacaLikeTransactorRaw struct {
	Contract *IAlpacaLikeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAlpacaLike creates a new instance of IAlpacaLike, bound to a specific deployed contract.
func NewIAlpacaLike(address common.Address, backend bind.ContractBackend) (*IAlpacaLike, error) {
	contract, err := bindIAlpacaLike(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAlpacaLike{IAlpacaLikeCaller: IAlpacaLikeCaller{contract: contract}, IAlpacaLikeTransactor: IAlpacaLikeTransactor{contract: contract}, IAlpacaLikeFilterer: IAlpacaLikeFilterer{contract: contract}}, nil
}

// NewIAlpacaLikeCaller creates a new read-only instance of IAlpacaLike, bound to a specific deployed contract.
func NewIAlpacaLikeCaller(address common.Address, caller bind.ContractCaller) (*IAlpacaLikeCaller, error) {
	contract, err := bindIAlpacaLike(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAlpacaLikeCaller{contract: contract}, nil
}

// NewIAlpacaLikeTransactor creates a new write-only instance of IAlpacaLike, bound to a specific deployed contract.
func NewIAlpacaLikeTransactor(address common.Address, transactor bind.ContractTransactor) (*IAlpacaLikeTransactor, error) {
	contract, err := bindIAlpacaLike(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAlpacaLikeTransactor{contract: contract}, nil
}

// NewIAlpacaLikeFilterer creates a new log filterer instance of IAlpacaLike, bound to a specific deployed contract.
func NewIAlpacaLikeFilterer(address common.Address, filterer bind.ContractFilterer) (*IAlpacaLikeFilterer, error) {
	contract, err := bindIAlpacaLike(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAlpacaLikeFilterer{contract: contract}, nil
}

// bindIAlpacaLike binds a generic wrapper to an already deployed contract.
func bindIAlpacaLike(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAlpacaLikeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAlpacaLike *IAlpacaLikeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAlpacaLike.Contract.IAlpacaLikeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAlpacaLike *IAlpacaLikeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAlpacaLike.Contract.IAlpacaLikeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAlpacaLike *IAlpacaLikeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAlpacaLike.Contract.IAlpacaLikeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAlpacaLike *IAlpacaLikeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAlpacaLike.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAlpacaLike *IAlpacaLikeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAlpacaLike.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAlpacaLike *IAlpacaLikeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAlpacaLike.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address amount) returns()
func (_IAlpacaLike *IAlpacaLikeTransactor) BalanceOf(opts *bind.TransactOpts, amount common.Address) (*types.Transaction, error) {
	return _IAlpacaLike.contract.Transact(opts, "balanceOf", amount)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address amount) returns()
func (_IAlpacaLike *IAlpacaLikeSession) BalanceOf(amount common.Address) (*types.Transaction, error) {
	return _IAlpacaLike.Contract.BalanceOf(&_IAlpacaLike.TransactOpts, amount)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address amount) returns()
func (_IAlpacaLike *IAlpacaLikeTransactorSession) BalanceOf(amount common.Address) (*types.Transaction, error) {
	return _IAlpacaLike.Contract.BalanceOf(&_IAlpacaLike.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_IAlpacaLike *IAlpacaLikeTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IAlpacaLike.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_IAlpacaLike *IAlpacaLikeSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _IAlpacaLike.Contract.Deposit(&_IAlpacaLike.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_IAlpacaLike *IAlpacaLikeTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _IAlpacaLike.Contract.Deposit(&_IAlpacaLike.TransactOpts, amount)
}

// PendingInterest is a paid mutator transaction binding the contract method 0x2fc11c0f.
//
// Solidity: function pendingInterest(uint256 value) returns(uint256)
func (_IAlpacaLike *IAlpacaLikeTransactor) PendingInterest(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _IAlpacaLike.contract.Transact(opts, "pendingInterest", value)
}

// PendingInterest is a paid mutator transaction binding the contract method 0x2fc11c0f.
//
// Solidity: function pendingInterest(uint256 value) returns(uint256)
func (_IAlpacaLike *IAlpacaLikeSession) PendingInterest(value *big.Int) (*types.Transaction, error) {
	return _IAlpacaLike.Contract.PendingInterest(&_IAlpacaLike.TransactOpts, value)
}

// PendingInterest is a paid mutator transaction binding the contract method 0x2fc11c0f.
//
// Solidity: function pendingInterest(uint256 value) returns(uint256)
func (_IAlpacaLike *IAlpacaLikeTransactorSession) PendingInterest(value *big.Int) (*types.Transaction, error) {
	return _IAlpacaLike.Contract.PendingInterest(&_IAlpacaLike.TransactOpts, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 share) returns()
func (_IAlpacaLike *IAlpacaLikeTransactor) Withdraw(opts *bind.TransactOpts, share *big.Int) (*types.Transaction, error) {
	return _IAlpacaLike.contract.Transact(opts, "withdraw", share)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 share) returns()
func (_IAlpacaLike *IAlpacaLikeSession) Withdraw(share *big.Int) (*types.Transaction, error) {
	return _IAlpacaLike.Contract.Withdraw(&_IAlpacaLike.TransactOpts, share)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 share) returns()
func (_IAlpacaLike *IAlpacaLikeTransactorSession) Withdraw(share *big.Int) (*types.Transaction, error) {
	return _IAlpacaLike.Contract.Withdraw(&_IAlpacaLike.TransactOpts, share)
}
