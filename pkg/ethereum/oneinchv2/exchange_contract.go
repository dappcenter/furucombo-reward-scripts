// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oneinchv2

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

// ExchangeContractABI is the input ABI used to generate the binding from.
const ExchangeContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"spentAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"guaranteedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"Swapped\",\"type\":\"event\"}]"

// ExchangeContract is an auto generated Go binding around an Ethereum contract.
type ExchangeContract struct {
	ExchangeContractCaller     // Read-only binding to the contract
	ExchangeContractTransactor // Write-only binding to the contract
	ExchangeContractFilterer   // Log filterer for contract events
}

// ExchangeContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeContractSession struct {
	Contract     *ExchangeContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeContractCallerSession struct {
	Contract *ExchangeContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ExchangeContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeContractTransactorSession struct {
	Contract     *ExchangeContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ExchangeContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeContractRaw struct {
	Contract *ExchangeContract // Generic contract binding to access the raw methods on
}

// ExchangeContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeContractCallerRaw struct {
	Contract *ExchangeContractCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeContractTransactorRaw struct {
	Contract *ExchangeContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchangeContract creates a new instance of ExchangeContract, bound to a specific deployed contract.
func NewExchangeContract(address common.Address, backend bind.ContractBackend) (*ExchangeContract, error) {
	contract, err := bindExchangeContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExchangeContract{ExchangeContractCaller: ExchangeContractCaller{contract: contract}, ExchangeContractTransactor: ExchangeContractTransactor{contract: contract}, ExchangeContractFilterer: ExchangeContractFilterer{contract: contract}}, nil
}

// NewExchangeContractCaller creates a new read-only instance of ExchangeContract, bound to a specific deployed contract.
func NewExchangeContractCaller(address common.Address, caller bind.ContractCaller) (*ExchangeContractCaller, error) {
	contract, err := bindExchangeContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeContractCaller{contract: contract}, nil
}

// NewExchangeContractTransactor creates a new write-only instance of ExchangeContract, bound to a specific deployed contract.
func NewExchangeContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeContractTransactor, error) {
	contract, err := bindExchangeContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeContractTransactor{contract: contract}, nil
}

// NewExchangeContractFilterer creates a new log filterer instance of ExchangeContract, bound to a specific deployed contract.
func NewExchangeContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeContractFilterer, error) {
	contract, err := bindExchangeContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeContractFilterer{contract: contract}, nil
}

// bindExchangeContract binds a generic wrapper to an already deployed contract.
func bindExchangeContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeContract *ExchangeContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExchangeContract.Contract.ExchangeContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeContract *ExchangeContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeContract.Contract.ExchangeContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeContract *ExchangeContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeContract.Contract.ExchangeContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeContract *ExchangeContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExchangeContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeContract *ExchangeContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeContract *ExchangeContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeContract.Contract.contract.Transact(opts, method, params...)
}

// ExchangeContractSwappedIterator is returned from FilterSwapped and is used to iterate over the raw logs and unpacked data for Swapped events raised by the ExchangeContract contract.
type ExchangeContractSwappedIterator struct {
	Event *ExchangeContractSwapped // Event containing the contract specifics and raw log

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
func (it *ExchangeContractSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeContractSwapped)
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
		it.Event = new(ExchangeContractSwapped)
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
func (it *ExchangeContractSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeContractSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeContractSwapped represents a Swapped event raised by the ExchangeContract contract.
type ExchangeContractSwapped struct {
	Sender           common.Address
	SrcToken         common.Address
	DstToken         common.Address
	DstReceiver      common.Address
	Amount           *big.Int
	SpentAmount      *big.Int
	ReturnAmount     *big.Int
	MinReturnAmount  *big.Int
	GuaranteedAmount *big.Int
	Referrer         common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSwapped is a free log retrieval operation binding the contract event 0x76af224a143865a50b41496e1a73622698692c565c1214bc862f18e22d829c5e.
//
// Solidity: event Swapped(address indexed sender, address indexed srcToken, address indexed dstToken, address dstReceiver, uint256 amount, uint256 spentAmount, uint256 returnAmount, uint256 minReturnAmount, uint256 guaranteedAmount, address referrer)
func (_ExchangeContract *ExchangeContractFilterer) FilterSwapped(opts *bind.FilterOpts, sender []common.Address, srcToken []common.Address, dstToken []common.Address) (*ExchangeContractSwappedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var srcTokenRule []interface{}
	for _, srcTokenItem := range srcToken {
		srcTokenRule = append(srcTokenRule, srcTokenItem)
	}
	var dstTokenRule []interface{}
	for _, dstTokenItem := range dstToken {
		dstTokenRule = append(dstTokenRule, dstTokenItem)
	}

	logs, sub, err := _ExchangeContract.contract.FilterLogs(opts, "Swapped", senderRule, srcTokenRule, dstTokenRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeContractSwappedIterator{contract: _ExchangeContract.contract, event: "Swapped", logs: logs, sub: sub}, nil
}

// WatchSwapped is a free log subscription operation binding the contract event 0x76af224a143865a50b41496e1a73622698692c565c1214bc862f18e22d829c5e.
//
// Solidity: event Swapped(address indexed sender, address indexed srcToken, address indexed dstToken, address dstReceiver, uint256 amount, uint256 spentAmount, uint256 returnAmount, uint256 minReturnAmount, uint256 guaranteedAmount, address referrer)
func (_ExchangeContract *ExchangeContractFilterer) WatchSwapped(opts *bind.WatchOpts, sink chan<- *ExchangeContractSwapped, sender []common.Address, srcToken []common.Address, dstToken []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var srcTokenRule []interface{}
	for _, srcTokenItem := range srcToken {
		srcTokenRule = append(srcTokenRule, srcTokenItem)
	}
	var dstTokenRule []interface{}
	for _, dstTokenItem := range dstToken {
		dstTokenRule = append(dstTokenRule, dstTokenItem)
	}

	logs, sub, err := _ExchangeContract.contract.WatchLogs(opts, "Swapped", senderRule, srcTokenRule, dstTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeContractSwapped)
				if err := _ExchangeContract.contract.UnpackLog(event, "Swapped", log); err != nil {
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

// ParseSwapped is a log parse operation binding the contract event 0x76af224a143865a50b41496e1a73622698692c565c1214bc862f18e22d829c5e.
//
// Solidity: event Swapped(address indexed sender, address indexed srcToken, address indexed dstToken, address dstReceiver, uint256 amount, uint256 spentAmount, uint256 returnAmount, uint256 minReturnAmount, uint256 guaranteedAmount, address referrer)
func (_ExchangeContract *ExchangeContractFilterer) ParseSwapped(log types.Log) (*ExchangeContractSwapped, error) {
	event := new(ExchangeContractSwapped)
	if err := _ExchangeContract.contract.UnpackLog(event, "Swapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
