// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// EventFactoryAttendeeTicket is an auto generated low-level Go binding around an user-defined struct.
type EventFactoryAttendeeTicket struct {
	TokenId       *big.Int
	TicketAddress common.Address
	EventAddress  common.Address
}

// EventFactoryMetaData contains all meta data concerning the EventFactory contract.
var EventFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_eventImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ticketImplementation\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"eventHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"clone\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"EventCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ticketHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"clone\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"TicketCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ticketAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"eventAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"TicketMinted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"attendeeToAllTickets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ticketAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eventAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"attendeeToTickets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventHash\",\"type\":\"string\"}],\"name\":\"createEvent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"eventAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ticketHash\",\"type\":\"string\"}],\"name\":\"createTicket\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"deployedEvents\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eventImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"attendee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eventAddress\",\"type\":\"address\"}],\"name\":\"getAttendeeTicketsForEvent\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeployedEvents\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getTicketOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"attendee\",\"type\":\"address\"}],\"name\":\"getTicketsOfAttendee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ticketAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eventAddress\",\"type\":\"address\"}],\"internalType\":\"structEventFactory.AttendeeTicket[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ticketAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eventAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"recordMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ticketImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ticketToAttendee\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// EventFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use EventFactoryMetaData.ABI instead.
var EventFactoryABI = EventFactoryMetaData.ABI

// EventFactory is an auto generated Go binding around an Ethereum contract.
type EventFactory struct {
	EventFactoryCaller     // Read-only binding to the contract
	EventFactoryTransactor // Write-only binding to the contract
	EventFactoryFilterer   // Log filterer for contract events
}

// EventFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventFactorySession struct {
	Contract     *EventFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventFactoryCallerSession struct {
	Contract *EventFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// EventFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventFactoryTransactorSession struct {
	Contract     *EventFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// EventFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventFactoryRaw struct {
	Contract *EventFactory // Generic contract binding to access the raw methods on
}

// EventFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventFactoryCallerRaw struct {
	Contract *EventFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// EventFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventFactoryTransactorRaw struct {
	Contract *EventFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEventFactory creates a new instance of EventFactory, bound to a specific deployed contract.
func NewEventFactory(address common.Address, backend bind.ContractBackend) (*EventFactory, error) {
	contract, err := bindEventFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EventFactory{EventFactoryCaller: EventFactoryCaller{contract: contract}, EventFactoryTransactor: EventFactoryTransactor{contract: contract}, EventFactoryFilterer: EventFactoryFilterer{contract: contract}}, nil
}

// NewEventFactoryCaller creates a new read-only instance of EventFactory, bound to a specific deployed contract.
func NewEventFactoryCaller(address common.Address, caller bind.ContractCaller) (*EventFactoryCaller, error) {
	contract, err := bindEventFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventFactoryCaller{contract: contract}, nil
}

// NewEventFactoryTransactor creates a new write-only instance of EventFactory, bound to a specific deployed contract.
func NewEventFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*EventFactoryTransactor, error) {
	contract, err := bindEventFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventFactoryTransactor{contract: contract}, nil
}

// NewEventFactoryFilterer creates a new log filterer instance of EventFactory, bound to a specific deployed contract.
func NewEventFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*EventFactoryFilterer, error) {
	contract, err := bindEventFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventFactoryFilterer{contract: contract}, nil
}

// bindEventFactory binds a generic wrapper to an already deployed contract.
func bindEventFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EventFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventFactory *EventFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventFactory.Contract.EventFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventFactory *EventFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventFactory.Contract.EventFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventFactory *EventFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventFactory.Contract.EventFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventFactory *EventFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventFactory *EventFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventFactory *EventFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventFactory.Contract.contract.Transact(opts, method, params...)
}

// AttendeeToAllTickets is a free data retrieval call binding the contract method 0x27aa6552.
//
// Solidity: function attendeeToAllTickets(address , uint256 ) view returns(uint256 tokenId, address ticketAddress, address eventAddress)
func (_EventFactory *EventFactoryCaller) AttendeeToAllTickets(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	TokenId       *big.Int
	TicketAddress common.Address
	EventAddress  common.Address
}, error) {
	var out []interface{}
	err := _EventFactory.contract.Call(opts, &out, "attendeeToAllTickets", arg0, arg1)

	outstruct := new(struct {
		TokenId       *big.Int
		TicketAddress common.Address
		EventAddress  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TicketAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.EventAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// AttendeeToAllTickets is a free data retrieval call binding the contract method 0x27aa6552.
//
// Solidity: function attendeeToAllTickets(address , uint256 ) view returns(uint256 tokenId, address ticketAddress, address eventAddress)
func (_EventFactory *EventFactorySession) AttendeeToAllTickets(arg0 common.Address, arg1 *big.Int) (struct {
	TokenId       *big.Int
	TicketAddress common.Address
	EventAddress  common.Address
}, error) {
	return _EventFactory.Contract.AttendeeToAllTickets(&_EventFactory.CallOpts, arg0, arg1)
}

// AttendeeToAllTickets is a free data retrieval call binding the contract method 0x27aa6552.
//
// Solidity: function attendeeToAllTickets(address , uint256 ) view returns(uint256 tokenId, address ticketAddress, address eventAddress)
func (_EventFactory *EventFactoryCallerSession) AttendeeToAllTickets(arg0 common.Address, arg1 *big.Int) (struct {
	TokenId       *big.Int
	TicketAddress common.Address
	EventAddress  common.Address
}, error) {
	return _EventFactory.Contract.AttendeeToAllTickets(&_EventFactory.CallOpts, arg0, arg1)
}

// AttendeeToTickets is a free data retrieval call binding the contract method 0x7231f6cf.
//
// Solidity: function attendeeToTickets(address , address , uint256 ) view returns(uint256)
func (_EventFactory *EventFactoryCaller) AttendeeToTickets(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EventFactory.contract.Call(opts, &out, "attendeeToTickets", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttendeeToTickets is a free data retrieval call binding the contract method 0x7231f6cf.
//
// Solidity: function attendeeToTickets(address , address , uint256 ) view returns(uint256)
func (_EventFactory *EventFactorySession) AttendeeToTickets(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	return _EventFactory.Contract.AttendeeToTickets(&_EventFactory.CallOpts, arg0, arg1, arg2)
}

// AttendeeToTickets is a free data retrieval call binding the contract method 0x7231f6cf.
//
// Solidity: function attendeeToTickets(address , address , uint256 ) view returns(uint256)
func (_EventFactory *EventFactoryCallerSession) AttendeeToTickets(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	return _EventFactory.Contract.AttendeeToTickets(&_EventFactory.CallOpts, arg0, arg1, arg2)
}

// DeployedEvents is a free data retrieval call binding the contract method 0xeade22f6.
//
// Solidity: function deployedEvents(uint256 ) view returns(address)
func (_EventFactory *EventFactoryCaller) DeployedEvents(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EventFactory.contract.Call(opts, &out, "deployedEvents", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DeployedEvents is a free data retrieval call binding the contract method 0xeade22f6.
//
// Solidity: function deployedEvents(uint256 ) view returns(address)
func (_EventFactory *EventFactorySession) DeployedEvents(arg0 *big.Int) (common.Address, error) {
	return _EventFactory.Contract.DeployedEvents(&_EventFactory.CallOpts, arg0)
}

// DeployedEvents is a free data retrieval call binding the contract method 0xeade22f6.
//
// Solidity: function deployedEvents(uint256 ) view returns(address)
func (_EventFactory *EventFactoryCallerSession) DeployedEvents(arg0 *big.Int) (common.Address, error) {
	return _EventFactory.Contract.DeployedEvents(&_EventFactory.CallOpts, arg0)
}

// EventImplementation is a free data retrieval call binding the contract method 0xcd8425be.
//
// Solidity: function eventImplementation() view returns(address)
func (_EventFactory *EventFactoryCaller) EventImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EventFactory.contract.Call(opts, &out, "eventImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EventImplementation is a free data retrieval call binding the contract method 0xcd8425be.
//
// Solidity: function eventImplementation() view returns(address)
func (_EventFactory *EventFactorySession) EventImplementation() (common.Address, error) {
	return _EventFactory.Contract.EventImplementation(&_EventFactory.CallOpts)
}

// EventImplementation is a free data retrieval call binding the contract method 0xcd8425be.
//
// Solidity: function eventImplementation() view returns(address)
func (_EventFactory *EventFactoryCallerSession) EventImplementation() (common.Address, error) {
	return _EventFactory.Contract.EventImplementation(&_EventFactory.CallOpts)
}

// GetAttendeeTicketsForEvent is a free data retrieval call binding the contract method 0x3678c7e8.
//
// Solidity: function getAttendeeTicketsForEvent(address attendee, address eventAddress) view returns(uint256[])
func (_EventFactory *EventFactoryCaller) GetAttendeeTicketsForEvent(opts *bind.CallOpts, attendee common.Address, eventAddress common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _EventFactory.contract.Call(opts, &out, "getAttendeeTicketsForEvent", attendee, eventAddress)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAttendeeTicketsForEvent is a free data retrieval call binding the contract method 0x3678c7e8.
//
// Solidity: function getAttendeeTicketsForEvent(address attendee, address eventAddress) view returns(uint256[])
func (_EventFactory *EventFactorySession) GetAttendeeTicketsForEvent(attendee common.Address, eventAddress common.Address) ([]*big.Int, error) {
	return _EventFactory.Contract.GetAttendeeTicketsForEvent(&_EventFactory.CallOpts, attendee, eventAddress)
}

// GetAttendeeTicketsForEvent is a free data retrieval call binding the contract method 0x3678c7e8.
//
// Solidity: function getAttendeeTicketsForEvent(address attendee, address eventAddress) view returns(uint256[])
func (_EventFactory *EventFactoryCallerSession) GetAttendeeTicketsForEvent(attendee common.Address, eventAddress common.Address) ([]*big.Int, error) {
	return _EventFactory.Contract.GetAttendeeTicketsForEvent(&_EventFactory.CallOpts, attendee, eventAddress)
}

// GetDeployedEvents is a free data retrieval call binding the contract method 0x3812783e.
//
// Solidity: function getDeployedEvents() view returns(address[])
func (_EventFactory *EventFactoryCaller) GetDeployedEvents(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EventFactory.contract.Call(opts, &out, "getDeployedEvents")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDeployedEvents is a free data retrieval call binding the contract method 0x3812783e.
//
// Solidity: function getDeployedEvents() view returns(address[])
func (_EventFactory *EventFactorySession) GetDeployedEvents() ([]common.Address, error) {
	return _EventFactory.Contract.GetDeployedEvents(&_EventFactory.CallOpts)
}

// GetDeployedEvents is a free data retrieval call binding the contract method 0x3812783e.
//
// Solidity: function getDeployedEvents() view returns(address[])
func (_EventFactory *EventFactoryCallerSession) GetDeployedEvents() ([]common.Address, error) {
	return _EventFactory.Contract.GetDeployedEvents(&_EventFactory.CallOpts)
}

// GetTicketOwner is a free data retrieval call binding the contract method 0xad093409.
//
// Solidity: function getTicketOwner(uint256 tokenId) view returns(address)
func (_EventFactory *EventFactoryCaller) GetTicketOwner(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EventFactory.contract.Call(opts, &out, "getTicketOwner", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTicketOwner is a free data retrieval call binding the contract method 0xad093409.
//
// Solidity: function getTicketOwner(uint256 tokenId) view returns(address)
func (_EventFactory *EventFactorySession) GetTicketOwner(tokenId *big.Int) (common.Address, error) {
	return _EventFactory.Contract.GetTicketOwner(&_EventFactory.CallOpts, tokenId)
}

// GetTicketOwner is a free data retrieval call binding the contract method 0xad093409.
//
// Solidity: function getTicketOwner(uint256 tokenId) view returns(address)
func (_EventFactory *EventFactoryCallerSession) GetTicketOwner(tokenId *big.Int) (common.Address, error) {
	return _EventFactory.Contract.GetTicketOwner(&_EventFactory.CallOpts, tokenId)
}

// GetTicketsOfAttendee is a free data retrieval call binding the contract method 0xc7f43db6.
//
// Solidity: function getTicketsOfAttendee(address attendee) view returns((uint256,address,address)[])
func (_EventFactory *EventFactoryCaller) GetTicketsOfAttendee(opts *bind.CallOpts, attendee common.Address) ([]EventFactoryAttendeeTicket, error) {
	var out []interface{}
	err := _EventFactory.contract.Call(opts, &out, "getTicketsOfAttendee", attendee)

	if err != nil {
		return *new([]EventFactoryAttendeeTicket), err
	}

	out0 := *abi.ConvertType(out[0], new([]EventFactoryAttendeeTicket)).(*[]EventFactoryAttendeeTicket)

	return out0, err

}

// GetTicketsOfAttendee is a free data retrieval call binding the contract method 0xc7f43db6.
//
// Solidity: function getTicketsOfAttendee(address attendee) view returns((uint256,address,address)[])
func (_EventFactory *EventFactorySession) GetTicketsOfAttendee(attendee common.Address) ([]EventFactoryAttendeeTicket, error) {
	return _EventFactory.Contract.GetTicketsOfAttendee(&_EventFactory.CallOpts, attendee)
}

// GetTicketsOfAttendee is a free data retrieval call binding the contract method 0xc7f43db6.
//
// Solidity: function getTicketsOfAttendee(address attendee) view returns((uint256,address,address)[])
func (_EventFactory *EventFactoryCallerSession) GetTicketsOfAttendee(attendee common.Address) ([]EventFactoryAttendeeTicket, error) {
	return _EventFactory.Contract.GetTicketsOfAttendee(&_EventFactory.CallOpts, attendee)
}

// TicketImplementation is a free data retrieval call binding the contract method 0x3a72eeab.
//
// Solidity: function ticketImplementation() view returns(address)
func (_EventFactory *EventFactoryCaller) TicketImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EventFactory.contract.Call(opts, &out, "ticketImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TicketImplementation is a free data retrieval call binding the contract method 0x3a72eeab.
//
// Solidity: function ticketImplementation() view returns(address)
func (_EventFactory *EventFactorySession) TicketImplementation() (common.Address, error) {
	return _EventFactory.Contract.TicketImplementation(&_EventFactory.CallOpts)
}

// TicketImplementation is a free data retrieval call binding the contract method 0x3a72eeab.
//
// Solidity: function ticketImplementation() view returns(address)
func (_EventFactory *EventFactoryCallerSession) TicketImplementation() (common.Address, error) {
	return _EventFactory.Contract.TicketImplementation(&_EventFactory.CallOpts)
}

// TicketToAttendee is a free data retrieval call binding the contract method 0x5de7bdc7.
//
// Solidity: function ticketToAttendee(uint256 ) view returns(address)
func (_EventFactory *EventFactoryCaller) TicketToAttendee(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EventFactory.contract.Call(opts, &out, "ticketToAttendee", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TicketToAttendee is a free data retrieval call binding the contract method 0x5de7bdc7.
//
// Solidity: function ticketToAttendee(uint256 ) view returns(address)
func (_EventFactory *EventFactorySession) TicketToAttendee(arg0 *big.Int) (common.Address, error) {
	return _EventFactory.Contract.TicketToAttendee(&_EventFactory.CallOpts, arg0)
}

// TicketToAttendee is a free data retrieval call binding the contract method 0x5de7bdc7.
//
// Solidity: function ticketToAttendee(uint256 ) view returns(address)
func (_EventFactory *EventFactoryCallerSession) TicketToAttendee(arg0 *big.Int) (common.Address, error) {
	return _EventFactory.Contract.TicketToAttendee(&_EventFactory.CallOpts, arg0)
}

// CreateEvent is a paid mutator transaction binding the contract method 0x9bc2bc71.
//
// Solidity: function createEvent(string eventHash) returns(address)
func (_EventFactory *EventFactoryTransactor) CreateEvent(opts *bind.TransactOpts, eventHash string) (*types.Transaction, error) {
	return _EventFactory.contract.Transact(opts, "createEvent", eventHash)
}

// CreateEvent is a paid mutator transaction binding the contract method 0x9bc2bc71.
//
// Solidity: function createEvent(string eventHash) returns(address)
func (_EventFactory *EventFactorySession) CreateEvent(eventHash string) (*types.Transaction, error) {
	return _EventFactory.Contract.CreateEvent(&_EventFactory.TransactOpts, eventHash)
}

// CreateEvent is a paid mutator transaction binding the contract method 0x9bc2bc71.
//
// Solidity: function createEvent(string eventHash) returns(address)
func (_EventFactory *EventFactoryTransactorSession) CreateEvent(eventHash string) (*types.Transaction, error) {
	return _EventFactory.Contract.CreateEvent(&_EventFactory.TransactOpts, eventHash)
}

// CreateTicket is a paid mutator transaction binding the contract method 0x12c359bb.
//
// Solidity: function createTicket(address eventAddress, string name, uint256 maxSupply, uint256 mintPrice, string ticketHash) returns(address)
func (_EventFactory *EventFactoryTransactor) CreateTicket(opts *bind.TransactOpts, eventAddress common.Address, name string, maxSupply *big.Int, mintPrice *big.Int, ticketHash string) (*types.Transaction, error) {
	return _EventFactory.contract.Transact(opts, "createTicket", eventAddress, name, maxSupply, mintPrice, ticketHash)
}

// CreateTicket is a paid mutator transaction binding the contract method 0x12c359bb.
//
// Solidity: function createTicket(address eventAddress, string name, uint256 maxSupply, uint256 mintPrice, string ticketHash) returns(address)
func (_EventFactory *EventFactorySession) CreateTicket(eventAddress common.Address, name string, maxSupply *big.Int, mintPrice *big.Int, ticketHash string) (*types.Transaction, error) {
	return _EventFactory.Contract.CreateTicket(&_EventFactory.TransactOpts, eventAddress, name, maxSupply, mintPrice, ticketHash)
}

// CreateTicket is a paid mutator transaction binding the contract method 0x12c359bb.
//
// Solidity: function createTicket(address eventAddress, string name, uint256 maxSupply, uint256 mintPrice, string ticketHash) returns(address)
func (_EventFactory *EventFactoryTransactorSession) CreateTicket(eventAddress common.Address, name string, maxSupply *big.Int, mintPrice *big.Int, ticketHash string) (*types.Transaction, error) {
	return _EventFactory.Contract.CreateTicket(&_EventFactory.TransactOpts, eventAddress, name, maxSupply, mintPrice, ticketHash)
}

// RecordMint is a paid mutator transaction binding the contract method 0x37715400.
//
// Solidity: function recordMint(uint256 tokenId, address ticketAddress, address eventAddress, address minter) returns()
func (_EventFactory *EventFactoryTransactor) RecordMint(opts *bind.TransactOpts, tokenId *big.Int, ticketAddress common.Address, eventAddress common.Address, minter common.Address) (*types.Transaction, error) {
	return _EventFactory.contract.Transact(opts, "recordMint", tokenId, ticketAddress, eventAddress, minter)
}

// RecordMint is a paid mutator transaction binding the contract method 0x37715400.
//
// Solidity: function recordMint(uint256 tokenId, address ticketAddress, address eventAddress, address minter) returns()
func (_EventFactory *EventFactorySession) RecordMint(tokenId *big.Int, ticketAddress common.Address, eventAddress common.Address, minter common.Address) (*types.Transaction, error) {
	return _EventFactory.Contract.RecordMint(&_EventFactory.TransactOpts, tokenId, ticketAddress, eventAddress, minter)
}

// RecordMint is a paid mutator transaction binding the contract method 0x37715400.
//
// Solidity: function recordMint(uint256 tokenId, address ticketAddress, address eventAddress, address minter) returns()
func (_EventFactory *EventFactoryTransactorSession) RecordMint(tokenId *big.Int, ticketAddress common.Address, eventAddress common.Address, minter common.Address) (*types.Transaction, error) {
	return _EventFactory.Contract.RecordMint(&_EventFactory.TransactOpts, tokenId, ticketAddress, eventAddress, minter)
}

// EventFactoryEventCreatedIterator is returned from FilterEventCreated and is used to iterate over the raw logs and unpacked data for EventCreated events raised by the EventFactory contract.
type EventFactoryEventCreatedIterator struct {
	Event *EventFactoryEventCreated // Event containing the contract specifics and raw log

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
func (it *EventFactoryEventCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventFactoryEventCreated)
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
		it.Event = new(EventFactoryEventCreated)
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
func (it *EventFactoryEventCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventFactoryEventCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventFactoryEventCreated represents a EventCreated event raised by the EventFactory contract.
type EventFactoryEventCreated struct {
	EventHash string
	Clone     common.Address
	Sender    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEventCreated is a free log retrieval operation binding the contract event 0x83de87bf204a6b6ef0a90d876d58169771736fe338aad2c42516db0a334802fa.
//
// Solidity: event EventCreated(string eventHash, address clone, address sender)
func (_EventFactory *EventFactoryFilterer) FilterEventCreated(opts *bind.FilterOpts) (*EventFactoryEventCreatedIterator, error) {

	logs, sub, err := _EventFactory.contract.FilterLogs(opts, "EventCreated")
	if err != nil {
		return nil, err
	}
	return &EventFactoryEventCreatedIterator{contract: _EventFactory.contract, event: "EventCreated", logs: logs, sub: sub}, nil
}

// WatchEventCreated is a free log subscription operation binding the contract event 0x83de87bf204a6b6ef0a90d876d58169771736fe338aad2c42516db0a334802fa.
//
// Solidity: event EventCreated(string eventHash, address clone, address sender)
func (_EventFactory *EventFactoryFilterer) WatchEventCreated(opts *bind.WatchOpts, sink chan<- *EventFactoryEventCreated) (event.Subscription, error) {

	logs, sub, err := _EventFactory.contract.WatchLogs(opts, "EventCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventFactoryEventCreated)
				if err := _EventFactory.contract.UnpackLog(event, "EventCreated", log); err != nil {
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

// ParseEventCreated is a log parse operation binding the contract event 0x83de87bf204a6b6ef0a90d876d58169771736fe338aad2c42516db0a334802fa.
//
// Solidity: event EventCreated(string eventHash, address clone, address sender)
func (_EventFactory *EventFactoryFilterer) ParseEventCreated(log types.Log) (*EventFactoryEventCreated, error) {
	event := new(EventFactoryEventCreated)
	if err := _EventFactory.contract.UnpackLog(event, "EventCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventFactoryTicketCreatedIterator is returned from FilterTicketCreated and is used to iterate over the raw logs and unpacked data for TicketCreated events raised by the EventFactory contract.
type EventFactoryTicketCreatedIterator struct {
	Event *EventFactoryTicketCreated // Event containing the contract specifics and raw log

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
func (it *EventFactoryTicketCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventFactoryTicketCreated)
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
		it.Event = new(EventFactoryTicketCreated)
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
func (it *EventFactoryTicketCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventFactoryTicketCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventFactoryTicketCreated represents a TicketCreated event raised by the EventFactory contract.
type EventFactoryTicketCreated struct {
	TicketHash string
	Clone      common.Address
	Sender     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTicketCreated is a free log retrieval operation binding the contract event 0x3bb616572b8688b40a323300eb4f02fb815b05c7fa7e3227991d403fce0eb184.
//
// Solidity: event TicketCreated(string ticketHash, address clone, address sender)
func (_EventFactory *EventFactoryFilterer) FilterTicketCreated(opts *bind.FilterOpts) (*EventFactoryTicketCreatedIterator, error) {

	logs, sub, err := _EventFactory.contract.FilterLogs(opts, "TicketCreated")
	if err != nil {
		return nil, err
	}
	return &EventFactoryTicketCreatedIterator{contract: _EventFactory.contract, event: "TicketCreated", logs: logs, sub: sub}, nil
}

// WatchTicketCreated is a free log subscription operation binding the contract event 0x3bb616572b8688b40a323300eb4f02fb815b05c7fa7e3227991d403fce0eb184.
//
// Solidity: event TicketCreated(string ticketHash, address clone, address sender)
func (_EventFactory *EventFactoryFilterer) WatchTicketCreated(opts *bind.WatchOpts, sink chan<- *EventFactoryTicketCreated) (event.Subscription, error) {

	logs, sub, err := _EventFactory.contract.WatchLogs(opts, "TicketCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventFactoryTicketCreated)
				if err := _EventFactory.contract.UnpackLog(event, "TicketCreated", log); err != nil {
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

// ParseTicketCreated is a log parse operation binding the contract event 0x3bb616572b8688b40a323300eb4f02fb815b05c7fa7e3227991d403fce0eb184.
//
// Solidity: event TicketCreated(string ticketHash, address clone, address sender)
func (_EventFactory *EventFactoryFilterer) ParseTicketCreated(log types.Log) (*EventFactoryTicketCreated, error) {
	event := new(EventFactoryTicketCreated)
	if err := _EventFactory.contract.UnpackLog(event, "TicketCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventFactoryTicketMintedIterator is returned from FilterTicketMinted and is used to iterate over the raw logs and unpacked data for TicketMinted events raised by the EventFactory contract.
type EventFactoryTicketMintedIterator struct {
	Event *EventFactoryTicketMinted // Event containing the contract specifics and raw log

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
func (it *EventFactoryTicketMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventFactoryTicketMinted)
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
		it.Event = new(EventFactoryTicketMinted)
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
func (it *EventFactoryTicketMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventFactoryTicketMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventFactoryTicketMinted represents a TicketMinted event raised by the EventFactory contract.
type EventFactoryTicketMinted struct {
	TokenId       *big.Int
	TicketAddress common.Address
	EventAddress  common.Address
	BlockNumber   *big.Int
	Minter        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTicketMinted is a free log retrieval operation binding the contract event 0x559db1a551c77fc0584b30434dfaf17be73a7fb313edc3ce2577b2fd08b11270.
//
// Solidity: event TicketMinted(uint256 indexed tokenId, address indexed ticketAddress, address indexed eventAddress, uint256 blockNumber, address minter)
func (_EventFactory *EventFactoryFilterer) FilterTicketMinted(opts *bind.FilterOpts, tokenId []*big.Int, ticketAddress []common.Address, eventAddress []common.Address) (*EventFactoryTicketMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ticketAddressRule []interface{}
	for _, ticketAddressItem := range ticketAddress {
		ticketAddressRule = append(ticketAddressRule, ticketAddressItem)
	}
	var eventAddressRule []interface{}
	for _, eventAddressItem := range eventAddress {
		eventAddressRule = append(eventAddressRule, eventAddressItem)
	}

	logs, sub, err := _EventFactory.contract.FilterLogs(opts, "TicketMinted", tokenIdRule, ticketAddressRule, eventAddressRule)
	if err != nil {
		return nil, err
	}
	return &EventFactoryTicketMintedIterator{contract: _EventFactory.contract, event: "TicketMinted", logs: logs, sub: sub}, nil
}

// WatchTicketMinted is a free log subscription operation binding the contract event 0x559db1a551c77fc0584b30434dfaf17be73a7fb313edc3ce2577b2fd08b11270.
//
// Solidity: event TicketMinted(uint256 indexed tokenId, address indexed ticketAddress, address indexed eventAddress, uint256 blockNumber, address minter)
func (_EventFactory *EventFactoryFilterer) WatchTicketMinted(opts *bind.WatchOpts, sink chan<- *EventFactoryTicketMinted, tokenId []*big.Int, ticketAddress []common.Address, eventAddress []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ticketAddressRule []interface{}
	for _, ticketAddressItem := range ticketAddress {
		ticketAddressRule = append(ticketAddressRule, ticketAddressItem)
	}
	var eventAddressRule []interface{}
	for _, eventAddressItem := range eventAddress {
		eventAddressRule = append(eventAddressRule, eventAddressItem)
	}

	logs, sub, err := _EventFactory.contract.WatchLogs(opts, "TicketMinted", tokenIdRule, ticketAddressRule, eventAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventFactoryTicketMinted)
				if err := _EventFactory.contract.UnpackLog(event, "TicketMinted", log); err != nil {
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

// ParseTicketMinted is a log parse operation binding the contract event 0x559db1a551c77fc0584b30434dfaf17be73a7fb313edc3ce2577b2fd08b11270.
//
// Solidity: event TicketMinted(uint256 indexed tokenId, address indexed ticketAddress, address indexed eventAddress, uint256 blockNumber, address minter)
func (_EventFactory *EventFactoryFilterer) ParseTicketMinted(log types.Log) (*EventFactoryTicketMinted, error) {
	event := new(EventFactoryTicketMinted)
	if err := _EventFactory.contract.UnpackLog(event, "TicketMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
