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

// IERC20MintBurnABI is the input ABI used to generate the binding from.
const IERC20MintBurnABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20MintBurn is an auto generated Go binding around an Ethereum contract.
type IERC20MintBurn struct {
	IERC20MintBurnCaller     // Read-only binding to the contract
	IERC20MintBurnTransactor // Write-only binding to the contract
	IERC20MintBurnFilterer   // Log filterer for contract events
}

// IERC20MintBurnCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20MintBurnCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MintBurnTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20MintBurnTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MintBurnFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20MintBurnFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MintBurnSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20MintBurnSession struct {
	Contract     *IERC20MintBurn   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20MintBurnCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20MintBurnCallerSession struct {
	Contract *IERC20MintBurnCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IERC20MintBurnTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20MintBurnTransactorSession struct {
	Contract     *IERC20MintBurnTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IERC20MintBurnRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20MintBurnRaw struct {
	Contract *IERC20MintBurn // Generic contract binding to access the raw methods on
}

// IERC20MintBurnCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20MintBurnCallerRaw struct {
	Contract *IERC20MintBurnCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20MintBurnTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20MintBurnTransactorRaw struct {
	Contract *IERC20MintBurnTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20MintBurn creates a new instance of IERC20MintBurn, bound to a specific deployed contract.
func NewIERC20MintBurn(address common.Address, backend bind.ContractBackend) (*IERC20MintBurn, error) {
	contract, err := bindIERC20MintBurn(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20MintBurn{IERC20MintBurnCaller: IERC20MintBurnCaller{contract: contract}, IERC20MintBurnTransactor: IERC20MintBurnTransactor{contract: contract}, IERC20MintBurnFilterer: IERC20MintBurnFilterer{contract: contract}}, nil
}

// NewIERC20MintBurnCaller creates a new read-only instance of IERC20MintBurn, bound to a specific deployed contract.
func NewIERC20MintBurnCaller(address common.Address, caller bind.ContractCaller) (*IERC20MintBurnCaller, error) {
	contract, err := bindIERC20MintBurn(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MintBurnCaller{contract: contract}, nil
}

// NewIERC20MintBurnTransactor creates a new write-only instance of IERC20MintBurn, bound to a specific deployed contract.
func NewIERC20MintBurnTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20MintBurnTransactor, error) {
	contract, err := bindIERC20MintBurn(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MintBurnTransactor{contract: contract}, nil
}

// NewIERC20MintBurnFilterer creates a new log filterer instance of IERC20MintBurn, bound to a specific deployed contract.
func NewIERC20MintBurnFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20MintBurnFilterer, error) {
	contract, err := bindIERC20MintBurn(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20MintBurnFilterer{contract: contract}, nil
}

// bindIERC20MintBurn binds a generic wrapper to an already deployed contract.
func bindIERC20MintBurn(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20MintBurnABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20MintBurn *IERC20MintBurnRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20MintBurn.Contract.IERC20MintBurnCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20MintBurn *IERC20MintBurnRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20MintBurn.Contract.IERC20MintBurnTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20MintBurn *IERC20MintBurnRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20MintBurn.Contract.IERC20MintBurnTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20MintBurn *IERC20MintBurnCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20MintBurn.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20MintBurn *IERC20MintBurnTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20MintBurn.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20MintBurn *IERC20MintBurnTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20MintBurn.Contract.contract.Transact(opts, method, params...)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address , uint256 ) returns()
func (_IERC20MintBurn *IERC20MintBurnTransactor) BurnFrom(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IERC20MintBurn.contract.Transact(opts, "burnFrom", arg0, arg1)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address , uint256 ) returns()
func (_IERC20MintBurn *IERC20MintBurnSession) BurnFrom(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IERC20MintBurn.Contract.BurnFrom(&_IERC20MintBurn.TransactOpts, arg0, arg1)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address , uint256 ) returns()
func (_IERC20MintBurn *IERC20MintBurnTransactorSession) BurnFrom(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IERC20MintBurn.Contract.BurnFrom(&_IERC20MintBurn.TransactOpts, arg0, arg1)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address , uint256 ) returns()
func (_IERC20MintBurn *IERC20MintBurnTransactor) Mint(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IERC20MintBurn.contract.Transact(opts, "mint", arg0, arg1)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address , uint256 ) returns()
func (_IERC20MintBurn *IERC20MintBurnSession) Mint(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IERC20MintBurn.Contract.Mint(&_IERC20MintBurn.TransactOpts, arg0, arg1)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address , uint256 ) returns()
func (_IERC20MintBurn *IERC20MintBurnTransactorSession) Mint(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IERC20MintBurn.Contract.Mint(&_IERC20MintBurn.TransactOpts, arg0, arg1)
}
