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

// ECVerifyABI is the input ABI used to generate the binding from.
const ECVerifyABI = "[]"

// ECVerifyBin is the compiled bytecode used for deploying new contracts.
const ECVerifyBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058208377d8353eea6842efdc9fb1868e4fa1cf41cdd5d066638b36f4e3cdc2919d2d0029`

// DeployECVerify deploys a new Ethereum contract, binding an instance of ECVerify to it.
func DeployECVerify(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECVerify, error) {
	parsed, err := abi.JSON(strings.NewReader(ECVerifyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ECVerifyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECVerify{ECVerifyCaller: ECVerifyCaller{contract: contract}, ECVerifyTransactor: ECVerifyTransactor{contract: contract}, ECVerifyFilterer: ECVerifyFilterer{contract: contract}}, nil
}

// ECVerify is an auto generated Go binding around an Ethereum contract.
type ECVerify struct {
	ECVerifyCaller     // Read-only binding to the contract
	ECVerifyTransactor // Write-only binding to the contract
	ECVerifyFilterer   // Log filterer for contract events
}

// ECVerifyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ECVerifyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECVerifyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECVerifyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECVerifyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECVerifyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECVerifySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECVerifySession struct {
	Contract     *ECVerify         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECVerifyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECVerifyCallerSession struct {
	Contract *ECVerifyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ECVerifyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECVerifyTransactorSession struct {
	Contract     *ECVerifyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ECVerifyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ECVerifyRaw struct {
	Contract *ECVerify // Generic contract binding to access the raw methods on
}

// ECVerifyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECVerifyCallerRaw struct {
	Contract *ECVerifyCaller // Generic read-only contract binding to access the raw methods on
}

// ECVerifyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECVerifyTransactorRaw struct {
	Contract *ECVerifyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewECVerify creates a new instance of ECVerify, bound to a specific deployed contract.
func NewECVerify(address common.Address, backend bind.ContractBackend) (*ECVerify, error) {
	contract, err := bindECVerify(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECVerify{ECVerifyCaller: ECVerifyCaller{contract: contract}, ECVerifyTransactor: ECVerifyTransactor{contract: contract}, ECVerifyFilterer: ECVerifyFilterer{contract: contract}}, nil
}

// NewECVerifyCaller creates a new read-only instance of ECVerify, bound to a specific deployed contract.
func NewECVerifyCaller(address common.Address, caller bind.ContractCaller) (*ECVerifyCaller, error) {
	contract, err := bindECVerify(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECVerifyCaller{contract: contract}, nil
}

// NewECVerifyTransactor creates a new write-only instance of ECVerify, bound to a specific deployed contract.
func NewECVerifyTransactor(address common.Address, transactor bind.ContractTransactor) (*ECVerifyTransactor, error) {
	contract, err := bindECVerify(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECVerifyTransactor{contract: contract}, nil
}

// NewECVerifyFilterer creates a new log filterer instance of ECVerify, bound to a specific deployed contract.
func NewECVerifyFilterer(address common.Address, filterer bind.ContractFilterer) (*ECVerifyFilterer, error) {
	contract, err := bindECVerify(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECVerifyFilterer{contract: contract}, nil
}

// bindECVerify binds a generic wrapper to an already deployed contract.
func bindECVerify(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECVerifyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECVerify *ECVerifyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECVerify.Contract.ECVerifyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECVerify *ECVerifyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECVerify.Contract.ECVerifyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECVerify *ECVerifyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECVerify.Contract.ECVerifyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECVerify *ECVerifyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECVerify.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECVerify *ECVerifyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECVerify.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECVerify *ECVerifyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECVerify.Contract.contract.Transact(opts, method, params...)
}

// SecretRegistryABI is the input ABI used to generate the binding from.
const SecretRegistryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"secret\",\"type\":\"bytes32\"}],\"name\":\"registerSecret\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"secrethash_to_block\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contract_version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"secrethash\",\"type\":\"bytes32\"}],\"name\":\"getSecretRevealBlockHeight\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"secrethash\",\"type\":\"bytes32\"}],\"name\":\"SecretRevealed\",\"type\":\"event\"}]"

// SecretRegistryBin is the compiled bytecode used for deploying new contracts.
const SecretRegistryBin = `0x608060405234801561001057600080fd5b506102cd806100206000396000f3006080604052600436106100615763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166312ad8bfc81146100665780639734030914610092578063b32c65c8146100bc578063c1f6294614610146575b600080fd5b34801561007257600080fd5b5061007e60043561015e565b604080519115158252519081900360200190f35b34801561009e57600080fd5b506100aa600435610246565b60408051918252519081900360200190f35b3480156100c857600080fd5b506100d1610258565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561010b5781810151838201526020016100f3565b50505050905090810190601f1680156101385780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561015257600080fd5b506100aa60043561028f565b604080516020808201849052825180830382018152918301928390528151600093849392909182918401908083835b602083106101ac5780518252601f19909201916020918201910161018d565b5181516020939093036101000a60001901801990911692169190911790526040519201829003909120935050841591508190506101f55750600081815260208190526040812054115b156102035760009150610240565b6000818152602081905260408082204390555182917f9b7ddc883342824bd7ddbff103e7a69f8f2e60b96c075cd1b8b8b9713ecc75a491a2600191505b50919050565b60006020819052908152604090205481565b60408051808201909152600581527f302e332e5f000000000000000000000000000000000000000000000000000000602082015281565b600090815260208190526040902054905600a165627a7a72305820f69c8bac516d7b33d97cfff0956fa761274de6891a330e80786dd2e9e978b5e10029`

// DeploySecretRegistry deploys a new Ethereum contract, binding an instance of SecretRegistry to it.
func DeploySecretRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SecretRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(SecretRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SecretRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SecretRegistry{SecretRegistryCaller: SecretRegistryCaller{contract: contract}, SecretRegistryTransactor: SecretRegistryTransactor{contract: contract}, SecretRegistryFilterer: SecretRegistryFilterer{contract: contract}}, nil
}

// SecretRegistry is an auto generated Go binding around an Ethereum contract.
type SecretRegistry struct {
	SecretRegistryCaller     // Read-only binding to the contract
	SecretRegistryTransactor // Write-only binding to the contract
	SecretRegistryFilterer   // Log filterer for contract events
}

// SecretRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SecretRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SecretRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SecretRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SecretRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SecretRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SecretRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SecretRegistrySession struct {
	Contract     *SecretRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SecretRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SecretRegistryCallerSession struct {
	Contract *SecretRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SecretRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SecretRegistryTransactorSession struct {
	Contract     *SecretRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SecretRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SecretRegistryRaw struct {
	Contract *SecretRegistry // Generic contract binding to access the raw methods on
}

// SecretRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SecretRegistryCallerRaw struct {
	Contract *SecretRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// SecretRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SecretRegistryTransactorRaw struct {
	Contract *SecretRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSecretRegistry creates a new instance of SecretRegistry, bound to a specific deployed contract.
func NewSecretRegistry(address common.Address, backend bind.ContractBackend) (*SecretRegistry, error) {
	contract, err := bindSecretRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SecretRegistry{SecretRegistryCaller: SecretRegistryCaller{contract: contract}, SecretRegistryTransactor: SecretRegistryTransactor{contract: contract}, SecretRegistryFilterer: SecretRegistryFilterer{contract: contract}}, nil
}

// NewSecretRegistryCaller creates a new read-only instance of SecretRegistry, bound to a specific deployed contract.
func NewSecretRegistryCaller(address common.Address, caller bind.ContractCaller) (*SecretRegistryCaller, error) {
	contract, err := bindSecretRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SecretRegistryCaller{contract: contract}, nil
}

// NewSecretRegistryTransactor creates a new write-only instance of SecretRegistry, bound to a specific deployed contract.
func NewSecretRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*SecretRegistryTransactor, error) {
	contract, err := bindSecretRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SecretRegistryTransactor{contract: contract}, nil
}

// NewSecretRegistryFilterer creates a new log filterer instance of SecretRegistry, bound to a specific deployed contract.
func NewSecretRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*SecretRegistryFilterer, error) {
	contract, err := bindSecretRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SecretRegistryFilterer{contract: contract}, nil
}

// bindSecretRegistry binds a generic wrapper to an already deployed contract.
func bindSecretRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SecretRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SecretRegistry *SecretRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SecretRegistry.Contract.SecretRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SecretRegistry *SecretRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SecretRegistry.Contract.SecretRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SecretRegistry *SecretRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SecretRegistry.Contract.SecretRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SecretRegistry *SecretRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SecretRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SecretRegistry *SecretRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SecretRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SecretRegistry *SecretRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SecretRegistry.Contract.contract.Transact(opts, method, params...)
}

// Contract_version is a free data retrieval call binding the contract method 0xb32c65c8.
//
// Solidity: function contract_version() constant returns(string)
func (_SecretRegistry *SecretRegistryCaller) Contract_version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SecretRegistry.contract.Call(opts, out, "contract_version")
	return *ret0, err
}

// Contract_version is a free data retrieval call binding the contract method 0xb32c65c8.
//
// Solidity: function contract_version() constant returns(string)
func (_SecretRegistry *SecretRegistrySession) Contract_version() (string, error) {
	return _SecretRegistry.Contract.Contract_version(&_SecretRegistry.CallOpts)
}

// Contract_version is a free data retrieval call binding the contract method 0xb32c65c8.
//
// Solidity: function contract_version() constant returns(string)
func (_SecretRegistry *SecretRegistryCallerSession) Contract_version() (string, error) {
	return _SecretRegistry.Contract.Contract_version(&_SecretRegistry.CallOpts)
}

// GetSecretRevealBlockHeight is a free data retrieval call binding the contract method 0xc1f62946.
//
// Solidity: function getSecretRevealBlockHeight(secrethash bytes32) constant returns(uint256)
func (_SecretRegistry *SecretRegistryCaller) GetSecretRevealBlockHeight(opts *bind.CallOpts, secrethash [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SecretRegistry.contract.Call(opts, out, "getSecretRevealBlockHeight", secrethash)
	return *ret0, err
}

// GetSecretRevealBlockHeight is a free data retrieval call binding the contract method 0xc1f62946.
//
// Solidity: function getSecretRevealBlockHeight(secrethash bytes32) constant returns(uint256)
func (_SecretRegistry *SecretRegistrySession) GetSecretRevealBlockHeight(secrethash [32]byte) (*big.Int, error) {
	return _SecretRegistry.Contract.GetSecretRevealBlockHeight(&_SecretRegistry.CallOpts, secrethash)
}

// GetSecretRevealBlockHeight is a free data retrieval call binding the contract method 0xc1f62946.
//
// Solidity: function getSecretRevealBlockHeight(secrethash bytes32) constant returns(uint256)
func (_SecretRegistry *SecretRegistryCallerSession) GetSecretRevealBlockHeight(secrethash [32]byte) (*big.Int, error) {
	return _SecretRegistry.Contract.GetSecretRevealBlockHeight(&_SecretRegistry.CallOpts, secrethash)
}

// Secrethash_to_block is a free data retrieval call binding the contract method 0x97340309.
//
// Solidity: function secrethash_to_block( bytes32) constant returns(uint256)
func (_SecretRegistry *SecretRegistryCaller) Secrethash_to_block(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SecretRegistry.contract.Call(opts, out, "secrethash_to_block", arg0)
	return *ret0, err
}

// Secrethash_to_block is a free data retrieval call binding the contract method 0x97340309.
//
// Solidity: function secrethash_to_block( bytes32) constant returns(uint256)
func (_SecretRegistry *SecretRegistrySession) Secrethash_to_block(arg0 [32]byte) (*big.Int, error) {
	return _SecretRegistry.Contract.Secrethash_to_block(&_SecretRegistry.CallOpts, arg0)
}

// Secrethash_to_block is a free data retrieval call binding the contract method 0x97340309.
//
// Solidity: function secrethash_to_block( bytes32) constant returns(uint256)
func (_SecretRegistry *SecretRegistryCallerSession) Secrethash_to_block(arg0 [32]byte) (*big.Int, error) {
	return _SecretRegistry.Contract.Secrethash_to_block(&_SecretRegistry.CallOpts, arg0)
}

// RegisterSecret is a paid mutator transaction binding the contract method 0x12ad8bfc.
//
// Solidity: function registerSecret(secret bytes32) returns(bool)
func (_SecretRegistry *SecretRegistryTransactor) RegisterSecret(opts *bind.TransactOpts, secret [32]byte) (*types.Transaction, error) {
	return _SecretRegistry.contract.Transact(opts, "registerSecret", secret)
}

// RegisterSecret is a paid mutator transaction binding the contract method 0x12ad8bfc.
//
// Solidity: function registerSecret(secret bytes32) returns(bool)
func (_SecretRegistry *SecretRegistrySession) RegisterSecret(secret [32]byte) (*types.Transaction, error) {
	return _SecretRegistry.Contract.RegisterSecret(&_SecretRegistry.TransactOpts, secret)
}

// RegisterSecret is a paid mutator transaction binding the contract method 0x12ad8bfc.
//
// Solidity: function registerSecret(secret bytes32) returns(bool)
func (_SecretRegistry *SecretRegistryTransactorSession) RegisterSecret(secret [32]byte) (*types.Transaction, error) {
	return _SecretRegistry.Contract.RegisterSecret(&_SecretRegistry.TransactOpts, secret)
}

// SecretRegistrySecretRevealedIterator is returned from FilterSecretRevealed and is used to iterate over the raw logs and unpacked data for SecretRevealed events raised by the SecretRegistry contract.
type SecretRegistrySecretRevealedIterator struct {
	Event *SecretRegistrySecretRevealed // Event containing the contract specifics and raw log

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
func (it *SecretRegistrySecretRevealedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecretRegistrySecretRevealed)
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
		it.Event = new(SecretRegistrySecretRevealed)
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
func (it *SecretRegistrySecretRevealedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SecretRegistrySecretRevealedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SecretRegistrySecretRevealed represents a SecretRevealed event raised by the SecretRegistry contract.
type SecretRegistrySecretRevealed struct {
	Secrethash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSecretRevealed is a free log retrieval operation binding the contract event 0x9b7ddc883342824bd7ddbff103e7a69f8f2e60b96c075cd1b8b8b9713ecc75a4.
//
// Solidity: event SecretRevealed(secrethash indexed bytes32)
func (_SecretRegistry *SecretRegistryFilterer) FilterSecretRevealed(opts *bind.FilterOpts, secrethash [][32]byte) (*SecretRegistrySecretRevealedIterator, error) {

	var secrethashRule []interface{}
	for _, secrethashItem := range secrethash {
		secrethashRule = append(secrethashRule, secrethashItem)
	}

	logs, sub, err := _SecretRegistry.contract.FilterLogs(opts, "SecretRevealed", secrethashRule)
	if err != nil {
		return nil, err
	}
	return &SecretRegistrySecretRevealedIterator{contract: _SecretRegistry.contract, event: "SecretRevealed", logs: logs, sub: sub}, nil
}

// WatchSecretRevealed is a free log subscription operation binding the contract event 0x9b7ddc883342824bd7ddbff103e7a69f8f2e60b96c075cd1b8b8b9713ecc75a4.
//
// Solidity: event SecretRevealed(secrethash indexed bytes32)
func (_SecretRegistry *SecretRegistryFilterer) WatchSecretRevealed(opts *bind.WatchOpts, sink chan<- *SecretRegistrySecretRevealed, secrethash [][32]byte) (event.Subscription, error) {

	var secrethashRule []interface{}
	for _, secrethashItem := range secrethash {
		secrethashRule = append(secrethashRule, secrethashItem)
	}

	logs, sub, err := _SecretRegistry.contract.WatchLogs(opts, "SecretRevealed", secrethashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SecretRegistrySecretRevealed)
				if err := _SecretRegistry.contract.UnpackLog(event, "SecretRevealed", log); err != nil {
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

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"supply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `0x`

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Token *TokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _value)
}

// TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Token contract.
type TokenApprovalIterator struct {
	Event *TokenApproval // Event containing the contract specifics and raw log

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
func (it *TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenApproval)
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
		it.Event = new(TokenApproval)
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
func (it *TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenApproval represents a Approval event raised by the Token contract.
type TokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_Token *TokenFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*TokenApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenApprovalIterator{contract: _Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_Token *TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenApproval)
				if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
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

// TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Token contract.
type TokenTransferIterator struct {
	Event *TokenTransfer // Event containing the contract specifics and raw log

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
func (it *TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransfer)
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
		it.Event = new(TokenTransfer)
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
func (it *TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransfer represents a Transfer event raised by the Token contract.
type TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(_from indexed address, _to indexed address, _value uint256)
func (_Token *TokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*TokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransferIterator{contract: _Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(_from indexed address, _to indexed address, _value uint256)
func (_Token *TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransfer)
				if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// TokenNetworkABI is the input ABI used to generate the binding from.
const TokenNetworkABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"participant1\",\"type\":\"address\"},{\"name\":\"participant2\",\"type\":\"address\"},{\"name\":\"settle_timeout\",\"type\":\"uint256\"}],\"name\":\"openChannel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"last_channel_index\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"openedchannels\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"secret_registry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"channel_identifier\",\"type\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\"},{\"name\":\"lockhash\",\"type\":\"bytes32\"},{\"name\":\"additional_hash\",\"type\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\"},{\"name\":\"merkle_proof\",\"type\":\"bytes\"}],\"name\":\"punishObsoleteUnlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"participant1\",\"type\":\"address\"},{\"name\":\"participant2\",\"type\":\"address\"},{\"name\":\"participant1_deposit\",\"type\":\"uint256\"},{\"name\":\"participant2_deposit\",\"type\":\"uint256\"},{\"name\":\"participant1_withdraw\",\"type\":\"uint256\"},{\"name\":\"participant2_withdraw\",\"type\":\"uint256\"},{\"name\":\"participant1_signature\",\"type\":\"bytes\"},{\"name\":\"participant2_signature\",\"type\":\"bytes\"},{\"name\":\"channel_identifier\",\"type\":\"uint256\"}],\"name\":\"withDraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chain_id\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"contract_address\",\"type\":\"address\"}],\"name\":\"contractExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"channel_identifier\",\"type\":\"uint256\"},{\"name\":\"participant\",\"type\":\"address\"},{\"name\":\"total_deposit\",\"type\":\"uint256\"}],\"name\":\"setTotalDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"channel_identifier\",\"type\":\"uint256\"},{\"name\":\"transferred_amount\",\"type\":\"uint256\"},{\"name\":\"locksroot\",\"type\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"additional_hash\",\"type\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"closeChannel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"channel_identifier\",\"type\":\"uint256\"},{\"name\":\"participant\",\"type\":\"address\"}],\"name\":\"getChannelParticipantInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"participant1_address\",\"type\":\"address\"},{\"name\":\"participant1_balance\",\"type\":\"uint256\"},{\"name\":\"participant2_address\",\"type\":\"address\"},{\"name\":\"participant2_balance\",\"type\":\"uint256\"},{\"name\":\"participant1_signature\",\"type\":\"bytes\"},{\"name\":\"participant2_signature\",\"type\":\"bytes\"}],\"name\":\"cooperativeSettle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"channel_identifier\",\"type\":\"uint256\"},{\"name\":\"non_closing_participant\",\"type\":\"address\"},{\"name\":\"locksroot\",\"type\":\"bytes32\"},{\"name\":\"transferred_amount\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"additional_hash\",\"type\":\"bytes32\"},{\"name\":\"closing_signature\",\"type\":\"bytes\"},{\"name\":\"non_closing_signature\",\"type\":\"bytes\"}],\"name\":\"updateNonClosingBalanceProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contract_version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"participant\",\"type\":\"address\"},{\"name\":\"partner\",\"type\":\"address\"}],\"name\":\"getChannelHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"participant1\",\"type\":\"address\"},{\"name\":\"participant2\",\"type\":\"address\"}],\"name\":\"settleChannel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"channels\",\"outputs\":[{\"name\":\"settle_block_number\",\"type\":\"uint256\"},{\"name\":\"closing_participant\",\"type\":\"address\"},{\"name\":\"state\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"channel_identifier\",\"type\":\"uint256\"},{\"name\":\"participant\",\"type\":\"address\"},{\"name\":\"merkle_tree_leaves\",\"type\":\"bytes\"}],\"name\":\"unlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"participant1\",\"type\":\"address\"},{\"name\":\"participant2\",\"type\":\"address\"}],\"name\":\"getChannelInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token_address\",\"type\":\"address\"},{\"name\":\"_secret_registry\",\"type\":\"address\"},{\"name\":\"_chain_id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"channel_identifier\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"participant1\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"participant2\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"settle_timeout\",\"type\":\"uint256\"}],\"name\":\"ChannelOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"channel_identifier\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"total_deposit\",\"type\":\"uint256\"}],\"name\":\"ChannelNewDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"channel_identifier\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"closing_participant\",\"type\":\"address\"}],\"name\":\"ChannelClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"channel_identifier\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"payer_participant\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"locskroot\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"transferred_amount\",\"type\":\"uint256\"}],\"name\":\"ChannelUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"channel_identifier\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"closing_participant\",\"type\":\"address\"}],\"name\":\"NonClosingBalanceProofUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"channel_identifier\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"participant1_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"participant2_amount\",\"type\":\"uint256\"}],\"name\":\"ChannelSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"channel_identifier\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"participant1_deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"participant2_deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"participant1_withdraw\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"participant2_withdraw\",\"type\":\"uint256\"}],\"name\":\"Channelwithdraw\",\"type\":\"event\"}]"

// TokenNetworkBin is the compiled bytecode used for deploying new contracts.
const TokenNetworkBin = `0x608060405260006005553480156200001657600080fd5b5060405160608062002d3f833981016040908152815160208301519190920151600160a060020a03831615156200004c57600080fd5b600160a060020a03821615156200006257600080fd5b600081116200007057600080fd5b62000084836401000000006200017c810204565b15156200009057600080fd5b620000a4826401000000006200017c810204565b1515620000b057600080fd5b60008054600160a060020a03808616600160a060020a031992831617808455600180548784169416939093179092556002849055604080517f18160ddd000000000000000000000000000000000000000000000000000000008152905192909116916318160ddd9160048082019260209290919082900301818787803b1580156200013a57600080fd5b505af11580156200014f573d6000803e3d6000fd5b505050506040513d60208110156200016657600080fd5b5051116200017357600080fd5b50505062000184565b6000903b1190565b612bab80620001946000396000f3006080604052600436106100f85763ffffffff60e060020a6000350416630a798f2481146100fd5780631f466acf146101295780632419a2e01461015057806324d73a931461016857806328a8d70f1461019957806335fefe7d1461024b5780633af973b11461030d5780637709bc78146103225780637944bd1f146103575780637c0f097c1461037e5780637fb5885e146103ea5780638568536a1461043b5780638d373f51146104f2578063b32c65c8146105ae578063bb94ccc814610638578063dee8bd0a1461065f578063e5949b5d14610686578063f3da17e8146106c7578063f94c9e1314610730578063fc0c546a14610788575b600080fd5b34801561010957600080fd5b50610127600160a060020a036004358116906024351660443561079d565b005b34801561013557600080fd5b5061013e61096f565b60408051918252519081900360200190f35b34801561015c57600080fd5b5061013e600435610975565b34801561017457600080fd5b5061017d610987565b60408051600160a060020a039092168252519081900360200190f35b3480156101a557600080fd5b50604080516020600460843581810135601f81018490048402850184019095528484526101279482359460248035600160a060020a03169560443595606435953695919460a494909391019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506109969650505050505050565b34801561025757600080fd5b50604080516020601f60c43560048181013592830184900484028501840190955281845261012794600160a060020a038135811695602480359092169560443595606435956084359560a435953695919460e49492939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497505093359450610aa09350505050565b34801561031957600080fd5b5061013e611172565b34801561032e57600080fd5b50610343600160a060020a0360043516611178565b604080519115158252519081900360200190f35b34801561036357600080fd5b50610127600435600160a060020a0360243516604435611180565b34801561038a57600080fd5b50604080516020600460a43581810135601f81018490048402850184019095528484526101279482359460248035956044359560643595608435953695929460c494920191819084018382808284375094975061130b9650505050505050565b3480156103f657600080fd5b5061040e600435600160a060020a0360243516611452565b60408051951515865260208601949094528484019290925260608401526080830152519081900360a00190f35b34801561044757600080fd5b50604080516020601f60843560048181013592830184900484028501840190955281845261012794600160a060020a0381358116956024803596604435909316956064359536959460a49493919091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506114999650505050505050565b3480156104fe57600080fd5b50604080516020600460c43581810135601f81018490048402850184019095528484526101279482359460248035600160a060020a03169560443595606435956084359560a435953695939460e4949193919091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506118a59650505050505050565b3480156105ba57600080fd5b506105c3611996565b6040805160208082528351818301528351919283929083019185019080838360005b838110156105fd5781810151838201526020016105e5565b50505050905090810190601f16801561062a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561064457600080fd5b5061013e600160a060020a03600435811690602435166119cd565b34801561066b57600080fd5b50610127600160a060020a0360043581169060243516611b30565b34801561069257600080fd5b5061069e600435611f16565b60408051938452600160a060020a03909216602084015260ff1682820152519081900360600190f35b3480156106d357600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526101279482359460248035600160a060020a031695369594606494920191908190840183828082843750949750611f449650505050505050565b34801561073c57600080fd5b50610757600160a060020a0360043581169060243516612051565b604080519485526020850193909352600160a060020a039091168383015260ff166060830152519081900360800190f35b34801561079457600080fd5b5061017d6120b2565b600080600080600085600681101580156107ba5750622932e08111155b15156107c557600080fd5b600160a060020a03891615156107da57600080fd5b600160a060020a03881615156107ef57600080fd5b600160a060020a03898116908916141561080857600080fd5b600160056000828254019250508190555060055495506003600087815260200190815260200160002093508360020160008a600160a060020a0316600160a060020a03168152602001908152602001600020925083600201600089600160a060020a0316600160a060020a03168152602001908152602001600020915086846000018190555060018360040160006101000a81548160ff02191690831515021790555060018260040160006101000a81548160ff02191690831515021790555060018460010160146101000a81548160ff021916908360ff1602179055506108f089896119cd565b6000818152600460205260409020549095501561090c57600080fd5b60008581526004602090815260409182902088905581518981529151600160a060020a03808c1693908d16928a927f669a4b0ac0b9994c0f82ed4dbe07bb421fe74e5951725af4f139c7443ebf049d9281900390910190a4505050505050505050565b60055481565b60046020526000908152604090205481565b600154600160a060020a031681565b60008681526003602052604081206001015481908190819081908b9060a060020a900460ff166002146109c857600080fd5b60008c8152600360209081526040808320600160a060020a038f168452600281019092529091206004810154919550935060ff161515610a0757600080fd5b60018301541515610a1757600080fd5b610a238c8b8b8b6120c1565b600160a060020a038116600090815260028601602052604090206004810154919750925060ff161515610a5557600080fd5b600160a060020a03868116908c161415610a6e57600080fd5b610a788a88612178565b60018401549095508514610a8b57600080fd5b50506000600290910155505050505050505050565b600081815260036020526040812060018082015483918291829160a060020a90910460ff1614610acf57600080fd5b8d8c8e8d8d8a306002546040516020018089600160a060020a0316600160a060020a0316606060020a02815260140188815260200187600160a060020a0316600160a060020a0316606060020a02815260140186815260200185815260200184815260200183600160a060020a0316600160a060020a0316606060020a028152601401828152602001985050505050505050506040516020818303038152906040526040518082805190602001908083835b60208310610ba05780518252601f199092019160209182019101610b81565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209250610bd983896122c9565b600160a060020a038f8116911614610bf057600080fd5b8d8c8e8d8d8d8b30600254604051602001808a600160a060020a0316600160a060020a0316606060020a02815260140189815260200188600160a060020a0316600160a060020a0316606060020a02815260140187815260200186815260200185815260200184815260200183600160a060020a0316600160a060020a0316606060020a02815260140182815260200199505050505050505050506040516020818303038152906040526040518082805190602001908083835b60208310610cc95780518252601f199092019160209182019101610caa565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209250610d0283886122c9565b600160a060020a038e8116911614610d1957600080fd5b5050600160a060020a03808d166000908152600284016020526040808220928e1682529020600482015460ff161515610d5157600080fd5b600481015460ff161515610d6457600080fd5b80548254019450848c1115610d7857600080fd5b848b1115610d8557600080fd5b8b8b018514610d9357600080fd5b60008a1115610e4a576000809054906101000a9004600160a060020a0316600160a060020a031663a9059cbb8f8c6040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b158015610e1357600080fd5b505af1158015610e27573d6000803e3d6000fd5b505050506040513d6020811015610e3d57600080fd5b50511515610e4a57600080fd5b6000891115610f01576000809054906101000a9004600160a060020a0316600160a060020a031663a9059cbb8e8b6040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b158015610eca57600080fd5b505af1158015610ede573d6000803e3d6000fd5b505050506040513d6020811015610ef457600080fd5b50511515610f0157600080fd5b8b8a1115610f0e57600080fd5b8a891115610f1b57600080fd5b898c039b50888b039a508b82600001819055508a81600001819055508360020160008f600160a060020a0316600160a060020a031681526020019081526020016000206000808201600090556001820160009055600282016000905560038201600090556004820160006101000a81549060ff0219169055600582016000905550508360020160008e600160a060020a0316600160a060020a031681526020019081526020016000206000808201600090556001820160009055600282016000905560038201600090556004820160006101000a81549060ff021916905560058201600090555050600360008781526020019081526020016000206000808201600090556001820160006101000a815490600160a060020a0302191690556001820160146101000a81549060ff02191690555050600160056000828254019250508190555060055495508360036000888152602001908152602001600020600082015481600001556001820160009054906101000a9004600160a060020a03168160010160006101000a815481600160a060020a030219169083600160a060020a031602179055506001820160149054906101000a900460ff168160010160146101000a81548160ff021916908360ff1602179055509050506110fe8e8e6119cd565b60008181526004602090815260409182902089905581518981529081018f90528082018e9052606081018d9052608081018c905290519194507f8edee97a023db4a4c6f8f411985d0d55e73a88b28c85867e661fd85d793ece7d919081900360a00190a15050505050505050505050505050565b60025481565b6000903b1190565b600083815260036020526040812060019081015482918291879160a060020a90910460ff16146111af57600080fd5b600085116111bc57600080fd5b6000878152600360209081526040808320600160a060020a038a168452600281019092529091206004810154919450925060ff1615156111fb57600080fd5b8154851161120857600080fd5b8154808603908101835560008054604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018590529051939750600160a060020a03909116926323b872dd92606480840193602093929083900390910190829087803b15801561128a57600080fd5b505af115801561129e573d6000803e3d6000fd5b505050506040513d60208110156112b457600080fd5b505115156112c157600080fd5b81546040805191825251600160a060020a0388169189917f2b55547a3b586ab51f65ee9ce4927fa6d25191388299988e89e059a02f9dd4459181900360200190a350505050505050565b60008681526003602052604081206001908101548291899160a060020a900460ff161461133757600080fd5b600089815260036020908152604080832060018101805474ff0000000000000000000000000000000000000000191674020000000000000000000000000000000000000000179055338452600281019092529091206004015490925060ff1615156113a157600080fd5b60018201805473ffffffffffffffffffffffffffffffffffffffff191633179055815443018255600086111561141a576113df8989898989896123a9565b92506113ee8984888a8c612473565b600160a060020a038316600090815260028301602052604090206004015460ff16151561141a57600080fd5b60405133908a907fa8621c489a70a0a06448f2b4e3477913a3744d5f27a380e5f0d8db13837ce7c690600090a3505050505050505050565b6000918252600360208181526040808520600160a060020a03949094168552600293840190915290922060048101548154600183015493830154929094015460ff90911694565b60008060008060008060006114ae8d8c6119cd565b6000818152600460209081526040808320548084526003909252909120919650945092506114e0848e8e8e8e8e6124c4565b9650600160a060020a038d8116908816146114fa57600080fd5b611508848e8e8e8e8d6124c4565b9650600160a060020a038b81169088161461152257600080fd5b5050600160a060020a03808c166000908152600283016020526040808220928c1682529020600482015460ff16151561155a57600080fd5b600481015460ff16151561156d57600080fd5b60018381015460a060020a900460ff161461158757600080fd5b806000015482600001540195508260020160008e600160a060020a0316600160a060020a031681526020019081526020016000206000808201600090556001820160009055600282016000905560038201600090556004820160006101000a81549060ff0219169055600582016000905550508260020160008c600160a060020a0316600160a060020a031681526020019081526020016000206000808201600090556001820160009055600282016000905560038201600090556004820160006101000a81549060ff021916905560058201600090555050600360008581526020019081526020016000206000808201600090556001820160006101000a815490600160a060020a0302191690556001820160146101000a81549060ff021916905550506004600086600019166000191681526020019081526020016000206000905560008c1115611782576000809054906101000a9004600160a060020a0316600160a060020a031663a9059cbb8e8e6040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b15801561174b57600080fd5b505af115801561175f573d6000803e3d6000fd5b505050506040513d602081101561177557600080fd5b5051151561178257600080fd5b60008a11156118345760008054604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a038f81166004830152602482018f90529151919092169263a9059cbb92604480820193602093909283900390910190829087803b1580156117fd57600080fd5b505af1158015611811573d6000803e3d6000fd5b505050506040513d602081101561182757600080fd5b5051151561183457600080fd5b8b8a01861461184257600080fd5b8b86101561184f57600080fd5b8986101561185c57600080fd5b604080518d8152602081018c9052815186927f0e239ef20c651bd0bc45e6f6a5fd46252d77d39d6602103e347add00cabdb0b4928290030190a250505050505050505050505050565b6000808086116118b457600080fd5b505060008881526003602052604090206001810154600160a060020a0316906118e28a888a8989898961255d565b600160a060020a038a81169116146118f957600080fd5b6119078a888a8989896123a9565b600160a060020a0383811691161461191e57600080fd5b61192b8a83888b8b612473565b604051600160a060020a038316908b907fe5ccf2144fc46e5dbed7d342686643a19421a58ee918650b15f766f645b8ff0790600090a3600181015460a060020a900460ff1660021461197c57600080fd5b805443111561198a57600080fd5b50505050505050505050565b60408051808201909152600581527f302e332e5f000000000000000000000000000000000000000000000000000000602082015281565b600081600160a060020a031683600160a060020a03161015611aa45782826040516020018083600160a060020a0316600160a060020a0316606060020a02815260140182600160a060020a0316600160a060020a0316606060020a028152601401925050506040516020818303038152906040526040518082805190602001908083835b60208310611a705780518252601f199092019160209182019101611a51565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209050611b2a565b81836040516020018083600160a060020a0316600160a060020a0316606060020a02815260140182600160a060020a0316600160a060020a0316606060020a0281526014019250505060405160208183030381529060405260405180828051906020019080838360208310611a705780518252601f199092019160209182019101611a51565b92915050565b600080600080600080600080600080611b498c8c6119cd565b6000818152600460209081526040808320548084526003909252909120600181015491975091955090935060a060020a900460ff16600214611b8a57600080fd5b82544311611b9757600080fd5b5050600160a060020a03808b166000908152600283016020526040808220928c1682529020600482015460ff161515611bcf57600080fd5b600481015460ff161515611be257600080fd5b8054825460058085015460028087015492860154908601549290910191018083018281039e50919b50995091019650881115611c1d57600099505b611c278a8761269e565b995089860398508260020160008d600160a060020a0316600160a060020a031681526020019081526020016000206000808201600090556001820160009055600282016000905560038201600090556004820160006101000a81549060ff0219169055600582016000905550508260020160008c600160a060020a0316600160a060020a031681526020019081526020016000206000808201600090556001820160009055600282016000905560038201600090556004820160006101000a81549060ff021916905560058201600090555050600360008681526020019081526020016000206000808201600090556001820160006101000a815490600160a060020a0302191690556001820160146101000a81549060ff021916905550506004600085600019166000191681526020019081526020016000206000905560008a1115611e1c576000809054906101000a9004600160a060020a0316600160a060020a031663a9059cbb8d8c6040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b158015611de557600080fd5b505af1158015611df9573d6000803e3d6000fd5b505050506040513d6020811015611e0f57600080fd5b50511515611e1c57600080fd5b6000891115611ece5760008054604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a038f81166004830152602482018e90529151919092169263a9059cbb92604480820193602093909283900390910190829087803b158015611e9757600080fd5b505af1158015611eab573d6000803e3d6000fd5b505050506040513d6020811015611ec157600080fd5b50511515611ece57600080fd5b604080518b8152602081018b9052815187927f0e239ef20c651bd0bc45e6f6a5fd46252d77d39d6602103e347add00cabdb0b4928290030190a2505050505050505050505050565b60036020526000908152604090208054600190910154600160a060020a0381169060a060020a900460ff1683565b600083815260036020526040812060010154819081908190879060a060020a900460ff16600214611f7457600080fd5b8551600010611f8257600080fd5b60008881526003602052604090208054909350431115611fa157600080fd5b600160a060020a0387166000908152600284016020526040902060018101549092501515611fce57600080fd5b611fd7866126b6565b600184015491965094508514611fec57600080fd5b600582018490556002820154604080518a8152600160a060020a038a1660208201528082018890526060810192909252517f5842365ce79285f3d49939b96a9866db3c41c82754215e2f6e11bcba3d1a61b79181900360800190a15050505050505050565b600080600080600080600061206689896119cd565b60009081526004602090815260408083205480845260039092529091208054600190910154919b909a50600160a060020a038216995060a060020a90910460ff16975095505050505050565b600054600160a060020a031681565b600254604080516020808201879052818301889052606060020a30026060830152607482019390935260948082018690528251808303909101815260b490910191829052805160009384939182918401908083835b602083106121355780518252601f199092019160209182019101612116565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020905061216e81846122c9565b9695505050505050565b60008060006020845181151561218a57fe5b061561219557600080fd5b602091505b835182116122c05750828101518085101561223457604080516020808201889052818301849052825180830384018152606090920192839052815191929182918401908083835b602083106122005780518252601f1990920191602091820191016121e1565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902094506122b5565b604080516020808201849052818301889052825180830384018152606090920192839052815191929182918401908083835b602083106122855780518252601f199092019160209182019101612266565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902094505b60208201915061219a565b50929392505050565b600080600080845160411415156122df57600080fd5b50505060208201516040830151606084015160001a601b60ff8216101561230457601b015b8060ff16601b148061231957508060ff16601c145b151561232457600080fd5b60408051600080825260208083018085528a905260ff8516838501526060830187905260808301869052925160019360a0808501949193601f19840193928390039091019190865af115801561237e573d6000803e3d6000fd5b5050604051601f190151945050600160a060020a03841615156123a057600080fd5b50505092915050565b600254604080516020808201899052818301889052606082018790526080820186905260a082018a9052606060020a300260c083015260d4808301949094528251808303909401845260f490910191829052825160009384939092909182918401908083835b6020831061242e5780518252601f19909201916020918201910161240f565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020905061246781846122c9565b98975050505050505050565b6000858152600360208181526040808420600160a060020a0389168552600281019092529092209081015485116124a957600080fd5b60038101949094555060018301919091556002909101555050565b60025460408051606060020a600160a060020a03808a168202602080850191909152603484018a90529088168202605484015260688301879052608883018b9052309190910260a883015260bc808301949094528251808303909401845260dc90910191829052825160009384939092909182918401908083836020831061242e5780518252601f19909201916020918201910161240f565b600080878787878c306002548a604051602001808981526020018860001916600019168152602001878152602001866000191660001916815260200185815260200184600160a060020a0316600160a060020a0316606060020a02815260140183815260200182805190602001908083835b602083106125ee5780518252601f1990920191602091820191016125cf565b6001836020036101000a038019825116818451168082178552505050505050905001985050505050505050506040516020818303038152906040526040518082805190602001908083835b602083106126585780518252601f199092019160209182019101612639565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020905061269181846122c9565b9998505050505050505050565b60008183116126ad57826126af565b815b9392505050565b8051600090819081808080806060808706156126d157600080fd5b60608704600101604051908082528060200260200182016040528015612701578160200160208202803883390190505b509050602095505b8686101561274b5761271b8a87612a0a565b958601959450925082816060880481518110151561273557fe5b6020908102909101015260609590950194612709565b6060870496505b60018711156129e157600287061561279f57806001880381518110151561277557fe5b90602001906020020151818881518110151561278d57fe5b60209081029091010152600196909601955b600095505b600187038610156129d65780866001018151811015156127c057fe5b6020908102909101015181518290889081106127d857fe5b6020908102909101015114156128075780868151811015156127f657fe5b9060200190602002015192506129ae565b808660010181518110151561281857fe5b60209081029091010151815182908890811061283057fe5b6020908102909101015110156128f957808681518110151561284e57fe5b90602001906020020151818760010181518110151561286957fe5b6020908102909101810151604080518084019490945283810191909152805180840382018152606090930190819052825190918291908401908083835b602083106128c55780518252601f1990920191602091820191016128a6565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902092506129ae565b808660010181518110151561290a57fe5b90602001906020020151818781518110151561292257fe5b6020908102909101810151604080518084019490945283810191909152805180840382018152606090930190819052825190918291908401908083835b6020831061297e5780518252601f19909201916020918201910161295f565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902092505b8281600288048151811015156129c057fe5b60209081029091010152600295909501946127a4565b600286049650612752565b8060008151811015156129f057fe5b602090810290910101519a94995093975050505050505050565b6000806000806000806000878951111515612a2b5795506000945085612b73565b888801805160208083015160409384015184518084018590528086018390526060808201839052865180830390910181526080909101958690528051949a509198509550929182918401908083835b60208310612a995780518252601f199092019160209182019101612a7a565b51815160209384036101000a6000190180199092169116179052604080519290940182900382206001547fc1f62946000000000000000000000000000000000000000000000000000000008452600484018a90529451909750600160a060020a03909416955063c1f62946945060248083019491935090918290030181600087803b158015612b2757600080fd5b505af1158015612b3b573d6000803e3d6000fd5b505050506040513d6020811015612b5157600080fd5b50519250821580612b625750828511155b15612b6c57600093505b8084965096505b505050505092509290505600a165627a7a723058203ecb22ff97e8f409cff9befc0018d56a3b104a0b4254814e20c238d739a3fb260029`

// DeployTokenNetwork deploys a new Ethereum contract, binding an instance of TokenNetwork to it.
func DeployTokenNetwork(auth *bind.TransactOpts, backend bind.ContractBackend, _token_address common.Address, _secret_registry common.Address, _chain_id *big.Int) (common.Address, *types.Transaction, *TokenNetwork, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenNetworkABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenNetworkBin), backend, _token_address, _secret_registry, _chain_id)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenNetwork{TokenNetworkCaller: TokenNetworkCaller{contract: contract}, TokenNetworkTransactor: TokenNetworkTransactor{contract: contract}, TokenNetworkFilterer: TokenNetworkFilterer{contract: contract}}, nil
}

// TokenNetwork is an auto generated Go binding around an Ethereum contract.
type TokenNetwork struct {
	TokenNetworkCaller     // Read-only binding to the contract
	TokenNetworkTransactor // Write-only binding to the contract
	TokenNetworkFilterer   // Log filterer for contract events
}

// TokenNetworkCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenNetworkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenNetworkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenNetworkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenNetworkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenNetworkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenNetworkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenNetworkSession struct {
	Contract     *TokenNetwork     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenNetworkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenNetworkCallerSession struct {
	Contract *TokenNetworkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TokenNetworkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenNetworkTransactorSession struct {
	Contract     *TokenNetworkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TokenNetworkRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenNetworkRaw struct {
	Contract *TokenNetwork // Generic contract binding to access the raw methods on
}

// TokenNetworkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenNetworkCallerRaw struct {
	Contract *TokenNetworkCaller // Generic read-only contract binding to access the raw methods on
}

// TokenNetworkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenNetworkTransactorRaw struct {
	Contract *TokenNetworkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenNetwork creates a new instance of TokenNetwork, bound to a specific deployed contract.
func NewTokenNetwork(address common.Address, backend bind.ContractBackend) (*TokenNetwork, error) {
	contract, err := bindTokenNetwork(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenNetwork{TokenNetworkCaller: TokenNetworkCaller{contract: contract}, TokenNetworkTransactor: TokenNetworkTransactor{contract: contract}, TokenNetworkFilterer: TokenNetworkFilterer{contract: contract}}, nil
}

// NewTokenNetworkCaller creates a new read-only instance of TokenNetwork, bound to a specific deployed contract.
func NewTokenNetworkCaller(address common.Address, caller bind.ContractCaller) (*TokenNetworkCaller, error) {
	contract, err := bindTokenNetwork(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkCaller{contract: contract}, nil
}

// NewTokenNetworkTransactor creates a new write-only instance of TokenNetwork, bound to a specific deployed contract.
func NewTokenNetworkTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenNetworkTransactor, error) {
	contract, err := bindTokenNetwork(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkTransactor{contract: contract}, nil
}

// NewTokenNetworkFilterer creates a new log filterer instance of TokenNetwork, bound to a specific deployed contract.
func NewTokenNetworkFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenNetworkFilterer, error) {
	contract, err := bindTokenNetwork(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkFilterer{contract: contract}, nil
}

// bindTokenNetwork binds a generic wrapper to an already deployed contract.
func bindTokenNetwork(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenNetworkABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenNetwork *TokenNetworkRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenNetwork.Contract.TokenNetworkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenNetwork *TokenNetworkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenNetwork.Contract.TokenNetworkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenNetwork *TokenNetworkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenNetwork.Contract.TokenNetworkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenNetwork *TokenNetworkCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenNetwork.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenNetwork *TokenNetworkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenNetwork.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenNetwork *TokenNetworkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenNetwork.Contract.contract.Transact(opts, method, params...)
}

// Chain_id is a free data retrieval call binding the contract method 0x3af973b1.
//
// Solidity: function chain_id() constant returns(uint256)
func (_TokenNetwork *TokenNetworkCaller) Chain_id(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenNetwork.contract.Call(opts, out, "chain_id")
	return *ret0, err
}

// Chain_id is a free data retrieval call binding the contract method 0x3af973b1.
//
// Solidity: function chain_id() constant returns(uint256)
func (_TokenNetwork *TokenNetworkSession) Chain_id() (*big.Int, error) {
	return _TokenNetwork.Contract.Chain_id(&_TokenNetwork.CallOpts)
}

// Chain_id is a free data retrieval call binding the contract method 0x3af973b1.
//
// Solidity: function chain_id() constant returns(uint256)
func (_TokenNetwork *TokenNetworkCallerSession) Chain_id() (*big.Int, error) {
	return _TokenNetwork.Contract.Chain_id(&_TokenNetwork.CallOpts)
}

// Channels is a free data retrieval call binding the contract method 0xe5949b5d.
//
// Solidity: function channels( uint256) constant returns(settle_block_number uint256, closing_participant address, state uint8)
func (_TokenNetwork *TokenNetworkCaller) Channels(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Settle_block_number *big.Int
	Closing_participant common.Address
	State               uint8
}, error) {
	ret := new(struct {
		Settle_block_number *big.Int
		Closing_participant common.Address
		State               uint8
	})
	out := ret
	err := _TokenNetwork.contract.Call(opts, out, "channels", arg0)
	return *ret, err
}

// Channels is a free data retrieval call binding the contract method 0xe5949b5d.
//
// Solidity: function channels( uint256) constant returns(settle_block_number uint256, closing_participant address, state uint8)
func (_TokenNetwork *TokenNetworkSession) Channels(arg0 *big.Int) (struct {
	Settle_block_number *big.Int
	Closing_participant common.Address
	State               uint8
}, error) {
	return _TokenNetwork.Contract.Channels(&_TokenNetwork.CallOpts, arg0)
}

// Channels is a free data retrieval call binding the contract method 0xe5949b5d.
//
// Solidity: function channels( uint256) constant returns(settle_block_number uint256, closing_participant address, state uint8)
func (_TokenNetwork *TokenNetworkCallerSession) Channels(arg0 *big.Int) (struct {
	Settle_block_number *big.Int
	Closing_participant common.Address
	State               uint8
}, error) {
	return _TokenNetwork.Contract.Channels(&_TokenNetwork.CallOpts, arg0)
}

// ContractExists is a free data retrieval call binding the contract method 0x7709bc78.
//
// Solidity: function contractExists(contract_address address) constant returns(bool)
func (_TokenNetwork *TokenNetworkCaller) ContractExists(opts *bind.CallOpts, contract_address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenNetwork.contract.Call(opts, out, "contractExists", contract_address)
	return *ret0, err
}

// ContractExists is a free data retrieval call binding the contract method 0x7709bc78.
//
// Solidity: function contractExists(contract_address address) constant returns(bool)
func (_TokenNetwork *TokenNetworkSession) ContractExists(contract_address common.Address) (bool, error) {
	return _TokenNetwork.Contract.ContractExists(&_TokenNetwork.CallOpts, contract_address)
}

// ContractExists is a free data retrieval call binding the contract method 0x7709bc78.
//
// Solidity: function contractExists(contract_address address) constant returns(bool)
func (_TokenNetwork *TokenNetworkCallerSession) ContractExists(contract_address common.Address) (bool, error) {
	return _TokenNetwork.Contract.ContractExists(&_TokenNetwork.CallOpts, contract_address)
}

// Contract_version is a free data retrieval call binding the contract method 0xb32c65c8.
//
// Solidity: function contract_version() constant returns(string)
func (_TokenNetwork *TokenNetworkCaller) Contract_version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TokenNetwork.contract.Call(opts, out, "contract_version")
	return *ret0, err
}

// Contract_version is a free data retrieval call binding the contract method 0xb32c65c8.
//
// Solidity: function contract_version() constant returns(string)
func (_TokenNetwork *TokenNetworkSession) Contract_version() (string, error) {
	return _TokenNetwork.Contract.Contract_version(&_TokenNetwork.CallOpts)
}

// Contract_version is a free data retrieval call binding the contract method 0xb32c65c8.
//
// Solidity: function contract_version() constant returns(string)
func (_TokenNetwork *TokenNetworkCallerSession) Contract_version() (string, error) {
	return _TokenNetwork.Contract.Contract_version(&_TokenNetwork.CallOpts)
}

// GetChannelHash is a free data retrieval call binding the contract method 0xbb94ccc8.
//
// Solidity: function getChannelHash(participant address, partner address) constant returns(bytes32)
func (_TokenNetwork *TokenNetworkCaller) GetChannelHash(opts *bind.CallOpts, participant common.Address, partner common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TokenNetwork.contract.Call(opts, out, "getChannelHash", participant, partner)
	return *ret0, err
}

// GetChannelHash is a free data retrieval call binding the contract method 0xbb94ccc8.
//
// Solidity: function getChannelHash(participant address, partner address) constant returns(bytes32)
func (_TokenNetwork *TokenNetworkSession) GetChannelHash(participant common.Address, partner common.Address) ([32]byte, error) {
	return _TokenNetwork.Contract.GetChannelHash(&_TokenNetwork.CallOpts, participant, partner)
}

// GetChannelHash is a free data retrieval call binding the contract method 0xbb94ccc8.
//
// Solidity: function getChannelHash(participant address, partner address) constant returns(bytes32)
func (_TokenNetwork *TokenNetworkCallerSession) GetChannelHash(participant common.Address, partner common.Address) ([32]byte, error) {
	return _TokenNetwork.Contract.GetChannelHash(&_TokenNetwork.CallOpts, participant, partner)
}

// GetChannelInfo is a free data retrieval call binding the contract method 0xf94c9e13.
//
// Solidity: function getChannelInfo(participant1 address, participant2 address) constant returns(uint256, uint256, address, uint8)
func (_TokenNetwork *TokenNetworkCaller) GetChannelInfo(opts *bind.CallOpts, participant1 common.Address, participant2 common.Address) (*big.Int, *big.Int, common.Address, uint8, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(common.Address)
		ret3 = new(uint8)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _TokenNetwork.contract.Call(opts, out, "getChannelInfo", participant1, participant2)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetChannelInfo is a free data retrieval call binding the contract method 0xf94c9e13.
//
// Solidity: function getChannelInfo(participant1 address, participant2 address) constant returns(uint256, uint256, address, uint8)
func (_TokenNetwork *TokenNetworkSession) GetChannelInfo(participant1 common.Address, participant2 common.Address) (*big.Int, *big.Int, common.Address, uint8, error) {
	return _TokenNetwork.Contract.GetChannelInfo(&_TokenNetwork.CallOpts, participant1, participant2)
}

// GetChannelInfo is a free data retrieval call binding the contract method 0xf94c9e13.
//
// Solidity: function getChannelInfo(participant1 address, participant2 address) constant returns(uint256, uint256, address, uint8)
func (_TokenNetwork *TokenNetworkCallerSession) GetChannelInfo(participant1 common.Address, participant2 common.Address) (*big.Int, *big.Int, common.Address, uint8, error) {
	return _TokenNetwork.Contract.GetChannelInfo(&_TokenNetwork.CallOpts, participant1, participant2)
}

// GetChannelParticipantInfo is a free data retrieval call binding the contract method 0x7fb5885e.
//
// Solidity: function getChannelParticipantInfo(channel_identifier uint256, participant address) constant returns(bool, uint256, bytes32, uint256, uint256)
func (_TokenNetwork *TokenNetworkCaller) GetChannelParticipantInfo(opts *bind.CallOpts, channel_identifier *big.Int, participant common.Address) (bool, *big.Int, [32]byte, *big.Int, *big.Int, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
		ret2 = new([32]byte)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _TokenNetwork.contract.Call(opts, out, "getChannelParticipantInfo", channel_identifier, participant)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetChannelParticipantInfo is a free data retrieval call binding the contract method 0x7fb5885e.
//
// Solidity: function getChannelParticipantInfo(channel_identifier uint256, participant address) constant returns(bool, uint256, bytes32, uint256, uint256)
func (_TokenNetwork *TokenNetworkSession) GetChannelParticipantInfo(channel_identifier *big.Int, participant common.Address) (bool, *big.Int, [32]byte, *big.Int, *big.Int, error) {
	return _TokenNetwork.Contract.GetChannelParticipantInfo(&_TokenNetwork.CallOpts, channel_identifier, participant)
}

// GetChannelParticipantInfo is a free data retrieval call binding the contract method 0x7fb5885e.
//
// Solidity: function getChannelParticipantInfo(channel_identifier uint256, participant address) constant returns(bool, uint256, bytes32, uint256, uint256)
func (_TokenNetwork *TokenNetworkCallerSession) GetChannelParticipantInfo(channel_identifier *big.Int, participant common.Address) (bool, *big.Int, [32]byte, *big.Int, *big.Int, error) {
	return _TokenNetwork.Contract.GetChannelParticipantInfo(&_TokenNetwork.CallOpts, channel_identifier, participant)
}

// Last_channel_index is a free data retrieval call binding the contract method 0x1f466acf.
//
// Solidity: function last_channel_index() constant returns(uint256)
func (_TokenNetwork *TokenNetworkCaller) Last_channel_index(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenNetwork.contract.Call(opts, out, "last_channel_index")
	return *ret0, err
}

// Last_channel_index is a free data retrieval call binding the contract method 0x1f466acf.
//
// Solidity: function last_channel_index() constant returns(uint256)
func (_TokenNetwork *TokenNetworkSession) Last_channel_index() (*big.Int, error) {
	return _TokenNetwork.Contract.Last_channel_index(&_TokenNetwork.CallOpts)
}

// Last_channel_index is a free data retrieval call binding the contract method 0x1f466acf.
//
// Solidity: function last_channel_index() constant returns(uint256)
func (_TokenNetwork *TokenNetworkCallerSession) Last_channel_index() (*big.Int, error) {
	return _TokenNetwork.Contract.Last_channel_index(&_TokenNetwork.CallOpts)
}

// Openedchannels is a free data retrieval call binding the contract method 0x2419a2e0.
//
// Solidity: function openedchannels( bytes32) constant returns(uint256)
func (_TokenNetwork *TokenNetworkCaller) Openedchannels(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenNetwork.contract.Call(opts, out, "openedchannels", arg0)
	return *ret0, err
}

// Openedchannels is a free data retrieval call binding the contract method 0x2419a2e0.
//
// Solidity: function openedchannels( bytes32) constant returns(uint256)
func (_TokenNetwork *TokenNetworkSession) Openedchannels(arg0 [32]byte) (*big.Int, error) {
	return _TokenNetwork.Contract.Openedchannels(&_TokenNetwork.CallOpts, arg0)
}

// Openedchannels is a free data retrieval call binding the contract method 0x2419a2e0.
//
// Solidity: function openedchannels( bytes32) constant returns(uint256)
func (_TokenNetwork *TokenNetworkCallerSession) Openedchannels(arg0 [32]byte) (*big.Int, error) {
	return _TokenNetwork.Contract.Openedchannels(&_TokenNetwork.CallOpts, arg0)
}

// Secret_registry is a free data retrieval call binding the contract method 0x24d73a93.
//
// Solidity: function secret_registry() constant returns(address)
func (_TokenNetwork *TokenNetworkCaller) Secret_registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenNetwork.contract.Call(opts, out, "secret_registry")
	return *ret0, err
}

// Secret_registry is a free data retrieval call binding the contract method 0x24d73a93.
//
// Solidity: function secret_registry() constant returns(address)
func (_TokenNetwork *TokenNetworkSession) Secret_registry() (common.Address, error) {
	return _TokenNetwork.Contract.Secret_registry(&_TokenNetwork.CallOpts)
}

// Secret_registry is a free data retrieval call binding the contract method 0x24d73a93.
//
// Solidity: function secret_registry() constant returns(address)
func (_TokenNetwork *TokenNetworkCallerSession) Secret_registry() (common.Address, error) {
	return _TokenNetwork.Contract.Secret_registry(&_TokenNetwork.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_TokenNetwork *TokenNetworkCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenNetwork.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_TokenNetwork *TokenNetworkSession) Token() (common.Address, error) {
	return _TokenNetwork.Contract.Token(&_TokenNetwork.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_TokenNetwork *TokenNetworkCallerSession) Token() (common.Address, error) {
	return _TokenNetwork.Contract.Token(&_TokenNetwork.CallOpts)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x7c0f097c.
//
// Solidity: function closeChannel(channel_identifier uint256, transferred_amount uint256, locksroot bytes32, nonce uint256, additional_hash bytes32, signature bytes) returns()
func (_TokenNetwork *TokenNetworkTransactor) CloseChannel(opts *bind.TransactOpts, channel_identifier *big.Int, transferred_amount *big.Int, locksroot [32]byte, nonce *big.Int, additional_hash [32]byte, signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.contract.Transact(opts, "closeChannel", channel_identifier, transferred_amount, locksroot, nonce, additional_hash, signature)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x7c0f097c.
//
// Solidity: function closeChannel(channel_identifier uint256, transferred_amount uint256, locksroot bytes32, nonce uint256, additional_hash bytes32, signature bytes) returns()
func (_TokenNetwork *TokenNetworkSession) CloseChannel(channel_identifier *big.Int, transferred_amount *big.Int, locksroot [32]byte, nonce *big.Int, additional_hash [32]byte, signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.CloseChannel(&_TokenNetwork.TransactOpts, channel_identifier, transferred_amount, locksroot, nonce, additional_hash, signature)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x7c0f097c.
//
// Solidity: function closeChannel(channel_identifier uint256, transferred_amount uint256, locksroot bytes32, nonce uint256, additional_hash bytes32, signature bytes) returns()
func (_TokenNetwork *TokenNetworkTransactorSession) CloseChannel(channel_identifier *big.Int, transferred_amount *big.Int, locksroot [32]byte, nonce *big.Int, additional_hash [32]byte, signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.CloseChannel(&_TokenNetwork.TransactOpts, channel_identifier, transferred_amount, locksroot, nonce, additional_hash, signature)
}

// CooperativeSettle is a paid mutator transaction binding the contract method 0x8568536a.
//
// Solidity: function cooperativeSettle(participant1_address address, participant1_balance uint256, participant2_address address, participant2_balance uint256, participant1_signature bytes, participant2_signature bytes) returns()
func (_TokenNetwork *TokenNetworkTransactor) CooperativeSettle(opts *bind.TransactOpts, participant1_address common.Address, participant1_balance *big.Int, participant2_address common.Address, participant2_balance *big.Int, participant1_signature []byte, participant2_signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.contract.Transact(opts, "cooperativeSettle", participant1_address, participant1_balance, participant2_address, participant2_balance, participant1_signature, participant2_signature)
}

// CooperativeSettle is a paid mutator transaction binding the contract method 0x8568536a.
//
// Solidity: function cooperativeSettle(participant1_address address, participant1_balance uint256, participant2_address address, participant2_balance uint256, participant1_signature bytes, participant2_signature bytes) returns()
func (_TokenNetwork *TokenNetworkSession) CooperativeSettle(participant1_address common.Address, participant1_balance *big.Int, participant2_address common.Address, participant2_balance *big.Int, participant1_signature []byte, participant2_signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.CooperativeSettle(&_TokenNetwork.TransactOpts, participant1_address, participant1_balance, participant2_address, participant2_balance, participant1_signature, participant2_signature)
}

// CooperativeSettle is a paid mutator transaction binding the contract method 0x8568536a.
//
// Solidity: function cooperativeSettle(participant1_address address, participant1_balance uint256, participant2_address address, participant2_balance uint256, participant1_signature bytes, participant2_signature bytes) returns()
func (_TokenNetwork *TokenNetworkTransactorSession) CooperativeSettle(participant1_address common.Address, participant1_balance *big.Int, participant2_address common.Address, participant2_balance *big.Int, participant1_signature []byte, participant2_signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.CooperativeSettle(&_TokenNetwork.TransactOpts, participant1_address, participant1_balance, participant2_address, participant2_balance, participant1_signature, participant2_signature)
}

// OpenChannel is a paid mutator transaction binding the contract method 0x0a798f24.
//
// Solidity: function openChannel(participant1 address, participant2 address, settle_timeout uint256) returns()
func (_TokenNetwork *TokenNetworkTransactor) OpenChannel(opts *bind.TransactOpts, participant1 common.Address, participant2 common.Address, settle_timeout *big.Int) (*types.Transaction, error) {
	return _TokenNetwork.contract.Transact(opts, "openChannel", participant1, participant2, settle_timeout)
}

// OpenChannel is a paid mutator transaction binding the contract method 0x0a798f24.
//
// Solidity: function openChannel(participant1 address, participant2 address, settle_timeout uint256) returns()
func (_TokenNetwork *TokenNetworkSession) OpenChannel(participant1 common.Address, participant2 common.Address, settle_timeout *big.Int) (*types.Transaction, error) {
	return _TokenNetwork.Contract.OpenChannel(&_TokenNetwork.TransactOpts, participant1, participant2, settle_timeout)
}

// OpenChannel is a paid mutator transaction binding the contract method 0x0a798f24.
//
// Solidity: function openChannel(participant1 address, participant2 address, settle_timeout uint256) returns()
func (_TokenNetwork *TokenNetworkTransactorSession) OpenChannel(participant1 common.Address, participant2 common.Address, settle_timeout *big.Int) (*types.Transaction, error) {
	return _TokenNetwork.Contract.OpenChannel(&_TokenNetwork.TransactOpts, participant1, participant2, settle_timeout)
}

// PunishObsoleteUnlock is a paid mutator transaction binding the contract method 0x28a8d70f.
//
// Solidity: function punishObsoleteUnlock(channel_identifier uint256, beneficiary address, lockhash bytes32, additional_hash bytes32, signature bytes, merkle_proof bytes) returns()
func (_TokenNetwork *TokenNetworkTransactor) PunishObsoleteUnlock(opts *bind.TransactOpts, channel_identifier *big.Int, beneficiary common.Address, lockhash [32]byte, additional_hash [32]byte, signature []byte, merkle_proof []byte) (*types.Transaction, error) {
	return _TokenNetwork.contract.Transact(opts, "punishObsoleteUnlock", channel_identifier, beneficiary, lockhash, additional_hash, signature, merkle_proof)
}

// PunishObsoleteUnlock is a paid mutator transaction binding the contract method 0x28a8d70f.
//
// Solidity: function punishObsoleteUnlock(channel_identifier uint256, beneficiary address, lockhash bytes32, additional_hash bytes32, signature bytes, merkle_proof bytes) returns()
func (_TokenNetwork *TokenNetworkSession) PunishObsoleteUnlock(channel_identifier *big.Int, beneficiary common.Address, lockhash [32]byte, additional_hash [32]byte, signature []byte, merkle_proof []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.PunishObsoleteUnlock(&_TokenNetwork.TransactOpts, channel_identifier, beneficiary, lockhash, additional_hash, signature, merkle_proof)
}

// PunishObsoleteUnlock is a paid mutator transaction binding the contract method 0x28a8d70f.
//
// Solidity: function punishObsoleteUnlock(channel_identifier uint256, beneficiary address, lockhash bytes32, additional_hash bytes32, signature bytes, merkle_proof bytes) returns()
func (_TokenNetwork *TokenNetworkTransactorSession) PunishObsoleteUnlock(channel_identifier *big.Int, beneficiary common.Address, lockhash [32]byte, additional_hash [32]byte, signature []byte, merkle_proof []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.PunishObsoleteUnlock(&_TokenNetwork.TransactOpts, channel_identifier, beneficiary, lockhash, additional_hash, signature, merkle_proof)
}

// SetTotalDeposit is a paid mutator transaction binding the contract method 0x7944bd1f.
//
// Solidity: function setTotalDeposit(channel_identifier uint256, participant address, total_deposit uint256) returns()
func (_TokenNetwork *TokenNetworkTransactor) SetTotalDeposit(opts *bind.TransactOpts, channel_identifier *big.Int, participant common.Address, total_deposit *big.Int) (*types.Transaction, error) {
	return _TokenNetwork.contract.Transact(opts, "setTotalDeposit", channel_identifier, participant, total_deposit)
}

// SetTotalDeposit is a paid mutator transaction binding the contract method 0x7944bd1f.
//
// Solidity: function setTotalDeposit(channel_identifier uint256, participant address, total_deposit uint256) returns()
func (_TokenNetwork *TokenNetworkSession) SetTotalDeposit(channel_identifier *big.Int, participant common.Address, total_deposit *big.Int) (*types.Transaction, error) {
	return _TokenNetwork.Contract.SetTotalDeposit(&_TokenNetwork.TransactOpts, channel_identifier, participant, total_deposit)
}

// SetTotalDeposit is a paid mutator transaction binding the contract method 0x7944bd1f.
//
// Solidity: function setTotalDeposit(channel_identifier uint256, participant address, total_deposit uint256) returns()
func (_TokenNetwork *TokenNetworkTransactorSession) SetTotalDeposit(channel_identifier *big.Int, participant common.Address, total_deposit *big.Int) (*types.Transaction, error) {
	return _TokenNetwork.Contract.SetTotalDeposit(&_TokenNetwork.TransactOpts, channel_identifier, participant, total_deposit)
}

// SettleChannel is a paid mutator transaction binding the contract method 0xdee8bd0a.
//
// Solidity: function settleChannel(participant1 address, participant2 address) returns()
func (_TokenNetwork *TokenNetworkTransactor) SettleChannel(opts *bind.TransactOpts, participant1 common.Address, participant2 common.Address) (*types.Transaction, error) {
	return _TokenNetwork.contract.Transact(opts, "settleChannel", participant1, participant2)
}

// SettleChannel is a paid mutator transaction binding the contract method 0xdee8bd0a.
//
// Solidity: function settleChannel(participant1 address, participant2 address) returns()
func (_TokenNetwork *TokenNetworkSession) SettleChannel(participant1 common.Address, participant2 common.Address) (*types.Transaction, error) {
	return _TokenNetwork.Contract.SettleChannel(&_TokenNetwork.TransactOpts, participant1, participant2)
}

// SettleChannel is a paid mutator transaction binding the contract method 0xdee8bd0a.
//
// Solidity: function settleChannel(participant1 address, participant2 address) returns()
func (_TokenNetwork *TokenNetworkTransactorSession) SettleChannel(participant1 common.Address, participant2 common.Address) (*types.Transaction, error) {
	return _TokenNetwork.Contract.SettleChannel(&_TokenNetwork.TransactOpts, participant1, participant2)
}

// Unlock is a paid mutator transaction binding the contract method 0xf3da17e8.
//
// Solidity: function unlock(channel_identifier uint256, participant address, merkle_tree_leaves bytes) returns()
func (_TokenNetwork *TokenNetworkTransactor) Unlock(opts *bind.TransactOpts, channel_identifier *big.Int, participant common.Address, merkle_tree_leaves []byte) (*types.Transaction, error) {
	return _TokenNetwork.contract.Transact(opts, "unlock", channel_identifier, participant, merkle_tree_leaves)
}

// Unlock is a paid mutator transaction binding the contract method 0xf3da17e8.
//
// Solidity: function unlock(channel_identifier uint256, participant address, merkle_tree_leaves bytes) returns()
func (_TokenNetwork *TokenNetworkSession) Unlock(channel_identifier *big.Int, participant common.Address, merkle_tree_leaves []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.Unlock(&_TokenNetwork.TransactOpts, channel_identifier, participant, merkle_tree_leaves)
}

// Unlock is a paid mutator transaction binding the contract method 0xf3da17e8.
//
// Solidity: function unlock(channel_identifier uint256, participant address, merkle_tree_leaves bytes) returns()
func (_TokenNetwork *TokenNetworkTransactorSession) Unlock(channel_identifier *big.Int, participant common.Address, merkle_tree_leaves []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.Unlock(&_TokenNetwork.TransactOpts, channel_identifier, participant, merkle_tree_leaves)
}

// UpdateNonClosingBalanceProof is a paid mutator transaction binding the contract method 0x8d373f51.
//
// Solidity: function updateNonClosingBalanceProof(channel_identifier uint256, non_closing_participant address, locksroot bytes32, transferred_amount uint256, nonce uint256, additional_hash bytes32, closing_signature bytes, non_closing_signature bytes) returns()
func (_TokenNetwork *TokenNetworkTransactor) UpdateNonClosingBalanceProof(opts *bind.TransactOpts, channel_identifier *big.Int, non_closing_participant common.Address, locksroot [32]byte, transferred_amount *big.Int, nonce *big.Int, additional_hash [32]byte, closing_signature []byte, non_closing_signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.contract.Transact(opts, "updateNonClosingBalanceProof", channel_identifier, non_closing_participant, locksroot, transferred_amount, nonce, additional_hash, closing_signature, non_closing_signature)
}

// UpdateNonClosingBalanceProof is a paid mutator transaction binding the contract method 0x8d373f51.
//
// Solidity: function updateNonClosingBalanceProof(channel_identifier uint256, non_closing_participant address, locksroot bytes32, transferred_amount uint256, nonce uint256, additional_hash bytes32, closing_signature bytes, non_closing_signature bytes) returns()
func (_TokenNetwork *TokenNetworkSession) UpdateNonClosingBalanceProof(channel_identifier *big.Int, non_closing_participant common.Address, locksroot [32]byte, transferred_amount *big.Int, nonce *big.Int, additional_hash [32]byte, closing_signature []byte, non_closing_signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.UpdateNonClosingBalanceProof(&_TokenNetwork.TransactOpts, channel_identifier, non_closing_participant, locksroot, transferred_amount, nonce, additional_hash, closing_signature, non_closing_signature)
}

// UpdateNonClosingBalanceProof is a paid mutator transaction binding the contract method 0x8d373f51.
//
// Solidity: function updateNonClosingBalanceProof(channel_identifier uint256, non_closing_participant address, locksroot bytes32, transferred_amount uint256, nonce uint256, additional_hash bytes32, closing_signature bytes, non_closing_signature bytes) returns()
func (_TokenNetwork *TokenNetworkTransactorSession) UpdateNonClosingBalanceProof(channel_identifier *big.Int, non_closing_participant common.Address, locksroot [32]byte, transferred_amount *big.Int, nonce *big.Int, additional_hash [32]byte, closing_signature []byte, non_closing_signature []byte) (*types.Transaction, error) {
	return _TokenNetwork.Contract.UpdateNonClosingBalanceProof(&_TokenNetwork.TransactOpts, channel_identifier, non_closing_participant, locksroot, transferred_amount, nonce, additional_hash, closing_signature, non_closing_signature)
}

// WithDraw is a paid mutator transaction binding the contract method 0x35fefe7d.
//
// Solidity: function withDraw(participant1 address, participant2 address, participant1_deposit uint256, participant2_deposit uint256, participant1_withdraw uint256, participant2_withdraw uint256, participant1_signature bytes, participant2_signature bytes, channel_identifier uint256) returns()
func (_TokenNetwork *TokenNetworkTransactor) WithDraw(opts *bind.TransactOpts, participant1 common.Address, participant2 common.Address, participant1_deposit *big.Int, participant2_deposit *big.Int, participant1_withdraw *big.Int, participant2_withdraw *big.Int, participant1_signature []byte, participant2_signature []byte, channel_identifier *big.Int) (*types.Transaction, error) {
	return _TokenNetwork.contract.Transact(opts, "withDraw", participant1, participant2, participant1_deposit, participant2_deposit, participant1_withdraw, participant2_withdraw, participant1_signature, participant2_signature, channel_identifier)
}

// WithDraw is a paid mutator transaction binding the contract method 0x35fefe7d.
//
// Solidity: function withDraw(participant1 address, participant2 address, participant1_deposit uint256, participant2_deposit uint256, participant1_withdraw uint256, participant2_withdraw uint256, participant1_signature bytes, participant2_signature bytes, channel_identifier uint256) returns()
func (_TokenNetwork *TokenNetworkSession) WithDraw(participant1 common.Address, participant2 common.Address, participant1_deposit *big.Int, participant2_deposit *big.Int, participant1_withdraw *big.Int, participant2_withdraw *big.Int, participant1_signature []byte, participant2_signature []byte, channel_identifier *big.Int) (*types.Transaction, error) {
	return _TokenNetwork.Contract.WithDraw(&_TokenNetwork.TransactOpts, participant1, participant2, participant1_deposit, participant2_deposit, participant1_withdraw, participant2_withdraw, participant1_signature, participant2_signature, channel_identifier)
}

// WithDraw is a paid mutator transaction binding the contract method 0x35fefe7d.
//
// Solidity: function withDraw(participant1 address, participant2 address, participant1_deposit uint256, participant2_deposit uint256, participant1_withdraw uint256, participant2_withdraw uint256, participant1_signature bytes, participant2_signature bytes, channel_identifier uint256) returns()
func (_TokenNetwork *TokenNetworkTransactorSession) WithDraw(participant1 common.Address, participant2 common.Address, participant1_deposit *big.Int, participant2_deposit *big.Int, participant1_withdraw *big.Int, participant2_withdraw *big.Int, participant1_signature []byte, participant2_signature []byte, channel_identifier *big.Int) (*types.Transaction, error) {
	return _TokenNetwork.Contract.WithDraw(&_TokenNetwork.TransactOpts, participant1, participant2, participant1_deposit, participant2_deposit, participant1_withdraw, participant2_withdraw, participant1_signature, participant2_signature, channel_identifier)
}

// TokenNetworkChannelClosedIterator is returned from FilterChannelClosed and is used to iterate over the raw logs and unpacked data for ChannelClosed events raised by the TokenNetwork contract.
type TokenNetworkChannelClosedIterator struct {
	Event *TokenNetworkChannelClosed // Event containing the contract specifics and raw log

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
func (it *TokenNetworkChannelClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenNetworkChannelClosed)
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
		it.Event = new(TokenNetworkChannelClosed)
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
func (it *TokenNetworkChannelClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenNetworkChannelClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenNetworkChannelClosed represents a ChannelClosed event raised by the TokenNetwork contract.
type TokenNetworkChannelClosed struct {
	Channel_identifier  *big.Int
	Closing_participant common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterChannelClosed is a free log retrieval operation binding the contract event 0xa8621c489a70a0a06448f2b4e3477913a3744d5f27a380e5f0d8db13837ce7c6.
//
// Solidity: event ChannelClosed(channel_identifier indexed uint256, closing_participant indexed address)
func (_TokenNetwork *TokenNetworkFilterer) FilterChannelClosed(opts *bind.FilterOpts, channel_identifier []*big.Int, closing_participant []common.Address) (*TokenNetworkChannelClosedIterator, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}
	var closing_participantRule []interface{}
	for _, closing_participantItem := range closing_participant {
		closing_participantRule = append(closing_participantRule, closing_participantItem)
	}

	logs, sub, err := _TokenNetwork.contract.FilterLogs(opts, "ChannelClosed", channel_identifierRule, closing_participantRule)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkChannelClosedIterator{contract: _TokenNetwork.contract, event: "ChannelClosed", logs: logs, sub: sub}, nil
}

// WatchChannelClosed is a free log subscription operation binding the contract event 0xa8621c489a70a0a06448f2b4e3477913a3744d5f27a380e5f0d8db13837ce7c6.
//
// Solidity: event ChannelClosed(channel_identifier indexed uint256, closing_participant indexed address)
func (_TokenNetwork *TokenNetworkFilterer) WatchChannelClosed(opts *bind.WatchOpts, sink chan<- *TokenNetworkChannelClosed, channel_identifier []*big.Int, closing_participant []common.Address) (event.Subscription, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}
	var closing_participantRule []interface{}
	for _, closing_participantItem := range closing_participant {
		closing_participantRule = append(closing_participantRule, closing_participantItem)
	}

	logs, sub, err := _TokenNetwork.contract.WatchLogs(opts, "ChannelClosed", channel_identifierRule, closing_participantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenNetworkChannelClosed)
				if err := _TokenNetwork.contract.UnpackLog(event, "ChannelClosed", log); err != nil {
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

// TokenNetworkChannelNewDepositIterator is returned from FilterChannelNewDeposit and is used to iterate over the raw logs and unpacked data for ChannelNewDeposit events raised by the TokenNetwork contract.
type TokenNetworkChannelNewDepositIterator struct {
	Event *TokenNetworkChannelNewDeposit // Event containing the contract specifics and raw log

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
func (it *TokenNetworkChannelNewDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenNetworkChannelNewDeposit)
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
		it.Event = new(TokenNetworkChannelNewDeposit)
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
func (it *TokenNetworkChannelNewDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenNetworkChannelNewDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenNetworkChannelNewDeposit represents a ChannelNewDeposit event raised by the TokenNetwork contract.
type TokenNetworkChannelNewDeposit struct {
	Channel_identifier *big.Int
	Participant        common.Address
	Total_deposit      *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterChannelNewDeposit is a free log retrieval operation binding the contract event 0x2b55547a3b586ab51f65ee9ce4927fa6d25191388299988e89e059a02f9dd445.
//
// Solidity: event ChannelNewDeposit(channel_identifier indexed uint256, participant indexed address, total_deposit uint256)
func (_TokenNetwork *TokenNetworkFilterer) FilterChannelNewDeposit(opts *bind.FilterOpts, channel_identifier []*big.Int, participant []common.Address) (*TokenNetworkChannelNewDepositIterator, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}
	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _TokenNetwork.contract.FilterLogs(opts, "ChannelNewDeposit", channel_identifierRule, participantRule)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkChannelNewDepositIterator{contract: _TokenNetwork.contract, event: "ChannelNewDeposit", logs: logs, sub: sub}, nil
}

// WatchChannelNewDeposit is a free log subscription operation binding the contract event 0x2b55547a3b586ab51f65ee9ce4927fa6d25191388299988e89e059a02f9dd445.
//
// Solidity: event ChannelNewDeposit(channel_identifier indexed uint256, participant indexed address, total_deposit uint256)
func (_TokenNetwork *TokenNetworkFilterer) WatchChannelNewDeposit(opts *bind.WatchOpts, sink chan<- *TokenNetworkChannelNewDeposit, channel_identifier []*big.Int, participant []common.Address) (event.Subscription, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}
	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _TokenNetwork.contract.WatchLogs(opts, "ChannelNewDeposit", channel_identifierRule, participantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenNetworkChannelNewDeposit)
				if err := _TokenNetwork.contract.UnpackLog(event, "ChannelNewDeposit", log); err != nil {
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

// TokenNetworkChannelOpenedIterator is returned from FilterChannelOpened and is used to iterate over the raw logs and unpacked data for ChannelOpened events raised by the TokenNetwork contract.
type TokenNetworkChannelOpenedIterator struct {
	Event *TokenNetworkChannelOpened // Event containing the contract specifics and raw log

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
func (it *TokenNetworkChannelOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenNetworkChannelOpened)
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
		it.Event = new(TokenNetworkChannelOpened)
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
func (it *TokenNetworkChannelOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenNetworkChannelOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenNetworkChannelOpened represents a ChannelOpened event raised by the TokenNetwork contract.
type TokenNetworkChannelOpened struct {
	Channel_identifier *big.Int
	Participant1       common.Address
	Participant2       common.Address
	Settle_timeout     *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterChannelOpened is a free log retrieval operation binding the contract event 0x669a4b0ac0b9994c0f82ed4dbe07bb421fe74e5951725af4f139c7443ebf049d.
//
// Solidity: event ChannelOpened(channel_identifier indexed uint256, participant1 indexed address, participant2 indexed address, settle_timeout uint256)
func (_TokenNetwork *TokenNetworkFilterer) FilterChannelOpened(opts *bind.FilterOpts, channel_identifier []*big.Int, participant1 []common.Address, participant2 []common.Address) (*TokenNetworkChannelOpenedIterator, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}
	var participant1Rule []interface{}
	for _, participant1Item := range participant1 {
		participant1Rule = append(participant1Rule, participant1Item)
	}
	var participant2Rule []interface{}
	for _, participant2Item := range participant2 {
		participant2Rule = append(participant2Rule, participant2Item)
	}

	logs, sub, err := _TokenNetwork.contract.FilterLogs(opts, "ChannelOpened", channel_identifierRule, participant1Rule, participant2Rule)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkChannelOpenedIterator{contract: _TokenNetwork.contract, event: "ChannelOpened", logs: logs, sub: sub}, nil
}

// WatchChannelOpened is a free log subscription operation binding the contract event 0x669a4b0ac0b9994c0f82ed4dbe07bb421fe74e5951725af4f139c7443ebf049d.
//
// Solidity: event ChannelOpened(channel_identifier indexed uint256, participant1 indexed address, participant2 indexed address, settle_timeout uint256)
func (_TokenNetwork *TokenNetworkFilterer) WatchChannelOpened(opts *bind.WatchOpts, sink chan<- *TokenNetworkChannelOpened, channel_identifier []*big.Int, participant1 []common.Address, participant2 []common.Address) (event.Subscription, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}
	var participant1Rule []interface{}
	for _, participant1Item := range participant1 {
		participant1Rule = append(participant1Rule, participant1Item)
	}
	var participant2Rule []interface{}
	for _, participant2Item := range participant2 {
		participant2Rule = append(participant2Rule, participant2Item)
	}

	logs, sub, err := _TokenNetwork.contract.WatchLogs(opts, "ChannelOpened", channel_identifierRule, participant1Rule, participant2Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenNetworkChannelOpened)
				if err := _TokenNetwork.contract.UnpackLog(event, "ChannelOpened", log); err != nil {
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

// TokenNetworkChannelSettledIterator is returned from FilterChannelSettled and is used to iterate over the raw logs and unpacked data for ChannelSettled events raised by the TokenNetwork contract.
type TokenNetworkChannelSettledIterator struct {
	Event *TokenNetworkChannelSettled // Event containing the contract specifics and raw log

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
func (it *TokenNetworkChannelSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenNetworkChannelSettled)
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
		it.Event = new(TokenNetworkChannelSettled)
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
func (it *TokenNetworkChannelSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenNetworkChannelSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenNetworkChannelSettled represents a ChannelSettled event raised by the TokenNetwork contract.
type TokenNetworkChannelSettled struct {
	Channel_identifier  *big.Int
	Participant1_amount *big.Int
	Participant2_amount *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterChannelSettled is a free log retrieval operation binding the contract event 0x0e239ef20c651bd0bc45e6f6a5fd46252d77d39d6602103e347add00cabdb0b4.
//
// Solidity: event ChannelSettled(channel_identifier indexed uint256, participant1_amount uint256, participant2_amount uint256)
func (_TokenNetwork *TokenNetworkFilterer) FilterChannelSettled(opts *bind.FilterOpts, channel_identifier []*big.Int) (*TokenNetworkChannelSettledIterator, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}

	logs, sub, err := _TokenNetwork.contract.FilterLogs(opts, "ChannelSettled", channel_identifierRule)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkChannelSettledIterator{contract: _TokenNetwork.contract, event: "ChannelSettled", logs: logs, sub: sub}, nil
}

// WatchChannelSettled is a free log subscription operation binding the contract event 0x0e239ef20c651bd0bc45e6f6a5fd46252d77d39d6602103e347add00cabdb0b4.
//
// Solidity: event ChannelSettled(channel_identifier indexed uint256, participant1_amount uint256, participant2_amount uint256)
func (_TokenNetwork *TokenNetworkFilterer) WatchChannelSettled(opts *bind.WatchOpts, sink chan<- *TokenNetworkChannelSettled, channel_identifier []*big.Int) (event.Subscription, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}

	logs, sub, err := _TokenNetwork.contract.WatchLogs(opts, "ChannelSettled", channel_identifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenNetworkChannelSettled)
				if err := _TokenNetwork.contract.UnpackLog(event, "ChannelSettled", log); err != nil {
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

// TokenNetworkChannelUnlockedIterator is returned from FilterChannelUnlocked and is used to iterate over the raw logs and unpacked data for ChannelUnlocked events raised by the TokenNetwork contract.
type TokenNetworkChannelUnlockedIterator struct {
	Event *TokenNetworkChannelUnlocked // Event containing the contract specifics and raw log

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
func (it *TokenNetworkChannelUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenNetworkChannelUnlocked)
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
		it.Event = new(TokenNetworkChannelUnlocked)
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
func (it *TokenNetworkChannelUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenNetworkChannelUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenNetworkChannelUnlocked represents a ChannelUnlocked event raised by the TokenNetwork contract.
type TokenNetworkChannelUnlocked struct {
	Channel_identifier *big.Int
	Payer_participant  common.Address
	Locskroot          [32]byte
	Transferred_amount *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterChannelUnlocked is a free log retrieval operation binding the contract event 0x5842365ce79285f3d49939b96a9866db3c41c82754215e2f6e11bcba3d1a61b7.
//
// Solidity: event ChannelUnlocked(channel_identifier uint256, payer_participant address, locskroot bytes32, transferred_amount uint256)
func (_TokenNetwork *TokenNetworkFilterer) FilterChannelUnlocked(opts *bind.FilterOpts) (*TokenNetworkChannelUnlockedIterator, error) {

	logs, sub, err := _TokenNetwork.contract.FilterLogs(opts, "ChannelUnlocked")
	if err != nil {
		return nil, err
	}
	return &TokenNetworkChannelUnlockedIterator{contract: _TokenNetwork.contract, event: "ChannelUnlocked", logs: logs, sub: sub}, nil
}

// WatchChannelUnlocked is a free log subscription operation binding the contract event 0x5842365ce79285f3d49939b96a9866db3c41c82754215e2f6e11bcba3d1a61b7.
//
// Solidity: event ChannelUnlocked(channel_identifier uint256, payer_participant address, locskroot bytes32, transferred_amount uint256)
func (_TokenNetwork *TokenNetworkFilterer) WatchChannelUnlocked(opts *bind.WatchOpts, sink chan<- *TokenNetworkChannelUnlocked) (event.Subscription, error) {

	logs, sub, err := _TokenNetwork.contract.WatchLogs(opts, "ChannelUnlocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenNetworkChannelUnlocked)
				if err := _TokenNetwork.contract.UnpackLog(event, "ChannelUnlocked", log); err != nil {
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

// TokenNetworkChannelwithdrawIterator is returned from FilterChannelwithdraw and is used to iterate over the raw logs and unpacked data for Channelwithdraw events raised by the TokenNetwork contract.
type TokenNetworkChannelwithdrawIterator struct {
	Event *TokenNetworkChannelwithdraw // Event containing the contract specifics and raw log

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
func (it *TokenNetworkChannelwithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenNetworkChannelwithdraw)
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
		it.Event = new(TokenNetworkChannelwithdraw)
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
func (it *TokenNetworkChannelwithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenNetworkChannelwithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenNetworkChannelwithdraw represents a Channelwithdraw event raised by the TokenNetwork contract.
type TokenNetworkChannelwithdraw struct {
	Channel_identifier    *big.Int
	Participant1_deposit  *big.Int
	Participant2_deposit  *big.Int
	Participant1_withdraw *big.Int
	Participant2_withdraw *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterChannelwithdraw is a free log retrieval operation binding the contract event 0x8edee97a023db4a4c6f8f411985d0d55e73a88b28c85867e661fd85d793ece7d.
//
// Solidity: event Channelwithdraw(channel_identifier uint256, participant1_deposit uint256, participant2_deposit uint256, participant1_withdraw uint256, participant2_withdraw uint256)
func (_TokenNetwork *TokenNetworkFilterer) FilterChannelwithdraw(opts *bind.FilterOpts) (*TokenNetworkChannelwithdrawIterator, error) {

	logs, sub, err := _TokenNetwork.contract.FilterLogs(opts, "Channelwithdraw")
	if err != nil {
		return nil, err
	}
	return &TokenNetworkChannelwithdrawIterator{contract: _TokenNetwork.contract, event: "Channelwithdraw", logs: logs, sub: sub}, nil
}

// WatchChannelwithdraw is a free log subscription operation binding the contract event 0x8edee97a023db4a4c6f8f411985d0d55e73a88b28c85867e661fd85d793ece7d.
//
// Solidity: event Channelwithdraw(channel_identifier uint256, participant1_deposit uint256, participant2_deposit uint256, participant1_withdraw uint256, participant2_withdraw uint256)
func (_TokenNetwork *TokenNetworkFilterer) WatchChannelwithdraw(opts *bind.WatchOpts, sink chan<- *TokenNetworkChannelwithdraw) (event.Subscription, error) {

	logs, sub, err := _TokenNetwork.contract.WatchLogs(opts, "Channelwithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenNetworkChannelwithdraw)
				if err := _TokenNetwork.contract.UnpackLog(event, "Channelwithdraw", log); err != nil {
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

// TokenNetworkNonClosingBalanceProofUpdatedIterator is returned from FilterNonClosingBalanceProofUpdated and is used to iterate over the raw logs and unpacked data for NonClosingBalanceProofUpdated events raised by the TokenNetwork contract.
type TokenNetworkNonClosingBalanceProofUpdatedIterator struct {
	Event *TokenNetworkNonClosingBalanceProofUpdated // Event containing the contract specifics and raw log

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
func (it *TokenNetworkNonClosingBalanceProofUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenNetworkNonClosingBalanceProofUpdated)
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
		it.Event = new(TokenNetworkNonClosingBalanceProofUpdated)
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
func (it *TokenNetworkNonClosingBalanceProofUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenNetworkNonClosingBalanceProofUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenNetworkNonClosingBalanceProofUpdated represents a NonClosingBalanceProofUpdated event raised by the TokenNetwork contract.
type TokenNetworkNonClosingBalanceProofUpdated struct {
	Channel_identifier  *big.Int
	Closing_participant common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNonClosingBalanceProofUpdated is a free log retrieval operation binding the contract event 0xe5ccf2144fc46e5dbed7d342686643a19421a58ee918650b15f766f645b8ff07.
//
// Solidity: event NonClosingBalanceProofUpdated(channel_identifier indexed uint256, closing_participant indexed address)
func (_TokenNetwork *TokenNetworkFilterer) FilterNonClosingBalanceProofUpdated(opts *bind.FilterOpts, channel_identifier []*big.Int, closing_participant []common.Address) (*TokenNetworkNonClosingBalanceProofUpdatedIterator, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}
	var closing_participantRule []interface{}
	for _, closing_participantItem := range closing_participant {
		closing_participantRule = append(closing_participantRule, closing_participantItem)
	}

	logs, sub, err := _TokenNetwork.contract.FilterLogs(opts, "NonClosingBalanceProofUpdated", channel_identifierRule, closing_participantRule)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkNonClosingBalanceProofUpdatedIterator{contract: _TokenNetwork.contract, event: "NonClosingBalanceProofUpdated", logs: logs, sub: sub}, nil
}

// WatchNonClosingBalanceProofUpdated is a free log subscription operation binding the contract event 0xe5ccf2144fc46e5dbed7d342686643a19421a58ee918650b15f766f645b8ff07.
//
// Solidity: event NonClosingBalanceProofUpdated(channel_identifier indexed uint256, closing_participant indexed address)
func (_TokenNetwork *TokenNetworkFilterer) WatchNonClosingBalanceProofUpdated(opts *bind.WatchOpts, sink chan<- *TokenNetworkNonClosingBalanceProofUpdated, channel_identifier []*big.Int, closing_participant []common.Address) (event.Subscription, error) {

	var channel_identifierRule []interface{}
	for _, channel_identifierItem := range channel_identifier {
		channel_identifierRule = append(channel_identifierRule, channel_identifierItem)
	}
	var closing_participantRule []interface{}
	for _, closing_participantItem := range closing_participant {
		closing_participantRule = append(closing_participantRule, closing_participantItem)
	}

	logs, sub, err := _TokenNetwork.contract.WatchLogs(opts, "NonClosingBalanceProofUpdated", channel_identifierRule, closing_participantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenNetworkNonClosingBalanceProofUpdated)
				if err := _TokenNetwork.contract.UnpackLog(event, "NonClosingBalanceProofUpdated", log); err != nil {
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

// TokenNetworkRegistryABI is the input ABI used to generate the binding from.
const TokenNetworkRegistryABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"token_to_token_networks\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chain_id\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token_address\",\"type\":\"address\"}],\"name\":\"createERC20TokenNetwork\",\"outputs\":[{\"name\":\"token_network_address\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"contract_address\",\"type\":\"address\"}],\"name\":\"contractExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contract_version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"secret_registry_address\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_secret_registry_address\",\"type\":\"address\"},{\"name\":\"_chain_id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token_address\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token_network_address\",\"type\":\"address\"}],\"name\":\"TokenNetworkCreated\",\"type\":\"event\"}]"

// TokenNetworkRegistryBin is the compiled bytecode used for deploying new contracts.
const TokenNetworkRegistryBin = `0x608060405234801561001057600080fd5b506040516040806131528339810160405280516020909101516000811161003657600080fd5b600160a060020a038216151561004b57600080fd5b61005d82640100000000610091810204565b151561006857600080fd5b60008054600160a060020a031916600160a060020a039390931692909217909155600155610099565b6000903b1190565b6130aa806100a86000396000f3006080604052600436106100775763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630fabd9e7811461007c5780633af973b1146100b95780634cf71a04146100e05780637709bc7814610101578063b32c65c814610136578063d0ad4bec146101c0575b600080fd5b34801561008857600080fd5b5061009d600160a060020a03600435166101d5565b60408051600160a060020a039092168252519081900360200190f35b3480156100c557600080fd5b506100ce6101f0565b60408051918252519081900360200190f35b3480156100ec57600080fd5b5061009d600160a060020a03600435166101f6565b34801561010d57600080fd5b50610122600160a060020a03600435166102e1565b604080519115158252519081900360200190f35b34801561014257600080fd5b5061014b6102e9565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561018557818101518382015260200161016d565b50505050905090810190601f1680156101b25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101cc57600080fd5b5061009d610320565b600260205260009081526040902054600160a060020a031681565b60015481565b600160a060020a038082166000908152600260205260408120549091161561021d57600080fd5b6000546001548391600160a060020a03169061023761032f565b600160a060020a039384168152919092166020820152604080820192909252905190819003606001906000f080158015610275573d6000803e3d6000fd5b50600160a060020a03838116600081815260026020526040808220805473ffffffffffffffffffffffffffffffffffffffff1916948616948517905551939450919290917ff11a7558a113d9627989c5edf26cbd19143b7375248e621c8e30ac9e0847dc3f91a3919050565b6000903b1190565b60408051808201909152600581527f302e332e5f000000000000000000000000000000000000000000000000000000602082015281565b600054600160a060020a031681565b604051612d3f80610340833901905600608060405260006005553480156200001657600080fd5b5060405160608062002d3f833981016040908152815160208301519190920151600160a060020a03831615156200004c57600080fd5b600160a060020a03821615156200006257600080fd5b600081116200007057600080fd5b62000084836401000000006200017c810204565b15156200009057600080fd5b620000a4826401000000006200017c810204565b1515620000b057600080fd5b60008054600160a060020a03808616600160a060020a031992831617808455600180548784169416939093179092556002849055604080517f18160ddd000000000000000000000000000000000000000000000000000000008152905192909116916318160ddd9160048082019260209290919082900301818787803b1580156200013a57600080fd5b505af11580156200014f573d6000803e3d6000fd5b505050506040513d60208110156200016657600080fd5b5051116200017357600080fd5b50505062000184565b6000903b1190565b612bab80620001946000396000f3006080604052600436106100f85763ffffffff60e060020a6000350416630a798f2481146100fd5780631f466acf146101295780632419a2e01461015057806324d73a931461016857806328a8d70f1461019957806335fefe7d1461024b5780633af973b11461030d5780637709bc78146103225780637944bd1f146103575780637c0f097c1461037e5780637fb5885e146103ea5780638568536a1461043b5780638d373f51146104f2578063b32c65c8146105ae578063bb94ccc814610638578063dee8bd0a1461065f578063e5949b5d14610686578063f3da17e8146106c7578063f94c9e1314610730578063fc0c546a14610788575b600080fd5b34801561010957600080fd5b50610127600160a060020a036004358116906024351660443561079d565b005b34801561013557600080fd5b5061013e61096f565b60408051918252519081900360200190f35b34801561015c57600080fd5b5061013e600435610975565b34801561017457600080fd5b5061017d610987565b60408051600160a060020a039092168252519081900360200190f35b3480156101a557600080fd5b50604080516020600460843581810135601f81018490048402850184019095528484526101279482359460248035600160a060020a03169560443595606435953695919460a494909391019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506109969650505050505050565b34801561025757600080fd5b50604080516020601f60c43560048181013592830184900484028501840190955281845261012794600160a060020a038135811695602480359092169560443595606435956084359560a435953695919460e49492939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497505093359450610aa09350505050565b34801561031957600080fd5b5061013e611172565b34801561032e57600080fd5b50610343600160a060020a0360043516611178565b604080519115158252519081900360200190f35b34801561036357600080fd5b50610127600435600160a060020a0360243516604435611180565b34801561038a57600080fd5b50604080516020600460a43581810135601f81018490048402850184019095528484526101279482359460248035956044359560643595608435953695929460c494920191819084018382808284375094975061130b9650505050505050565b3480156103f657600080fd5b5061040e600435600160a060020a0360243516611452565b60408051951515865260208601949094528484019290925260608401526080830152519081900360a00190f35b34801561044757600080fd5b50604080516020601f60843560048181013592830184900484028501840190955281845261012794600160a060020a0381358116956024803596604435909316956064359536959460a49493919091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506114999650505050505050565b3480156104fe57600080fd5b50604080516020600460c43581810135601f81018490048402850184019095528484526101279482359460248035600160a060020a03169560443595606435956084359560a435953695939460e4949193919091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506118a59650505050505050565b3480156105ba57600080fd5b506105c3611996565b6040805160208082528351818301528351919283929083019185019080838360005b838110156105fd5781810151838201526020016105e5565b50505050905090810190601f16801561062a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561064457600080fd5b5061013e600160a060020a03600435811690602435166119cd565b34801561066b57600080fd5b50610127600160a060020a0360043581169060243516611b30565b34801561069257600080fd5b5061069e600435611f16565b60408051938452600160a060020a03909216602084015260ff1682820152519081900360600190f35b3480156106d357600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526101279482359460248035600160a060020a031695369594606494920191908190840183828082843750949750611f449650505050505050565b34801561073c57600080fd5b50610757600160a060020a0360043581169060243516612051565b604080519485526020850193909352600160a060020a039091168383015260ff166060830152519081900360800190f35b34801561079457600080fd5b5061017d6120b2565b600080600080600085600681101580156107ba5750622932e08111155b15156107c557600080fd5b600160a060020a03891615156107da57600080fd5b600160a060020a03881615156107ef57600080fd5b600160a060020a03898116908916141561080857600080fd5b600160056000828254019250508190555060055495506003600087815260200190815260200160002093508360020160008a600160a060020a0316600160a060020a03168152602001908152602001600020925083600201600089600160a060020a0316600160a060020a03168152602001908152602001600020915086846000018190555060018360040160006101000a81548160ff02191690831515021790555060018260040160006101000a81548160ff02191690831515021790555060018460010160146101000a81548160ff021916908360ff1602179055506108f089896119cd565b6000818152600460205260409020549095501561090c57600080fd5b60008581526004602090815260409182902088905581518981529151600160a060020a03808c1693908d16928a927f669a4b0ac0b9994c0f82ed4dbe07bb421fe74e5951725af4f139c7443ebf049d9281900390910190a4505050505050505050565b60055481565b60046020526000908152604090205481565b600154600160a060020a031681565b60008681526003602052604081206001015481908190819081908b9060a060020a900460ff166002146109c857600080fd5b60008c8152600360209081526040808320600160a060020a038f168452600281019092529091206004810154919550935060ff161515610a0757600080fd5b60018301541515610a1757600080fd5b610a238c8b8b8b6120c1565b600160a060020a038116600090815260028601602052604090206004810154919750925060ff161515610a5557600080fd5b600160a060020a03868116908c161415610a6e57600080fd5b610a788a88612178565b60018401549095508514610a8b57600080fd5b50506000600290910155505050505050505050565b600081815260036020526040812060018082015483918291829160a060020a90910460ff1614610acf57600080fd5b8d8c8e8d8d8a306002546040516020018089600160a060020a0316600160a060020a0316606060020a02815260140188815260200187600160a060020a0316600160a060020a0316606060020a02815260140186815260200185815260200184815260200183600160a060020a0316600160a060020a0316606060020a028152601401828152602001985050505050505050506040516020818303038152906040526040518082805190602001908083835b60208310610ba05780518252601f199092019160209182019101610b81565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209250610bd983896122c9565b600160a060020a038f8116911614610bf057600080fd5b8d8c8e8d8d8d8b30600254604051602001808a600160a060020a0316600160a060020a0316606060020a02815260140189815260200188600160a060020a0316600160a060020a0316606060020a02815260140187815260200186815260200185815260200184815260200183600160a060020a0316600160a060020a0316606060020a02815260140182815260200199505050505050505050506040516020818303038152906040526040518082805190602001908083835b60208310610cc95780518252601f199092019160209182019101610caa565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209250610d0283886122c9565b600160a060020a038e8116911614610d1957600080fd5b5050600160a060020a03808d166000908152600284016020526040808220928e1682529020600482015460ff161515610d5157600080fd5b600481015460ff161515610d6457600080fd5b80548254019450848c1115610d7857600080fd5b848b1115610d8557600080fd5b8b8b018514610d9357600080fd5b60008a1115610e4a576000809054906101000a9004600160a060020a0316600160a060020a031663a9059cbb8f8c6040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b158015610e1357600080fd5b505af1158015610e27573d6000803e3d6000fd5b505050506040513d6020811015610e3d57600080fd5b50511515610e4a57600080fd5b6000891115610f01576000809054906101000a9004600160a060020a0316600160a060020a031663a9059cbb8e8b6040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b158015610eca57600080fd5b505af1158015610ede573d6000803e3d6000fd5b505050506040513d6020811015610ef457600080fd5b50511515610f0157600080fd5b8b8a1115610f0e57600080fd5b8a891115610f1b57600080fd5b898c039b50888b039a508b82600001819055508a81600001819055508360020160008f600160a060020a0316600160a060020a031681526020019081526020016000206000808201600090556001820160009055600282016000905560038201600090556004820160006101000a81549060ff0219169055600582016000905550508360020160008e600160a060020a0316600160a060020a031681526020019081526020016000206000808201600090556001820160009055600282016000905560038201600090556004820160006101000a81549060ff021916905560058201600090555050600360008781526020019081526020016000206000808201600090556001820160006101000a815490600160a060020a0302191690556001820160146101000a81549060ff02191690555050600160056000828254019250508190555060055495508360036000888152602001908152602001600020600082015481600001556001820160009054906101000a9004600160a060020a03168160010160006101000a815481600160a060020a030219169083600160a060020a031602179055506001820160149054906101000a900460ff168160010160146101000a81548160ff021916908360ff1602179055509050506110fe8e8e6119cd565b60008181526004602090815260409182902089905581518981529081018f90528082018e9052606081018d9052608081018c905290519194507f8edee97a023db4a4c6f8f411985d0d55e73a88b28c85867e661fd85d793ece7d919081900360a00190a15050505050505050505050505050565b60025481565b6000903b1190565b600083815260036020526040812060019081015482918291879160a060020a90910460ff16146111af57600080fd5b600085116111bc57600080fd5b6000878152600360209081526040808320600160a060020a038a168452600281019092529091206004810154919450925060ff1615156111fb57600080fd5b8154851161120857600080fd5b8154808603908101835560008054604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018590529051939750600160a060020a03909116926323b872dd92606480840193602093929083900390910190829087803b15801561128a57600080fd5b505af115801561129e573d6000803e3d6000fd5b505050506040513d60208110156112b457600080fd5b505115156112c157600080fd5b81546040805191825251600160a060020a0388169189917f2b55547a3b586ab51f65ee9ce4927fa6d25191388299988e89e059a02f9dd4459181900360200190a350505050505050565b60008681526003602052604081206001908101548291899160a060020a900460ff161461133757600080fd5b600089815260036020908152604080832060018101805474ff0000000000000000000000000000000000000000191674020000000000000000000000000000000000000000179055338452600281019092529091206004015490925060ff1615156113a157600080fd5b60018201805473ffffffffffffffffffffffffffffffffffffffff191633179055815443018255600086111561141a576113df8989898989896123a9565b92506113ee8984888a8c612473565b600160a060020a038316600090815260028301602052604090206004015460ff16151561141a57600080fd5b60405133908a907fa8621c489a70a0a06448f2b4e3477913a3744d5f27a380e5f0d8db13837ce7c690600090a3505050505050505050565b6000918252600360208181526040808520600160a060020a03949094168552600293840190915290922060048101548154600183015493830154929094015460ff90911694565b60008060008060008060006114ae8d8c6119cd565b6000818152600460209081526040808320548084526003909252909120919650945092506114e0848e8e8e8e8e6124c4565b9650600160a060020a038d8116908816146114fa57600080fd5b611508848e8e8e8e8d6124c4565b9650600160a060020a038b81169088161461152257600080fd5b5050600160a060020a03808c166000908152600283016020526040808220928c1682529020600482015460ff16151561155a57600080fd5b600481015460ff16151561156d57600080fd5b60018381015460a060020a900460ff161461158757600080fd5b806000015482600001540195508260020160008e600160a060020a0316600160a060020a031681526020019081526020016000206000808201600090556001820160009055600282016000905560038201600090556004820160006101000a81549060ff0219169055600582016000905550508260020160008c600160a060020a0316600160a060020a031681526020019081526020016000206000808201600090556001820160009055600282016000905560038201600090556004820160006101000a81549060ff021916905560058201600090555050600360008581526020019081526020016000206000808201600090556001820160006101000a815490600160a060020a0302191690556001820160146101000a81549060ff021916905550506004600086600019166000191681526020019081526020016000206000905560008c1115611782576000809054906101000a9004600160a060020a0316600160a060020a031663a9059cbb8e8e6040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b15801561174b57600080fd5b505af115801561175f573d6000803e3d6000fd5b505050506040513d602081101561177557600080fd5b5051151561178257600080fd5b60008a11156118345760008054604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a038f81166004830152602482018f90529151919092169263a9059cbb92604480820193602093909283900390910190829087803b1580156117fd57600080fd5b505af1158015611811573d6000803e3d6000fd5b505050506040513d602081101561182757600080fd5b5051151561183457600080fd5b8b8a01861461184257600080fd5b8b86101561184f57600080fd5b8986101561185c57600080fd5b604080518d8152602081018c9052815186927f0e239ef20c651bd0bc45e6f6a5fd46252d77d39d6602103e347add00cabdb0b4928290030190a250505050505050505050505050565b6000808086116118b457600080fd5b505060008881526003602052604090206001810154600160a060020a0316906118e28a888a8989898961255d565b600160a060020a038a81169116146118f957600080fd5b6119078a888a8989896123a9565b600160a060020a0383811691161461191e57600080fd5b61192b8a83888b8b612473565b604051600160a060020a038316908b907fe5ccf2144fc46e5dbed7d342686643a19421a58ee918650b15f766f645b8ff0790600090a3600181015460a060020a900460ff1660021461197c57600080fd5b805443111561198a57600080fd5b50505050505050505050565b60408051808201909152600581527f302e332e5f000000000000000000000000000000000000000000000000000000602082015281565b600081600160a060020a031683600160a060020a03161015611aa45782826040516020018083600160a060020a0316600160a060020a0316606060020a02815260140182600160a060020a0316600160a060020a0316606060020a028152601401925050506040516020818303038152906040526040518082805190602001908083835b60208310611a705780518252601f199092019160209182019101611a51565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209050611b2a565b81836040516020018083600160a060020a0316600160a060020a0316606060020a02815260140182600160a060020a0316600160a060020a0316606060020a0281526014019250505060405160208183030381529060405260405180828051906020019080838360208310611a705780518252601f199092019160209182019101611a51565b92915050565b600080600080600080600080600080611b498c8c6119cd565b6000818152600460209081526040808320548084526003909252909120600181015491975091955090935060a060020a900460ff16600214611b8a57600080fd5b82544311611b9757600080fd5b5050600160a060020a03808b166000908152600283016020526040808220928c1682529020600482015460ff161515611bcf57600080fd5b600481015460ff161515611be257600080fd5b8054825460058085015460028087015492860154908601549290910191018083018281039e50919b50995091019650881115611c1d57600099505b611c278a8761269e565b995089860398508260020160008d600160a060020a0316600160a060020a031681526020019081526020016000206000808201600090556001820160009055600282016000905560038201600090556004820160006101000a81549060ff0219169055600582016000905550508260020160008c600160a060020a0316600160a060020a031681526020019081526020016000206000808201600090556001820160009055600282016000905560038201600090556004820160006101000a81549060ff021916905560058201600090555050600360008681526020019081526020016000206000808201600090556001820160006101000a815490600160a060020a0302191690556001820160146101000a81549060ff021916905550506004600085600019166000191681526020019081526020016000206000905560008a1115611e1c576000809054906101000a9004600160a060020a0316600160a060020a031663a9059cbb8d8c6040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b158015611de557600080fd5b505af1158015611df9573d6000803e3d6000fd5b505050506040513d6020811015611e0f57600080fd5b50511515611e1c57600080fd5b6000891115611ece5760008054604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a038f81166004830152602482018e90529151919092169263a9059cbb92604480820193602093909283900390910190829087803b158015611e9757600080fd5b505af1158015611eab573d6000803e3d6000fd5b505050506040513d6020811015611ec157600080fd5b50511515611ece57600080fd5b604080518b8152602081018b9052815187927f0e239ef20c651bd0bc45e6f6a5fd46252d77d39d6602103e347add00cabdb0b4928290030190a2505050505050505050505050565b60036020526000908152604090208054600190910154600160a060020a0381169060a060020a900460ff1683565b600083815260036020526040812060010154819081908190879060a060020a900460ff16600214611f7457600080fd5b8551600010611f8257600080fd5b60008881526003602052604090208054909350431115611fa157600080fd5b600160a060020a0387166000908152600284016020526040902060018101549092501515611fce57600080fd5b611fd7866126b6565b600184015491965094508514611fec57600080fd5b600582018490556002820154604080518a8152600160a060020a038a1660208201528082018890526060810192909252517f5842365ce79285f3d49939b96a9866db3c41c82754215e2f6e11bcba3d1a61b79181900360800190a15050505050505050565b600080600080600080600061206689896119cd565b60009081526004602090815260408083205480845260039092529091208054600190910154919b909a50600160a060020a038216995060a060020a90910460ff16975095505050505050565b600054600160a060020a031681565b600254604080516020808201879052818301889052606060020a30026060830152607482019390935260948082018690528251808303909101815260b490910191829052805160009384939182918401908083835b602083106121355780518252601f199092019160209182019101612116565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020905061216e81846122c9565b9695505050505050565b60008060006020845181151561218a57fe5b061561219557600080fd5b602091505b835182116122c05750828101518085101561223457604080516020808201889052818301849052825180830384018152606090920192839052815191929182918401908083835b602083106122005780518252601f1990920191602091820191016121e1565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902094506122b5565b604080516020808201849052818301889052825180830384018152606090920192839052815191929182918401908083835b602083106122855780518252601f199092019160209182019101612266565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902094505b60208201915061219a565b50929392505050565b600080600080845160411415156122df57600080fd5b50505060208201516040830151606084015160001a601b60ff8216101561230457601b015b8060ff16601b148061231957508060ff16601c145b151561232457600080fd5b60408051600080825260208083018085528a905260ff8516838501526060830187905260808301869052925160019360a0808501949193601f19840193928390039091019190865af115801561237e573d6000803e3d6000fd5b5050604051601f190151945050600160a060020a03841615156123a057600080fd5b50505092915050565b600254604080516020808201899052818301889052606082018790526080820186905260a082018a9052606060020a300260c083015260d4808301949094528251808303909401845260f490910191829052825160009384939092909182918401908083835b6020831061242e5780518252601f19909201916020918201910161240f565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020905061246781846122c9565b98975050505050505050565b6000858152600360208181526040808420600160a060020a0389168552600281019092529092209081015485116124a957600080fd5b60038101949094555060018301919091556002909101555050565b60025460408051606060020a600160a060020a03808a168202602080850191909152603484018a90529088168202605484015260688301879052608883018b9052309190910260a883015260bc808301949094528251808303909401845260dc90910191829052825160009384939092909182918401908083836020831061242e5780518252601f19909201916020918201910161240f565b600080878787878c306002548a604051602001808981526020018860001916600019168152602001878152602001866000191660001916815260200185815260200184600160a060020a0316600160a060020a0316606060020a02815260140183815260200182805190602001908083835b602083106125ee5780518252601f1990920191602091820191016125cf565b6001836020036101000a038019825116818451168082178552505050505050905001985050505050505050506040516020818303038152906040526040518082805190602001908083835b602083106126585780518252601f199092019160209182019101612639565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020905061269181846122c9565b9998505050505050505050565b60008183116126ad57826126af565b815b9392505050565b8051600090819081808080806060808706156126d157600080fd5b60608704600101604051908082528060200260200182016040528015612701578160200160208202803883390190505b509050602095505b8686101561274b5761271b8a87612a0a565b958601959450925082816060880481518110151561273557fe5b6020908102909101015260609590950194612709565b6060870496505b60018711156129e157600287061561279f57806001880381518110151561277557fe5b90602001906020020151818881518110151561278d57fe5b60209081029091010152600196909601955b600095505b600187038610156129d65780866001018151811015156127c057fe5b6020908102909101015181518290889081106127d857fe5b6020908102909101015114156128075780868151811015156127f657fe5b9060200190602002015192506129ae565b808660010181518110151561281857fe5b60209081029091010151815182908890811061283057fe5b6020908102909101015110156128f957808681518110151561284e57fe5b90602001906020020151818760010181518110151561286957fe5b6020908102909101810151604080518084019490945283810191909152805180840382018152606090930190819052825190918291908401908083835b602083106128c55780518252601f1990920191602091820191016128a6565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902092506129ae565b808660010181518110151561290a57fe5b90602001906020020151818781518110151561292257fe5b6020908102909101810151604080518084019490945283810191909152805180840382018152606090930190819052825190918291908401908083835b6020831061297e5780518252601f19909201916020918201910161295f565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902092505b8281600288048151811015156129c057fe5b60209081029091010152600295909501946127a4565b600286049650612752565b8060008151811015156129f057fe5b602090810290910101519a94995093975050505050505050565b6000806000806000806000878951111515612a2b5795506000945085612b73565b888801805160208083015160409384015184518084018590528086018390526060808201839052865180830390910181526080909101958690528051949a509198509550929182918401908083835b60208310612a995780518252601f199092019160209182019101612a7a565b51815160209384036101000a6000190180199092169116179052604080519290940182900382206001547fc1f62946000000000000000000000000000000000000000000000000000000008452600484018a90529451909750600160a060020a03909416955063c1f62946945060248083019491935090918290030181600087803b158015612b2757600080fd5b505af1158015612b3b573d6000803e3d6000fd5b505050506040513d6020811015612b5157600080fd5b50519250821580612b625750828511155b15612b6c57600093505b8084965096505b505050505092509290505600a165627a7a723058203ecb22ff97e8f409cff9befc0018d56a3b104a0b4254814e20c238d739a3fb260029a165627a7a723058200333bb46cb780d63557dccdc9c31b975cc7d06ba3dff471e16792844619e958d0029`

// DeployTokenNetworkRegistry deploys a new Ethereum contract, binding an instance of TokenNetworkRegistry to it.
func DeployTokenNetworkRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _secret_registry_address common.Address, _chain_id *big.Int) (common.Address, *types.Transaction, *TokenNetworkRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenNetworkRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenNetworkRegistryBin), backend, _secret_registry_address, _chain_id)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenNetworkRegistry{TokenNetworkRegistryCaller: TokenNetworkRegistryCaller{contract: contract}, TokenNetworkRegistryTransactor: TokenNetworkRegistryTransactor{contract: contract}, TokenNetworkRegistryFilterer: TokenNetworkRegistryFilterer{contract: contract}}, nil
}

// TokenNetworkRegistry is an auto generated Go binding around an Ethereum contract.
type TokenNetworkRegistry struct {
	TokenNetworkRegistryCaller     // Read-only binding to the contract
	TokenNetworkRegistryTransactor // Write-only binding to the contract
	TokenNetworkRegistryFilterer   // Log filterer for contract events
}

// TokenNetworkRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenNetworkRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenNetworkRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenNetworkRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenNetworkRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenNetworkRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenNetworkRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenNetworkRegistrySession struct {
	Contract     *TokenNetworkRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TokenNetworkRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenNetworkRegistryCallerSession struct {
	Contract *TokenNetworkRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// TokenNetworkRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenNetworkRegistryTransactorSession struct {
	Contract     *TokenNetworkRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// TokenNetworkRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenNetworkRegistryRaw struct {
	Contract *TokenNetworkRegistry // Generic contract binding to access the raw methods on
}

// TokenNetworkRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenNetworkRegistryCallerRaw struct {
	Contract *TokenNetworkRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// TokenNetworkRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenNetworkRegistryTransactorRaw struct {
	Contract *TokenNetworkRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenNetworkRegistry creates a new instance of TokenNetworkRegistry, bound to a specific deployed contract.
func NewTokenNetworkRegistry(address common.Address, backend bind.ContractBackend) (*TokenNetworkRegistry, error) {
	contract, err := bindTokenNetworkRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkRegistry{TokenNetworkRegistryCaller: TokenNetworkRegistryCaller{contract: contract}, TokenNetworkRegistryTransactor: TokenNetworkRegistryTransactor{contract: contract}, TokenNetworkRegistryFilterer: TokenNetworkRegistryFilterer{contract: contract}}, nil
}

// NewTokenNetworkRegistryCaller creates a new read-only instance of TokenNetworkRegistry, bound to a specific deployed contract.
func NewTokenNetworkRegistryCaller(address common.Address, caller bind.ContractCaller) (*TokenNetworkRegistryCaller, error) {
	contract, err := bindTokenNetworkRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkRegistryCaller{contract: contract}, nil
}

// NewTokenNetworkRegistryTransactor creates a new write-only instance of TokenNetworkRegistry, bound to a specific deployed contract.
func NewTokenNetworkRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenNetworkRegistryTransactor, error) {
	contract, err := bindTokenNetworkRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkRegistryTransactor{contract: contract}, nil
}

// NewTokenNetworkRegistryFilterer creates a new log filterer instance of TokenNetworkRegistry, bound to a specific deployed contract.
func NewTokenNetworkRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenNetworkRegistryFilterer, error) {
	contract, err := bindTokenNetworkRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkRegistryFilterer{contract: contract}, nil
}

// bindTokenNetworkRegistry binds a generic wrapper to an already deployed contract.
func bindTokenNetworkRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenNetworkRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenNetworkRegistry *TokenNetworkRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenNetworkRegistry.Contract.TokenNetworkRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenNetworkRegistry *TokenNetworkRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenNetworkRegistry.Contract.TokenNetworkRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenNetworkRegistry *TokenNetworkRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenNetworkRegistry.Contract.TokenNetworkRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenNetworkRegistry *TokenNetworkRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenNetworkRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenNetworkRegistry *TokenNetworkRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenNetworkRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenNetworkRegistry *TokenNetworkRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenNetworkRegistry.Contract.contract.Transact(opts, method, params...)
}

// Chain_id is a free data retrieval call binding the contract method 0x3af973b1.
//
// Solidity: function chain_id() constant returns(uint256)
func (_TokenNetworkRegistry *TokenNetworkRegistryCaller) Chain_id(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenNetworkRegistry.contract.Call(opts, out, "chain_id")
	return *ret0, err
}

// Chain_id is a free data retrieval call binding the contract method 0x3af973b1.
//
// Solidity: function chain_id() constant returns(uint256)
func (_TokenNetworkRegistry *TokenNetworkRegistrySession) Chain_id() (*big.Int, error) {
	return _TokenNetworkRegistry.Contract.Chain_id(&_TokenNetworkRegistry.CallOpts)
}

// Chain_id is a free data retrieval call binding the contract method 0x3af973b1.
//
// Solidity: function chain_id() constant returns(uint256)
func (_TokenNetworkRegistry *TokenNetworkRegistryCallerSession) Chain_id() (*big.Int, error) {
	return _TokenNetworkRegistry.Contract.Chain_id(&_TokenNetworkRegistry.CallOpts)
}

// ContractExists is a free data retrieval call binding the contract method 0x7709bc78.
//
// Solidity: function contractExists(contract_address address) constant returns(bool)
func (_TokenNetworkRegistry *TokenNetworkRegistryCaller) ContractExists(opts *bind.CallOpts, contract_address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenNetworkRegistry.contract.Call(opts, out, "contractExists", contract_address)
	return *ret0, err
}

// ContractExists is a free data retrieval call binding the contract method 0x7709bc78.
//
// Solidity: function contractExists(contract_address address) constant returns(bool)
func (_TokenNetworkRegistry *TokenNetworkRegistrySession) ContractExists(contract_address common.Address) (bool, error) {
	return _TokenNetworkRegistry.Contract.ContractExists(&_TokenNetworkRegistry.CallOpts, contract_address)
}

// ContractExists is a free data retrieval call binding the contract method 0x7709bc78.
//
// Solidity: function contractExists(contract_address address) constant returns(bool)
func (_TokenNetworkRegistry *TokenNetworkRegistryCallerSession) ContractExists(contract_address common.Address) (bool, error) {
	return _TokenNetworkRegistry.Contract.ContractExists(&_TokenNetworkRegistry.CallOpts, contract_address)
}

// Contract_version is a free data retrieval call binding the contract method 0xb32c65c8.
//
// Solidity: function contract_version() constant returns(string)
func (_TokenNetworkRegistry *TokenNetworkRegistryCaller) Contract_version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TokenNetworkRegistry.contract.Call(opts, out, "contract_version")
	return *ret0, err
}

// Contract_version is a free data retrieval call binding the contract method 0xb32c65c8.
//
// Solidity: function contract_version() constant returns(string)
func (_TokenNetworkRegistry *TokenNetworkRegistrySession) Contract_version() (string, error) {
	return _TokenNetworkRegistry.Contract.Contract_version(&_TokenNetworkRegistry.CallOpts)
}

// Contract_version is a free data retrieval call binding the contract method 0xb32c65c8.
//
// Solidity: function contract_version() constant returns(string)
func (_TokenNetworkRegistry *TokenNetworkRegistryCallerSession) Contract_version() (string, error) {
	return _TokenNetworkRegistry.Contract.Contract_version(&_TokenNetworkRegistry.CallOpts)
}

// Secret_registry_address is a free data retrieval call binding the contract method 0xd0ad4bec.
//
// Solidity: function secret_registry_address() constant returns(address)
func (_TokenNetworkRegistry *TokenNetworkRegistryCaller) Secret_registry_address(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenNetworkRegistry.contract.Call(opts, out, "secret_registry_address")
	return *ret0, err
}

// Secret_registry_address is a free data retrieval call binding the contract method 0xd0ad4bec.
//
// Solidity: function secret_registry_address() constant returns(address)
func (_TokenNetworkRegistry *TokenNetworkRegistrySession) Secret_registry_address() (common.Address, error) {
	return _TokenNetworkRegistry.Contract.Secret_registry_address(&_TokenNetworkRegistry.CallOpts)
}

// Secret_registry_address is a free data retrieval call binding the contract method 0xd0ad4bec.
//
// Solidity: function secret_registry_address() constant returns(address)
func (_TokenNetworkRegistry *TokenNetworkRegistryCallerSession) Secret_registry_address() (common.Address, error) {
	return _TokenNetworkRegistry.Contract.Secret_registry_address(&_TokenNetworkRegistry.CallOpts)
}

// Token_to_token_networks is a free data retrieval call binding the contract method 0x0fabd9e7.
//
// Solidity: function token_to_token_networks( address) constant returns(address)
func (_TokenNetworkRegistry *TokenNetworkRegistryCaller) Token_to_token_networks(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenNetworkRegistry.contract.Call(opts, out, "token_to_token_networks", arg0)
	return *ret0, err
}

// Token_to_token_networks is a free data retrieval call binding the contract method 0x0fabd9e7.
//
// Solidity: function token_to_token_networks( address) constant returns(address)
func (_TokenNetworkRegistry *TokenNetworkRegistrySession) Token_to_token_networks(arg0 common.Address) (common.Address, error) {
	return _TokenNetworkRegistry.Contract.Token_to_token_networks(&_TokenNetworkRegistry.CallOpts, arg0)
}

// Token_to_token_networks is a free data retrieval call binding the contract method 0x0fabd9e7.
//
// Solidity: function token_to_token_networks( address) constant returns(address)
func (_TokenNetworkRegistry *TokenNetworkRegistryCallerSession) Token_to_token_networks(arg0 common.Address) (common.Address, error) {
	return _TokenNetworkRegistry.Contract.Token_to_token_networks(&_TokenNetworkRegistry.CallOpts, arg0)
}

// CreateERC20TokenNetwork is a paid mutator transaction binding the contract method 0x4cf71a04.
//
// Solidity: function createERC20TokenNetwork(_token_address address) returns(token_network_address address)
func (_TokenNetworkRegistry *TokenNetworkRegistryTransactor) CreateERC20TokenNetwork(opts *bind.TransactOpts, _token_address common.Address) (*types.Transaction, error) {
	return _TokenNetworkRegistry.contract.Transact(opts, "createERC20TokenNetwork", _token_address)
}

// CreateERC20TokenNetwork is a paid mutator transaction binding the contract method 0x4cf71a04.
//
// Solidity: function createERC20TokenNetwork(_token_address address) returns(token_network_address address)
func (_TokenNetworkRegistry *TokenNetworkRegistrySession) CreateERC20TokenNetwork(_token_address common.Address) (*types.Transaction, error) {
	return _TokenNetworkRegistry.Contract.CreateERC20TokenNetwork(&_TokenNetworkRegistry.TransactOpts, _token_address)
}

// CreateERC20TokenNetwork is a paid mutator transaction binding the contract method 0x4cf71a04.
//
// Solidity: function createERC20TokenNetwork(_token_address address) returns(token_network_address address)
func (_TokenNetworkRegistry *TokenNetworkRegistryTransactorSession) CreateERC20TokenNetwork(_token_address common.Address) (*types.Transaction, error) {
	return _TokenNetworkRegistry.Contract.CreateERC20TokenNetwork(&_TokenNetworkRegistry.TransactOpts, _token_address)
}

// TokenNetworkRegistryTokenNetworkCreatedIterator is returned from FilterTokenNetworkCreated and is used to iterate over the raw logs and unpacked data for TokenNetworkCreated events raised by the TokenNetworkRegistry contract.
type TokenNetworkRegistryTokenNetworkCreatedIterator struct {
	Event *TokenNetworkRegistryTokenNetworkCreated // Event containing the contract specifics and raw log

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
func (it *TokenNetworkRegistryTokenNetworkCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenNetworkRegistryTokenNetworkCreated)
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
		it.Event = new(TokenNetworkRegistryTokenNetworkCreated)
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
func (it *TokenNetworkRegistryTokenNetworkCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenNetworkRegistryTokenNetworkCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenNetworkRegistryTokenNetworkCreated represents a TokenNetworkCreated event raised by the TokenNetworkRegistry contract.
type TokenNetworkRegistryTokenNetworkCreated struct {
	Token_address         common.Address
	Token_network_address common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterTokenNetworkCreated is a free log retrieval operation binding the contract event 0xf11a7558a113d9627989c5edf26cbd19143b7375248e621c8e30ac9e0847dc3f.
//
// Solidity: event TokenNetworkCreated(token_address indexed address, token_network_address indexed address)
func (_TokenNetworkRegistry *TokenNetworkRegistryFilterer) FilterTokenNetworkCreated(opts *bind.FilterOpts, token_address []common.Address, token_network_address []common.Address) (*TokenNetworkRegistryTokenNetworkCreatedIterator, error) {

	var token_addressRule []interface{}
	for _, token_addressItem := range token_address {
		token_addressRule = append(token_addressRule, token_addressItem)
	}
	var token_network_addressRule []interface{}
	for _, token_network_addressItem := range token_network_address {
		token_network_addressRule = append(token_network_addressRule, token_network_addressItem)
	}

	logs, sub, err := _TokenNetworkRegistry.contract.FilterLogs(opts, "TokenNetworkCreated", token_addressRule, token_network_addressRule)
	if err != nil {
		return nil, err
	}
	return &TokenNetworkRegistryTokenNetworkCreatedIterator{contract: _TokenNetworkRegistry.contract, event: "TokenNetworkCreated", logs: logs, sub: sub}, nil
}

// WatchTokenNetworkCreated is a free log subscription operation binding the contract event 0xf11a7558a113d9627989c5edf26cbd19143b7375248e621c8e30ac9e0847dc3f.
//
// Solidity: event TokenNetworkCreated(token_address indexed address, token_network_address indexed address)
func (_TokenNetworkRegistry *TokenNetworkRegistryFilterer) WatchTokenNetworkCreated(opts *bind.WatchOpts, sink chan<- *TokenNetworkRegistryTokenNetworkCreated, token_address []common.Address, token_network_address []common.Address) (event.Subscription, error) {

	var token_addressRule []interface{}
	for _, token_addressItem := range token_address {
		token_addressRule = append(token_addressRule, token_addressItem)
	}
	var token_network_addressRule []interface{}
	for _, token_network_addressItem := range token_network_address {
		token_network_addressRule = append(token_network_addressRule, token_network_addressItem)
	}

	logs, sub, err := _TokenNetworkRegistry.contract.WatchLogs(opts, "TokenNetworkCreated", token_addressRule, token_network_addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenNetworkRegistryTokenNetworkCreated)
				if err := _TokenNetworkRegistry.contract.UnpackLog(event, "TokenNetworkCreated", log); err != nil {
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

// UtilsABI is the input ABI used to generate the binding from.
const UtilsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"contract_address\",\"type\":\"address\"}],\"name\":\"contractExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contract_version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// UtilsBin is the compiled bytecode used for deploying new contracts.
const UtilsBin = `0x608060405234801561001057600080fd5b50610187806100206000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416637709bc788114610050578063b32c65c814610092575b600080fd5b34801561005c57600080fd5b5061007e73ffffffffffffffffffffffffffffffffffffffff6004351661011c565b604080519115158252519081900360200190f35b34801561009e57600080fd5b506100a7610124565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100e15781810151838201526020016100c9565b50505050905090810190601f16801561010e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6000903b1190565b60408051808201909152600581527f302e332e5f0000000000000000000000000000000000000000000000000000006020820152815600a165627a7a7230582063bfdb817be794b9a6e6367ae4c51f7c7e599062ea99b601f169f887d4d8ca1b0029`

// DeployUtils deploys a new Ethereum contract, binding an instance of Utils to it.
func DeployUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Utils, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// Utils is an auto generated Go binding around an Ethereum contract.
type Utils struct {
	UtilsCaller     // Read-only binding to the contract
	UtilsTransactor // Write-only binding to the contract
	UtilsFilterer   // Log filterer for contract events
}

// UtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type UtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UtilsSession struct {
	Contract     *Utils            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UtilsCallerSession struct {
	Contract *UtilsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UtilsTransactorSession struct {
	Contract     *UtilsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type UtilsRaw struct {
	Contract *Utils // Generic contract binding to access the raw methods on
}

// UtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UtilsCallerRaw struct {
	Contract *UtilsCaller // Generic read-only contract binding to access the raw methods on
}

// UtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UtilsTransactorRaw struct {
	Contract *UtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUtils creates a new instance of Utils, bound to a specific deployed contract.
func NewUtils(address common.Address, backend bind.ContractBackend) (*Utils, error) {
	contract, err := bindUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// NewUtilsCaller creates a new read-only instance of Utils, bound to a specific deployed contract.
func NewUtilsCaller(address common.Address, caller bind.ContractCaller) (*UtilsCaller, error) {
	contract, err := bindUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsCaller{contract: contract}, nil
}

// NewUtilsTransactor creates a new write-only instance of Utils, bound to a specific deployed contract.
func NewUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*UtilsTransactor, error) {
	contract, err := bindUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsTransactor{contract: contract}, nil
}

// NewUtilsFilterer creates a new log filterer instance of Utils, bound to a specific deployed contract.
func NewUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*UtilsFilterer, error) {
	contract, err := bindUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UtilsFilterer{contract: contract}, nil
}

// bindUtils binds a generic wrapper to an already deployed contract.
func bindUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.UtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transact(opts, method, params...)
}

// ContractExists is a free data retrieval call binding the contract method 0x7709bc78.
//
// Solidity: function contractExists(contract_address address) constant returns(bool)
func (_Utils *UtilsCaller) ContractExists(opts *bind.CallOpts, contract_address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Utils.contract.Call(opts, out, "contractExists", contract_address)
	return *ret0, err
}

// ContractExists is a free data retrieval call binding the contract method 0x7709bc78.
//
// Solidity: function contractExists(contract_address address) constant returns(bool)
func (_Utils *UtilsSession) ContractExists(contract_address common.Address) (bool, error) {
	return _Utils.Contract.ContractExists(&_Utils.CallOpts, contract_address)
}

// ContractExists is a free data retrieval call binding the contract method 0x7709bc78.
//
// Solidity: function contractExists(contract_address address) constant returns(bool)
func (_Utils *UtilsCallerSession) ContractExists(contract_address common.Address) (bool, error) {
	return _Utils.Contract.ContractExists(&_Utils.CallOpts, contract_address)
}

// Contract_version is a free data retrieval call binding the contract method 0xb32c65c8.
//
// Solidity: function contract_version() constant returns(string)
func (_Utils *UtilsCaller) Contract_version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Utils.contract.Call(opts, out, "contract_version")
	return *ret0, err
}

// Contract_version is a free data retrieval call binding the contract method 0xb32c65c8.
//
// Solidity: function contract_version() constant returns(string)
func (_Utils *UtilsSession) Contract_version() (string, error) {
	return _Utils.Contract.Contract_version(&_Utils.CallOpts)
}

// Contract_version is a free data retrieval call binding the contract method 0xb32c65c8.
//
// Solidity: function contract_version() constant returns(string)
func (_Utils *UtilsCallerSession) Contract_version() (string, error) {
	return _Utils.Contract.Contract_version(&_Utils.CallOpts)
}