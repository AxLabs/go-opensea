package opensea

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

type Order struct {
	ID    int64 `json:"id" bson:"id"`
	Asset Asset `json:"asset" bson:"asset"`
	// AssetBundle          interface{}          `json:"asset_bundle" bson:"asset_bundle"`
	CreatedDate *TimeNano `json:"created_date" bson:"created_date"`
	ClosingDate *TimeNano `json:"closing_date" bson:"closing_date"`
	// ClosingExtendable bool      `json:"closing_extendable" bson:"closing_extendable"`
	ExpirationTime int64 `json:"expiration_time" bson:"expiration_time"`
	ListingTime    int64 `json:"listing_time" bson:"listing_time"`
	// OrderHash            string               `json:"order_hash" bson:"order_hash"`
	// Metadata Metadata `json:"metadata" bson:"metadata"`
	Exchange     Address `json:"exchange" bson:"exchange"`
	Maker        Account `json:"maker" bson:"maker"`
	Taker        Account `json:"taker" bson:"taker"`
	CurrentPrice Number  `json:"current_price" bson:"current_price"`
	// CurrentBounty        string               `json:"current_bounty" bson:"current_bounty"`
	// BountyMultiple       string               `json:"bounty_multiple" bson:"bounty_multiple"`
	MakerRelayerFee    Number    `json:"maker_relayer_fee" bson:"maker_relayer_fee"`
	TakerRelayerFee    Number    `json:"taker_relayer_fee" bson:"taker_relayer_fee"`
	MakerProtocolFee   Number    `json:"maker_protocol_fee" bson:"maker_protocol_fee"`
	TakerProtocolFee   Number    `json:"taker_protocol_fee" bson:"taker_protocol_fee"`
	MakerReferrerFee   Number    `json:"maker_referrer_fee" bson:"maker_referrer_fee"`
	FeeRecipient       Account   `json:"fee_recipient" bson:"fee_recipient"`
	FeeMethod          FeeMethod `json:"fee_method" bson:"fee_method"`
	Side               Side      `json:"side" bson:"side"` // 0 for buy orders and 1 for sell orders.
	SaleKind           SaleKind  `json:"sale_kind" bson:"sale_kind"`
	Target             Address   `json:"target" bson:"target"`
	HowToCall          HowToCall `json:"how_to_call" bson:"how_to_call"`
	Calldata           Bytes     `json:"calldata" bson:"calldata"`
	ReplacementPattern Bytes     `json:"replacement_pattern" bson:"replacement_pattern"`
	StaticTarget       Address   `json:"static_target" bson:"static_target"`
	StaticExtradata    Bytes     `json:"static_extradata" bson:"static_extradata"`
	PaymentToken       Address   `json:"payment_token" bson:"payment_token"`
	// PaymentTokenContract PaymentTokenContract `json:"payment_token_contract" bson:"payment_token_contract"`
	BasePrice       Number `json:"base_price" bson:"base_price"`
	Extra           Number `json:"extra" bson:"extra"`
	Quantity        string `json:"quantity" bson:"quantity"`
	Salt            Number `json:"salt" bson:"salt"`
	V               *uint8 `json:"v" bson:"v"`
	R               *Bytes `json:"r" bson:"r"`
	S               *Bytes `json:"s" bson:"s"`
	ApprovedOnChain bool   `json:"approved_on_chain" bson:"approved_on_chain"`
	Cancelled       bool   `json:"cancelled" bson:"cancelled"`
	Finalized       bool   `json:"finalized" bson:"finalized"`
	MarkedInvalid   bool   `json:"marked_invalid" bson:"marked_invalid"`
	// PrefixedHash         string               `json:"prefixed_hash" bson:"prefixed_hash"`
}

func (o Order) IsPrivate() bool {
	if o.Taker.Address != NullAddress {
		return true
	}
	return false
}

type Side uint8

const (
	Buy Side = iota
	Sell
)

type SaleKind uint8

const (
	FixedOrMinBit SaleKind = iota // 0 for fixed-price sales or min-bid auctions
	DutchAuctions                 // 1 for declining-price Dutch Auctions
)

type HowToCall uint8

const (
	Call HowToCall = iota
	DelegateCall
)

type FeeMethod uint8

const (
	ProtocolFee FeeMethod = iota
	SplitFee
)

// type Metadata struct {
// 	Asset  MetadataAsset `json:"asset"`
// 	Schema string        `json:"schema"`
// }

// type MetadataAsset struct {
// 	ID      string `json:"id"`
// 	Address string `json:"address"`
// }

func (o Opensea) GetOrders(assetContractAddress string, listedAfter int64) ([]*Order, error) {
	ctx := context.TODO()
	return o.GetOrdersWithContext(ctx, assetContractAddress, listedAfter)
}

func (o Opensea) GetOrdersWithContext(ctx context.Context, assetContractAddress string, listedAfter int64) (orders []*Order, err error) {
	offset := 0
	limit := 100

	q := url.Values{}
	q.Set("asset_contract_address", assetContractAddress)
	q.Set("listed_after", fmt.Sprintf("%d", listedAfter))
	q.Set("limit", fmt.Sprintf("%d", limit))
	q.Set("order_by", "created_date")
	q.Set("order_direction", "asc")

	orders = []*Order{}

	for true {
		q.Set("offset", fmt.Sprintf("%d", offset))
		path := "/wyvern/v1/orders?" + q.Encode()
		b, err := o.GetPath(ctx, path)
		if err != nil {
			return nil, err
		}

		out := &struct {
			Count  int64    `json:"count"`
			Orders []*Order `json:"orders"`
		}{}

		err = json.Unmarshal(b, out)
		if err != nil {
			return nil, err
		}
		orders = append(orders, out.Orders...)

		if len(out.Orders) < limit {
			break
		}
		offset += limit
	}

	return
}
