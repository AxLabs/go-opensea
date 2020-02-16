// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wyvernExchange

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// WyvernExchangeABI is the input ABI used to generate the binding from.
const WyvernExchangeABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenTransferProxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"calldata\",\"type\":\"bytes\"},{\"name\":\"extradata\",\"type\":\"bytes\"}],\"name\":\"staticCall\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newMinimumMakerProtocolFee\",\"type\":\"uint256\"}],\"name\":\"changeMinimumMakerProtocolFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newMinimumTakerProtocolFee\",\"type\":\"uint256\"}],\"name\":\"changeMinimumTakerProtocolFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"array\",\"type\":\"bytes\"},{\"name\":\"desired\",\"type\":\"bytes\"},{\"name\":\"mask\",\"type\":\"bytes\"}],\"name\":\"guardedArrayReplace\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimumTakerProtocolFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"codename\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"testCopyAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arrToCopy\",\"type\":\"bytes\"}],\"name\":\"testCopy\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[7]\"},{\"name\":\"uints\",\"type\":\"uint256[9]\"},{\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"name\":\"side\",\"type\":\"uint8\"},{\"name\":\"saleKind\",\"type\":\"uint8\"},{\"name\":\"howToCall\",\"type\":\"uint8\"},{\"name\":\"calldata\",\"type\":\"bytes\"},{\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"name\":\"staticExtradata\",\"type\":\"bytes\"}],\"name\":\"calculateCurrentPrice_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newProtocolFeeRecipient\",\"type\":\"address\"}],\"name\":\"changeProtocolFeeRecipient\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"buyCalldata\",\"type\":\"bytes\"},{\"name\":\"buyReplacementPattern\",\"type\":\"bytes\"},{\"name\":\"sellCalldata\",\"type\":\"bytes\"},{\"name\":\"sellReplacementPattern\",\"type\":\"bytes\"}],\"name\":\"orderCalldataCanMatch\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[7]\"},{\"name\":\"uints\",\"type\":\"uint256[9]\"},{\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"name\":\"side\",\"type\":\"uint8\"},{\"name\":\"saleKind\",\"type\":\"uint8\"},{\"name\":\"howToCall\",\"type\":\"uint8\"},{\"name\":\"calldata\",\"type\":\"bytes\"},{\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"name\":\"staticExtradata\",\"type\":\"bytes\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"validateOrder_\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"side\",\"type\":\"uint8\"},{\"name\":\"saleKind\",\"type\":\"uint8\"},{\"name\":\"basePrice\",\"type\":\"uint256\"},{\"name\":\"extra\",\"type\":\"uint256\"},{\"name\":\"listingTime\",\"type\":\"uint256\"},{\"name\":\"expirationTime\",\"type\":\"uint256\"}],\"name\":\"calculateFinalPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"protocolFeeRecipient\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[7]\"},{\"name\":\"uints\",\"type\":\"uint256[9]\"},{\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"name\":\"side\",\"type\":\"uint8\"},{\"name\":\"saleKind\",\"type\":\"uint8\"},{\"name\":\"howToCall\",\"type\":\"uint8\"},{\"name\":\"calldata\",\"type\":\"bytes\"},{\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"name\":\"staticExtradata\",\"type\":\"bytes\"}],\"name\":\"hashOrder_\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[14]\"},{\"name\":\"uints\",\"type\":\"uint256[18]\"},{\"name\":\"feeMethodsSidesKindsHowToCalls\",\"type\":\"uint8[8]\"},{\"name\":\"calldataBuy\",\"type\":\"bytes\"},{\"name\":\"calldataSell\",\"type\":\"bytes\"},{\"name\":\"replacementPatternBuy\",\"type\":\"bytes\"},{\"name\":\"replacementPatternSell\",\"type\":\"bytes\"},{\"name\":\"staticExtradataBuy\",\"type\":\"bytes\"},{\"name\":\"staticExtradataSell\",\"type\":\"bytes\"}],\"name\":\"ordersCanMatch_\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[7]\"},{\"name\":\"uints\",\"type\":\"uint256[9]\"},{\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"name\":\"side\",\"type\":\"uint8\"},{\"name\":\"saleKind\",\"type\":\"uint8\"},{\"name\":\"howToCall\",\"type\":\"uint8\"},{\"name\":\"calldata\",\"type\":\"bytes\"},{\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"name\":\"staticExtradata\",\"type\":\"bytes\"},{\"name\":\"orderbookInclusionDesired\",\"type\":\"bool\"}],\"name\":\"approveOrder_\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimumMakerProtocolFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[7]\"},{\"name\":\"uints\",\"type\":\"uint256[9]\"},{\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"name\":\"side\",\"type\":\"uint8\"},{\"name\":\"saleKind\",\"type\":\"uint8\"},{\"name\":\"howToCall\",\"type\":\"uint8\"},{\"name\":\"calldata\",\"type\":\"bytes\"},{\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"name\":\"staticExtradata\",\"type\":\"bytes\"}],\"name\":\"hashToSign_\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cancelledOrFinalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchangeToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[7]\"},{\"name\":\"uints\",\"type\":\"uint256[9]\"},{\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"name\":\"side\",\"type\":\"uint8\"},{\"name\":\"saleKind\",\"type\":\"uint8\"},{\"name\":\"howToCall\",\"type\":\"uint8\"},{\"name\":\"calldata\",\"type\":\"bytes\"},{\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"name\":\"staticExtradata\",\"type\":\"bytes\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"cancelOrder_\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[14]\"},{\"name\":\"uints\",\"type\":\"uint256[18]\"},{\"name\":\"feeMethodsSidesKindsHowToCalls\",\"type\":\"uint8[8]\"},{\"name\":\"calldataBuy\",\"type\":\"bytes\"},{\"name\":\"calldataSell\",\"type\":\"bytes\"},{\"name\":\"replacementPatternBuy\",\"type\":\"bytes\"},{\"name\":\"replacementPatternSell\",\"type\":\"bytes\"},{\"name\":\"staticExtradataBuy\",\"type\":\"bytes\"},{\"name\":\"staticExtradataSell\",\"type\":\"bytes\"},{\"name\":\"vs\",\"type\":\"uint8[2]\"},{\"name\":\"rssMetadata\",\"type\":\"bytes32[5]\"}],\"name\":\"atomicMatch_\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[7]\"},{\"name\":\"uints\",\"type\":\"uint256[9]\"},{\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"name\":\"side\",\"type\":\"uint8\"},{\"name\":\"saleKind\",\"type\":\"uint8\"},{\"name\":\"howToCall\",\"type\":\"uint8\"},{\"name\":\"calldata\",\"type\":\"bytes\"},{\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"name\":\"staticExtradata\",\"type\":\"bytes\"}],\"name\":\"validateOrderParameters_\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INVERSE_BASIS_POINT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[14]\"},{\"name\":\"uints\",\"type\":\"uint256[18]\"},{\"name\":\"feeMethodsSidesKindsHowToCalls\",\"type\":\"uint8[8]\"},{\"name\":\"calldataBuy\",\"type\":\"bytes\"},{\"name\":\"calldataSell\",\"type\":\"bytes\"},{\"name\":\"replacementPatternBuy\",\"type\":\"bytes\"},{\"name\":\"replacementPatternSell\",\"type\":\"bytes\"},{\"name\":\"staticExtradataBuy\",\"type\":\"bytes\"},{\"name\":\"staticExtradataSell\",\"type\":\"bytes\"}],\"name\":\"calculateMatchPrice_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"approvedOrders\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"registryAddress\",\"type\":\"address\"},{\"name\":\"tokenTransferProxyAddress\",\"type\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\"},{\"name\":\"protocolFeeAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"exchange\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"makerRelayerFee\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"takerRelayerFee\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"makerProtocolFee\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"takerProtocolFee\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"feeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"side\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"saleKind\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"target\",\"type\":\"address\"}],\"name\":\"OrderApprovedPartOne\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"howToCall\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"calldata\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"staticTarget\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"staticExtradata\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"paymentToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"basePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"extra\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"listingTime\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"salt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"orderbookInclusionDesired\",\"type\":\"bool\"}],\"name\":\"OrderApprovedPartTwo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"buyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"sellHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"metadata\",\"type\":\"bytes32\"}],\"name\":\"OrdersMatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// WyvernExchange is an auto generated Go binding around an Ethereum contract.
type WyvernExchange struct {
	WyvernExchangeCaller     // Read-only binding to the contract
	WyvernExchangeTransactor // Write-only binding to the contract
	WyvernExchangeFilterer   // Log filterer for contract events
}

// WyvernExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type WyvernExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WyvernExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WyvernExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WyvernExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WyvernExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WyvernExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WyvernExchangeSession struct {
	Contract     *WyvernExchange   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WyvernExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WyvernExchangeCallerSession struct {
	Contract *WyvernExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// WyvernExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WyvernExchangeTransactorSession struct {
	Contract     *WyvernExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// WyvernExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type WyvernExchangeRaw struct {
	Contract *WyvernExchange // Generic contract binding to access the raw methods on
}

// WyvernExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WyvernExchangeCallerRaw struct {
	Contract *WyvernExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// WyvernExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WyvernExchangeTransactorRaw struct {
	Contract *WyvernExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWyvernExchange creates a new instance of WyvernExchange, bound to a specific deployed contract.
func NewWyvernExchange(address common.Address, backend bind.ContractBackend) (*WyvernExchange, error) {
	contract, err := bindWyvernExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WyvernExchange{WyvernExchangeCaller: WyvernExchangeCaller{contract: contract}, WyvernExchangeTransactor: WyvernExchangeTransactor{contract: contract}, WyvernExchangeFilterer: WyvernExchangeFilterer{contract: contract}}, nil
}

// NewWyvernExchangeCaller creates a new read-only instance of WyvernExchange, bound to a specific deployed contract.
func NewWyvernExchangeCaller(address common.Address, caller bind.ContractCaller) (*WyvernExchangeCaller, error) {
	contract, err := bindWyvernExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WyvernExchangeCaller{contract: contract}, nil
}

// NewWyvernExchangeTransactor creates a new write-only instance of WyvernExchange, bound to a specific deployed contract.
func NewWyvernExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*WyvernExchangeTransactor, error) {
	contract, err := bindWyvernExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WyvernExchangeTransactor{contract: contract}, nil
}

// NewWyvernExchangeFilterer creates a new log filterer instance of WyvernExchange, bound to a specific deployed contract.
func NewWyvernExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*WyvernExchangeFilterer, error) {
	contract, err := bindWyvernExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WyvernExchangeFilterer{contract: contract}, nil
}

// bindWyvernExchange binds a generic wrapper to an already deployed contract.
func bindWyvernExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WyvernExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WyvernExchange *WyvernExchangeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WyvernExchange.Contract.WyvernExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WyvernExchange *WyvernExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WyvernExchange.Contract.WyvernExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WyvernExchange *WyvernExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WyvernExchange.Contract.WyvernExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WyvernExchange *WyvernExchangeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WyvernExchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WyvernExchange *WyvernExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WyvernExchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WyvernExchange *WyvernExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WyvernExchange.Contract.contract.Transact(opts, method, params...)
}

// INVERSEBASISPOINT is a free data retrieval call binding the contract method 0xcae6047f.
//
// Solidity: function INVERSE_BASIS_POINT() constant returns(uint256)
func (_WyvernExchange *WyvernExchangeCaller) INVERSEBASISPOINT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WyvernExchange.contract.Call(opts, out, "INVERSE_BASIS_POINT")
	return *ret0, err
}

// INVERSEBASISPOINT is a free data retrieval call binding the contract method 0xcae6047f.
//
// Solidity: function INVERSE_BASIS_POINT() constant returns(uint256)
func (_WyvernExchange *WyvernExchangeSession) INVERSEBASISPOINT() (*big.Int, error) {
	return _WyvernExchange.Contract.INVERSEBASISPOINT(&_WyvernExchange.CallOpts)
}

// INVERSEBASISPOINT is a free data retrieval call binding the contract method 0xcae6047f.
//
// Solidity: function INVERSE_BASIS_POINT() constant returns(uint256)
func (_WyvernExchange *WyvernExchangeCallerSession) INVERSEBASISPOINT() (*big.Int, error) {
	return _WyvernExchange.Contract.INVERSEBASISPOINT(&_WyvernExchange.CallOpts)
}

// ApprovedOrders is a free data retrieval call binding the contract method 0xe57d4adb.
//
// Solidity: function approvedOrders(bytes32 ) constant returns(bool)
func (_WyvernExchange *WyvernExchangeCaller) ApprovedOrders(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WyvernExchange.contract.Call(opts, out, "approvedOrders", arg0)
	return *ret0, err
}

// ApprovedOrders is a free data retrieval call binding the contract method 0xe57d4adb.
//
// Solidity: function approvedOrders(bytes32 ) constant returns(bool)
func (_WyvernExchange *WyvernExchangeSession) ApprovedOrders(arg0 [32]byte) (bool, error) {
	return _WyvernExchange.Contract.ApprovedOrders(&_WyvernExchange.CallOpts, arg0)
}

// ApprovedOrders is a free data retrieval call binding the contract method 0xe57d4adb.
//
// Solidity: function approvedOrders(bytes32 ) constant returns(bool)
func (_WyvernExchange *WyvernExchangeCallerSession) ApprovedOrders(arg0 [32]byte) (bool, error) {
	return _WyvernExchange.Contract.ApprovedOrders(&_WyvernExchange.CallOpts, arg0)
}

// CalculateCurrentPrice is a free data retrieval call binding the contract method 0x3f67ee0d.
//
// Solidity: function calculateCurrentPrice_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata) constant returns(uint256)
func (_WyvernExchange *WyvernExchangeCaller) CalculateCurrentPrice(opts *bind.CallOpts, addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WyvernExchange.contract.Call(opts, out, "calculateCurrentPrice_", addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata)
	return *ret0, err
}

// CalculateCurrentPrice is a free data retrieval call binding the contract method 0x3f67ee0d.
//
// Solidity: function calculateCurrentPrice_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata) constant returns(uint256)
func (_WyvernExchange *WyvernExchangeSession) CalculateCurrentPrice(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte) (*big.Int, error) {
	return _WyvernExchange.Contract.CalculateCurrentPrice(&_WyvernExchange.CallOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata)
}

// CalculateCurrentPrice is a free data retrieval call binding the contract method 0x3f67ee0d.
//
// Solidity: function calculateCurrentPrice_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata) constant returns(uint256)
func (_WyvernExchange *WyvernExchangeCallerSession) CalculateCurrentPrice(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte) (*big.Int, error) {
	return _WyvernExchange.Contract.CalculateCurrentPrice(&_WyvernExchange.CallOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata)
}

// CalculateFinalPrice is a free data retrieval call binding the contract method 0x63d36c0b.
//
// Solidity: function calculateFinalPrice(uint8 side, uint8 saleKind, uint256 basePrice, uint256 extra, uint256 listingTime, uint256 expirationTime) constant returns(uint256)
func (_WyvernExchange *WyvernExchangeCaller) CalculateFinalPrice(opts *bind.CallOpts, side uint8, saleKind uint8, basePrice *big.Int, extra *big.Int, listingTime *big.Int, expirationTime *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WyvernExchange.contract.Call(opts, out, "calculateFinalPrice", side, saleKind, basePrice, extra, listingTime, expirationTime)
	return *ret0, err
}

// CalculateFinalPrice is a free data retrieval call binding the contract method 0x63d36c0b.
//
// Solidity: function calculateFinalPrice(uint8 side, uint8 saleKind, uint256 basePrice, uint256 extra, uint256 listingTime, uint256 expirationTime) constant returns(uint256)
func (_WyvernExchange *WyvernExchangeSession) CalculateFinalPrice(side uint8, saleKind uint8, basePrice *big.Int, extra *big.Int, listingTime *big.Int, expirationTime *big.Int) (*big.Int, error) {
	return _WyvernExchange.Contract.CalculateFinalPrice(&_WyvernExchange.CallOpts, side, saleKind, basePrice, extra, listingTime, expirationTime)
}

// CalculateFinalPrice is a free data retrieval call binding the contract method 0x63d36c0b.
//
// Solidity: function calculateFinalPrice(uint8 side, uint8 saleKind, uint256 basePrice, uint256 extra, uint256 listingTime, uint256 expirationTime) constant returns(uint256)
func (_WyvernExchange *WyvernExchangeCallerSession) CalculateFinalPrice(side uint8, saleKind uint8, basePrice *big.Int, extra *big.Int, listingTime *big.Int, expirationTime *big.Int) (*big.Int, error) {
	return _WyvernExchange.Contract.CalculateFinalPrice(&_WyvernExchange.CallOpts, side, saleKind, basePrice, extra, listingTime, expirationTime)
}

// CalculateMatchPrice is a free data retrieval call binding the contract method 0xd537e131.
//
// Solidity: function calculateMatchPrice_(address[14] addrs, uint256[18] uints, uint8[8] feeMethodsSidesKindsHowToCalls, bytes calldataBuy, bytes calldataSell, bytes replacementPatternBuy, bytes replacementPatternSell, bytes staticExtradataBuy, bytes staticExtradataSell) constant returns(uint256)
func (_WyvernExchange *WyvernExchangeCaller) CalculateMatchPrice(opts *bind.CallOpts, addrs [14]common.Address, uints [18]*big.Int, feeMethodsSidesKindsHowToCalls [8]uint8, calldataBuy []byte, calldataSell []byte, replacementPatternBuy []byte, replacementPatternSell []byte, staticExtradataBuy []byte, staticExtradataSell []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WyvernExchange.contract.Call(opts, out, "calculateMatchPrice_", addrs, uints, feeMethodsSidesKindsHowToCalls, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell)
	return *ret0, err
}

// CalculateMatchPrice is a free data retrieval call binding the contract method 0xd537e131.
//
// Solidity: function calculateMatchPrice_(address[14] addrs, uint256[18] uints, uint8[8] feeMethodsSidesKindsHowToCalls, bytes calldataBuy, bytes calldataSell, bytes replacementPatternBuy, bytes replacementPatternSell, bytes staticExtradataBuy, bytes staticExtradataSell) constant returns(uint256)
func (_WyvernExchange *WyvernExchangeSession) CalculateMatchPrice(addrs [14]common.Address, uints [18]*big.Int, feeMethodsSidesKindsHowToCalls [8]uint8, calldataBuy []byte, calldataSell []byte, replacementPatternBuy []byte, replacementPatternSell []byte, staticExtradataBuy []byte, staticExtradataSell []byte) (*big.Int, error) {
	return _WyvernExchange.Contract.CalculateMatchPrice(&_WyvernExchange.CallOpts, addrs, uints, feeMethodsSidesKindsHowToCalls, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell)
}

// CalculateMatchPrice is a free data retrieval call binding the contract method 0xd537e131.
//
// Solidity: function calculateMatchPrice_(address[14] addrs, uint256[18] uints, uint8[8] feeMethodsSidesKindsHowToCalls, bytes calldataBuy, bytes calldataSell, bytes replacementPatternBuy, bytes replacementPatternSell, bytes staticExtradataBuy, bytes staticExtradataSell) constant returns(uint256)
func (_WyvernExchange *WyvernExchangeCallerSession) CalculateMatchPrice(addrs [14]common.Address, uints [18]*big.Int, feeMethodsSidesKindsHowToCalls [8]uint8, calldataBuy []byte, calldataSell []byte, replacementPatternBuy []byte, replacementPatternSell []byte, staticExtradataBuy []byte, staticExtradataSell []byte) (*big.Int, error) {
	return _WyvernExchange.Contract.CalculateMatchPrice(&_WyvernExchange.CallOpts, addrs, uints, feeMethodsSidesKindsHowToCalls, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell)
}

// CancelledOrFinalized is a free data retrieval call binding the contract method 0x8076f005.
//
// Solidity: function cancelledOrFinalized(bytes32 ) constant returns(bool)
func (_WyvernExchange *WyvernExchangeCaller) CancelledOrFinalized(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WyvernExchange.contract.Call(opts, out, "cancelledOrFinalized", arg0)
	return *ret0, err
}

// CancelledOrFinalized is a free data retrieval call binding the contract method 0x8076f005.
//
// Solidity: function cancelledOrFinalized(bytes32 ) constant returns(bool)
func (_WyvernExchange *WyvernExchangeSession) CancelledOrFinalized(arg0 [32]byte) (bool, error) {
	return _WyvernExchange.Contract.CancelledOrFinalized(&_WyvernExchange.CallOpts, arg0)
}

// CancelledOrFinalized is a free data retrieval call binding the contract method 0x8076f005.
//
// Solidity: function cancelledOrFinalized(bytes32 ) constant returns(bool)
func (_WyvernExchange *WyvernExchangeCallerSession) CancelledOrFinalized(arg0 [32]byte) (bool, error) {
	return _WyvernExchange.Contract.CancelledOrFinalized(&_WyvernExchange.CallOpts, arg0)
}

// Codename is a free data retrieval call binding the contract method 0x31e63199.
//
// Solidity: function codename() constant returns(string)
func (_WyvernExchange *WyvernExchangeCaller) Codename(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _WyvernExchange.contract.Call(opts, out, "codename")
	return *ret0, err
}

// Codename is a free data retrieval call binding the contract method 0x31e63199.
//
// Solidity: function codename() constant returns(string)
func (_WyvernExchange *WyvernExchangeSession) Codename() (string, error) {
	return _WyvernExchange.Contract.Codename(&_WyvernExchange.CallOpts)
}

// Codename is a free data retrieval call binding the contract method 0x31e63199.
//
// Solidity: function codename() constant returns(string)
func (_WyvernExchange *WyvernExchangeCallerSession) Codename() (string, error) {
	return _WyvernExchange.Contract.Codename(&_WyvernExchange.CallOpts)
}

// ExchangeToken is a free data retrieval call binding the contract method 0xa25eb5d9.
//
// Solidity: function exchangeToken() constant returns(address)
func (_WyvernExchange *WyvernExchangeCaller) ExchangeToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WyvernExchange.contract.Call(opts, out, "exchangeToken")
	return *ret0, err
}

// ExchangeToken is a free data retrieval call binding the contract method 0xa25eb5d9.
//
// Solidity: function exchangeToken() constant returns(address)
func (_WyvernExchange *WyvernExchangeSession) ExchangeToken() (common.Address, error) {
	return _WyvernExchange.Contract.ExchangeToken(&_WyvernExchange.CallOpts)
}

// ExchangeToken is a free data retrieval call binding the contract method 0xa25eb5d9.
//
// Solidity: function exchangeToken() constant returns(address)
func (_WyvernExchange *WyvernExchangeCallerSession) ExchangeToken() (common.Address, error) {
	return _WyvernExchange.Contract.ExchangeToken(&_WyvernExchange.CallOpts)
}

// GuardedArrayReplace is a free data retrieval call binding the contract method 0x239e83df.
//
// Solidity: function guardedArrayReplace(bytes array, bytes desired, bytes mask) constant returns(bytes)
func (_WyvernExchange *WyvernExchangeCaller) GuardedArrayReplace(opts *bind.CallOpts, array []byte, desired []byte, mask []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _WyvernExchange.contract.Call(opts, out, "guardedArrayReplace", array, desired, mask)
	return *ret0, err
}

// GuardedArrayReplace is a free data retrieval call binding the contract method 0x239e83df.
//
// Solidity: function guardedArrayReplace(bytes array, bytes desired, bytes mask) constant returns(bytes)
func (_WyvernExchange *WyvernExchangeSession) GuardedArrayReplace(array []byte, desired []byte, mask []byte) ([]byte, error) {
	return _WyvernExchange.Contract.GuardedArrayReplace(&_WyvernExchange.CallOpts, array, desired, mask)
}

// GuardedArrayReplace is a free data retrieval call binding the contract method 0x239e83df.
//
// Solidity: function guardedArrayReplace(bytes array, bytes desired, bytes mask) constant returns(bytes)
func (_WyvernExchange *WyvernExchangeCallerSession) GuardedArrayReplace(array []byte, desired []byte, mask []byte) ([]byte, error) {
	return _WyvernExchange.Contract.GuardedArrayReplace(&_WyvernExchange.CallOpts, array, desired, mask)
}

// HashOrder is a free data retrieval call binding the contract method 0x71d02b38.
//
// Solidity: function hashOrder_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata) constant returns(bytes32)
func (_WyvernExchange *WyvernExchangeCaller) HashOrder(opts *bind.CallOpts, addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _WyvernExchange.contract.Call(opts, out, "hashOrder_", addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata)
	return *ret0, err
}

// HashOrder is a free data retrieval call binding the contract method 0x71d02b38.
//
// Solidity: function hashOrder_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata) constant returns(bytes32)
func (_WyvernExchange *WyvernExchangeSession) HashOrder(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte) ([32]byte, error) {
	return _WyvernExchange.Contract.HashOrder(&_WyvernExchange.CallOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata)
}

// HashOrder is a free data retrieval call binding the contract method 0x71d02b38.
//
// Solidity: function hashOrder_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata) constant returns(bytes32)
func (_WyvernExchange *WyvernExchangeCallerSession) HashOrder(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte) ([32]byte, error) {
	return _WyvernExchange.Contract.HashOrder(&_WyvernExchange.CallOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata)
}

// HashToSign is a free data retrieval call binding the contract method 0x7d766981.
//
// Solidity: function hashToSign_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata) constant returns(bytes32)
func (_WyvernExchange *WyvernExchangeCaller) HashToSign(opts *bind.CallOpts, addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _WyvernExchange.contract.Call(opts, out, "hashToSign_", addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata)
	return *ret0, err
}

// HashToSign is a free data retrieval call binding the contract method 0x7d766981.
//
// Solidity: function hashToSign_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata) constant returns(bytes32)
func (_WyvernExchange *WyvernExchangeSession) HashToSign(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte) ([32]byte, error) {
	return _WyvernExchange.Contract.HashToSign(&_WyvernExchange.CallOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata)
}

// HashToSign is a free data retrieval call binding the contract method 0x7d766981.
//
// Solidity: function hashToSign_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata) constant returns(bytes32)
func (_WyvernExchange *WyvernExchangeCallerSession) HashToSign(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte) ([32]byte, error) {
	return _WyvernExchange.Contract.HashToSign(&_WyvernExchange.CallOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata)
}

// MinimumMakerProtocolFee is a free data retrieval call binding the contract method 0x7ccefc52.
//
// Solidity: function minimumMakerProtocolFee() constant returns(uint256)
func (_WyvernExchange *WyvernExchangeCaller) MinimumMakerProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WyvernExchange.contract.Call(opts, out, "minimumMakerProtocolFee")
	return *ret0, err
}

// MinimumMakerProtocolFee is a free data retrieval call binding the contract method 0x7ccefc52.
//
// Solidity: function minimumMakerProtocolFee() constant returns(uint256)
func (_WyvernExchange *WyvernExchangeSession) MinimumMakerProtocolFee() (*big.Int, error) {
	return _WyvernExchange.Contract.MinimumMakerProtocolFee(&_WyvernExchange.CallOpts)
}

// MinimumMakerProtocolFee is a free data retrieval call binding the contract method 0x7ccefc52.
//
// Solidity: function minimumMakerProtocolFee() constant returns(uint256)
func (_WyvernExchange *WyvernExchangeCallerSession) MinimumMakerProtocolFee() (*big.Int, error) {
	return _WyvernExchange.Contract.MinimumMakerProtocolFee(&_WyvernExchange.CallOpts)
}

// MinimumTakerProtocolFee is a free data retrieval call binding the contract method 0x28a8ee68.
//
// Solidity: function minimumTakerProtocolFee() constant returns(uint256)
func (_WyvernExchange *WyvernExchangeCaller) MinimumTakerProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WyvernExchange.contract.Call(opts, out, "minimumTakerProtocolFee")
	return *ret0, err
}

// MinimumTakerProtocolFee is a free data retrieval call binding the contract method 0x28a8ee68.
//
// Solidity: function minimumTakerProtocolFee() constant returns(uint256)
func (_WyvernExchange *WyvernExchangeSession) MinimumTakerProtocolFee() (*big.Int, error) {
	return _WyvernExchange.Contract.MinimumTakerProtocolFee(&_WyvernExchange.CallOpts)
}

// MinimumTakerProtocolFee is a free data retrieval call binding the contract method 0x28a8ee68.
//
// Solidity: function minimumTakerProtocolFee() constant returns(uint256)
func (_WyvernExchange *WyvernExchangeCallerSession) MinimumTakerProtocolFee() (*big.Int, error) {
	return _WyvernExchange.Contract.MinimumTakerProtocolFee(&_WyvernExchange.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_WyvernExchange *WyvernExchangeCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _WyvernExchange.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_WyvernExchange *WyvernExchangeSession) Name() (string, error) {
	return _WyvernExchange.Contract.Name(&_WyvernExchange.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_WyvernExchange *WyvernExchangeCallerSession) Name() (string, error) {
	return _WyvernExchange.Contract.Name(&_WyvernExchange.CallOpts)
}

// OrderCalldataCanMatch is a free data retrieval call binding the contract method 0x562b2ebc.
//
// Solidity: function orderCalldataCanMatch(bytes buyCalldata, bytes buyReplacementPattern, bytes sellCalldata, bytes sellReplacementPattern) constant returns(bool)
func (_WyvernExchange *WyvernExchangeCaller) OrderCalldataCanMatch(opts *bind.CallOpts, buyCalldata []byte, buyReplacementPattern []byte, sellCalldata []byte, sellReplacementPattern []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WyvernExchange.contract.Call(opts, out, "orderCalldataCanMatch", buyCalldata, buyReplacementPattern, sellCalldata, sellReplacementPattern)
	return *ret0, err
}

// OrderCalldataCanMatch is a free data retrieval call binding the contract method 0x562b2ebc.
//
// Solidity: function orderCalldataCanMatch(bytes buyCalldata, bytes buyReplacementPattern, bytes sellCalldata, bytes sellReplacementPattern) constant returns(bool)
func (_WyvernExchange *WyvernExchangeSession) OrderCalldataCanMatch(buyCalldata []byte, buyReplacementPattern []byte, sellCalldata []byte, sellReplacementPattern []byte) (bool, error) {
	return _WyvernExchange.Contract.OrderCalldataCanMatch(&_WyvernExchange.CallOpts, buyCalldata, buyReplacementPattern, sellCalldata, sellReplacementPattern)
}

// OrderCalldataCanMatch is a free data retrieval call binding the contract method 0x562b2ebc.
//
// Solidity: function orderCalldataCanMatch(bytes buyCalldata, bytes buyReplacementPattern, bytes sellCalldata, bytes sellReplacementPattern) constant returns(bool)
func (_WyvernExchange *WyvernExchangeCallerSession) OrderCalldataCanMatch(buyCalldata []byte, buyReplacementPattern []byte, sellCalldata []byte, sellReplacementPattern []byte) (bool, error) {
	return _WyvernExchange.Contract.OrderCalldataCanMatch(&_WyvernExchange.CallOpts, buyCalldata, buyReplacementPattern, sellCalldata, sellReplacementPattern)
}

// OrdersCanMatch is a free data retrieval call binding the contract method 0x72593b4c.
//
// Solidity: function ordersCanMatch_(address[14] addrs, uint256[18] uints, uint8[8] feeMethodsSidesKindsHowToCalls, bytes calldataBuy, bytes calldataSell, bytes replacementPatternBuy, bytes replacementPatternSell, bytes staticExtradataBuy, bytes staticExtradataSell) constant returns(bool)
func (_WyvernExchange *WyvernExchangeCaller) OrdersCanMatch(opts *bind.CallOpts, addrs [14]common.Address, uints [18]*big.Int, feeMethodsSidesKindsHowToCalls [8]uint8, calldataBuy []byte, calldataSell []byte, replacementPatternBuy []byte, replacementPatternSell []byte, staticExtradataBuy []byte, staticExtradataSell []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WyvernExchange.contract.Call(opts, out, "ordersCanMatch_", addrs, uints, feeMethodsSidesKindsHowToCalls, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell)
	return *ret0, err
}

// OrdersCanMatch is a free data retrieval call binding the contract method 0x72593b4c.
//
// Solidity: function ordersCanMatch_(address[14] addrs, uint256[18] uints, uint8[8] feeMethodsSidesKindsHowToCalls, bytes calldataBuy, bytes calldataSell, bytes replacementPatternBuy, bytes replacementPatternSell, bytes staticExtradataBuy, bytes staticExtradataSell) constant returns(bool)
func (_WyvernExchange *WyvernExchangeSession) OrdersCanMatch(addrs [14]common.Address, uints [18]*big.Int, feeMethodsSidesKindsHowToCalls [8]uint8, calldataBuy []byte, calldataSell []byte, replacementPatternBuy []byte, replacementPatternSell []byte, staticExtradataBuy []byte, staticExtradataSell []byte) (bool, error) {
	return _WyvernExchange.Contract.OrdersCanMatch(&_WyvernExchange.CallOpts, addrs, uints, feeMethodsSidesKindsHowToCalls, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell)
}

// OrdersCanMatch is a free data retrieval call binding the contract method 0x72593b4c.
//
// Solidity: function ordersCanMatch_(address[14] addrs, uint256[18] uints, uint8[8] feeMethodsSidesKindsHowToCalls, bytes calldataBuy, bytes calldataSell, bytes replacementPatternBuy, bytes replacementPatternSell, bytes staticExtradataBuy, bytes staticExtradataSell) constant returns(bool)
func (_WyvernExchange *WyvernExchangeCallerSession) OrdersCanMatch(addrs [14]common.Address, uints [18]*big.Int, feeMethodsSidesKindsHowToCalls [8]uint8, calldataBuy []byte, calldataSell []byte, replacementPatternBuy []byte, replacementPatternSell []byte, staticExtradataBuy []byte, staticExtradataSell []byte) (bool, error) {
	return _WyvernExchange.Contract.OrdersCanMatch(&_WyvernExchange.CallOpts, addrs, uints, feeMethodsSidesKindsHowToCalls, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_WyvernExchange *WyvernExchangeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WyvernExchange.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_WyvernExchange *WyvernExchangeSession) Owner() (common.Address, error) {
	return _WyvernExchange.Contract.Owner(&_WyvernExchange.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_WyvernExchange *WyvernExchangeCallerSession) Owner() (common.Address, error) {
	return _WyvernExchange.Contract.Owner(&_WyvernExchange.CallOpts)
}

// ProtocolFeeRecipient is a free data retrieval call binding the contract method 0x64df049e.
//
// Solidity: function protocolFeeRecipient() constant returns(address)
func (_WyvernExchange *WyvernExchangeCaller) ProtocolFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WyvernExchange.contract.Call(opts, out, "protocolFeeRecipient")
	return *ret0, err
}

// ProtocolFeeRecipient is a free data retrieval call binding the contract method 0x64df049e.
//
// Solidity: function protocolFeeRecipient() constant returns(address)
func (_WyvernExchange *WyvernExchangeSession) ProtocolFeeRecipient() (common.Address, error) {
	return _WyvernExchange.Contract.ProtocolFeeRecipient(&_WyvernExchange.CallOpts)
}

// ProtocolFeeRecipient is a free data retrieval call binding the contract method 0x64df049e.
//
// Solidity: function protocolFeeRecipient() constant returns(address)
func (_WyvernExchange *WyvernExchangeCallerSession) ProtocolFeeRecipient() (common.Address, error) {
	return _WyvernExchange.Contract.ProtocolFeeRecipient(&_WyvernExchange.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_WyvernExchange *WyvernExchangeCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WyvernExchange.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_WyvernExchange *WyvernExchangeSession) Registry() (common.Address, error) {
	return _WyvernExchange.Contract.Registry(&_WyvernExchange.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_WyvernExchange *WyvernExchangeCallerSession) Registry() (common.Address, error) {
	return _WyvernExchange.Contract.Registry(&_WyvernExchange.CallOpts)
}

// StaticCall is a free data retrieval call binding the contract method 0x10796a47.
//
// Solidity: function staticCall(address target, bytes calldata, bytes extradata) constant returns(bool result)
func (_WyvernExchange *WyvernExchangeCaller) StaticCall(opts *bind.CallOpts, target common.Address, calldata []byte, extradata []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WyvernExchange.contract.Call(opts, out, "staticCall", target, calldata, extradata)
	return *ret0, err
}

// StaticCall is a free data retrieval call binding the contract method 0x10796a47.
//
// Solidity: function staticCall(address target, bytes calldata, bytes extradata) constant returns(bool result)
func (_WyvernExchange *WyvernExchangeSession) StaticCall(target common.Address, calldata []byte, extradata []byte) (bool, error) {
	return _WyvernExchange.Contract.StaticCall(&_WyvernExchange.CallOpts, target, calldata, extradata)
}

// StaticCall is a free data retrieval call binding the contract method 0x10796a47.
//
// Solidity: function staticCall(address target, bytes calldata, bytes extradata) constant returns(bool result)
func (_WyvernExchange *WyvernExchangeCallerSession) StaticCall(target common.Address, calldata []byte, extradata []byte) (bool, error) {
	return _WyvernExchange.Contract.StaticCall(&_WyvernExchange.CallOpts, target, calldata, extradata)
}

// TestCopy is a free data retrieval call binding the contract method 0x3e1e292a.
//
// Solidity: function testCopy(bytes arrToCopy) constant returns(bytes)
func (_WyvernExchange *WyvernExchangeCaller) TestCopy(opts *bind.CallOpts, arrToCopy []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _WyvernExchange.contract.Call(opts, out, "testCopy", arrToCopy)
	return *ret0, err
}

// TestCopy is a free data retrieval call binding the contract method 0x3e1e292a.
//
// Solidity: function testCopy(bytes arrToCopy) constant returns(bytes)
func (_WyvernExchange *WyvernExchangeSession) TestCopy(arrToCopy []byte) ([]byte, error) {
	return _WyvernExchange.Contract.TestCopy(&_WyvernExchange.CallOpts, arrToCopy)
}

// TestCopy is a free data retrieval call binding the contract method 0x3e1e292a.
//
// Solidity: function testCopy(bytes arrToCopy) constant returns(bytes)
func (_WyvernExchange *WyvernExchangeCallerSession) TestCopy(arrToCopy []byte) ([]byte, error) {
	return _WyvernExchange.Contract.TestCopy(&_WyvernExchange.CallOpts, arrToCopy)
}

// TestCopyAddress is a free data retrieval call binding the contract method 0x3464af6a.
//
// Solidity: function testCopyAddress(address addr) constant returns(bytes)
func (_WyvernExchange *WyvernExchangeCaller) TestCopyAddress(opts *bind.CallOpts, addr common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _WyvernExchange.contract.Call(opts, out, "testCopyAddress", addr)
	return *ret0, err
}

// TestCopyAddress is a free data retrieval call binding the contract method 0x3464af6a.
//
// Solidity: function testCopyAddress(address addr) constant returns(bytes)
func (_WyvernExchange *WyvernExchangeSession) TestCopyAddress(addr common.Address) ([]byte, error) {
	return _WyvernExchange.Contract.TestCopyAddress(&_WyvernExchange.CallOpts, addr)
}

// TestCopyAddress is a free data retrieval call binding the contract method 0x3464af6a.
//
// Solidity: function testCopyAddress(address addr) constant returns(bytes)
func (_WyvernExchange *WyvernExchangeCallerSession) TestCopyAddress(addr common.Address) ([]byte, error) {
	return _WyvernExchange.Contract.TestCopyAddress(&_WyvernExchange.CallOpts, addr)
}

// TokenTransferProxy is a free data retrieval call binding the contract method 0x0eefdbad.
//
// Solidity: function tokenTransferProxy() constant returns(address)
func (_WyvernExchange *WyvernExchangeCaller) TokenTransferProxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WyvernExchange.contract.Call(opts, out, "tokenTransferProxy")
	return *ret0, err
}

// TokenTransferProxy is a free data retrieval call binding the contract method 0x0eefdbad.
//
// Solidity: function tokenTransferProxy() constant returns(address)
func (_WyvernExchange *WyvernExchangeSession) TokenTransferProxy() (common.Address, error) {
	return _WyvernExchange.Contract.TokenTransferProxy(&_WyvernExchange.CallOpts)
}

// TokenTransferProxy is a free data retrieval call binding the contract method 0x0eefdbad.
//
// Solidity: function tokenTransferProxy() constant returns(address)
func (_WyvernExchange *WyvernExchangeCallerSession) TokenTransferProxy() (common.Address, error) {
	return _WyvernExchange.Contract.TokenTransferProxy(&_WyvernExchange.CallOpts)
}

// ValidateOrderParameters is a free data retrieval call binding the contract method 0xca595b9a.
//
// Solidity: function validateOrderParameters_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata) constant returns(bool)
func (_WyvernExchange *WyvernExchangeCaller) ValidateOrderParameters(opts *bind.CallOpts, addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WyvernExchange.contract.Call(opts, out, "validateOrderParameters_", addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata)
	return *ret0, err
}

// ValidateOrderParameters is a free data retrieval call binding the contract method 0xca595b9a.
//
// Solidity: function validateOrderParameters_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata) constant returns(bool)
func (_WyvernExchange *WyvernExchangeSession) ValidateOrderParameters(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte) (bool, error) {
	return _WyvernExchange.Contract.ValidateOrderParameters(&_WyvernExchange.CallOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata)
}

// ValidateOrderParameters is a free data retrieval call binding the contract method 0xca595b9a.
//
// Solidity: function validateOrderParameters_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata) constant returns(bool)
func (_WyvernExchange *WyvernExchangeCallerSession) ValidateOrderParameters(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte) (bool, error) {
	return _WyvernExchange.Contract.ValidateOrderParameters(&_WyvernExchange.CallOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata)
}

// ValidateOrder is a free data retrieval call binding the contract method 0x60bef33a.
//
// Solidity: function validateOrder_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata, uint8 v, bytes32 r, bytes32 s) constant returns(bool)
func (_WyvernExchange *WyvernExchangeCaller) ValidateOrder(opts *bind.CallOpts, addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WyvernExchange.contract.Call(opts, out, "validateOrder_", addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata, v, r, s)
	return *ret0, err
}

// ValidateOrder is a free data retrieval call binding the contract method 0x60bef33a.
//
// Solidity: function validateOrder_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata, uint8 v, bytes32 r, bytes32 s) constant returns(bool)
func (_WyvernExchange *WyvernExchangeSession) ValidateOrder(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _WyvernExchange.Contract.ValidateOrder(&_WyvernExchange.CallOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata, v, r, s)
}

// ValidateOrder is a free data retrieval call binding the contract method 0x60bef33a.
//
// Solidity: function validateOrder_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata, uint8 v, bytes32 r, bytes32 s) constant returns(bool)
func (_WyvernExchange *WyvernExchangeCallerSession) ValidateOrder(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _WyvernExchange.Contract.ValidateOrder(&_WyvernExchange.CallOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata, v, r, s)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_WyvernExchange *WyvernExchangeCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _WyvernExchange.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_WyvernExchange *WyvernExchangeSession) Version() (string, error) {
	return _WyvernExchange.Contract.Version(&_WyvernExchange.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_WyvernExchange *WyvernExchangeCallerSession) Version() (string, error) {
	return _WyvernExchange.Contract.Version(&_WyvernExchange.CallOpts)
}

// ApproveOrder is a paid mutator transaction binding the contract method 0x79666868.
//
// Solidity: function approveOrder_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata, bool orderbookInclusionDesired) returns()
func (_WyvernExchange *WyvernExchangeTransactor) ApproveOrder(opts *bind.TransactOpts, addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte, orderbookInclusionDesired bool) (*types.Transaction, error) {
	return _WyvernExchange.contract.Transact(opts, "approveOrder_", addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata, orderbookInclusionDesired)
}

// ApproveOrder is a paid mutator transaction binding the contract method 0x79666868.
//
// Solidity: function approveOrder_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata, bool orderbookInclusionDesired) returns()
func (_WyvernExchange *WyvernExchangeSession) ApproveOrder(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte, orderbookInclusionDesired bool) (*types.Transaction, error) {
	return _WyvernExchange.Contract.ApproveOrder(&_WyvernExchange.TransactOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata, orderbookInclusionDesired)
}

// ApproveOrder is a paid mutator transaction binding the contract method 0x79666868.
//
// Solidity: function approveOrder_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata, bool orderbookInclusionDesired) returns()
func (_WyvernExchange *WyvernExchangeTransactorSession) ApproveOrder(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte, orderbookInclusionDesired bool) (*types.Transaction, error) {
	return _WyvernExchange.Contract.ApproveOrder(&_WyvernExchange.TransactOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata, orderbookInclusionDesired)
}

// AtomicMatch is a paid mutator transaction binding the contract method 0xab834bab.
//
// Solidity: function atomicMatch_(address[14] addrs, uint256[18] uints, uint8[8] feeMethodsSidesKindsHowToCalls, bytes calldataBuy, bytes calldataSell, bytes replacementPatternBuy, bytes replacementPatternSell, bytes staticExtradataBuy, bytes staticExtradataSell, uint8[2] vs, bytes32[5] rssMetadata) returns()
func (_WyvernExchange *WyvernExchangeTransactor) AtomicMatch(opts *bind.TransactOpts, addrs [14]common.Address, uints [18]*big.Int, feeMethodsSidesKindsHowToCalls [8]uint8, calldataBuy []byte, calldataSell []byte, replacementPatternBuy []byte, replacementPatternSell []byte, staticExtradataBuy []byte, staticExtradataSell []byte, vs [2]uint8, rssMetadata [5][32]byte) (*types.Transaction, error) {
	return _WyvernExchange.contract.Transact(opts, "atomicMatch_", addrs, uints, feeMethodsSidesKindsHowToCalls, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell, vs, rssMetadata)
}

// AtomicMatch is a paid mutator transaction binding the contract method 0xab834bab.
//
// Solidity: function atomicMatch_(address[14] addrs, uint256[18] uints, uint8[8] feeMethodsSidesKindsHowToCalls, bytes calldataBuy, bytes calldataSell, bytes replacementPatternBuy, bytes replacementPatternSell, bytes staticExtradataBuy, bytes staticExtradataSell, uint8[2] vs, bytes32[5] rssMetadata) returns()
func (_WyvernExchange *WyvernExchangeSession) AtomicMatch(addrs [14]common.Address, uints [18]*big.Int, feeMethodsSidesKindsHowToCalls [8]uint8, calldataBuy []byte, calldataSell []byte, replacementPatternBuy []byte, replacementPatternSell []byte, staticExtradataBuy []byte, staticExtradataSell []byte, vs [2]uint8, rssMetadata [5][32]byte) (*types.Transaction, error) {
	return _WyvernExchange.Contract.AtomicMatch(&_WyvernExchange.TransactOpts, addrs, uints, feeMethodsSidesKindsHowToCalls, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell, vs, rssMetadata)
}

// AtomicMatch is a paid mutator transaction binding the contract method 0xab834bab.
//
// Solidity: function atomicMatch_(address[14] addrs, uint256[18] uints, uint8[8] feeMethodsSidesKindsHowToCalls, bytes calldataBuy, bytes calldataSell, bytes replacementPatternBuy, bytes replacementPatternSell, bytes staticExtradataBuy, bytes staticExtradataSell, uint8[2] vs, bytes32[5] rssMetadata) returns()
func (_WyvernExchange *WyvernExchangeTransactorSession) AtomicMatch(addrs [14]common.Address, uints [18]*big.Int, feeMethodsSidesKindsHowToCalls [8]uint8, calldataBuy []byte, calldataSell []byte, replacementPatternBuy []byte, replacementPatternSell []byte, staticExtradataBuy []byte, staticExtradataSell []byte, vs [2]uint8, rssMetadata [5][32]byte) (*types.Transaction, error) {
	return _WyvernExchange.Contract.AtomicMatch(&_WyvernExchange.TransactOpts, addrs, uints, feeMethodsSidesKindsHowToCalls, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell, vs, rssMetadata)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xa8a41c70.
//
// Solidity: function cancelOrder_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata, uint8 v, bytes32 r, bytes32 s) returns()
func (_WyvernExchange *WyvernExchangeTransactor) CancelOrder(opts *bind.TransactOpts, addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WyvernExchange.contract.Transact(opts, "cancelOrder_", addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata, v, r, s)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xa8a41c70.
//
// Solidity: function cancelOrder_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata, uint8 v, bytes32 r, bytes32 s) returns()
func (_WyvernExchange *WyvernExchangeSession) CancelOrder(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WyvernExchange.Contract.CancelOrder(&_WyvernExchange.TransactOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata, v, r, s)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xa8a41c70.
//
// Solidity: function cancelOrder_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata, uint8 v, bytes32 r, bytes32 s) returns()
func (_WyvernExchange *WyvernExchangeTransactorSession) CancelOrder(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WyvernExchange.Contract.CancelOrder(&_WyvernExchange.TransactOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata, v, r, s)
}

// ChangeMinimumMakerProtocolFee is a paid mutator transaction binding the contract method 0x14350c24.
//
// Solidity: function changeMinimumMakerProtocolFee(uint256 newMinimumMakerProtocolFee) returns()
func (_WyvernExchange *WyvernExchangeTransactor) ChangeMinimumMakerProtocolFee(opts *bind.TransactOpts, newMinimumMakerProtocolFee *big.Int) (*types.Transaction, error) {
	return _WyvernExchange.contract.Transact(opts, "changeMinimumMakerProtocolFee", newMinimumMakerProtocolFee)
}

// ChangeMinimumMakerProtocolFee is a paid mutator transaction binding the contract method 0x14350c24.
//
// Solidity: function changeMinimumMakerProtocolFee(uint256 newMinimumMakerProtocolFee) returns()
func (_WyvernExchange *WyvernExchangeSession) ChangeMinimumMakerProtocolFee(newMinimumMakerProtocolFee *big.Int) (*types.Transaction, error) {
	return _WyvernExchange.Contract.ChangeMinimumMakerProtocolFee(&_WyvernExchange.TransactOpts, newMinimumMakerProtocolFee)
}

// ChangeMinimumMakerProtocolFee is a paid mutator transaction binding the contract method 0x14350c24.
//
// Solidity: function changeMinimumMakerProtocolFee(uint256 newMinimumMakerProtocolFee) returns()
func (_WyvernExchange *WyvernExchangeTransactorSession) ChangeMinimumMakerProtocolFee(newMinimumMakerProtocolFee *big.Int) (*types.Transaction, error) {
	return _WyvernExchange.Contract.ChangeMinimumMakerProtocolFee(&_WyvernExchange.TransactOpts, newMinimumMakerProtocolFee)
}

// ChangeMinimumTakerProtocolFee is a paid mutator transaction binding the contract method 0x1a6b13e2.
//
// Solidity: function changeMinimumTakerProtocolFee(uint256 newMinimumTakerProtocolFee) returns()
func (_WyvernExchange *WyvernExchangeTransactor) ChangeMinimumTakerProtocolFee(opts *bind.TransactOpts, newMinimumTakerProtocolFee *big.Int) (*types.Transaction, error) {
	return _WyvernExchange.contract.Transact(opts, "changeMinimumTakerProtocolFee", newMinimumTakerProtocolFee)
}

// ChangeMinimumTakerProtocolFee is a paid mutator transaction binding the contract method 0x1a6b13e2.
//
// Solidity: function changeMinimumTakerProtocolFee(uint256 newMinimumTakerProtocolFee) returns()
func (_WyvernExchange *WyvernExchangeSession) ChangeMinimumTakerProtocolFee(newMinimumTakerProtocolFee *big.Int) (*types.Transaction, error) {
	return _WyvernExchange.Contract.ChangeMinimumTakerProtocolFee(&_WyvernExchange.TransactOpts, newMinimumTakerProtocolFee)
}

// ChangeMinimumTakerProtocolFee is a paid mutator transaction binding the contract method 0x1a6b13e2.
//
// Solidity: function changeMinimumTakerProtocolFee(uint256 newMinimumTakerProtocolFee) returns()
func (_WyvernExchange *WyvernExchangeTransactorSession) ChangeMinimumTakerProtocolFee(newMinimumTakerProtocolFee *big.Int) (*types.Transaction, error) {
	return _WyvernExchange.Contract.ChangeMinimumTakerProtocolFee(&_WyvernExchange.TransactOpts, newMinimumTakerProtocolFee)
}

// ChangeProtocolFeeRecipient is a paid mutator transaction binding the contract method 0x514f0330.
//
// Solidity: function changeProtocolFeeRecipient(address newProtocolFeeRecipient) returns()
func (_WyvernExchange *WyvernExchangeTransactor) ChangeProtocolFeeRecipient(opts *bind.TransactOpts, newProtocolFeeRecipient common.Address) (*types.Transaction, error) {
	return _WyvernExchange.contract.Transact(opts, "changeProtocolFeeRecipient", newProtocolFeeRecipient)
}

// ChangeProtocolFeeRecipient is a paid mutator transaction binding the contract method 0x514f0330.
//
// Solidity: function changeProtocolFeeRecipient(address newProtocolFeeRecipient) returns()
func (_WyvernExchange *WyvernExchangeSession) ChangeProtocolFeeRecipient(newProtocolFeeRecipient common.Address) (*types.Transaction, error) {
	return _WyvernExchange.Contract.ChangeProtocolFeeRecipient(&_WyvernExchange.TransactOpts, newProtocolFeeRecipient)
}

// ChangeProtocolFeeRecipient is a paid mutator transaction binding the contract method 0x514f0330.
//
// Solidity: function changeProtocolFeeRecipient(address newProtocolFeeRecipient) returns()
func (_WyvernExchange *WyvernExchangeTransactorSession) ChangeProtocolFeeRecipient(newProtocolFeeRecipient common.Address) (*types.Transaction, error) {
	return _WyvernExchange.Contract.ChangeProtocolFeeRecipient(&_WyvernExchange.TransactOpts, newProtocolFeeRecipient)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WyvernExchange *WyvernExchangeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WyvernExchange.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WyvernExchange *WyvernExchangeSession) RenounceOwnership() (*types.Transaction, error) {
	return _WyvernExchange.Contract.RenounceOwnership(&_WyvernExchange.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WyvernExchange *WyvernExchangeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WyvernExchange.Contract.RenounceOwnership(&_WyvernExchange.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WyvernExchange *WyvernExchangeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WyvernExchange.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WyvernExchange *WyvernExchangeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WyvernExchange.Contract.TransferOwnership(&_WyvernExchange.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WyvernExchange *WyvernExchangeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WyvernExchange.Contract.TransferOwnership(&_WyvernExchange.TransactOpts, newOwner)
}

// WyvernExchangeOrderApprovedPartOneIterator is returned from FilterOrderApprovedPartOne and is used to iterate over the raw logs and unpacked data for OrderApprovedPartOne events raised by the WyvernExchange contract.
type WyvernExchangeOrderApprovedPartOneIterator struct {
	Event *WyvernExchangeOrderApprovedPartOne // Event containing the contract specifics and raw log

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
func (it *WyvernExchangeOrderApprovedPartOneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WyvernExchangeOrderApprovedPartOne)
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
		it.Event = new(WyvernExchangeOrderApprovedPartOne)
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
func (it *WyvernExchangeOrderApprovedPartOneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WyvernExchangeOrderApprovedPartOneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WyvernExchangeOrderApprovedPartOne represents a OrderApprovedPartOne event raised by the WyvernExchange contract.
type WyvernExchangeOrderApprovedPartOne struct {
	Hash             [32]byte
	Exchange         common.Address
	Maker            common.Address
	Taker            common.Address
	MakerRelayerFee  *big.Int
	TakerRelayerFee  *big.Int
	MakerProtocolFee *big.Int
	TakerProtocolFee *big.Int
	FeeRecipient     common.Address
	FeeMethod        uint8
	Side             uint8
	SaleKind         uint8
	Target           common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOrderApprovedPartOne is a free log retrieval operation binding the contract event 0x90c7f9f5b58c15f0f635bfb99f55d3d78fdbef3559e7d8abf5c81052a5276622.
//
// Solidity: event OrderApprovedPartOne(bytes32 indexed hash, address exchange, address indexed maker, address taker, uint256 makerRelayerFee, uint256 takerRelayerFee, uint256 makerProtocolFee, uint256 takerProtocolFee, address indexed feeRecipient, uint8 feeMethod, uint8 side, uint8 saleKind, address target)
func (_WyvernExchange *WyvernExchangeFilterer) FilterOrderApprovedPartOne(opts *bind.FilterOpts, hash [][32]byte, maker []common.Address, feeRecipient []common.Address) (*WyvernExchangeOrderApprovedPartOneIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}

	logs, sub, err := _WyvernExchange.contract.FilterLogs(opts, "OrderApprovedPartOne", hashRule, makerRule, feeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &WyvernExchangeOrderApprovedPartOneIterator{contract: _WyvernExchange.contract, event: "OrderApprovedPartOne", logs: logs, sub: sub}, nil
}

// WatchOrderApprovedPartOne is a free log subscription operation binding the contract event 0x90c7f9f5b58c15f0f635bfb99f55d3d78fdbef3559e7d8abf5c81052a5276622.
//
// Solidity: event OrderApprovedPartOne(bytes32 indexed hash, address exchange, address indexed maker, address taker, uint256 makerRelayerFee, uint256 takerRelayerFee, uint256 makerProtocolFee, uint256 takerProtocolFee, address indexed feeRecipient, uint8 feeMethod, uint8 side, uint8 saleKind, address target)
func (_WyvernExchange *WyvernExchangeFilterer) WatchOrderApprovedPartOne(opts *bind.WatchOpts, sink chan<- *WyvernExchangeOrderApprovedPartOne, hash [][32]byte, maker []common.Address, feeRecipient []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}

	logs, sub, err := _WyvernExchange.contract.WatchLogs(opts, "OrderApprovedPartOne", hashRule, makerRule, feeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WyvernExchangeOrderApprovedPartOne)
				if err := _WyvernExchange.contract.UnpackLog(event, "OrderApprovedPartOne", log); err != nil {
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

// WyvernExchangeOrderApprovedPartTwoIterator is returned from FilterOrderApprovedPartTwo and is used to iterate over the raw logs and unpacked data for OrderApprovedPartTwo events raised by the WyvernExchange contract.
type WyvernExchangeOrderApprovedPartTwoIterator struct {
	Event *WyvernExchangeOrderApprovedPartTwo // Event containing the contract specifics and raw log

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
func (it *WyvernExchangeOrderApprovedPartTwoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WyvernExchangeOrderApprovedPartTwo)
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
		it.Event = new(WyvernExchangeOrderApprovedPartTwo)
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
func (it *WyvernExchangeOrderApprovedPartTwoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WyvernExchangeOrderApprovedPartTwoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WyvernExchangeOrderApprovedPartTwo represents a OrderApprovedPartTwo event raised by the WyvernExchange contract.
type WyvernExchangeOrderApprovedPartTwo struct {
	Hash                      [32]byte
	HowToCall                 uint8
	Calldata                  []byte
	ReplacementPattern        []byte
	StaticTarget              common.Address
	StaticExtradata           []byte
	PaymentToken              common.Address
	BasePrice                 *big.Int
	Extra                     *big.Int
	ListingTime               *big.Int
	ExpirationTime            *big.Int
	Salt                      *big.Int
	OrderbookInclusionDesired bool
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterOrderApprovedPartTwo is a free log retrieval operation binding the contract event 0xe55393c778364e440d958b39ac1debd99dcfae3775a8a04d1e79124adf6a2d08.
//
// Solidity: event OrderApprovedPartTwo(bytes32 indexed hash, uint8 howToCall, bytes calldata, bytes replacementPattern, address staticTarget, bytes staticExtradata, address paymentToken, uint256 basePrice, uint256 extra, uint256 listingTime, uint256 expirationTime, uint256 salt, bool orderbookInclusionDesired)
func (_WyvernExchange *WyvernExchangeFilterer) FilterOrderApprovedPartTwo(opts *bind.FilterOpts, hash [][32]byte) (*WyvernExchangeOrderApprovedPartTwoIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _WyvernExchange.contract.FilterLogs(opts, "OrderApprovedPartTwo", hashRule)
	if err != nil {
		return nil, err
	}
	return &WyvernExchangeOrderApprovedPartTwoIterator{contract: _WyvernExchange.contract, event: "OrderApprovedPartTwo", logs: logs, sub: sub}, nil
}

// WatchOrderApprovedPartTwo is a free log subscription operation binding the contract event 0xe55393c778364e440d958b39ac1debd99dcfae3775a8a04d1e79124adf6a2d08.
//
// Solidity: event OrderApprovedPartTwo(bytes32 indexed hash, uint8 howToCall, bytes calldata, bytes replacementPattern, address staticTarget, bytes staticExtradata, address paymentToken, uint256 basePrice, uint256 extra, uint256 listingTime, uint256 expirationTime, uint256 salt, bool orderbookInclusionDesired)
func (_WyvernExchange *WyvernExchangeFilterer) WatchOrderApprovedPartTwo(opts *bind.WatchOpts, sink chan<- *WyvernExchangeOrderApprovedPartTwo, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _WyvernExchange.contract.WatchLogs(opts, "OrderApprovedPartTwo", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WyvernExchangeOrderApprovedPartTwo)
				if err := _WyvernExchange.contract.UnpackLog(event, "OrderApprovedPartTwo", log); err != nil {
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

// WyvernExchangeOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the WyvernExchange contract.
type WyvernExchangeOrderCancelledIterator struct {
	Event *WyvernExchangeOrderCancelled // Event containing the contract specifics and raw log

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
func (it *WyvernExchangeOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WyvernExchangeOrderCancelled)
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
		it.Event = new(WyvernExchangeOrderCancelled)
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
func (it *WyvernExchangeOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WyvernExchangeOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WyvernExchangeOrderCancelled represents a OrderCancelled event raised by the WyvernExchange contract.
type WyvernExchangeOrderCancelled struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 indexed hash)
func (_WyvernExchange *WyvernExchangeFilterer) FilterOrderCancelled(opts *bind.FilterOpts, hash [][32]byte) (*WyvernExchangeOrderCancelledIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _WyvernExchange.contract.FilterLogs(opts, "OrderCancelled", hashRule)
	if err != nil {
		return nil, err
	}
	return &WyvernExchangeOrderCancelledIterator{contract: _WyvernExchange.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 indexed hash)
func (_WyvernExchange *WyvernExchangeFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *WyvernExchangeOrderCancelled, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _WyvernExchange.contract.WatchLogs(opts, "OrderCancelled", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WyvernExchangeOrderCancelled)
				if err := _WyvernExchange.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// WyvernExchangeOrdersMatchedIterator is returned from FilterOrdersMatched and is used to iterate over the raw logs and unpacked data for OrdersMatched events raised by the WyvernExchange contract.
type WyvernExchangeOrdersMatchedIterator struct {
	Event *WyvernExchangeOrdersMatched // Event containing the contract specifics and raw log

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
func (it *WyvernExchangeOrdersMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WyvernExchangeOrdersMatched)
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
		it.Event = new(WyvernExchangeOrdersMatched)
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
func (it *WyvernExchangeOrdersMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WyvernExchangeOrdersMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WyvernExchangeOrdersMatched represents a OrdersMatched event raised by the WyvernExchange contract.
type WyvernExchangeOrdersMatched struct {
	BuyHash  [32]byte
	SellHash [32]byte
	Maker    common.Address
	Taker    common.Address
	Price    *big.Int
	Metadata [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOrdersMatched is a free log retrieval operation binding the contract event 0xc4109843e0b7d514e4c093114b863f8e7d8d9a458c372cd51bfe526b588006c9.
//
// Solidity: event OrdersMatched(bytes32 buyHash, bytes32 sellHash, address indexed maker, address indexed taker, uint256 price, bytes32 indexed metadata)
func (_WyvernExchange *WyvernExchangeFilterer) FilterOrdersMatched(opts *bind.FilterOpts, maker []common.Address, taker []common.Address, metadata [][32]byte) (*WyvernExchangeOrdersMatchedIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	var metadataRule []interface{}
	for _, metadataItem := range metadata {
		metadataRule = append(metadataRule, metadataItem)
	}

	logs, sub, err := _WyvernExchange.contract.FilterLogs(opts, "OrdersMatched", makerRule, takerRule, metadataRule)
	if err != nil {
		return nil, err
	}
	return &WyvernExchangeOrdersMatchedIterator{contract: _WyvernExchange.contract, event: "OrdersMatched", logs: logs, sub: sub}, nil
}

// WatchOrdersMatched is a free log subscription operation binding the contract event 0xc4109843e0b7d514e4c093114b863f8e7d8d9a458c372cd51bfe526b588006c9.
//
// Solidity: event OrdersMatched(bytes32 buyHash, bytes32 sellHash, address indexed maker, address indexed taker, uint256 price, bytes32 indexed metadata)
func (_WyvernExchange *WyvernExchangeFilterer) WatchOrdersMatched(opts *bind.WatchOpts, sink chan<- *WyvernExchangeOrdersMatched, maker []common.Address, taker []common.Address, metadata [][32]byte) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	var metadataRule []interface{}
	for _, metadataItem := range metadata {
		metadataRule = append(metadataRule, metadataItem)
	}

	logs, sub, err := _WyvernExchange.contract.WatchLogs(opts, "OrdersMatched", makerRule, takerRule, metadataRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WyvernExchangeOrdersMatched)
				if err := _WyvernExchange.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
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

// WyvernExchangeOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the WyvernExchange contract.
type WyvernExchangeOwnershipRenouncedIterator struct {
	Event *WyvernExchangeOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *WyvernExchangeOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WyvernExchangeOwnershipRenounced)
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
		it.Event = new(WyvernExchangeOwnershipRenounced)
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
func (it *WyvernExchangeOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WyvernExchangeOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WyvernExchangeOwnershipRenounced represents a OwnershipRenounced event raised by the WyvernExchange contract.
type WyvernExchangeOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_WyvernExchange *WyvernExchangeFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*WyvernExchangeOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _WyvernExchange.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WyvernExchangeOwnershipRenouncedIterator{contract: _WyvernExchange.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_WyvernExchange *WyvernExchangeFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *WyvernExchangeOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _WyvernExchange.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WyvernExchangeOwnershipRenounced)
				if err := _WyvernExchange.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// WyvernExchangeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WyvernExchange contract.
type WyvernExchangeOwnershipTransferredIterator struct {
	Event *WyvernExchangeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WyvernExchangeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WyvernExchangeOwnershipTransferred)
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
		it.Event = new(WyvernExchangeOwnershipTransferred)
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
func (it *WyvernExchangeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WyvernExchangeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WyvernExchangeOwnershipTransferred represents a OwnershipTransferred event raised by the WyvernExchange contract.
type WyvernExchangeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WyvernExchange *WyvernExchangeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WyvernExchangeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WyvernExchange.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WyvernExchangeOwnershipTransferredIterator{contract: _WyvernExchange.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WyvernExchange *WyvernExchangeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WyvernExchangeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WyvernExchange.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WyvernExchangeOwnershipTransferred)
				if err := _WyvernExchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

