// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractPriceFeedAdapter

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

// ContractPriceFeedAdapterMetaData contains all meta data concerning the ContractPriceFeedAdapter contract.
var ContractPriceFeedAdapterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addFeed\",\"inputs\":[{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_feedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"diaOracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeds\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractAggregatorV3Interface\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestPrice\",\"inputs\":[{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPriceDia\",\"inputs\":[{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"Price\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeFeed\",\"inputs\":[{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDiaOracleAddress\",\"inputs\":[{\"name\":\"_diaOracle\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FeedAdded\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"feedAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeedRemoved\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"chainlinkAdapterPrice\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"_ChainlinkAdapterPrice\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"diaAdapterPrice\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"_DiaAdapterPrice\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"diaOracleAddress\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"_DiaOracleAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b610aab8061007e6000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80638da5cb5b116100665780638da5cb5b1461013757806390cab41614610148578063abd8a6c71461015b578063f2fde38b14610186578063f66a1b711461019957600080fd5b80632ce773cd146100a357806349b3ce56146100d3578063685e8d80146100e85780636a811e3b146100fb578063715018a61461012f575b600080fd5b6002546100b6906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6100e66100e136600461082e565b6101ba565b005b6100e66100f6366004610887565b6102a0565b6100b661010936600461082e565b80516020818301810180516001825292820191909301209152546001600160a01b031681565b6100e6610384565b6000546001600160a01b03166100b6565b6100e66101563660046108a2565b610398565b61016e61016936600461082e565b6104c3565b6040516001600160801b0390911681526020016100ca565b6100e6610194366004610887565b61053b565b6101ac6101a736600461082e565b6105b4565b6040519081526020016100ca565b6101c26106e1565b60006001600160a01b03166001826040516101dd9190610920565b908152604051908190036020019020546001600160a01b0316141561023b5760405162461bcd60e51b815260206004820152600f60248201526e2332b2b2103737ba103337bab7321760891b60448201526064015b60405180910390fd5b60018160405161024b9190610920565b90815260405190819003602001812080546001600160a01b03191690557f06aa2b76883f4d9b405ce5d8cb618072ed06ad9a9f4a3a638d6a76f995f950ce90610295908390610968565b60405180910390a150565b6102a86106e1565b6001600160a01b0381166103095760405162461bcd60e51b815260206004820152602260248201527f446961204f7261636c6520616464726573732063616e6e6f74206265207a6572604482015261379760f11b6064820152608401610232565b600280546001600160a01b0319166001600160a01b0383169081179091556040805181815260129181019190915271446961204f7261636c65204164647265737360701b606082015260208101919091527f34fa09dd73d22ca570e8b4a45ec050b8a0cbe4a14b26df98ddeef0016e87fc1790608001610295565b61038c6106e1565b610396600061073b565b565b6103a06106e1565b6001600160a01b0381166103f65760405162461bcd60e51b815260206004820152601c60248201527f4665656420616464726573732063616e6e6f74206265207a65726f2e000000006044820152606401610232565b60008251116104475760405162461bcd60e51b815260206004820152601760248201527f53796d626f6c2063616e6e6f7420626520656d7074792e0000000000000000006044820152606401610232565b806001836040516104589190610920565b90815260405190819003602001812080546001600160a01b03939093166001600160a01b0319909316929092179091557f917a6ab4f9d1b23331c7307351d940aa91024649e9745f66c8d7d9ccdb5f3704906104b7908490849061097b565b60405180910390a15050565b6002546040516304b01c2560e51b81526000916001600160a01b03169063960384a0906104f4908590600401610968565b6040805180830381865afa158015610510573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061053491906109bc565b9392505050565b6105436106e1565b6001600160a01b0381166105a85760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610232565b6105b18161073b565b50565b6000806001836040516105c79190610920565b908152604051908190036020019020546001600160a01b03169050806106215760405162461bcd60e51b815260206004820152600f60248201526e2332b2b2103737ba103337bab7321760891b6044820152606401610232565b600080826001600160a01b031663feaf968c6040518163ffffffff1660e01b815260040160a060405180830381865afa158015610662573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106869190610a00565b50935050925050610e10814261069c9190610a50565b106106d95760405162461bcd60e51b815260206004820152600d60248201526c44617461206973207374616c6560981b6044820152606401610232565b509392505050565b6000546001600160a01b031633146103965760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610232565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126107b257600080fd5b813567ffffffffffffffff808211156107cd576107cd61078b565b604051601f8301601f19908116603f011681019082821181831017156107f5576107f561078b565b8160405283815286602085880101111561080e57600080fd5b836020870160208301376000602085830101528094505050505092915050565b60006020828403121561084057600080fd5b813567ffffffffffffffff81111561085757600080fd5b610863848285016107a1565b949350505050565b80356001600160a01b038116811461088257600080fd5b919050565b60006020828403121561089957600080fd5b6105348261086b565b600080604083850312156108b557600080fd5b823567ffffffffffffffff8111156108cc57600080fd5b6108d8858286016107a1565b9250506108e76020840161086b565b90509250929050565b60005b8381101561090b5781810151838201526020016108f3565b8381111561091a576000848401525b50505050565b600082516109328184602087016108f0565b9190910192915050565b600081518084526109548160208601602086016108f0565b601f01601f19169290920160200192915050565b602081526000610534602083018461093c565b60408152600061098e604083018561093c565b905060018060a01b03831660208301529392505050565b80516001600160801b038116811461088257600080fd5b600080604083850312156109cf57600080fd5b6109d8836109a5565b91506108e7602084016109a5565b805169ffffffffffffffffffff8116811461088257600080fd5b600080600080600060a08688031215610a1857600080fd5b610a21866109e6565b9450602086015193506040860151925060608601519150610a44608087016109e6565b90509295509295909350565b600082821015610a7057634e487b7160e01b600052601160045260246000fd5b50039056fea26469706673582212209ca23d3736990c32bd215f69e759401bd81d1e5876227568ac79df3667a348a864736f6c634300080c0033",
}

// ContractPriceFeedAdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractPriceFeedAdapterMetaData.ABI instead.
var ContractPriceFeedAdapterABI = ContractPriceFeedAdapterMetaData.ABI

// ContractPriceFeedAdapterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractPriceFeedAdapterMetaData.Bin instead.
var ContractPriceFeedAdapterBin = ContractPriceFeedAdapterMetaData.Bin

// DeployContractPriceFeedAdapter deploys a new Ethereum contract, binding an instance of ContractPriceFeedAdapter to it.
func DeployContractPriceFeedAdapter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractPriceFeedAdapter, error) {
	parsed, err := ContractPriceFeedAdapterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractPriceFeedAdapterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractPriceFeedAdapter{ContractPriceFeedAdapterCaller: ContractPriceFeedAdapterCaller{contract: contract}, ContractPriceFeedAdapterTransactor: ContractPriceFeedAdapterTransactor{contract: contract}, ContractPriceFeedAdapterFilterer: ContractPriceFeedAdapterFilterer{contract: contract}}, nil
}

// ContractPriceFeedAdapter is an auto generated Go binding around an Ethereum contract.
type ContractPriceFeedAdapter struct {
	ContractPriceFeedAdapterCaller     // Read-only binding to the contract
	ContractPriceFeedAdapterTransactor // Write-only binding to the contract
	ContractPriceFeedAdapterFilterer   // Log filterer for contract events
}

// ContractPriceFeedAdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractPriceFeedAdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractPriceFeedAdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractPriceFeedAdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractPriceFeedAdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractPriceFeedAdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractPriceFeedAdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractPriceFeedAdapterSession struct {
	Contract     *ContractPriceFeedAdapter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ContractPriceFeedAdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractPriceFeedAdapterCallerSession struct {
	Contract *ContractPriceFeedAdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ContractPriceFeedAdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractPriceFeedAdapterTransactorSession struct {
	Contract     *ContractPriceFeedAdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ContractPriceFeedAdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractPriceFeedAdapterRaw struct {
	Contract *ContractPriceFeedAdapter // Generic contract binding to access the raw methods on
}

// ContractPriceFeedAdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractPriceFeedAdapterCallerRaw struct {
	Contract *ContractPriceFeedAdapterCaller // Generic read-only contract binding to access the raw methods on
}

// ContractPriceFeedAdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractPriceFeedAdapterTransactorRaw struct {
	Contract *ContractPriceFeedAdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractPriceFeedAdapter creates a new instance of ContractPriceFeedAdapter, bound to a specific deployed contract.
func NewContractPriceFeedAdapter(address common.Address, backend bind.ContractBackend) (*ContractPriceFeedAdapter, error) {
	contract, err := bindContractPriceFeedAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractPriceFeedAdapter{ContractPriceFeedAdapterCaller: ContractPriceFeedAdapterCaller{contract: contract}, ContractPriceFeedAdapterTransactor: ContractPriceFeedAdapterTransactor{contract: contract}, ContractPriceFeedAdapterFilterer: ContractPriceFeedAdapterFilterer{contract: contract}}, nil
}

// NewContractPriceFeedAdapterCaller creates a new read-only instance of ContractPriceFeedAdapter, bound to a specific deployed contract.
func NewContractPriceFeedAdapterCaller(address common.Address, caller bind.ContractCaller) (*ContractPriceFeedAdapterCaller, error) {
	contract, err := bindContractPriceFeedAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractPriceFeedAdapterCaller{contract: contract}, nil
}

// NewContractPriceFeedAdapterTransactor creates a new write-only instance of ContractPriceFeedAdapter, bound to a specific deployed contract.
func NewContractPriceFeedAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractPriceFeedAdapterTransactor, error) {
	contract, err := bindContractPriceFeedAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractPriceFeedAdapterTransactor{contract: contract}, nil
}

// NewContractPriceFeedAdapterFilterer creates a new log filterer instance of ContractPriceFeedAdapter, bound to a specific deployed contract.
func NewContractPriceFeedAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractPriceFeedAdapterFilterer, error) {
	contract, err := bindContractPriceFeedAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractPriceFeedAdapterFilterer{contract: contract}, nil
}

// bindContractPriceFeedAdapter binds a generic wrapper to an already deployed contract.
func bindContractPriceFeedAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractPriceFeedAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractPriceFeedAdapter.Contract.ContractPriceFeedAdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.Contract.ContractPriceFeedAdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.Contract.ContractPriceFeedAdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractPriceFeedAdapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.Contract.contract.Transact(opts, method, params...)
}

// DiaOracle is a free data retrieval call binding the contract method 0x2ce773cd.
//
// Solidity: function diaOracle() view returns(address)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterCaller) DiaOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractPriceFeedAdapter.contract.Call(opts, &out, "diaOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DiaOracle is a free data retrieval call binding the contract method 0x2ce773cd.
//
// Solidity: function diaOracle() view returns(address)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterSession) DiaOracle() (common.Address, error) {
	return _ContractPriceFeedAdapter.Contract.DiaOracle(&_ContractPriceFeedAdapter.CallOpts)
}

// DiaOracle is a free data retrieval call binding the contract method 0x2ce773cd.
//
// Solidity: function diaOracle() view returns(address)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterCallerSession) DiaOracle() (common.Address, error) {
	return _ContractPriceFeedAdapter.Contract.DiaOracle(&_ContractPriceFeedAdapter.CallOpts)
}

// Feeds is a free data retrieval call binding the contract method 0x6a811e3b.
//
// Solidity: function feeds(string ) view returns(address)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterCaller) Feeds(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _ContractPriceFeedAdapter.contract.Call(opts, &out, "feeds", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Feeds is a free data retrieval call binding the contract method 0x6a811e3b.
//
// Solidity: function feeds(string ) view returns(address)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterSession) Feeds(arg0 string) (common.Address, error) {
	return _ContractPriceFeedAdapter.Contract.Feeds(&_ContractPriceFeedAdapter.CallOpts, arg0)
}

// Feeds is a free data retrieval call binding the contract method 0x6a811e3b.
//
// Solidity: function feeds(string ) view returns(address)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterCallerSession) Feeds(arg0 string) (common.Address, error) {
	return _ContractPriceFeedAdapter.Contract.Feeds(&_ContractPriceFeedAdapter.CallOpts, arg0)
}

// GetLatestPrice is a free data retrieval call binding the contract method 0xf66a1b71.
//
// Solidity: function getLatestPrice(string _symbol) view returns(int256)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterCaller) GetLatestPrice(opts *bind.CallOpts, _symbol string) (*big.Int, error) {
	var out []interface{}
	err := _ContractPriceFeedAdapter.contract.Call(opts, &out, "getLatestPrice", _symbol)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPrice is a free data retrieval call binding the contract method 0xf66a1b71.
//
// Solidity: function getLatestPrice(string _symbol) view returns(int256)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterSession) GetLatestPrice(_symbol string) (*big.Int, error) {
	return _ContractPriceFeedAdapter.Contract.GetLatestPrice(&_ContractPriceFeedAdapter.CallOpts, _symbol)
}

// GetLatestPrice is a free data retrieval call binding the contract method 0xf66a1b71.
//
// Solidity: function getLatestPrice(string _symbol) view returns(int256)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterCallerSession) GetLatestPrice(_symbol string) (*big.Int, error) {
	return _ContractPriceFeedAdapter.Contract.GetLatestPrice(&_ContractPriceFeedAdapter.CallOpts, _symbol)
}

// GetPriceDia is a free data retrieval call binding the contract method 0xabd8a6c7.
//
// Solidity: function getPriceDia(string _symbol) view returns(uint128 Price)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterCaller) GetPriceDia(opts *bind.CallOpts, _symbol string) (*big.Int, error) {
	var out []interface{}
	err := _ContractPriceFeedAdapter.contract.Call(opts, &out, "getPriceDia", _symbol)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriceDia is a free data retrieval call binding the contract method 0xabd8a6c7.
//
// Solidity: function getPriceDia(string _symbol) view returns(uint128 Price)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterSession) GetPriceDia(_symbol string) (*big.Int, error) {
	return _ContractPriceFeedAdapter.Contract.GetPriceDia(&_ContractPriceFeedAdapter.CallOpts, _symbol)
}

// GetPriceDia is a free data retrieval call binding the contract method 0xabd8a6c7.
//
// Solidity: function getPriceDia(string _symbol) view returns(uint128 Price)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterCallerSession) GetPriceDia(_symbol string) (*big.Int, error) {
	return _ContractPriceFeedAdapter.Contract.GetPriceDia(&_ContractPriceFeedAdapter.CallOpts, _symbol)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractPriceFeedAdapter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterSession) Owner() (common.Address, error) {
	return _ContractPriceFeedAdapter.Contract.Owner(&_ContractPriceFeedAdapter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterCallerSession) Owner() (common.Address, error) {
	return _ContractPriceFeedAdapter.Contract.Owner(&_ContractPriceFeedAdapter.CallOpts)
}

// AddFeed is a paid mutator transaction binding the contract method 0x90cab416.
//
// Solidity: function addFeed(string _symbol, address _feedAddress) returns()
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterTransactor) AddFeed(opts *bind.TransactOpts, _symbol string, _feedAddress common.Address) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.contract.Transact(opts, "addFeed", _symbol, _feedAddress)
}

// AddFeed is a paid mutator transaction binding the contract method 0x90cab416.
//
// Solidity: function addFeed(string _symbol, address _feedAddress) returns()
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterSession) AddFeed(_symbol string, _feedAddress common.Address) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.Contract.AddFeed(&_ContractPriceFeedAdapter.TransactOpts, _symbol, _feedAddress)
}

// AddFeed is a paid mutator transaction binding the contract method 0x90cab416.
//
// Solidity: function addFeed(string _symbol, address _feedAddress) returns()
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterTransactorSession) AddFeed(_symbol string, _feedAddress common.Address) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.Contract.AddFeed(&_ContractPriceFeedAdapter.TransactOpts, _symbol, _feedAddress)
}

// RemoveFeed is a paid mutator transaction binding the contract method 0x49b3ce56.
//
// Solidity: function removeFeed(string _symbol) returns()
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterTransactor) RemoveFeed(opts *bind.TransactOpts, _symbol string) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.contract.Transact(opts, "removeFeed", _symbol)
}

// RemoveFeed is a paid mutator transaction binding the contract method 0x49b3ce56.
//
// Solidity: function removeFeed(string _symbol) returns()
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterSession) RemoveFeed(_symbol string) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.Contract.RemoveFeed(&_ContractPriceFeedAdapter.TransactOpts, _symbol)
}

// RemoveFeed is a paid mutator transaction binding the contract method 0x49b3ce56.
//
// Solidity: function removeFeed(string _symbol) returns()
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterTransactorSession) RemoveFeed(_symbol string) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.Contract.RemoveFeed(&_ContractPriceFeedAdapter.TransactOpts, _symbol)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.Contract.RenounceOwnership(&_ContractPriceFeedAdapter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.Contract.RenounceOwnership(&_ContractPriceFeedAdapter.TransactOpts)
}

// SetDiaOracleAddress is a paid mutator transaction binding the contract method 0x685e8d80.
//
// Solidity: function setDiaOracleAddress(address _diaOracle) returns()
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterTransactor) SetDiaOracleAddress(opts *bind.TransactOpts, _diaOracle common.Address) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.contract.Transact(opts, "setDiaOracleAddress", _diaOracle)
}

// SetDiaOracleAddress is a paid mutator transaction binding the contract method 0x685e8d80.
//
// Solidity: function setDiaOracleAddress(address _diaOracle) returns()
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterSession) SetDiaOracleAddress(_diaOracle common.Address) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.Contract.SetDiaOracleAddress(&_ContractPriceFeedAdapter.TransactOpts, _diaOracle)
}

// SetDiaOracleAddress is a paid mutator transaction binding the contract method 0x685e8d80.
//
// Solidity: function setDiaOracleAddress(address _diaOracle) returns()
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterTransactorSession) SetDiaOracleAddress(_diaOracle common.Address) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.Contract.SetDiaOracleAddress(&_ContractPriceFeedAdapter.TransactOpts, _diaOracle)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.Contract.TransferOwnership(&_ContractPriceFeedAdapter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.Contract.TransferOwnership(&_ContractPriceFeedAdapter.TransactOpts, newOwner)
}

// ContractPriceFeedAdapterFeedAddedIterator is returned from FilterFeedAdded and is used to iterate over the raw logs and unpacked data for FeedAdded events raised by the ContractPriceFeedAdapter contract.
type ContractPriceFeedAdapterFeedAddedIterator struct {
	Event *ContractPriceFeedAdapterFeedAdded // Event containing the contract specifics and raw log

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
func (it *ContractPriceFeedAdapterFeedAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPriceFeedAdapterFeedAdded)
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
		it.Event = new(ContractPriceFeedAdapterFeedAdded)
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
func (it *ContractPriceFeedAdapterFeedAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPriceFeedAdapterFeedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPriceFeedAdapterFeedAdded represents a FeedAdded event raised by the ContractPriceFeedAdapter contract.
type ContractPriceFeedAdapterFeedAdded struct {
	Symbol      string
	FeedAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeedAdded is a free log retrieval operation binding the contract event 0x917a6ab4f9d1b23331c7307351d940aa91024649e9745f66c8d7d9ccdb5f3704.
//
// Solidity: event FeedAdded(string symbol, address feedAddress)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterFilterer) FilterFeedAdded(opts *bind.FilterOpts) (*ContractPriceFeedAdapterFeedAddedIterator, error) {

	logs, sub, err := _ContractPriceFeedAdapter.contract.FilterLogs(opts, "FeedAdded")
	if err != nil {
		return nil, err
	}
	return &ContractPriceFeedAdapterFeedAddedIterator{contract: _ContractPriceFeedAdapter.contract, event: "FeedAdded", logs: logs, sub: sub}, nil
}

// WatchFeedAdded is a free log subscription operation binding the contract event 0x917a6ab4f9d1b23331c7307351d940aa91024649e9745f66c8d7d9ccdb5f3704.
//
// Solidity: event FeedAdded(string symbol, address feedAddress)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterFilterer) WatchFeedAdded(opts *bind.WatchOpts, sink chan<- *ContractPriceFeedAdapterFeedAdded) (event.Subscription, error) {

	logs, sub, err := _ContractPriceFeedAdapter.contract.WatchLogs(opts, "FeedAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPriceFeedAdapterFeedAdded)
				if err := _ContractPriceFeedAdapter.contract.UnpackLog(event, "FeedAdded", log); err != nil {
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

// ParseFeedAdded is a log parse operation binding the contract event 0x917a6ab4f9d1b23331c7307351d940aa91024649e9745f66c8d7d9ccdb5f3704.
//
// Solidity: event FeedAdded(string symbol, address feedAddress)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterFilterer) ParseFeedAdded(log types.Log) (*ContractPriceFeedAdapterFeedAdded, error) {
	event := new(ContractPriceFeedAdapterFeedAdded)
	if err := _ContractPriceFeedAdapter.contract.UnpackLog(event, "FeedAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPriceFeedAdapterFeedRemovedIterator is returned from FilterFeedRemoved and is used to iterate over the raw logs and unpacked data for FeedRemoved events raised by the ContractPriceFeedAdapter contract.
type ContractPriceFeedAdapterFeedRemovedIterator struct {
	Event *ContractPriceFeedAdapterFeedRemoved // Event containing the contract specifics and raw log

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
func (it *ContractPriceFeedAdapterFeedRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPriceFeedAdapterFeedRemoved)
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
		it.Event = new(ContractPriceFeedAdapterFeedRemoved)
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
func (it *ContractPriceFeedAdapterFeedRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPriceFeedAdapterFeedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPriceFeedAdapterFeedRemoved represents a FeedRemoved event raised by the ContractPriceFeedAdapter contract.
type ContractPriceFeedAdapterFeedRemoved struct {
	Symbol string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeedRemoved is a free log retrieval operation binding the contract event 0x06aa2b76883f4d9b405ce5d8cb618072ed06ad9a9f4a3a638d6a76f995f950ce.
//
// Solidity: event FeedRemoved(string symbol)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterFilterer) FilterFeedRemoved(opts *bind.FilterOpts) (*ContractPriceFeedAdapterFeedRemovedIterator, error) {

	logs, sub, err := _ContractPriceFeedAdapter.contract.FilterLogs(opts, "FeedRemoved")
	if err != nil {
		return nil, err
	}
	return &ContractPriceFeedAdapterFeedRemovedIterator{contract: _ContractPriceFeedAdapter.contract, event: "FeedRemoved", logs: logs, sub: sub}, nil
}

// WatchFeedRemoved is a free log subscription operation binding the contract event 0x06aa2b76883f4d9b405ce5d8cb618072ed06ad9a9f4a3a638d6a76f995f950ce.
//
// Solidity: event FeedRemoved(string symbol)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterFilterer) WatchFeedRemoved(opts *bind.WatchOpts, sink chan<- *ContractPriceFeedAdapterFeedRemoved) (event.Subscription, error) {

	logs, sub, err := _ContractPriceFeedAdapter.contract.WatchLogs(opts, "FeedRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPriceFeedAdapterFeedRemoved)
				if err := _ContractPriceFeedAdapter.contract.UnpackLog(event, "FeedRemoved", log); err != nil {
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

// ParseFeedRemoved is a log parse operation binding the contract event 0x06aa2b76883f4d9b405ce5d8cb618072ed06ad9a9f4a3a638d6a76f995f950ce.
//
// Solidity: event FeedRemoved(string symbol)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterFilterer) ParseFeedRemoved(log types.Log) (*ContractPriceFeedAdapterFeedRemoved, error) {
	event := new(ContractPriceFeedAdapterFeedRemoved)
	if err := _ContractPriceFeedAdapter.contract.UnpackLog(event, "FeedRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPriceFeedAdapterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractPriceFeedAdapter contract.
type ContractPriceFeedAdapterOwnershipTransferredIterator struct {
	Event *ContractPriceFeedAdapterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractPriceFeedAdapterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPriceFeedAdapterOwnershipTransferred)
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
		it.Event = new(ContractPriceFeedAdapterOwnershipTransferred)
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
func (it *ContractPriceFeedAdapterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPriceFeedAdapterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPriceFeedAdapterOwnershipTransferred represents a OwnershipTransferred event raised by the ContractPriceFeedAdapter contract.
type ContractPriceFeedAdapterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractPriceFeedAdapterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractPriceFeedAdapter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractPriceFeedAdapterOwnershipTransferredIterator{contract: _ContractPriceFeedAdapter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractPriceFeedAdapterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractPriceFeedAdapter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPriceFeedAdapterOwnershipTransferred)
				if err := _ContractPriceFeedAdapter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterFilterer) ParseOwnershipTransferred(log types.Log) (*ContractPriceFeedAdapterOwnershipTransferred, error) {
	event := new(ContractPriceFeedAdapterOwnershipTransferred)
	if err := _ContractPriceFeedAdapter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPriceFeedAdapterChainlinkAdapterPriceIterator is returned from FilterChainlinkAdapterPrice and is used to iterate over the raw logs and unpacked data for ChainlinkAdapterPrice events raised by the ContractPriceFeedAdapter contract.
type ContractPriceFeedAdapterChainlinkAdapterPriceIterator struct {
	Event *ContractPriceFeedAdapterChainlinkAdapterPrice // Event containing the contract specifics and raw log

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
func (it *ContractPriceFeedAdapterChainlinkAdapterPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPriceFeedAdapterChainlinkAdapterPrice)
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
		it.Event = new(ContractPriceFeedAdapterChainlinkAdapterPrice)
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
func (it *ContractPriceFeedAdapterChainlinkAdapterPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPriceFeedAdapterChainlinkAdapterPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPriceFeedAdapterChainlinkAdapterPrice represents a ChainlinkAdapterPrice event raised by the ContractPriceFeedAdapter contract.
type ContractPriceFeedAdapterChainlinkAdapterPrice struct {
	Symbol                string
	ChainlinkAdapterPrice *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterChainlinkAdapterPrice is a free log retrieval operation binding the contract event 0x35e7728e5ba1465def842af9ec022a54cdbeb36fe61d509976bd4de9142cbcbb.
//
// Solidity: event chainlinkAdapterPrice(string symbol, int256 _ChainlinkAdapterPrice)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterFilterer) FilterChainlinkAdapterPrice(opts *bind.FilterOpts) (*ContractPriceFeedAdapterChainlinkAdapterPriceIterator, error) {

	logs, sub, err := _ContractPriceFeedAdapter.contract.FilterLogs(opts, "chainlinkAdapterPrice")
	if err != nil {
		return nil, err
	}
	return &ContractPriceFeedAdapterChainlinkAdapterPriceIterator{contract: _ContractPriceFeedAdapter.contract, event: "chainlinkAdapterPrice", logs: logs, sub: sub}, nil
}

// WatchChainlinkAdapterPrice is a free log subscription operation binding the contract event 0x35e7728e5ba1465def842af9ec022a54cdbeb36fe61d509976bd4de9142cbcbb.
//
// Solidity: event chainlinkAdapterPrice(string symbol, int256 _ChainlinkAdapterPrice)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterFilterer) WatchChainlinkAdapterPrice(opts *bind.WatchOpts, sink chan<- *ContractPriceFeedAdapterChainlinkAdapterPrice) (event.Subscription, error) {

	logs, sub, err := _ContractPriceFeedAdapter.contract.WatchLogs(opts, "chainlinkAdapterPrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPriceFeedAdapterChainlinkAdapterPrice)
				if err := _ContractPriceFeedAdapter.contract.UnpackLog(event, "chainlinkAdapterPrice", log); err != nil {
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

// ParseChainlinkAdapterPrice is a log parse operation binding the contract event 0x35e7728e5ba1465def842af9ec022a54cdbeb36fe61d509976bd4de9142cbcbb.
//
// Solidity: event chainlinkAdapterPrice(string symbol, int256 _ChainlinkAdapterPrice)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterFilterer) ParseChainlinkAdapterPrice(log types.Log) (*ContractPriceFeedAdapterChainlinkAdapterPrice, error) {
	event := new(ContractPriceFeedAdapterChainlinkAdapterPrice)
	if err := _ContractPriceFeedAdapter.contract.UnpackLog(event, "chainlinkAdapterPrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPriceFeedAdapterDiaAdapterPriceIterator is returned from FilterDiaAdapterPrice and is used to iterate over the raw logs and unpacked data for DiaAdapterPrice events raised by the ContractPriceFeedAdapter contract.
type ContractPriceFeedAdapterDiaAdapterPriceIterator struct {
	Event *ContractPriceFeedAdapterDiaAdapterPrice // Event containing the contract specifics and raw log

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
func (it *ContractPriceFeedAdapterDiaAdapterPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPriceFeedAdapterDiaAdapterPrice)
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
		it.Event = new(ContractPriceFeedAdapterDiaAdapterPrice)
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
func (it *ContractPriceFeedAdapterDiaAdapterPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPriceFeedAdapterDiaAdapterPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPriceFeedAdapterDiaAdapterPrice represents a DiaAdapterPrice event raised by the ContractPriceFeedAdapter contract.
type ContractPriceFeedAdapterDiaAdapterPrice struct {
	Symbol          string
	DiaAdapterPrice *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDiaAdapterPrice is a free log retrieval operation binding the contract event 0x7f48874caeb1738a3f8498b97e6be3ba0ac0c2f79aeb8751d78df2a3a91f7a66.
//
// Solidity: event diaAdapterPrice(string symbol, uint128 _DiaAdapterPrice)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterFilterer) FilterDiaAdapterPrice(opts *bind.FilterOpts) (*ContractPriceFeedAdapterDiaAdapterPriceIterator, error) {

	logs, sub, err := _ContractPriceFeedAdapter.contract.FilterLogs(opts, "diaAdapterPrice")
	if err != nil {
		return nil, err
	}
	return &ContractPriceFeedAdapterDiaAdapterPriceIterator{contract: _ContractPriceFeedAdapter.contract, event: "diaAdapterPrice", logs: logs, sub: sub}, nil
}

// WatchDiaAdapterPrice is a free log subscription operation binding the contract event 0x7f48874caeb1738a3f8498b97e6be3ba0ac0c2f79aeb8751d78df2a3a91f7a66.
//
// Solidity: event diaAdapterPrice(string symbol, uint128 _DiaAdapterPrice)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterFilterer) WatchDiaAdapterPrice(opts *bind.WatchOpts, sink chan<- *ContractPriceFeedAdapterDiaAdapterPrice) (event.Subscription, error) {

	logs, sub, err := _ContractPriceFeedAdapter.contract.WatchLogs(opts, "diaAdapterPrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPriceFeedAdapterDiaAdapterPrice)
				if err := _ContractPriceFeedAdapter.contract.UnpackLog(event, "diaAdapterPrice", log); err != nil {
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

// ParseDiaAdapterPrice is a log parse operation binding the contract event 0x7f48874caeb1738a3f8498b97e6be3ba0ac0c2f79aeb8751d78df2a3a91f7a66.
//
// Solidity: event diaAdapterPrice(string symbol, uint128 _DiaAdapterPrice)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterFilterer) ParseDiaAdapterPrice(log types.Log) (*ContractPriceFeedAdapterDiaAdapterPrice, error) {
	event := new(ContractPriceFeedAdapterDiaAdapterPrice)
	if err := _ContractPriceFeedAdapter.contract.UnpackLog(event, "diaAdapterPrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPriceFeedAdapterDiaOracleAddressIterator is returned from FilterDiaOracleAddress and is used to iterate over the raw logs and unpacked data for DiaOracleAddress events raised by the ContractPriceFeedAdapter contract.
type ContractPriceFeedAdapterDiaOracleAddressIterator struct {
	Event *ContractPriceFeedAdapterDiaOracleAddress // Event containing the contract specifics and raw log

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
func (it *ContractPriceFeedAdapterDiaOracleAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPriceFeedAdapterDiaOracleAddress)
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
		it.Event = new(ContractPriceFeedAdapterDiaOracleAddress)
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
func (it *ContractPriceFeedAdapterDiaOracleAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPriceFeedAdapterDiaOracleAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPriceFeedAdapterDiaOracleAddress represents a DiaOracleAddress event raised by the ContractPriceFeedAdapter contract.
type ContractPriceFeedAdapterDiaOracleAddress struct {
	Symbol           string
	DiaOracleAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDiaOracleAddress is a free log retrieval operation binding the contract event 0x34fa09dd73d22ca570e8b4a45ec050b8a0cbe4a14b26df98ddeef0016e87fc17.
//
// Solidity: event diaOracleAddress(string symbol, address _DiaOracleAddress)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterFilterer) FilterDiaOracleAddress(opts *bind.FilterOpts) (*ContractPriceFeedAdapterDiaOracleAddressIterator, error) {

	logs, sub, err := _ContractPriceFeedAdapter.contract.FilterLogs(opts, "diaOracleAddress")
	if err != nil {
		return nil, err
	}
	return &ContractPriceFeedAdapterDiaOracleAddressIterator{contract: _ContractPriceFeedAdapter.contract, event: "diaOracleAddress", logs: logs, sub: sub}, nil
}

// WatchDiaOracleAddress is a free log subscription operation binding the contract event 0x34fa09dd73d22ca570e8b4a45ec050b8a0cbe4a14b26df98ddeef0016e87fc17.
//
// Solidity: event diaOracleAddress(string symbol, address _DiaOracleAddress)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterFilterer) WatchDiaOracleAddress(opts *bind.WatchOpts, sink chan<- *ContractPriceFeedAdapterDiaOracleAddress) (event.Subscription, error) {

	logs, sub, err := _ContractPriceFeedAdapter.contract.WatchLogs(opts, "diaOracleAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPriceFeedAdapterDiaOracleAddress)
				if err := _ContractPriceFeedAdapter.contract.UnpackLog(event, "diaOracleAddress", log); err != nil {
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

// ParseDiaOracleAddress is a log parse operation binding the contract event 0x34fa09dd73d22ca570e8b4a45ec050b8a0cbe4a14b26df98ddeef0016e87fc17.
//
// Solidity: event diaOracleAddress(string symbol, address _DiaOracleAddress)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterFilterer) ParseDiaOracleAddress(log types.Log) (*ContractPriceFeedAdapterDiaOracleAddress, error) {
	event := new(ContractPriceFeedAdapterDiaOracleAddress)
	if err := _ContractPriceFeedAdapter.contract.UnpackLog(event, "diaOracleAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
