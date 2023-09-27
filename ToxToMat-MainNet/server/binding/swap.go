// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package intoSwap

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// SwapMetaData contains all meta data concerning the Swap contract.
var SwapMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toxNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"matNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ctime\",\"type\":\"uint256\"}],\"name\":\"SwapRecord\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxToxPerTime_\",\"type\":\"uint256\"}],\"name\":\"adminSetMaxToxPerTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minToxPerTime_\",\"type\":\"uint256\"}],\"name\":\"adminSetMinToxPerTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate_\",\"type\":\"uint256\"}],\"name\":\"adminSetRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"swapTotalMax_\",\"type\":\"uint256\"}],\"name\":\"adminSetSwapTotalMax\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"toxAddress_\",\"type\":\"address\"}],\"name\":\"adminSetToxAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allAdmins\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"admins\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"amounts\",\"type\":\"address[]\"}],\"name\":\"batchAddAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMatBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxToxPerTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minToxPerTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toxNum_\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapTotalMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toxAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userSwapTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"withdrawMat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// SwapABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapMetaData.ABI instead.
var SwapABI = SwapMetaData.ABI

// Swap is an auto generated Go binding around an Ethereum contract.
type Swap struct {
	SwapCaller     // Read-only binding to the contract
	SwapTransactor // Write-only binding to the contract
	SwapFilterer   // Log filterer for contract events
}

// SwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapSession struct {
	Contract     *Swap             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapCallerSession struct {
	Contract *SwapCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapTransactorSession struct {
	Contract     *SwapTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapRaw struct {
	Contract *Swap // Generic contract binding to access the raw methods on
}

// SwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapCallerRaw struct {
	Contract *SwapCaller // Generic read-only contract binding to access the raw methods on
}

// SwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapTransactorRaw struct {
	Contract *SwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwap creates a new instance of Swap, bound to a specific deployed contract.
func NewSwap(address common.Address, backend bind.ContractBackend) (*Swap, error) {
	contract, err := bindSwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Swap{SwapCaller: SwapCaller{contract: contract}, SwapTransactor: SwapTransactor{contract: contract}, SwapFilterer: SwapFilterer{contract: contract}}, nil
}

// NewSwapCaller creates a new read-only instance of Swap, bound to a specific deployed contract.
func NewSwapCaller(address common.Address, caller bind.ContractCaller) (*SwapCaller, error) {
	contract, err := bindSwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapCaller{contract: contract}, nil
}

// NewSwapTransactor creates a new write-only instance of Swap, bound to a specific deployed contract.
func NewSwapTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapTransactor, error) {
	contract, err := bindSwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapTransactor{contract: contract}, nil
}

// NewSwapFilterer creates a new log filterer instance of Swap, bound to a specific deployed contract.
func NewSwapFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapFilterer, error) {
	contract, err := bindSwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapFilterer{contract: contract}, nil
}

// bindSwap binds a generic wrapper to an already deployed contract.
func bindSwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SwapMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swap *SwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swap.Contract.SwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swap *SwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.Contract.SwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swap *SwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swap.Contract.SwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swap *SwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swap *SwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swap *SwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swap.Contract.contract.Transact(opts, method, params...)
}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_Swap *SwapCaller) AllAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "allAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_Swap *SwapSession) AllAdmins() ([]common.Address, error) {
	return _Swap.Contract.AllAdmins(&_Swap.CallOpts)
}

// AllAdmins is a free data retrieval call binding the contract method 0x40f32be6.
//
// Solidity: function allAdmins() view returns(address[] admins)
func (_Swap *SwapCallerSession) AllAdmins() ([]common.Address, error) {
	return _Swap.Contract.AllAdmins(&_Swap.CallOpts)
}

// GetMatBalance is a free data retrieval call binding the contract method 0xbb7d2d42.
//
// Solidity: function getMatBalance() view returns(uint256)
func (_Swap *SwapCaller) GetMatBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "getMatBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMatBalance is a free data retrieval call binding the contract method 0xbb7d2d42.
//
// Solidity: function getMatBalance() view returns(uint256)
func (_Swap *SwapSession) GetMatBalance() (*big.Int, error) {
	return _Swap.Contract.GetMatBalance(&_Swap.CallOpts)
}

// GetMatBalance is a free data retrieval call binding the contract method 0xbb7d2d42.
//
// Solidity: function getMatBalance() view returns(uint256)
func (_Swap *SwapCallerSession) GetMatBalance() (*big.Int, error) {
	return _Swap.Contract.GetMatBalance(&_Swap.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_Swap *SwapCaller) IsAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "isAdmin", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_Swap *SwapSession) IsAdmin(account common.Address) (bool, error) {
	return _Swap.Contract.IsAdmin(&_Swap.CallOpts, account)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address account) view returns(bool)
func (_Swap *SwapCallerSession) IsAdmin(account common.Address) (bool, error) {
	return _Swap.Contract.IsAdmin(&_Swap.CallOpts, account)
}

// MaxToxPerTime is a free data retrieval call binding the contract method 0xc921bbe9.
//
// Solidity: function maxToxPerTime() view returns(uint256)
func (_Swap *SwapCaller) MaxToxPerTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "maxToxPerTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxToxPerTime is a free data retrieval call binding the contract method 0xc921bbe9.
//
// Solidity: function maxToxPerTime() view returns(uint256)
func (_Swap *SwapSession) MaxToxPerTime() (*big.Int, error) {
	return _Swap.Contract.MaxToxPerTime(&_Swap.CallOpts)
}

// MaxToxPerTime is a free data retrieval call binding the contract method 0xc921bbe9.
//
// Solidity: function maxToxPerTime() view returns(uint256)
func (_Swap *SwapCallerSession) MaxToxPerTime() (*big.Int, error) {
	return _Swap.Contract.MaxToxPerTime(&_Swap.CallOpts)
}

// MinToxPerTime is a free data retrieval call binding the contract method 0x48510822.
//
// Solidity: function minToxPerTime() view returns(uint256)
func (_Swap *SwapCaller) MinToxPerTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "minToxPerTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinToxPerTime is a free data retrieval call binding the contract method 0x48510822.
//
// Solidity: function minToxPerTime() view returns(uint256)
func (_Swap *SwapSession) MinToxPerTime() (*big.Int, error) {
	return _Swap.Contract.MinToxPerTime(&_Swap.CallOpts)
}

// MinToxPerTime is a free data retrieval call binding the contract method 0x48510822.
//
// Solidity: function minToxPerTime() view returns(uint256)
func (_Swap *SwapCallerSession) MinToxPerTime() (*big.Int, error) {
	return _Swap.Contract.MinToxPerTime(&_Swap.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Swap *SwapCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Swap *SwapSession) Rate() (*big.Int, error) {
	return _Swap.Contract.Rate(&_Swap.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Swap *SwapCallerSession) Rate() (*big.Int, error) {
	return _Swap.Contract.Rate(&_Swap.CallOpts)
}

// SwapTotalMax is a free data retrieval call binding the contract method 0xd0e43a18.
//
// Solidity: function swapTotalMax() view returns(uint256)
func (_Swap *SwapCaller) SwapTotalMax(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "swapTotalMax")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapTotalMax is a free data retrieval call binding the contract method 0xd0e43a18.
//
// Solidity: function swapTotalMax() view returns(uint256)
func (_Swap *SwapSession) SwapTotalMax() (*big.Int, error) {
	return _Swap.Contract.SwapTotalMax(&_Swap.CallOpts)
}

// SwapTotalMax is a free data retrieval call binding the contract method 0xd0e43a18.
//
// Solidity: function swapTotalMax() view returns(uint256)
func (_Swap *SwapCallerSession) SwapTotalMax() (*big.Int, error) {
	return _Swap.Contract.SwapTotalMax(&_Swap.CallOpts)
}

// ToxAddress is a free data retrieval call binding the contract method 0x9cf297e4.
//
// Solidity: function toxAddress() view returns(address)
func (_Swap *SwapCaller) ToxAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "toxAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ToxAddress is a free data retrieval call binding the contract method 0x9cf297e4.
//
// Solidity: function toxAddress() view returns(address)
func (_Swap *SwapSession) ToxAddress() (common.Address, error) {
	return _Swap.Contract.ToxAddress(&_Swap.CallOpts)
}

// ToxAddress is a free data retrieval call binding the contract method 0x9cf297e4.
//
// Solidity: function toxAddress() view returns(address)
func (_Swap *SwapCallerSession) ToxAddress() (common.Address, error) {
	return _Swap.Contract.ToxAddress(&_Swap.CallOpts)
}

// UserSwapTotal is a free data retrieval call binding the contract method 0xef703c8d.
//
// Solidity: function userSwapTotal(address ) view returns(uint256)
func (_Swap *SwapCaller) UserSwapTotal(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "userSwapTotal", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserSwapTotal is a free data retrieval call binding the contract method 0xef703c8d.
//
// Solidity: function userSwapTotal(address ) view returns(uint256)
func (_Swap *SwapSession) UserSwapTotal(arg0 common.Address) (*big.Int, error) {
	return _Swap.Contract.UserSwapTotal(&_Swap.CallOpts, arg0)
}

// UserSwapTotal is a free data retrieval call binding the contract method 0xef703c8d.
//
// Solidity: function userSwapTotal(address ) view returns(uint256)
func (_Swap *SwapCallerSession) UserSwapTotal(arg0 common.Address) (*big.Int, error) {
	return _Swap.Contract.UserSwapTotal(&_Swap.CallOpts, arg0)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_Swap *SwapTransactor) AddAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "addAdmin", account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_Swap *SwapSession) AddAdmin(account common.Address) (*types.Transaction, error) {
	return _Swap.Contract.AddAdmin(&_Swap.TransactOpts, account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address account) returns()
func (_Swap *SwapTransactorSession) AddAdmin(account common.Address) (*types.Transaction, error) {
	return _Swap.Contract.AddAdmin(&_Swap.TransactOpts, account)
}

// AdminSetMaxToxPerTime is a paid mutator transaction binding the contract method 0xf546b8d0.
//
// Solidity: function adminSetMaxToxPerTime(uint256 maxToxPerTime_) returns()
func (_Swap *SwapTransactor) AdminSetMaxToxPerTime(opts *bind.TransactOpts, maxToxPerTime_ *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "adminSetMaxToxPerTime", maxToxPerTime_)
}

// AdminSetMaxToxPerTime is a paid mutator transaction binding the contract method 0xf546b8d0.
//
// Solidity: function adminSetMaxToxPerTime(uint256 maxToxPerTime_) returns()
func (_Swap *SwapSession) AdminSetMaxToxPerTime(maxToxPerTime_ *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.AdminSetMaxToxPerTime(&_Swap.TransactOpts, maxToxPerTime_)
}

// AdminSetMaxToxPerTime is a paid mutator transaction binding the contract method 0xf546b8d0.
//
// Solidity: function adminSetMaxToxPerTime(uint256 maxToxPerTime_) returns()
func (_Swap *SwapTransactorSession) AdminSetMaxToxPerTime(maxToxPerTime_ *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.AdminSetMaxToxPerTime(&_Swap.TransactOpts, maxToxPerTime_)
}

// AdminSetMinToxPerTime is a paid mutator transaction binding the contract method 0x96e99381.
//
// Solidity: function adminSetMinToxPerTime(uint256 minToxPerTime_) returns()
func (_Swap *SwapTransactor) AdminSetMinToxPerTime(opts *bind.TransactOpts, minToxPerTime_ *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "adminSetMinToxPerTime", minToxPerTime_)
}

// AdminSetMinToxPerTime is a paid mutator transaction binding the contract method 0x96e99381.
//
// Solidity: function adminSetMinToxPerTime(uint256 minToxPerTime_) returns()
func (_Swap *SwapSession) AdminSetMinToxPerTime(minToxPerTime_ *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.AdminSetMinToxPerTime(&_Swap.TransactOpts, minToxPerTime_)
}

// AdminSetMinToxPerTime is a paid mutator transaction binding the contract method 0x96e99381.
//
// Solidity: function adminSetMinToxPerTime(uint256 minToxPerTime_) returns()
func (_Swap *SwapTransactorSession) AdminSetMinToxPerTime(minToxPerTime_ *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.AdminSetMinToxPerTime(&_Swap.TransactOpts, minToxPerTime_)
}

// AdminSetRate is a paid mutator transaction binding the contract method 0x56e92268.
//
// Solidity: function adminSetRate(uint256 rate_) returns()
func (_Swap *SwapTransactor) AdminSetRate(opts *bind.TransactOpts, rate_ *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "adminSetRate", rate_)
}

// AdminSetRate is a paid mutator transaction binding the contract method 0x56e92268.
//
// Solidity: function adminSetRate(uint256 rate_) returns()
func (_Swap *SwapSession) AdminSetRate(rate_ *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.AdminSetRate(&_Swap.TransactOpts, rate_)
}

// AdminSetRate is a paid mutator transaction binding the contract method 0x56e92268.
//
// Solidity: function adminSetRate(uint256 rate_) returns()
func (_Swap *SwapTransactorSession) AdminSetRate(rate_ *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.AdminSetRate(&_Swap.TransactOpts, rate_)
}

// AdminSetSwapTotalMax is a paid mutator transaction binding the contract method 0xacc2781b.
//
// Solidity: function adminSetSwapTotalMax(uint256 swapTotalMax_) returns()
func (_Swap *SwapTransactor) AdminSetSwapTotalMax(opts *bind.TransactOpts, swapTotalMax_ *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "adminSetSwapTotalMax", swapTotalMax_)
}

// AdminSetSwapTotalMax is a paid mutator transaction binding the contract method 0xacc2781b.
//
// Solidity: function adminSetSwapTotalMax(uint256 swapTotalMax_) returns()
func (_Swap *SwapSession) AdminSetSwapTotalMax(swapTotalMax_ *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.AdminSetSwapTotalMax(&_Swap.TransactOpts, swapTotalMax_)
}

// AdminSetSwapTotalMax is a paid mutator transaction binding the contract method 0xacc2781b.
//
// Solidity: function adminSetSwapTotalMax(uint256 swapTotalMax_) returns()
func (_Swap *SwapTransactorSession) AdminSetSwapTotalMax(swapTotalMax_ *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.AdminSetSwapTotalMax(&_Swap.TransactOpts, swapTotalMax_)
}

// AdminSetToxAddress is a paid mutator transaction binding the contract method 0xbb8641a0.
//
// Solidity: function adminSetToxAddress(address toxAddress_) returns()
func (_Swap *SwapTransactor) AdminSetToxAddress(opts *bind.TransactOpts, toxAddress_ common.Address) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "adminSetToxAddress", toxAddress_)
}

// AdminSetToxAddress is a paid mutator transaction binding the contract method 0xbb8641a0.
//
// Solidity: function adminSetToxAddress(address toxAddress_) returns()
func (_Swap *SwapSession) AdminSetToxAddress(toxAddress_ common.Address) (*types.Transaction, error) {
	return _Swap.Contract.AdminSetToxAddress(&_Swap.TransactOpts, toxAddress_)
}

// AdminSetToxAddress is a paid mutator transaction binding the contract method 0xbb8641a0.
//
// Solidity: function adminSetToxAddress(address toxAddress_) returns()
func (_Swap *SwapTransactorSession) AdminSetToxAddress(toxAddress_ common.Address) (*types.Transaction, error) {
	return _Swap.Contract.AdminSetToxAddress(&_Swap.TransactOpts, toxAddress_)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_Swap *SwapTransactor) BatchAddAdmin(opts *bind.TransactOpts, amounts []common.Address) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "batchAddAdmin", amounts)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_Swap *SwapSession) BatchAddAdmin(amounts []common.Address) (*types.Transaction, error) {
	return _Swap.Contract.BatchAddAdmin(&_Swap.TransactOpts, amounts)
}

// BatchAddAdmin is a paid mutator transaction binding the contract method 0x2c9ab42b.
//
// Solidity: function batchAddAdmin(address[] amounts) returns()
func (_Swap *SwapTransactorSession) BatchAddAdmin(amounts []common.Address) (*types.Transaction, error) {
	return _Swap.Contract.BatchAddAdmin(&_Swap.TransactOpts, amounts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Swap *SwapTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Swap *SwapSession) Initialize() (*types.Transaction, error) {
	return _Swap.Contract.Initialize(&_Swap.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Swap *SwapTransactorSession) Initialize() (*types.Transaction, error) {
	return _Swap.Contract.Initialize(&_Swap.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_Swap *SwapTransactor) RemoveAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "removeAdmin", account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_Swap *SwapSession) RemoveAdmin(account common.Address) (*types.Transaction, error) {
	return _Swap.Contract.RemoveAdmin(&_Swap.TransactOpts, account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address account) returns()
func (_Swap *SwapTransactorSession) RemoveAdmin(account common.Address) (*types.Transaction, error) {
	return _Swap.Contract.RemoveAdmin(&_Swap.TransactOpts, account)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_Swap *SwapTransactor) RenounceAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "renounceAdmin")
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_Swap *SwapSession) RenounceAdmin() (*types.Transaction, error) {
	return _Swap.Contract.RenounceAdmin(&_Swap.TransactOpts)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_Swap *SwapTransactorSession) RenounceAdmin() (*types.Transaction, error) {
	return _Swap.Contract.RenounceAdmin(&_Swap.TransactOpts)
}

// Swap is a paid mutator transaction binding the contract method 0x94b918de.
//
// Solidity: function swap(uint256 toxNum_) returns()
func (_Swap *SwapTransactor) Swap(opts *bind.TransactOpts, toxNum_ *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "swap", toxNum_)
}

// Swap is a paid mutator transaction binding the contract method 0x94b918de.
//
// Solidity: function swap(uint256 toxNum_) returns()
func (_Swap *SwapSession) Swap(toxNum_ *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.Swap(&_Swap.TransactOpts, toxNum_)
}

// Swap is a paid mutator transaction binding the contract method 0x94b918de.
//
// Solidity: function swap(uint256 toxNum_) returns()
func (_Swap *SwapTransactorSession) Swap(toxNum_ *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.Swap(&_Swap.TransactOpts, toxNum_)
}

// WithdrawMat is a paid mutator transaction binding the contract method 0xbd7f4c8d.
//
// Solidity: function withdrawMat(address user_, uint256 amount_) returns()
func (_Swap *SwapTransactor) WithdrawMat(opts *bind.TransactOpts, user_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "withdrawMat", user_, amount_)
}

// WithdrawMat is a paid mutator transaction binding the contract method 0xbd7f4c8d.
//
// Solidity: function withdrawMat(address user_, uint256 amount_) returns()
func (_Swap *SwapSession) WithdrawMat(user_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.WithdrawMat(&_Swap.TransactOpts, user_, amount_)
}

// WithdrawMat is a paid mutator transaction binding the contract method 0xbd7f4c8d.
//
// Solidity: function withdrawMat(address user_, uint256 amount_) returns()
func (_Swap *SwapTransactorSession) WithdrawMat(user_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.WithdrawMat(&_Swap.TransactOpts, user_, amount_)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Swap *SwapTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Swap *SwapSession) Receive() (*types.Transaction, error) {
	return _Swap.Contract.Receive(&_Swap.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Swap *SwapTransactorSession) Receive() (*types.Transaction, error) {
	return _Swap.Contract.Receive(&_Swap.TransactOpts)
}

// SwapAdminAddedIterator is returned from FilterAdminAdded and is used to iterate over the raw logs and unpacked data for AdminAdded events raised by the Swap contract.
type SwapAdminAddedIterator struct {
	Event *SwapAdminAdded // Event containing the contract specifics and raw log

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
func (it *SwapAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapAdminAdded)
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
		it.Event = new(SwapAdminAdded)
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
func (it *SwapAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapAdminAdded represents a AdminAdded event raised by the Swap contract.
type SwapAdminAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminAdded is a free log retrieval operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed account)
func (_Swap *SwapFilterer) FilterAdminAdded(opts *bind.FilterOpts, account []common.Address) (*SwapAdminAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Swap.contract.FilterLogs(opts, "AdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &SwapAdminAddedIterator{contract: _Swap.contract, event: "AdminAdded", logs: logs, sub: sub}, nil
}

// WatchAdminAdded is a free log subscription operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed account)
func (_Swap *SwapFilterer) WatchAdminAdded(opts *bind.WatchOpts, sink chan<- *SwapAdminAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Swap.contract.WatchLogs(opts, "AdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapAdminAdded)
				if err := _Swap.contract.UnpackLog(event, "AdminAdded", log); err != nil {
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

// ParseAdminAdded is a log parse operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed account)
func (_Swap *SwapFilterer) ParseAdminAdded(log types.Log) (*SwapAdminAdded, error) {
	event := new(SwapAdminAdded)
	if err := _Swap.contract.UnpackLog(event, "AdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the Swap contract.
type SwapAdminRemovedIterator struct {
	Event *SwapAdminRemoved // Event containing the contract specifics and raw log

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
func (it *SwapAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapAdminRemoved)
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
		it.Event = new(SwapAdminRemoved)
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
func (it *SwapAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapAdminRemoved represents a AdminRemoved event raised by the Swap contract.
type SwapAdminRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed account)
func (_Swap *SwapFilterer) FilterAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*SwapAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Swap.contract.FilterLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &SwapAdminRemovedIterator{contract: _Swap.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed account)
func (_Swap *SwapFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *SwapAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Swap.contract.WatchLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapAdminRemoved)
				if err := _Swap.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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

// ParseAdminRemoved is a log parse operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed account)
func (_Swap *SwapFilterer) ParseAdminRemoved(log types.Log) (*SwapAdminRemoved, error) {
	event := new(SwapAdminRemoved)
	if err := _Swap.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Swap contract.
type SwapInitializedIterator struct {
	Event *SwapInitialized // Event containing the contract specifics and raw log

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
func (it *SwapInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapInitialized)
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
		it.Event = new(SwapInitialized)
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
func (it *SwapInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapInitialized represents a Initialized event raised by the Swap contract.
type SwapInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Swap *SwapFilterer) FilterInitialized(opts *bind.FilterOpts) (*SwapInitializedIterator, error) {

	logs, sub, err := _Swap.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SwapInitializedIterator{contract: _Swap.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Swap *SwapFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SwapInitialized) (event.Subscription, error) {

	logs, sub, err := _Swap.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapInitialized)
				if err := _Swap.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Swap *SwapFilterer) ParseInitialized(log types.Log) (*SwapInitialized, error) {
	event := new(SwapInitialized)
	if err := _Swap.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapSwapRecordIterator is returned from FilterSwapRecord and is used to iterate over the raw logs and unpacked data for SwapRecord events raised by the Swap contract.
type SwapSwapRecordIterator struct {
	Event *SwapSwapRecord // Event containing the contract specifics and raw log

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
func (it *SwapSwapRecordIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapSwapRecord)
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
		it.Event = new(SwapSwapRecord)
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
func (it *SwapSwapRecordIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapSwapRecordIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapSwapRecord represents a SwapRecord event raised by the Swap contract.
type SwapSwapRecord struct {
	User   common.Address
	ToxNum *big.Int
	MatNum *big.Int
	Ctime  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSwapRecord is a free log retrieval operation binding the contract event 0xd9356407d335130ab3d2134530ed3183ba05b8858b771304002a5e55a1694154.
//
// Solidity: event SwapRecord(address user, uint256 toxNum, uint256 matNum, uint256 ctime)
func (_Swap *SwapFilterer) FilterSwapRecord(opts *bind.FilterOpts) (*SwapSwapRecordIterator, error) {

	logs, sub, err := _Swap.contract.FilterLogs(opts, "SwapRecord")
	if err != nil {
		return nil, err
	}
	return &SwapSwapRecordIterator{contract: _Swap.contract, event: "SwapRecord", logs: logs, sub: sub}, nil
}

// WatchSwapRecord is a free log subscription operation binding the contract event 0xd9356407d335130ab3d2134530ed3183ba05b8858b771304002a5e55a1694154.
//
// Solidity: event SwapRecord(address user, uint256 toxNum, uint256 matNum, uint256 ctime)
func (_Swap *SwapFilterer) WatchSwapRecord(opts *bind.WatchOpts, sink chan<- *SwapSwapRecord) (event.Subscription, error) {

	logs, sub, err := _Swap.contract.WatchLogs(opts, "SwapRecord")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapSwapRecord)
				if err := _Swap.contract.UnpackLog(event, "SwapRecord", log); err != nil {
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

// ParseSwapRecord is a log parse operation binding the contract event 0xd9356407d335130ab3d2134530ed3183ba05b8858b771304002a5e55a1694154.
//
// Solidity: event SwapRecord(address user, uint256 toxNum, uint256 matNum, uint256 ctime)
func (_Swap *SwapFilterer) ParseSwapRecord(log types.Log) (*SwapSwapRecord, error) {
	event := new(SwapSwapRecord)
	if err := _Swap.contract.UnpackLog(event, "SwapRecord", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
