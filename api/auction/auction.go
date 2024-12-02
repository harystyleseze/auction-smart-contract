// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package auction

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

// AuctionMetaData contains all meta data concerning the Auction contract.
var AuctionMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_nft\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nftId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allBids\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bid\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"end\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"endTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ended\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"highestBid\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"highestBidder\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nft\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC721\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nftId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"start\",\"inputs\":[{\"name\":\"_openingBid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"started\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Bid\",\"inputs\":[{\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"End\",\"inputs\":[{\"name\":\"highestBidder\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Start\",\"inputs\":[{\"name\":\"startTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60e060405234801561000f575f5ffd5b50604051610c2b380380610c2b83398101604081905261002e91610049565b336080526001600160a01b039190911660a05260c052610080565b5f5f6040838503121561005a575f5ffd5b82516001600160a01b0381168114610070575f5ffd5b6020939093015192949293505050565b60805160a05160c051610b496100e25f395f818161024601528181610718015261098101525f818161015f0152818161073e01526109a701525f81816101aa01528181610602015281816106ea015281816107f60152610a0e0152610b495ff3fe6080604052600436106100bf575f3560e01c80638da5cb5b1161007c578063bd20160711610057578063bd2016071461020a578063c6bc518214610235578063d57bde7914610268578063efbe1c1c1461027d575f5ffd5b80638da5cb5b146101995780638fb4b573146101cc57806391f90157146101eb575f5ffd5b806312fa6feb146100c35780631998aeef146100f55780631f2698ab146100ff5780633197cbb6146101175780633ccfd60b1461013a57806347ccca021461014e575b5f5ffd5b3480156100ce575f5ffd5b505f546100e090610100900460ff1681565b60405190151581526020015b60405180910390f35b6100fd610291565b005b34801561010a575f5ffd5b505f546100e09060ff1681565b348015610122575f5ffd5b5061012c60015481565b6040519081526020016100ec565b348015610145575f5ffd5b506100fd610423565b348015610159575f5ffd5b506101817f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100ec565b3480156101a4575f5ffd5b506101817f000000000000000000000000000000000000000000000000000000000000000081565b3480156101d7575f5ffd5b506100fd6101e6366004610aa1565b6105f7565b3480156101f6575f5ffd5b50600354610181906001600160a01b031681565b348015610215575f5ffd5b5061012c610224366004610ac1565b60046020525f908152604090205481565b348015610240575f5ffd5b5061012c7f000000000000000000000000000000000000000000000000000000000000000081565b348015610273575f5ffd5b5061012c60025481565b348015610288575f5ffd5b506100fd6107eb565b5f5460ff166102e15760405162461bcd60e51b8152602060048201526017602482015276105d58dd1a5bdb881a185cc81b9bdd081cdd185c9d1959604a1b60448201526064015b60405180910390fd5b600254341161034a5760405162461bcd60e51b815260206004820152602f60248201527f426964207072696365206973206c6f776572207468616e20746865206375727260448201526e195b9d081a1a59da195cdd08189a59608a1b60648201526084016102d8565b600154421061038f5760405162461bcd60e51b8152602060048201526011602482015270105d58dd1a5bdb881a185cc8195b991959607a1b60448201526064016102d8565b6003546001600160a01b0316156103d2576002546003546001600160a01b03165f90815260046020526040812080549091906103cc908490610aee565b90915550505b346002819055600380546001600160a01b03191633908117909155604051918252907fe684a55f31b79eca403df938249029212a5925ec6be8012e099b45bc1019e5d29060200160405180910390a2565b5f54610100900460ff166104715760405162461bcd60e51b8152602060048201526015602482015274105d58dd1a5bdb881b9bdd081e595d08195b991959605a1b60448201526064016102d8565b6003546001600160a01b031633036104d65760405162461bcd60e51b815260206004820152602260248201527f57696e6e65722063616e6e6f742077697468647261772077696e6e696e6720626044820152611a5960f21b60648201526084016102d8565b335f90815260046020526040902054806105285760405162461bcd60e51b81526020600482015260136024820152724e6f7468696e6720746f20776974686472617760681b60448201526064016102d8565b335f818152600460205260408082208290555190919083908381818185875af1925050503d805f8114610576576040519150601f19603f3d011682016040523d82523d5f602084013e61057b565b606091505b50509050806105be5760405162461bcd60e51b815260206004820152600f60248201526e151c985b9cd9995c8819985a5b1959608a1b60448201526064016102d8565b60405182815233907f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a94243649060200160405180910390a25050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461066f5760405162461bcd60e51b815260206004820181905260248201527f4f6e6c79206f776e65722063616e2063616c6c207468652066756e6374696f6e60448201526064016102d8565b5f5460ff16156106c15760405162461bcd60e51b815260206004820152601b60248201527f41756374696f6e2068617320616c72656164792073746172746564000000000060448201526064016102d8565b60028290556106d08142610aee565b6001556040516323b872dd60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811660048301523060248301527f000000000000000000000000000000000000000000000000000000000000000060448301527f000000000000000000000000000000000000000000000000000000000000000016906323b872dd906064015f604051808303815f87803b15801561077f575f5ffd5b505af1158015610791573d5f5f3e3d5ffd5b50505f805460ff19166001908117909155546040517f5a7d5f50ab70a39d193bf53f0fb995377776df93bb6fdcf1cd1868b7e0dd44ee93506107df9250429190918252602082015260400190565b60405180910390a15050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146108635760405162461bcd60e51b815260206004820181905260248201527f4f6e6c79206f776e65722063616e2063616c6c207468652066756e6374696f6e60448201526064016102d8565b5f5460ff166108ae5760405162461bcd60e51b8152602060048201526017602482015276105d58dd1a5bdb881a185cc81b9bdd081cdd185c9d1959604a1b60448201526064016102d8565b6001544210156108f45760405162461bcd60e51b8152602060048201526011602482015270105d58dd1a5bdb881a185cc8195b991959607a1b60448201526064016102d8565b5f54610100900460ff161561094b5760405162461bcd60e51b815260206004820152601960248201527f41756374696f6e2068617320616c726561647920656e6465640000000000000060448201526064016102d8565b5f805461ff0019166101001790556003546040516323b872dd60e01b81523060048201526001600160a01b0391821660248201527f000000000000000000000000000000000000000000000000000000000000000060448201527f0000000000000000000000000000000000000000000000000000000000000000909116906323b872dd906064015f604051808303815f87803b1580156109ea575f5ffd5b505af11580156109fc573d5f5f3e3d5ffd5b50506002546040516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016935081156108fc0292505f818181858888f19350505050158015610a54573d5f5f3e3d5ffd5b50600354600254604080516001600160a01b03909316835260208301919091527f7d7570b046e5ead3a4f3fbc9fa2445113625a1e79823776146316bafed6b8cc5910160405180910390a1565b5f5f60408385031215610ab2575f5ffd5b50508035926020909101359150565b5f60208284031215610ad1575f5ffd5b81356001600160a01b0381168114610ae7575f5ffd5b9392505050565b80820180821115610b0d57634e487b7160e01b5f52601160045260245ffd5b9291505056fea2646970667358221220cf7625216c86106afd97dc1c0c1fb3931dd5da124d2d81908259d563d59fca7864736f6c634300081c0033",
}

// AuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use AuctionMetaData.ABI instead.
var AuctionABI = AuctionMetaData.ABI

// AuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AuctionMetaData.Bin instead.
var AuctionBin = AuctionMetaData.Bin

// DeployAuction deploys a new Ethereum contract, binding an instance of Auction to it.
func DeployAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _nft common.Address, _nftId *big.Int) (common.Address, *types.Transaction, *Auction, error) {
	parsed, err := AuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AuctionBin), backend, _nft, _nftId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Auction{AuctionCaller: AuctionCaller{contract: contract}, AuctionTransactor: AuctionTransactor{contract: contract}, AuctionFilterer: AuctionFilterer{contract: contract}}, nil
}

// Auction is an auto generated Go binding around an Ethereum contract.
type Auction struct {
	AuctionCaller     // Read-only binding to the contract
	AuctionTransactor // Write-only binding to the contract
	AuctionFilterer   // Log filterer for contract events
}

// AuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuctionSession struct {
	Contract     *Auction          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuctionCallerSession struct {
	Contract *AuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuctionTransactorSession struct {
	Contract     *AuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuctionRaw struct {
	Contract *Auction // Generic contract binding to access the raw methods on
}

// AuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuctionCallerRaw struct {
	Contract *AuctionCaller // Generic read-only contract binding to access the raw methods on
}

// AuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuctionTransactorRaw struct {
	Contract *AuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuction creates a new instance of Auction, bound to a specific deployed contract.
func NewAuction(address common.Address, backend bind.ContractBackend) (*Auction, error) {
	contract, err := bindAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Auction{AuctionCaller: AuctionCaller{contract: contract}, AuctionTransactor: AuctionTransactor{contract: contract}, AuctionFilterer: AuctionFilterer{contract: contract}}, nil
}

// NewAuctionCaller creates a new read-only instance of Auction, bound to a specific deployed contract.
func NewAuctionCaller(address common.Address, caller bind.ContractCaller) (*AuctionCaller, error) {
	contract, err := bindAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionCaller{contract: contract}, nil
}

// NewAuctionTransactor creates a new write-only instance of Auction, bound to a specific deployed contract.
func NewAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*AuctionTransactor, error) {
	contract, err := bindAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionTransactor{contract: contract}, nil
}

// NewAuctionFilterer creates a new log filterer instance of Auction, bound to a specific deployed contract.
func NewAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*AuctionFilterer, error) {
	contract, err := bindAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuctionFilterer{contract: contract}, nil
}

// bindAuction binds a generic wrapper to an already deployed contract.
func bindAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AuctionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auction *AuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Auction.Contract.AuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auction *AuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.Contract.AuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auction *AuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auction.Contract.AuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auction *AuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Auction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auction *AuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auction *AuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auction.Contract.contract.Transact(opts, method, params...)
}

// AllBids is a free data retrieval call binding the contract method 0xbd201607.
//
// Solidity: function allBids(address ) view returns(uint256)
func (_Auction *AuctionCaller) AllBids(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "allBids", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllBids is a free data retrieval call binding the contract method 0xbd201607.
//
// Solidity: function allBids(address ) view returns(uint256)
func (_Auction *AuctionSession) AllBids(arg0 common.Address) (*big.Int, error) {
	return _Auction.Contract.AllBids(&_Auction.CallOpts, arg0)
}

// AllBids is a free data retrieval call binding the contract method 0xbd201607.
//
// Solidity: function allBids(address ) view returns(uint256)
func (_Auction *AuctionCallerSession) AllBids(arg0 common.Address) (*big.Int, error) {
	return _Auction.Contract.AllBids(&_Auction.CallOpts, arg0)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Auction *AuctionCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "endTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Auction *AuctionSession) EndTime() (*big.Int, error) {
	return _Auction.Contract.EndTime(&_Auction.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Auction *AuctionCallerSession) EndTime() (*big.Int, error) {
	return _Auction.Contract.EndTime(&_Auction.CallOpts)
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_Auction *AuctionCaller) Ended(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "ended")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_Auction *AuctionSession) Ended() (bool, error) {
	return _Auction.Contract.Ended(&_Auction.CallOpts)
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_Auction *AuctionCallerSession) Ended() (bool, error) {
	return _Auction.Contract.Ended(&_Auction.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Auction *AuctionCaller) HighestBid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "highestBid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Auction *AuctionSession) HighestBid() (*big.Int, error) {
	return _Auction.Contract.HighestBid(&_Auction.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Auction *AuctionCallerSession) HighestBid() (*big.Int, error) {
	return _Auction.Contract.HighestBid(&_Auction.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_Auction *AuctionCaller) HighestBidder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "highestBidder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_Auction *AuctionSession) HighestBidder() (common.Address, error) {
	return _Auction.Contract.HighestBidder(&_Auction.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_Auction *AuctionCallerSession) HighestBidder() (common.Address, error) {
	return _Auction.Contract.HighestBidder(&_Auction.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_Auction *AuctionCaller) Nft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "nft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_Auction *AuctionSession) Nft() (common.Address, error) {
	return _Auction.Contract.Nft(&_Auction.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_Auction *AuctionCallerSession) Nft() (common.Address, error) {
	return _Auction.Contract.Nft(&_Auction.CallOpts)
}

// NftId is a free data retrieval call binding the contract method 0xc6bc5182.
//
// Solidity: function nftId() view returns(uint256)
func (_Auction *AuctionCaller) NftId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "nftId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NftId is a free data retrieval call binding the contract method 0xc6bc5182.
//
// Solidity: function nftId() view returns(uint256)
func (_Auction *AuctionSession) NftId() (*big.Int, error) {
	return _Auction.Contract.NftId(&_Auction.CallOpts)
}

// NftId is a free data retrieval call binding the contract method 0xc6bc5182.
//
// Solidity: function nftId() view returns(uint256)
func (_Auction *AuctionCallerSession) NftId() (*big.Int, error) {
	return _Auction.Contract.NftId(&_Auction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Auction *AuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Auction *AuctionSession) Owner() (common.Address, error) {
	return _Auction.Contract.Owner(&_Auction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Auction *AuctionCallerSession) Owner() (common.Address, error) {
	return _Auction.Contract.Owner(&_Auction.CallOpts)
}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() view returns(bool)
func (_Auction *AuctionCaller) Started(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "started")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() view returns(bool)
func (_Auction *AuctionSession) Started() (bool, error) {
	return _Auction.Contract.Started(&_Auction.CallOpts)
}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() view returns(bool)
func (_Auction *AuctionCallerSession) Started() (bool, error) {
	return _Auction.Contract.Started(&_Auction.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Auction *AuctionTransactor) Bid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "bid")
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Auction *AuctionSession) Bid() (*types.Transaction, error) {
	return _Auction.Contract.Bid(&_Auction.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Auction *AuctionTransactorSession) Bid() (*types.Transaction, error) {
	return _Auction.Contract.Bid(&_Auction.TransactOpts)
}

// End is a paid mutator transaction binding the contract method 0xefbe1c1c.
//
// Solidity: function end() returns()
func (_Auction *AuctionTransactor) End(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "end")
}

// End is a paid mutator transaction binding the contract method 0xefbe1c1c.
//
// Solidity: function end() returns()
func (_Auction *AuctionSession) End() (*types.Transaction, error) {
	return _Auction.Contract.End(&_Auction.TransactOpts)
}

// End is a paid mutator transaction binding the contract method 0xefbe1c1c.
//
// Solidity: function end() returns()
func (_Auction *AuctionTransactorSession) End() (*types.Transaction, error) {
	return _Auction.Contract.End(&_Auction.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0x8fb4b573.
//
// Solidity: function start(uint256 _openingBid, uint256 _duration) returns()
func (_Auction *AuctionTransactor) Start(opts *bind.TransactOpts, _openingBid *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "start", _openingBid, _duration)
}

// Start is a paid mutator transaction binding the contract method 0x8fb4b573.
//
// Solidity: function start(uint256 _openingBid, uint256 _duration) returns()
func (_Auction *AuctionSession) Start(_openingBid *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.Start(&_Auction.TransactOpts, _openingBid, _duration)
}

// Start is a paid mutator transaction binding the contract method 0x8fb4b573.
//
// Solidity: function start(uint256 _openingBid, uint256 _duration) returns()
func (_Auction *AuctionTransactorSession) Start(_openingBid *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.Start(&_Auction.TransactOpts, _openingBid, _duration)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Auction *AuctionTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Auction *AuctionSession) Withdraw() (*types.Transaction, error) {
	return _Auction.Contract.Withdraw(&_Auction.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Auction *AuctionTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Auction.Contract.Withdraw(&_Auction.TransactOpts)
}

// AuctionBidIterator is returned from FilterBid and is used to iterate over the raw logs and unpacked data for Bid events raised by the Auction contract.
type AuctionBidIterator struct {
	Event *AuctionBid // Event containing the contract specifics and raw log

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
func (it *AuctionBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionBid)
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
		it.Event = new(AuctionBid)
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
func (it *AuctionBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionBid represents a Bid event raised by the Auction contract.
type AuctionBid struct {
	Bidder common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBid is a free log retrieval operation binding the contract event 0xe684a55f31b79eca403df938249029212a5925ec6be8012e099b45bc1019e5d2.
//
// Solidity: event Bid(address indexed bidder, uint256 value)
func (_Auction *AuctionFilterer) FilterBid(opts *bind.FilterOpts, bidder []common.Address) (*AuctionBidIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.FilterLogs(opts, "Bid", bidderRule)
	if err != nil {
		return nil, err
	}
	return &AuctionBidIterator{contract: _Auction.contract, event: "Bid", logs: logs, sub: sub}, nil
}

// WatchBid is a free log subscription operation binding the contract event 0xe684a55f31b79eca403df938249029212a5925ec6be8012e099b45bc1019e5d2.
//
// Solidity: event Bid(address indexed bidder, uint256 value)
func (_Auction *AuctionFilterer) WatchBid(opts *bind.WatchOpts, sink chan<- *AuctionBid, bidder []common.Address) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.WatchLogs(opts, "Bid", bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionBid)
				if err := _Auction.contract.UnpackLog(event, "Bid", log); err != nil {
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

// ParseBid is a log parse operation binding the contract event 0xe684a55f31b79eca403df938249029212a5925ec6be8012e099b45bc1019e5d2.
//
// Solidity: event Bid(address indexed bidder, uint256 value)
func (_Auction *AuctionFilterer) ParseBid(log types.Log) (*AuctionBid, error) {
	event := new(AuctionBid)
	if err := _Auction.contract.UnpackLog(event, "Bid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionEndIterator is returned from FilterEnd and is used to iterate over the raw logs and unpacked data for End events raised by the Auction contract.
type AuctionEndIterator struct {
	Event *AuctionEnd // Event containing the contract specifics and raw log

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
func (it *AuctionEndIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionEnd)
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
		it.Event = new(AuctionEnd)
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
func (it *AuctionEndIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionEndIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionEnd represents a End event raised by the Auction contract.
type AuctionEnd struct {
	HighestBidder common.Address
	Value         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEnd is a free log retrieval operation binding the contract event 0x7d7570b046e5ead3a4f3fbc9fa2445113625a1e79823776146316bafed6b8cc5.
//
// Solidity: event End(address highestBidder, uint256 value)
func (_Auction *AuctionFilterer) FilterEnd(opts *bind.FilterOpts) (*AuctionEndIterator, error) {

	logs, sub, err := _Auction.contract.FilterLogs(opts, "End")
	if err != nil {
		return nil, err
	}
	return &AuctionEndIterator{contract: _Auction.contract, event: "End", logs: logs, sub: sub}, nil
}

// WatchEnd is a free log subscription operation binding the contract event 0x7d7570b046e5ead3a4f3fbc9fa2445113625a1e79823776146316bafed6b8cc5.
//
// Solidity: event End(address highestBidder, uint256 value)
func (_Auction *AuctionFilterer) WatchEnd(opts *bind.WatchOpts, sink chan<- *AuctionEnd) (event.Subscription, error) {

	logs, sub, err := _Auction.contract.WatchLogs(opts, "End")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionEnd)
				if err := _Auction.contract.UnpackLog(event, "End", log); err != nil {
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

// ParseEnd is a log parse operation binding the contract event 0x7d7570b046e5ead3a4f3fbc9fa2445113625a1e79823776146316bafed6b8cc5.
//
// Solidity: event End(address highestBidder, uint256 value)
func (_Auction *AuctionFilterer) ParseEnd(log types.Log) (*AuctionEnd, error) {
	event := new(AuctionEnd)
	if err := _Auction.contract.UnpackLog(event, "End", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionStartIterator is returned from FilterStart and is used to iterate over the raw logs and unpacked data for Start events raised by the Auction contract.
type AuctionStartIterator struct {
	Event *AuctionStart // Event containing the contract specifics and raw log

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
func (it *AuctionStartIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionStart)
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
		it.Event = new(AuctionStart)
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
func (it *AuctionStartIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionStartIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionStart represents a Start event raised by the Auction contract.
type AuctionStart struct {
	StartTime *big.Int
	EndTime   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStart is a free log retrieval operation binding the contract event 0x5a7d5f50ab70a39d193bf53f0fb995377776df93bb6fdcf1cd1868b7e0dd44ee.
//
// Solidity: event Start(uint256 startTime, uint256 endTime)
func (_Auction *AuctionFilterer) FilterStart(opts *bind.FilterOpts) (*AuctionStartIterator, error) {

	logs, sub, err := _Auction.contract.FilterLogs(opts, "Start")
	if err != nil {
		return nil, err
	}
	return &AuctionStartIterator{contract: _Auction.contract, event: "Start", logs: logs, sub: sub}, nil
}

// WatchStart is a free log subscription operation binding the contract event 0x5a7d5f50ab70a39d193bf53f0fb995377776df93bb6fdcf1cd1868b7e0dd44ee.
//
// Solidity: event Start(uint256 startTime, uint256 endTime)
func (_Auction *AuctionFilterer) WatchStart(opts *bind.WatchOpts, sink chan<- *AuctionStart) (event.Subscription, error) {

	logs, sub, err := _Auction.contract.WatchLogs(opts, "Start")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionStart)
				if err := _Auction.contract.UnpackLog(event, "Start", log); err != nil {
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

// ParseStart is a log parse operation binding the contract event 0x5a7d5f50ab70a39d193bf53f0fb995377776df93bb6fdcf1cd1868b7e0dd44ee.
//
// Solidity: event Start(uint256 startTime, uint256 endTime)
func (_Auction *AuctionFilterer) ParseStart(log types.Log) (*AuctionStart, error) {
	event := new(AuctionStart)
	if err := _Auction.contract.UnpackLog(event, "Start", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Auction contract.
type AuctionWithdrawIterator struct {
	Event *AuctionWithdraw // Event containing the contract specifics and raw log

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
func (it *AuctionWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionWithdraw)
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
		it.Event = new(AuctionWithdraw)
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
func (it *AuctionWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionWithdraw represents a Withdraw event raised by the Auction contract.
type AuctionWithdraw struct {
	Bidder common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed bidder, uint256 value)
func (_Auction *AuctionFilterer) FilterWithdraw(opts *bind.FilterOpts, bidder []common.Address) (*AuctionWithdrawIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.FilterLogs(opts, "Withdraw", bidderRule)
	if err != nil {
		return nil, err
	}
	return &AuctionWithdrawIterator{contract: _Auction.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed bidder, uint256 value)
func (_Auction *AuctionFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *AuctionWithdraw, bidder []common.Address) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.WatchLogs(opts, "Withdraw", bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionWithdraw)
				if err := _Auction.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed bidder, uint256 value)
func (_Auction *AuctionFilterer) ParseWithdraw(log types.Log) (*AuctionWithdraw, error) {
	event := new(AuctionWithdraw)
	if err := _Auction.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
