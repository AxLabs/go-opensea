package opensea

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type Order struct {
	ID    int64 `json:"id"`
	Asset Asset `json:"asset"`
	// AssetBundle          interface{}          `json:"asset_bundle"`
	CreatedDate *TimeNano `json:"created_date"`
	ClosingDate *TimeNano `json:"closing_date"`
	// ClosingExtendable    bool                 `json:"closing_extendable"`
	ExpirationTime int64 `json:"expiration_time"`
	ListingTime    int64 `json:"listing_time"`
	// OrderHash            string               `json:"order_hash"`
	// Metadata Metadata `json:"metadata"`
	Exchange     Address `json:"exchange"`
	Maker        Account `json:"maker"`
	Taker        Account `json:"taker"`
	CurrentPrice Number  `json:"current_price"`
	// CurrentBounty        string               `json:"current_bounty"`
	// BountyMultiple       string               `json:"bounty_multiple"`
	MakerRelayerFee    Number    `json:"maker_relayer_fee"`
	TakerRelayerFee    Number    `json:"taker_relayer_fee"`
	MakerProtocolFee   Number    `json:"maker_protocol_fee"`
	TakerProtocolFee   Number    `json:"taker_protocol_fee"`
	MakerReferrerFee   Number    `json:"maker_referrer_fee"`
	FeeRecipient       Account   `json:"fee_recipient"`
	FeeMethod          FeeMethod `json:"fee_method"`
	Side               Side      `json:"side"` // 0 for buy orders and 1 for sell orders.
	SaleKind           SaleKind  `json:"sale_kind"`
	Target             Address   `json:"target"`
	HowToCall          HowToCall `json:"how_to_call"`
	Calldata           Bytes     `json:"calldata"`
	ReplacementPattern Bytes     `json:"replacement_pattern"`
	StaticTarget       Address   `json:"static_target"`
	StaticExtradata    Bytes     `json:"static_extradata"`
	PaymentToken       Address   `json:"payment_token"`
	// PaymentTokenContract PaymentTokenContract `json:"payment_token_contract"`
	BasePrice       Number `json:"base_price"`
	Extra           Number `json:"extra"`
	Quantity        string `json:"quantity"`
	Salt            Number `json:"salt"`
	V               uint8  `json:"v"`
	R               Bytes  `json:"r"`
	S               Bytes  `json:"s"`
	ApprovedOnChain bool   `json:"approved_on_chain"`
	Cancelled       bool   `json:"cancelled"`
	Finalized       bool   `json:"finalized"`
	MarkedInvalid   bool   `json:"marked_invalid"`
	// PrefixedHash         string               `json:"prefixed_hash"`
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

type TimeNano time.Time

func (t TimeNano) Time() time.Time {
	return time.Time(t)
}

func (t *TimeNano) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	tt, err := time.Parse("2006-01-02T15:04:05.999999", s)
	if err != nil {
		return err
	}
	*t = TimeNano(tt)
	return nil
}

func (t TimeNano) MarshalJSON() ([]byte, error) {
	s := t.Time().Format("2006-01-02T15:04:05.999999")
	s = strconv.Quote(s)
	return []byte(s), nil
}

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
		b, err := o.getPath(ctx, path)
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
