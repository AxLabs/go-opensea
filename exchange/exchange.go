package exchange

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	opensea "github.com/rmanzoku/go-opensea"
	"github.com/rmanzoku/go-opensea/exchange/wyvernExchange"
)

type Exchange struct {
	Address common.Address
	cli     *ethclient.Client
	we      *wyvernExchange.WyvernExchange
}

type Order struct {
	Exchange           common.Address
	Maker              common.Address
	Taker              common.Address
	MakerRelayerFee    int64
	TakerRelayerFee    int64
	MakerProtocolFee   int64
	TakerProtocolFee   int64
	FeeRecipient       common.Address
	FeeMethod          opensea.FeeMethod
	Side               opensea.Side
	SaleKind           opensea.SaleKind
	Target             common.Address
	HowToCall          opensea.HowToCall
	Calldata           []byte
	ReplacementPattern []byte
	StaticTarget       common.Address
	StaticExtradata    []byte
	PaymentToken       common.Address
	BasePrice          int64
	Extra              int64
	ListingTime        int64
	ExpirationTime     int64
	Salt               int64
}

type Sig struct {
	V uint8
	R [32]byte
	S [32]byte
}

func NewExchange(rpc string, exchangeAddress string) (*Exchange, error) {
	var err error
	ret := new(Exchange)
	ret.cli, err = ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}
	if !common.IsHexAddress(exchangeAddress) {
		return nil, errors.New("Invalid address: " + exchangeAddress)
	}

	ret.Address = common.HexToAddress(exchangeAddress)
	ret.we, err = wyvernExchange.NewWyvernExchange(ret.Address, ret.cli)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (e Exchange) AtomicMatch(opts *bind.TransactOpts, buy *Order, buySig *Sig, sell *Order, sellSig *Sig, metadata [32]byte) (*types.Transaction, error) {
	addrs := [14]common.Address{
		buy.Exchange,
		buy.Maker,
		buy.Taker,
		buy.FeeRecipient,
		buy.Target,
		buy.StaticTarget,
		buy.PaymentToken,
		sell.Exchange,
		sell.Maker,
		sell.Taker,
		sell.FeeRecipient,
		sell.Target,
		sell.StaticTarget,
		sell.PaymentToken,
	}

	uints := [18]*big.Int{
		big.NewInt(buy.MakerRelayerFee),
		big.NewInt(buy.TakerRelayerFee),
		big.NewInt(buy.MakerProtocolFee),
		big.NewInt(buy.TakerProtocolFee),
		big.NewInt(buy.BasePrice),
		big.NewInt(buy.Extra),
		big.NewInt(buy.ListingTime),
		big.NewInt(buy.ExpirationTime),
		big.NewInt(buy.Salt),
		big.NewInt(sell.MakerRelayerFee),
		big.NewInt(sell.TakerRelayerFee),
		big.NewInt(sell.MakerProtocolFee),
		big.NewInt(sell.TakerProtocolFee),
		big.NewInt(sell.BasePrice),
		big.NewInt(sell.Extra),
		big.NewInt(sell.ListingTime),
		big.NewInt(sell.ExpirationTime),
		big.NewInt(sell.Salt),
	}

	feeMethodsSidesKindsHowToCalls := [8]uint8{
		uint8(buy.FeeMethod),
		uint8(buy.Side),
		uint8(buy.SaleKind),
		uint8(buy.HowToCall),
		uint8(sell.FeeMethod),
		uint8(sell.Side),
		uint8(sell.SaleKind),
		uint8(sell.HowToCall),
	}

	calldataBuy := buy.Calldata
	calldataSell := sell.Calldata
	replacementPatternBuy := buy.ReplacementPattern
	replacementPatternSell := sell.ReplacementPattern
	staticExtradataBuy := buy.StaticExtradata
	staticExtradataSell := sell.StaticExtradata
	vs := [2]uint8{
		buySig.V,
		sellSig.V,
	}

	rssMetadata := [5][32]byte{
		buySig.R,
		buySig.S,
		sellSig.R,
		sellSig.S,
		metadata,
	}

	return e.we.AtomicMatch(opts, addrs, uints, feeMethodsSidesKindsHowToCalls, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell, vs, rssMetadata)
}
