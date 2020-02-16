package exchange

import (
	"math/big"
	"time"

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
	MakerRelayerFee    *big.Int
	TakerRelayerFee    *big.Int
	MakerProtocolFee   *big.Int
	TakerProtocolFee   *big.Int
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
	BasePrice          *big.Int
	Extra              *big.Int
	ListingTime        *big.Int
	ExpirationTime     *big.Int
	Salt               *big.Int
}

type Sig struct {
	V uint8
	R [32]byte
	S [32]byte
}

func toAddress(a opensea.Address) common.Address {
	return common.HexToAddress(string(a))
}

func Generate(opts *bind.TransactOpts, in *opensea.Order) (buy *Order, buySig *Sig, sell *Order, sellSig *Sig, metadata [32]byte, err error) {
	buy = &Order{
		toAddress(in.Exchange),
		opts.From,
		toAddress(in.Maker.Address),
		in.TakerRelayerFee.Big(),
		in.MakerRelayerFee.Big(),
		in.TakerProtocolFee.Big(),
		in.MakerProtocolFee.Big(),
		toAddress(in.FeeRecipient.Address),
		in.FeeMethod,
		opensea.Buy,
		in.SaleKind,
		toAddress(in.Target),
		in.HowToCall,
		[]byte{}, // todo
		[]byte{}, // todo
		toAddress(in.StaticTarget),
		[]byte(in.StaticExtradata),
		toAddress(in.PaymentToken),
		in.Extra.Big(),
		in.BasePrice.Big(),
		big.NewInt(time.Now().Unix()),
		big.NewInt(0),
		big.NewInt(0), // todo
	}
	sell = &Order{
		toAddress(in.Exchange),
		toAddress(in.Maker.Address),
		opts.From,
		in.MakerRelayerFee.Big(),
		in.TakerRelayerFee.Big(),
		in.MakerProtocolFee.Big(),
		in.TakerProtocolFee.Big(),
		toAddress(in.FeeRecipient.Address),
		in.FeeMethod,
		opensea.Sell,
		in.SaleKind,
		toAddress(in.Target),
		in.HowToCall,
		in.Calldata,
		in.ReplacementPattern,
		toAddress(in.StaticTarget),
		[]byte(in.StaticExtradata),
		toAddress(in.PaymentToken),
		in.Extra.Big(),
		in.BasePrice.Big(),
		big.NewInt(in.ListingTime),
		big.NewInt(in.ExpirationTime),
		in.Salt.Big(),
	}
	buySig = &Sig{
		in.V,
		in.R.Bytes32(),
		in.S.Bytes32(),
	}
	sellSig = &Sig{
		in.V,
		in.R.Bytes32(),
		in.S.Bytes32(),
	}
	metadata = [32]byte{}
	return
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
		buy.MakerRelayerFee,
		buy.TakerRelayerFee,
		buy.MakerProtocolFee,
		buy.TakerProtocolFee,
		buy.BasePrice,
		buy.Extra,
		buy.ListingTime,
		buy.ExpirationTime,
		buy.Salt,
		sell.MakerRelayerFee,
		sell.TakerRelayerFee,
		sell.MakerProtocolFee,
		sell.TakerProtocolFee,
		sell.BasePrice,
		sell.Extra,
		sell.ListingTime,
		sell.ExpirationTime,
		sell.Salt,
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
