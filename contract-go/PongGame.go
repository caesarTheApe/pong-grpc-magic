// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PongGame

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

// PongGameMetaData contains all meta data concerning the PongGame contract.
var PongGameMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"player1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"player2\",\"type\":\"address\"}],\"name\":\"FundsUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"GameCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"player1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"betAmount\",\"type\":\"uint256\"}],\"name\":\"GameCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"player2\",\"type\":\"address\"}],\"name\":\"PlayerJoined\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"gameIdCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"games\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"player1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"player2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"betAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isComplete\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"fundsLocked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_betAmount\",\"type\":\"uint256\"}],\"name\":\"lockOrJoinGame\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gameId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_winner\",\"type\":\"address\"}],\"name\":\"setWinner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gameId\",\"type\":\"uint256\"}],\"name\":\"unlockFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260015f55348015610013575f80fd5b50611221806100215f395ff3fe608060405260043610610049575f3560e01c8063117a5b901461004d5780631bb412d61461008e5780632b6ecf65146100b85780636894c950146100d45780639c623683146100fc575b5f80fd5b348015610058575f80fd5b50610073600480360381019061006e9190610c66565b610124565b60405161008596959493929190610cf9565b60405180910390f35b348015610099575f80fd5b506100a26101d3565b6040516100af9190610d58565b60405180910390f35b6100d260048036038101906100cd9190610c66565b6101d8565b005b3480156100df575f80fd5b506100fa60048036038101906100f59190610c66565b61054c565b005b348015610107575f80fd5b50610122600480360381019061011d9190610d9b565b61090e565b005b6001602052805f5260405f205f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806002015490806003015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160149054906101000a900460ff16908060030160159054906101000a900460ff16905086565b5f5481565b80341461021a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021190610e33565b60405180910390fd5b5f61022482610b4e565b90505f8114610353575f60015f8381526020019081526020015f2090505f73ffffffffffffffffffffffffffffffffffffffff16816001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146102d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102c990610e9b565b60405180910390fd5b33816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f87969bc7faf902221a147b95ceba76e011c5efb0339a0a8ee7a2bb82d9cfbbd68233604051610345929190610eb9565b60405180910390a150610548565b5f8081548092919061036490610f0d565b9190505590506040518060c001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020018381526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f151581526020016001151581525060015f8381526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600201556060820151816003015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160030160146101000a81548160ff02191690831515021790555060a08201518160030160156101000a81548160ff0219169083151502179055509050507f7dfb67e9ff596fca4da65c7eedb128cd1aac553af54b3c0cb733625a2480d8bd81338460405161053f93929190610f54565b60405180910390a15b5050565b5f60015f8381526020019081526020015f209050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061060a5750806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b610649576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064090610ff9565b60405180910390fd5b8060030160149054906101000a900460ff161561069b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161069290611061565b60405180910390fd5b8060030160159054906101000a900460ff166106ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e3906110c9565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff16815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146107ac57805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc826002015490811502906040515f60405180830381858888f193505050501580156107aa573d5f803e3d5ffd5b505b5f73ffffffffffffffffffffffffffffffffffffffff16816001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461086e57806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc826002015490811502906040515f60405180830381858888f1935050505015801561086c573d5f803e3d5ffd5b505b5f8160030160156101000a81548160ff0219169083151502179055507fb26745e998cfb2500d3c5012ac5af86763b7126681162572e89f1dfa23d7ffa082825f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051610902939291906110e7565b60405180910390a15050565b5f60015f8481526020019081526020015f2090508060030160149054906101000a900460ff1615610974576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161096b90611061565b60405180910390fd5b805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480610a1e5750806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b610a5d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a549061118c565b60405180910390fd5b81816003015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018160030160146101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff166108fc60028360020154610ae791906111aa565b90811502906040515f60405180830381858888f19350505050158015610b0f573d5f803e3d5ffd5b507f7f40df5127524300187445e867486da572aab86769c0212e0b392bd87c4ca97f8383604051610b41929190610eb9565b60405180910390a1505050565b5f80600190505b5f54811015610c25575f73ffffffffffffffffffffffffffffffffffffffff1660015f8381526020019081526020015f206001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16148015610bdf57508260015f8381526020019081526020015f2060020154145b8015610c0a575060015f8281526020019081526020015f2060030160149054906101000a900460ff16155b15610c185780915050610c2a565b8080600101915050610b55565b505f90505b919050565b5f80fd5b5f819050919050565b610c4581610c33565b8114610c4f575f80fd5b50565b5f81359050610c6081610c3c565b92915050565b5f60208284031215610c7b57610c7a610c2f565b5b5f610c8884828501610c52565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610cba82610c91565b9050919050565b610cca81610cb0565b82525050565b610cd981610c33565b82525050565b5f8115159050919050565b610cf381610cdf565b82525050565b5f60c082019050610d0c5f830189610cc1565b610d196020830188610cc1565b610d266040830187610cd0565b610d336060830186610cc1565b610d406080830185610cea565b610d4d60a0830184610cea565b979650505050505050565b5f602082019050610d6b5f830184610cd0565b92915050565b610d7a81610cb0565b8114610d84575f80fd5b50565b5f81359050610d9581610d71565b92915050565b5f8060408385031215610db157610db0610c2f565b5b5f610dbe85828601610c52565b9250506020610dcf85828601610d87565b9150509250929050565b5f82825260208201905092915050565b7f496e636f72726563742062657420616d6f756e740000000000000000000000005f82015250565b5f610e1d601483610dd9565b9150610e2882610de9565b602082019050919050565b5f6020820190508181035f830152610e4a81610e11565b9050919050565b7f47616d6520616c7265616479206861732074776f20706c6179657273000000005f82015250565b5f610e85601c83610dd9565b9150610e9082610e51565b602082019050919050565b5f6020820190508181035f830152610eb281610e79565b9050919050565b5f604082019050610ecc5f830185610cd0565b610ed96020830184610cc1565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610f1782610c33565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610f4957610f48610ee0565b5b600182019050919050565b5f606082019050610f675f830186610cd0565b610f746020830185610cc1565b610f816040830184610cd0565b949350505050565b7f4f6e6c792067616d65207061727469636970616e74732063616e20756e6c6f635f8201527f6b2066756e647300000000000000000000000000000000000000000000000000602082015250565b5f610fe3602783610dd9565b9150610fee82610f89565b604082019050919050565b5f6020820190508181035f83015261101081610fd7565b9050919050565b7f47616d6520697320616c726561647920636f6d706c65746500000000000000005f82015250565b5f61104b601883610dd9565b915061105682611017565b602082019050919050565b5f6020820190508181035f8301526110788161103f565b9050919050565b7f46756e647320617265206e6f74206c6f636b65640000000000000000000000005f82015250565b5f6110b3601483610dd9565b91506110be8261107f565b602082019050919050565b5f6020820190508181035f8301526110e0816110a7565b9050919050565b5f6060820190506110fa5f830186610cd0565b6111076020830185610cc1565b6111146040830184610cc1565b949350505050565b7f4f6e6c792067616d65207061727469636970616e74732063616e2073657420745f8201527f68652077696e6e65720000000000000000000000000000000000000000000000602082015250565b5f611176602983610dd9565b91506111818261111c565b604082019050919050565b5f6020820190508181035f8301526111a38161116a565b9050919050565b5f6111b482610c33565b91506111bf83610c33565b92508282026111cd81610c33565b915082820484148315176111e4576111e3610ee0565b5b509291505056fea26469706673582212206a309e687bfba0280ac44814868c343f299dd0f8c228c06bca44396a40c639fd64736f6c63430008180033",
}

// PongGameABI is the input ABI used to generate the binding from.
// Deprecated: Use PongGameMetaData.ABI instead.
var PongGameABI = PongGameMetaData.ABI

// PongGameBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PongGameMetaData.Bin instead.
var PongGameBin = PongGameMetaData.Bin

// DeployPongGame deploys a new Ethereum contract, binding an instance of PongGame to it.
func DeployPongGame(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PongGame, error) {
	parsed, err := PongGameMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PongGameBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PongGame{PongGameCaller: PongGameCaller{contract: contract}, PongGameTransactor: PongGameTransactor{contract: contract}, PongGameFilterer: PongGameFilterer{contract: contract}}, nil
}

// PongGame is an auto generated Go binding around an Ethereum contract.
type PongGame struct {
	PongGameCaller     // Read-only binding to the contract
	PongGameTransactor // Write-only binding to the contract
	PongGameFilterer   // Log filterer for contract events
}

// PongGameCaller is an auto generated read-only Go binding around an Ethereum contract.
type PongGameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PongGameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PongGameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PongGameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PongGameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PongGameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PongGameSession struct {
	Contract     *PongGame         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PongGameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PongGameCallerSession struct {
	Contract *PongGameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PongGameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PongGameTransactorSession struct {
	Contract     *PongGameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PongGameRaw is an auto generated low-level Go binding around an Ethereum contract.
type PongGameRaw struct {
	Contract *PongGame // Generic contract binding to access the raw methods on
}

// PongGameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PongGameCallerRaw struct {
	Contract *PongGameCaller // Generic read-only contract binding to access the raw methods on
}

// PongGameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PongGameTransactorRaw struct {
	Contract *PongGameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPongGame creates a new instance of PongGame, bound to a specific deployed contract.
func NewPongGame(address common.Address, backend bind.ContractBackend) (*PongGame, error) {
	contract, err := bindPongGame(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PongGame{PongGameCaller: PongGameCaller{contract: contract}, PongGameTransactor: PongGameTransactor{contract: contract}, PongGameFilterer: PongGameFilterer{contract: contract}}, nil
}

// NewPongGameCaller creates a new read-only instance of PongGame, bound to a specific deployed contract.
func NewPongGameCaller(address common.Address, caller bind.ContractCaller) (*PongGameCaller, error) {
	contract, err := bindPongGame(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PongGameCaller{contract: contract}, nil
}

// NewPongGameTransactor creates a new write-only instance of PongGame, bound to a specific deployed contract.
func NewPongGameTransactor(address common.Address, transactor bind.ContractTransactor) (*PongGameTransactor, error) {
	contract, err := bindPongGame(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PongGameTransactor{contract: contract}, nil
}

// NewPongGameFilterer creates a new log filterer instance of PongGame, bound to a specific deployed contract.
func NewPongGameFilterer(address common.Address, filterer bind.ContractFilterer) (*PongGameFilterer, error) {
	contract, err := bindPongGame(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PongGameFilterer{contract: contract}, nil
}

// bindPongGame binds a generic wrapper to an already deployed contract.
func bindPongGame(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PongGameMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PongGame *PongGameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PongGame.Contract.PongGameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PongGame *PongGameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PongGame.Contract.PongGameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PongGame *PongGameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PongGame.Contract.PongGameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PongGame *PongGameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PongGame.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PongGame *PongGameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PongGame.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PongGame *PongGameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PongGame.Contract.contract.Transact(opts, method, params...)
}

// GameIdCounter is a free data retrieval call binding the contract method 0x1bb412d6.
//
// Solidity: function gameIdCounter() view returns(uint256)
func (_PongGame *PongGameCaller) GameIdCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PongGame.contract.Call(opts, &out, "gameIdCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GameIdCounter is a free data retrieval call binding the contract method 0x1bb412d6.
//
// Solidity: function gameIdCounter() view returns(uint256)
func (_PongGame *PongGameSession) GameIdCounter() (*big.Int, error) {
	return _PongGame.Contract.GameIdCounter(&_PongGame.CallOpts)
}

// GameIdCounter is a free data retrieval call binding the contract method 0x1bb412d6.
//
// Solidity: function gameIdCounter() view returns(uint256)
func (_PongGame *PongGameCallerSession) GameIdCounter() (*big.Int, error) {
	return _PongGame.Contract.GameIdCounter(&_PongGame.CallOpts)
}

// Games is a free data retrieval call binding the contract method 0x117a5b90.
//
// Solidity: function games(uint256 ) view returns(address player1, address player2, uint256 betAmount, address winner, bool isComplete, bool fundsLocked)
func (_PongGame *PongGameCaller) Games(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Player1     common.Address
	Player2     common.Address
	BetAmount   *big.Int
	Winner      common.Address
	IsComplete  bool
	FundsLocked bool
}, error) {
	var out []interface{}
	err := _PongGame.contract.Call(opts, &out, "games", arg0)

	outstruct := new(struct {
		Player1     common.Address
		Player2     common.Address
		BetAmount   *big.Int
		Winner      common.Address
		IsComplete  bool
		FundsLocked bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Player1 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Player2 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.BetAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Winner = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.IsComplete = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.FundsLocked = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// Games is a free data retrieval call binding the contract method 0x117a5b90.
//
// Solidity: function games(uint256 ) view returns(address player1, address player2, uint256 betAmount, address winner, bool isComplete, bool fundsLocked)
func (_PongGame *PongGameSession) Games(arg0 *big.Int) (struct {
	Player1     common.Address
	Player2     common.Address
	BetAmount   *big.Int
	Winner      common.Address
	IsComplete  bool
	FundsLocked bool
}, error) {
	return _PongGame.Contract.Games(&_PongGame.CallOpts, arg0)
}

// Games is a free data retrieval call binding the contract method 0x117a5b90.
//
// Solidity: function games(uint256 ) view returns(address player1, address player2, uint256 betAmount, address winner, bool isComplete, bool fundsLocked)
func (_PongGame *PongGameCallerSession) Games(arg0 *big.Int) (struct {
	Player1     common.Address
	Player2     common.Address
	BetAmount   *big.Int
	Winner      common.Address
	IsComplete  bool
	FundsLocked bool
}, error) {
	return _PongGame.Contract.Games(&_PongGame.CallOpts, arg0)
}

// LockOrJoinGame is a paid mutator transaction binding the contract method 0x2b6ecf65.
//
// Solidity: function lockOrJoinGame(uint256 _betAmount) payable returns()
func (_PongGame *PongGameTransactor) LockOrJoinGame(opts *bind.TransactOpts, _betAmount *big.Int) (*types.Transaction, error) {
	return _PongGame.contract.Transact(opts, "lockOrJoinGame", _betAmount)
}

// LockOrJoinGame is a paid mutator transaction binding the contract method 0x2b6ecf65.
//
// Solidity: function lockOrJoinGame(uint256 _betAmount) payable returns()
func (_PongGame *PongGameSession) LockOrJoinGame(_betAmount *big.Int) (*types.Transaction, error) {
	return _PongGame.Contract.LockOrJoinGame(&_PongGame.TransactOpts, _betAmount)
}

// LockOrJoinGame is a paid mutator transaction binding the contract method 0x2b6ecf65.
//
// Solidity: function lockOrJoinGame(uint256 _betAmount) payable returns()
func (_PongGame *PongGameTransactorSession) LockOrJoinGame(_betAmount *big.Int) (*types.Transaction, error) {
	return _PongGame.Contract.LockOrJoinGame(&_PongGame.TransactOpts, _betAmount)
}

// SetWinner is a paid mutator transaction binding the contract method 0x9c623683.
//
// Solidity: function setWinner(uint256 _gameId, address _winner) returns()
func (_PongGame *PongGameTransactor) SetWinner(opts *bind.TransactOpts, _gameId *big.Int, _winner common.Address) (*types.Transaction, error) {
	return _PongGame.contract.Transact(opts, "setWinner", _gameId, _winner)
}

// SetWinner is a paid mutator transaction binding the contract method 0x9c623683.
//
// Solidity: function setWinner(uint256 _gameId, address _winner) returns()
func (_PongGame *PongGameSession) SetWinner(_gameId *big.Int, _winner common.Address) (*types.Transaction, error) {
	return _PongGame.Contract.SetWinner(&_PongGame.TransactOpts, _gameId, _winner)
}

// SetWinner is a paid mutator transaction binding the contract method 0x9c623683.
//
// Solidity: function setWinner(uint256 _gameId, address _winner) returns()
func (_PongGame *PongGameTransactorSession) SetWinner(_gameId *big.Int, _winner common.Address) (*types.Transaction, error) {
	return _PongGame.Contract.SetWinner(&_PongGame.TransactOpts, _gameId, _winner)
}

// UnlockFunds is a paid mutator transaction binding the contract method 0x6894c950.
//
// Solidity: function unlockFunds(uint256 _gameId) returns()
func (_PongGame *PongGameTransactor) UnlockFunds(opts *bind.TransactOpts, _gameId *big.Int) (*types.Transaction, error) {
	return _PongGame.contract.Transact(opts, "unlockFunds", _gameId)
}

// UnlockFunds is a paid mutator transaction binding the contract method 0x6894c950.
//
// Solidity: function unlockFunds(uint256 _gameId) returns()
func (_PongGame *PongGameSession) UnlockFunds(_gameId *big.Int) (*types.Transaction, error) {
	return _PongGame.Contract.UnlockFunds(&_PongGame.TransactOpts, _gameId)
}

// UnlockFunds is a paid mutator transaction binding the contract method 0x6894c950.
//
// Solidity: function unlockFunds(uint256 _gameId) returns()
func (_PongGame *PongGameTransactorSession) UnlockFunds(_gameId *big.Int) (*types.Transaction, error) {
	return _PongGame.Contract.UnlockFunds(&_PongGame.TransactOpts, _gameId)
}

// PongGameFundsUnlockedIterator is returned from FilterFundsUnlocked and is used to iterate over the raw logs and unpacked data for FundsUnlocked events raised by the PongGame contract.
type PongGameFundsUnlockedIterator struct {
	Event *PongGameFundsUnlocked // Event containing the contract specifics and raw log

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
func (it *PongGameFundsUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PongGameFundsUnlocked)
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
		it.Event = new(PongGameFundsUnlocked)
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
func (it *PongGameFundsUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PongGameFundsUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PongGameFundsUnlocked represents a FundsUnlocked event raised by the PongGame contract.
type PongGameFundsUnlocked struct {
	GameId  *big.Int
	Player1 common.Address
	Player2 common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFundsUnlocked is a free log retrieval operation binding the contract event 0xb26745e998cfb2500d3c5012ac5af86763b7126681162572e89f1dfa23d7ffa0.
//
// Solidity: event FundsUnlocked(uint256 gameId, address player1, address player2)
func (_PongGame *PongGameFilterer) FilterFundsUnlocked(opts *bind.FilterOpts) (*PongGameFundsUnlockedIterator, error) {

	logs, sub, err := _PongGame.contract.FilterLogs(opts, "FundsUnlocked")
	if err != nil {
		return nil, err
	}
	return &PongGameFundsUnlockedIterator{contract: _PongGame.contract, event: "FundsUnlocked", logs: logs, sub: sub}, nil
}

// WatchFundsUnlocked is a free log subscription operation binding the contract event 0xb26745e998cfb2500d3c5012ac5af86763b7126681162572e89f1dfa23d7ffa0.
//
// Solidity: event FundsUnlocked(uint256 gameId, address player1, address player2)
func (_PongGame *PongGameFilterer) WatchFundsUnlocked(opts *bind.WatchOpts, sink chan<- *PongGameFundsUnlocked) (event.Subscription, error) {

	logs, sub, err := _PongGame.contract.WatchLogs(opts, "FundsUnlocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PongGameFundsUnlocked)
				if err := _PongGame.contract.UnpackLog(event, "FundsUnlocked", log); err != nil {
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

// ParseFundsUnlocked is a log parse operation binding the contract event 0xb26745e998cfb2500d3c5012ac5af86763b7126681162572e89f1dfa23d7ffa0.
//
// Solidity: event FundsUnlocked(uint256 gameId, address player1, address player2)
func (_PongGame *PongGameFilterer) ParseFundsUnlocked(log types.Log) (*PongGameFundsUnlocked, error) {
	event := new(PongGameFundsUnlocked)
	if err := _PongGame.contract.UnpackLog(event, "FundsUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PongGameGameCompletedIterator is returned from FilterGameCompleted and is used to iterate over the raw logs and unpacked data for GameCompleted events raised by the PongGame contract.
type PongGameGameCompletedIterator struct {
	Event *PongGameGameCompleted // Event containing the contract specifics and raw log

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
func (it *PongGameGameCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PongGameGameCompleted)
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
		it.Event = new(PongGameGameCompleted)
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
func (it *PongGameGameCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PongGameGameCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PongGameGameCompleted represents a GameCompleted event raised by the PongGame contract.
type PongGameGameCompleted struct {
	GameId *big.Int
	Winner common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterGameCompleted is a free log retrieval operation binding the contract event 0x7f40df5127524300187445e867486da572aab86769c0212e0b392bd87c4ca97f.
//
// Solidity: event GameCompleted(uint256 gameId, address winner)
func (_PongGame *PongGameFilterer) FilterGameCompleted(opts *bind.FilterOpts) (*PongGameGameCompletedIterator, error) {

	logs, sub, err := _PongGame.contract.FilterLogs(opts, "GameCompleted")
	if err != nil {
		return nil, err
	}
	return &PongGameGameCompletedIterator{contract: _PongGame.contract, event: "GameCompleted", logs: logs, sub: sub}, nil
}

// WatchGameCompleted is a free log subscription operation binding the contract event 0x7f40df5127524300187445e867486da572aab86769c0212e0b392bd87c4ca97f.
//
// Solidity: event GameCompleted(uint256 gameId, address winner)
func (_PongGame *PongGameFilterer) WatchGameCompleted(opts *bind.WatchOpts, sink chan<- *PongGameGameCompleted) (event.Subscription, error) {

	logs, sub, err := _PongGame.contract.WatchLogs(opts, "GameCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PongGameGameCompleted)
				if err := _PongGame.contract.UnpackLog(event, "GameCompleted", log); err != nil {
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

// ParseGameCompleted is a log parse operation binding the contract event 0x7f40df5127524300187445e867486da572aab86769c0212e0b392bd87c4ca97f.
//
// Solidity: event GameCompleted(uint256 gameId, address winner)
func (_PongGame *PongGameFilterer) ParseGameCompleted(log types.Log) (*PongGameGameCompleted, error) {
	event := new(PongGameGameCompleted)
	if err := _PongGame.contract.UnpackLog(event, "GameCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PongGameGameCreatedIterator is returned from FilterGameCreated and is used to iterate over the raw logs and unpacked data for GameCreated events raised by the PongGame contract.
type PongGameGameCreatedIterator struct {
	Event *PongGameGameCreated // Event containing the contract specifics and raw log

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
func (it *PongGameGameCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PongGameGameCreated)
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
		it.Event = new(PongGameGameCreated)
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
func (it *PongGameGameCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PongGameGameCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PongGameGameCreated represents a GameCreated event raised by the PongGame contract.
type PongGameGameCreated struct {
	GameId    *big.Int
	Player1   common.Address
	BetAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGameCreated is a free log retrieval operation binding the contract event 0x7dfb67e9ff596fca4da65c7eedb128cd1aac553af54b3c0cb733625a2480d8bd.
//
// Solidity: event GameCreated(uint256 gameId, address player1, uint256 betAmount)
func (_PongGame *PongGameFilterer) FilterGameCreated(opts *bind.FilterOpts) (*PongGameGameCreatedIterator, error) {

	logs, sub, err := _PongGame.contract.FilterLogs(opts, "GameCreated")
	if err != nil {
		return nil, err
	}
	return &PongGameGameCreatedIterator{contract: _PongGame.contract, event: "GameCreated", logs: logs, sub: sub}, nil
}

// WatchGameCreated is a free log subscription operation binding the contract event 0x7dfb67e9ff596fca4da65c7eedb128cd1aac553af54b3c0cb733625a2480d8bd.
//
// Solidity: event GameCreated(uint256 gameId, address player1, uint256 betAmount)
func (_PongGame *PongGameFilterer) WatchGameCreated(opts *bind.WatchOpts, sink chan<- *PongGameGameCreated) (event.Subscription, error) {

	logs, sub, err := _PongGame.contract.WatchLogs(opts, "GameCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PongGameGameCreated)
				if err := _PongGame.contract.UnpackLog(event, "GameCreated", log); err != nil {
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

// ParseGameCreated is a log parse operation binding the contract event 0x7dfb67e9ff596fca4da65c7eedb128cd1aac553af54b3c0cb733625a2480d8bd.
//
// Solidity: event GameCreated(uint256 gameId, address player1, uint256 betAmount)
func (_PongGame *PongGameFilterer) ParseGameCreated(log types.Log) (*PongGameGameCreated, error) {
	event := new(PongGameGameCreated)
	if err := _PongGame.contract.UnpackLog(event, "GameCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PongGamePlayerJoinedIterator is returned from FilterPlayerJoined and is used to iterate over the raw logs and unpacked data for PlayerJoined events raised by the PongGame contract.
type PongGamePlayerJoinedIterator struct {
	Event *PongGamePlayerJoined // Event containing the contract specifics and raw log

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
func (it *PongGamePlayerJoinedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PongGamePlayerJoined)
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
		it.Event = new(PongGamePlayerJoined)
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
func (it *PongGamePlayerJoinedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PongGamePlayerJoinedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PongGamePlayerJoined represents a PlayerJoined event raised by the PongGame contract.
type PongGamePlayerJoined struct {
	GameId  *big.Int
	Player2 common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPlayerJoined is a free log retrieval operation binding the contract event 0x87969bc7faf902221a147b95ceba76e011c5efb0339a0a8ee7a2bb82d9cfbbd6.
//
// Solidity: event PlayerJoined(uint256 gameId, address player2)
func (_PongGame *PongGameFilterer) FilterPlayerJoined(opts *bind.FilterOpts) (*PongGamePlayerJoinedIterator, error) {

	logs, sub, err := _PongGame.contract.FilterLogs(opts, "PlayerJoined")
	if err != nil {
		return nil, err
	}
	return &PongGamePlayerJoinedIterator{contract: _PongGame.contract, event: "PlayerJoined", logs: logs, sub: sub}, nil
}

// WatchPlayerJoined is a free log subscription operation binding the contract event 0x87969bc7faf902221a147b95ceba76e011c5efb0339a0a8ee7a2bb82d9cfbbd6.
//
// Solidity: event PlayerJoined(uint256 gameId, address player2)
func (_PongGame *PongGameFilterer) WatchPlayerJoined(opts *bind.WatchOpts, sink chan<- *PongGamePlayerJoined) (event.Subscription, error) {

	logs, sub, err := _PongGame.contract.WatchLogs(opts, "PlayerJoined")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PongGamePlayerJoined)
				if err := _PongGame.contract.UnpackLog(event, "PlayerJoined", log); err != nil {
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

// ParsePlayerJoined is a log parse operation binding the contract event 0x87969bc7faf902221a147b95ceba76e011c5efb0339a0a8ee7a2bb82d9cfbbd6.
//
// Solidity: event PlayerJoined(uint256 gameId, address player2)
func (_PongGame *PongGameFilterer) ParsePlayerJoined(log types.Log) (*PongGamePlayerJoined, error) {
	event := new(PongGamePlayerJoined)
	if err := _PongGame.contract.UnpackLog(event, "PlayerJoined", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
