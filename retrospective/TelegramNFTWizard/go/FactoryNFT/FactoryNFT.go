// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package FactoryNFT

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
)

// FactoryNFTMetaData contains all meta data concerning the FactoryNFT contract.
var FactoryNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"}],\"name\":\"CollectionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Collections\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"smbl_\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"file_ids\",\"type\":\"string[]\"}],\"name\":\"CreateCollection\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasure_fund\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FactoryNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use FactoryNFTMetaData.ABI instead.
var FactoryNFTABI = FactoryNFTMetaData.ABI

// FactoryNFT is an auto generated Go binding around an Ethereum contract.
type FactoryNFT struct {
	FactoryNFTCaller     // Read-only binding to the contract
	FactoryNFTTransactor // Write-only binding to the contract
	FactoryNFTFilterer   // Log filterer for contract events
}

// FactoryNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type FactoryNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FactoryNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FactoryNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FactoryNFTSession struct {
	Contract     *FactoryNFT       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FactoryNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FactoryNFTCallerSession struct {
	Contract *FactoryNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FactoryNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FactoryNFTTransactorSession struct {
	Contract     *FactoryNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FactoryNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type FactoryNFTRaw struct {
	Contract *FactoryNFT // Generic contract binding to access the raw methods on
}

// FactoryNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FactoryNFTCallerRaw struct {
	Contract *FactoryNFTCaller // Generic read-only contract binding to access the raw methods on
}

// FactoryNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FactoryNFTTransactorRaw struct {
	Contract *FactoryNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFactoryNFT creates a new instance of FactoryNFT, bound to a specific deployed contract.
func NewFactoryNFT(address common.Address, backend bind.ContractBackend) (*FactoryNFT, error) {
	contract, err := bindFactoryNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FactoryNFT{FactoryNFTCaller: FactoryNFTCaller{contract: contract}, FactoryNFTTransactor: FactoryNFTTransactor{contract: contract}, FactoryNFTFilterer: FactoryNFTFilterer{contract: contract}}, nil
}

// NewFactoryNFTCaller creates a new read-only instance of FactoryNFT, bound to a specific deployed contract.
func NewFactoryNFTCaller(address common.Address, caller bind.ContractCaller) (*FactoryNFTCaller, error) {
	contract, err := bindFactoryNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryNFTCaller{contract: contract}, nil
}

// NewFactoryNFTTransactor creates a new write-only instance of FactoryNFT, bound to a specific deployed contract.
func NewFactoryNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*FactoryNFTTransactor, error) {
	contract, err := bindFactoryNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryNFTTransactor{contract: contract}, nil
}

// NewFactoryNFTFilterer creates a new log filterer instance of FactoryNFT, bound to a specific deployed contract.
func NewFactoryNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*FactoryNFTFilterer, error) {
	contract, err := bindFactoryNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FactoryNFTFilterer{contract: contract}, nil
}

// bindFactoryNFT binds a generic wrapper to an already deployed contract.
func bindFactoryNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FactoryNFTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FactoryNFT *FactoryNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FactoryNFT.Contract.FactoryNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FactoryNFT *FactoryNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FactoryNFT.Contract.FactoryNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FactoryNFT *FactoryNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FactoryNFT.Contract.FactoryNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FactoryNFT *FactoryNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FactoryNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FactoryNFT *FactoryNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FactoryNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FactoryNFT *FactoryNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FactoryNFT.Contract.contract.Transact(opts, method, params...)
}

// Collections is a free data retrieval call binding the contract method 0x144ee425.
//
// Solidity: function Collections(address ) view returns(address)
func (_FactoryNFT *FactoryNFTCaller) Collections(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _FactoryNFT.contract.Call(opts, &out, "Collections", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Collections is a free data retrieval call binding the contract method 0x144ee425.
//
// Solidity: function Collections(address ) view returns(address)
func (_FactoryNFT *FactoryNFTSession) Collections(arg0 common.Address) (common.Address, error) {
	return _FactoryNFT.Contract.Collections(&_FactoryNFT.CallOpts, arg0)
}

// Collections is a free data retrieval call binding the contract method 0x144ee425.
//
// Solidity: function Collections(address ) view returns(address)
func (_FactoryNFT *FactoryNFTCallerSession) Collections(arg0 common.Address) (common.Address, error) {
	return _FactoryNFT.Contract.Collections(&_FactoryNFT.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FactoryNFT *FactoryNFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FactoryNFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FactoryNFT *FactoryNFTSession) Owner() (common.Address, error) {
	return _FactoryNFT.Contract.Owner(&_FactoryNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FactoryNFT *FactoryNFTCallerSession) Owner() (common.Address, error) {
	return _FactoryNFT.Contract.Owner(&_FactoryNFT.CallOpts)
}

// TreasureFund is a free data retrieval call binding the contract method 0xf9a4bb1e.
//
// Solidity: function treasure_fund() view returns(address)
func (_FactoryNFT *FactoryNFTCaller) TreasureFund(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FactoryNFT.contract.Call(opts, &out, "treasure_fund")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasureFund is a free data retrieval call binding the contract method 0xf9a4bb1e.
//
// Solidity: function treasure_fund() view returns(address)
func (_FactoryNFT *FactoryNFTSession) TreasureFund() (common.Address, error) {
	return _FactoryNFT.Contract.TreasureFund(&_FactoryNFT.CallOpts)
}

// TreasureFund is a free data retrieval call binding the contract method 0xf9a4bb1e.
//
// Solidity: function treasure_fund() view returns(address)
func (_FactoryNFT *FactoryNFTCallerSession) TreasureFund() (common.Address, error) {
	return _FactoryNFT.Contract.TreasureFund(&_FactoryNFT.CallOpts)
}

// CreateCollection is a paid mutator transaction binding the contract method 0x0f1d60c9.
//
// Solidity: function CreateCollection(string name_, string smbl_, string[] file_ids) returns(address)
func (_FactoryNFT *FactoryNFTTransactor) CreateCollection(opts *bind.TransactOpts, name_ string, smbl_ string, file_ids []string) (*types.Transaction, error) {
	return _FactoryNFT.contract.Transact(opts, "CreateCollection", name_, smbl_, file_ids)
}

// CreateCollection is a paid mutator transaction binding the contract method 0x0f1d60c9.
//
// Solidity: function CreateCollection(string name_, string smbl_, string[] file_ids) returns(address)
func (_FactoryNFT *FactoryNFTSession) CreateCollection(name_ string, smbl_ string, file_ids []string) (*types.Transaction, error) {
	return _FactoryNFT.Contract.CreateCollection(&_FactoryNFT.TransactOpts, name_, smbl_, file_ids)
}

// CreateCollection is a paid mutator transaction binding the contract method 0x0f1d60c9.
//
// Solidity: function CreateCollection(string name_, string smbl_, string[] file_ids) returns(address)
func (_FactoryNFT *FactoryNFTTransactorSession) CreateCollection(name_ string, smbl_ string, file_ids []string) (*types.Transaction, error) {
	return _FactoryNFT.Contract.CreateCollection(&_FactoryNFT.TransactOpts, name_, smbl_, file_ids)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FactoryNFT *FactoryNFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FactoryNFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FactoryNFT *FactoryNFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _FactoryNFT.Contract.RenounceOwnership(&_FactoryNFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FactoryNFT *FactoryNFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FactoryNFT.Contract.RenounceOwnership(&_FactoryNFT.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FactoryNFT *FactoryNFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FactoryNFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FactoryNFT *FactoryNFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FactoryNFT.Contract.TransferOwnership(&_FactoryNFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FactoryNFT *FactoryNFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FactoryNFT.Contract.TransferOwnership(&_FactoryNFT.TransactOpts, newOwner)
}

// FactoryNFTCollectionCreatedIterator is returned from FilterCollectionCreated and is used to iterate over the raw logs and unpacked data for CollectionCreated events raised by the FactoryNFT contract.
type FactoryNFTCollectionCreatedIterator struct {
	Event *FactoryNFTCollectionCreated // Event containing the contract specifics and raw log

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
func (it *FactoryNFTCollectionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryNFTCollectionCreated)
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
		it.Event = new(FactoryNFTCollectionCreated)
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
func (it *FactoryNFTCollectionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryNFTCollectionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryNFTCollectionCreated represents a CollectionCreated event raised by the FactoryNFT contract.
type FactoryNFTCollectionCreated struct {
	Creator    common.Address
	Collection common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollectionCreated is a free log retrieval operation binding the contract event 0x5d0de243db1669e3a7056744cd715c625f0c1c348736c2c2d53d0ddebff1a6c7.
//
// Solidity: event CollectionCreated(address indexed creator, address collection)
func (_FactoryNFT *FactoryNFTFilterer) FilterCollectionCreated(opts *bind.FilterOpts, creator []common.Address) (*FactoryNFTCollectionCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _FactoryNFT.contract.FilterLogs(opts, "CollectionCreated", creatorRule)
	if err != nil {
		return nil, err
	}
	return &FactoryNFTCollectionCreatedIterator{contract: _FactoryNFT.contract, event: "CollectionCreated", logs: logs, sub: sub}, nil
}

// WatchCollectionCreated is a free log subscription operation binding the contract event 0x5d0de243db1669e3a7056744cd715c625f0c1c348736c2c2d53d0ddebff1a6c7.
//
// Solidity: event CollectionCreated(address indexed creator, address collection)
func (_FactoryNFT *FactoryNFTFilterer) WatchCollectionCreated(opts *bind.WatchOpts, sink chan<- *FactoryNFTCollectionCreated, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _FactoryNFT.contract.WatchLogs(opts, "CollectionCreated", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryNFTCollectionCreated)
				if err := _FactoryNFT.contract.UnpackLog(event, "CollectionCreated", log); err != nil {
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

// ParseCollectionCreated is a log parse operation binding the contract event 0x5d0de243db1669e3a7056744cd715c625f0c1c348736c2c2d53d0ddebff1a6c7.
//
// Solidity: event CollectionCreated(address indexed creator, address collection)
func (_FactoryNFT *FactoryNFTFilterer) ParseCollectionCreated(log types.Log) (*FactoryNFTCollectionCreated, error) {
	event := new(FactoryNFTCollectionCreated)
	if err := _FactoryNFT.contract.UnpackLog(event, "CollectionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FactoryNFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FactoryNFT contract.
type FactoryNFTOwnershipTransferredIterator struct {
	Event *FactoryNFTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FactoryNFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryNFTOwnershipTransferred)
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
		it.Event = new(FactoryNFTOwnershipTransferred)
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
func (it *FactoryNFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryNFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryNFTOwnershipTransferred represents a OwnershipTransferred event raised by the FactoryNFT contract.
type FactoryNFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FactoryNFT *FactoryNFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FactoryNFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FactoryNFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FactoryNFTOwnershipTransferredIterator{contract: _FactoryNFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FactoryNFT *FactoryNFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FactoryNFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FactoryNFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryNFTOwnershipTransferred)
				if err := _FactoryNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FactoryNFT *FactoryNFTFilterer) ParseOwnershipTransferred(log types.Log) (*FactoryNFTOwnershipTransferred, error) {
	event := new(FactoryNFTOwnershipTransferred)
	if err := _FactoryNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
