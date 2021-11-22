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

// ICompoundLikeABI is the input ABI used to generate the binding from.
const ICompoundLikeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOfUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ICompoundLike is an auto generated Go binding around an Ethereum contract.
type ICompoundLike struct {
	ICompoundLikeCaller     // Read-only binding to the contract
	ICompoundLikeTransactor // Write-only binding to the contract
	ICompoundLikeFilterer   // Log filterer for contract events
}

// ICompoundLikeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICompoundLikeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICompoundLikeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICompoundLikeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICompoundLikeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICompoundLikeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICompoundLikeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICompoundLikeSession struct {
	Contract     *ICompoundLike    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICompoundLikeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICompoundLikeCallerSession struct {
	Contract *ICompoundLikeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ICompoundLikeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICompoundLikeTransactorSession struct {
	Contract     *ICompoundLikeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ICompoundLikeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICompoundLikeRaw struct {
	Contract *ICompoundLike // Generic contract binding to access the raw methods on
}

// ICompoundLikeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICompoundLikeCallerRaw struct {
	Contract *ICompoundLikeCaller // Generic read-only contract binding to access the raw methods on
}

// ICompoundLikeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICompoundLikeTransactorRaw struct {
	Contract *ICompoundLikeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICompoundLike creates a new instance of ICompoundLike, bound to a specific deployed contract.
func NewICompoundLike(address common.Address, backend bind.ContractBackend) (*ICompoundLike, error) {
	contract, err := bindICompoundLike(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICompoundLike{ICompoundLikeCaller: ICompoundLikeCaller{contract: contract}, ICompoundLikeTransactor: ICompoundLikeTransactor{contract: contract}, ICompoundLikeFilterer: ICompoundLikeFilterer{contract: contract}}, nil
}

// NewICompoundLikeCaller creates a new read-only instance of ICompoundLike, bound to a specific deployed contract.
func NewICompoundLikeCaller(address common.Address, caller bind.ContractCaller) (*ICompoundLikeCaller, error) {
	contract, err := bindICompoundLike(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICompoundLikeCaller{contract: contract}, nil
}

// NewICompoundLikeTransactor creates a new write-only instance of ICompoundLike, bound to a specific deployed contract.
func NewICompoundLikeTransactor(address common.Address, transactor bind.ContractTransactor) (*ICompoundLikeTransactor, error) {
	contract, err := bindICompoundLike(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICompoundLikeTransactor{contract: contract}, nil
}

// NewICompoundLikeFilterer creates a new log filterer instance of ICompoundLike, bound to a specific deployed contract.
func NewICompoundLikeFilterer(address common.Address, filterer bind.ContractFilterer) (*ICompoundLikeFilterer, error) {
	contract, err := bindICompoundLike(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICompoundLikeFilterer{contract: contract}, nil
}

// bindICompoundLike binds a generic wrapper to an already deployed contract.
func bindICompoundLike(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICompoundLikeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICompoundLike *ICompoundLikeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICompoundLike.Contract.ICompoundLikeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICompoundLike *ICompoundLikeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICompoundLike.Contract.ICompoundLikeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICompoundLike *ICompoundLikeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICompoundLike.Contract.ICompoundLikeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICompoundLike *ICompoundLikeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICompoundLike.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICompoundLike *ICompoundLikeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICompoundLike.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICompoundLike *ICompoundLikeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICompoundLike.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) returns(uint256)
func (_ICompoundLike *ICompoundLikeTransactor) BalanceOf(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ICompoundLike.contract.Transact(opts, "balanceOf", account)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) returns(uint256)
func (_ICompoundLike *ICompoundLikeSession) BalanceOf(account common.Address) (*types.Transaction, error) {
	return _ICompoundLike.Contract.BalanceOf(&_ICompoundLike.TransactOpts, account)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) returns(uint256)
func (_ICompoundLike *ICompoundLikeTransactorSession) BalanceOf(account common.Address) (*types.Transaction, error) {
	return _ICompoundLike.Contract.BalanceOf(&_ICompoundLike.TransactOpts, account)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address account) returns(uint256)
func (_ICompoundLike *ICompoundLikeTransactor) BalanceOfUnderlying(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ICompoundLike.contract.Transact(opts, "balanceOfUnderlying", account)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address account) returns(uint256)
func (_ICompoundLike *ICompoundLikeSession) BalanceOfUnderlying(account common.Address) (*types.Transaction, error) {
	return _ICompoundLike.Contract.BalanceOfUnderlying(&_ICompoundLike.TransactOpts, account)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address account) returns(uint256)
func (_ICompoundLike *ICompoundLikeTransactorSession) BalanceOfUnderlying(account common.Address) (*types.Transaction, error) {
	return _ICompoundLike.Contract.BalanceOfUnderlying(&_ICompoundLike.TransactOpts, account)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(uint256)
func (_ICompoundLike *ICompoundLikeTransactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ICompoundLike.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(uint256)
func (_ICompoundLike *ICompoundLikeSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _ICompoundLike.Contract.Mint(&_ICompoundLike.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(uint256)
func (_ICompoundLike *ICompoundLikeTransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _ICompoundLike.Contract.Mint(&_ICompoundLike.TransactOpts, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_ICompoundLike *ICompoundLikeTransactor) Redeem(opts *bind.TransactOpts, redeemTokens *big.Int) (*types.Transaction, error) {
	return _ICompoundLike.contract.Transact(opts, "redeem", redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_ICompoundLike *ICompoundLikeSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _ICompoundLike.Contract.Redeem(&_ICompoundLike.TransactOpts, redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_ICompoundLike *ICompoundLikeTransactorSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _ICompoundLike.Contract.Redeem(&_ICompoundLike.TransactOpts, redeemTokens)
}
