package opensea

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type Event struct {
	ID                  int64         `json:"id"`
	Transaction         *Transaction  `json:"transaction"`
	PaymentToken        *PaymentToken `json:"payment_token"`
	Asset               *Asset        `json:"asset"`
	AssetBundle         *AssetBundle  `json:"asset_bundle"`
	WinnerAccount       *Account      `json:"winner_account"`
	FromAccount         *Account      `json:"from_account"`
	ToAccount           *Account      `json:"to_account"`
	OwnerAccount        *Account      `json:"owner_account"`
	ApprovedAccount     *Account      `json:"approved_account"`
	Seller              *Account      `json:"seller"`
	DevFeePaymentEvent  interface{}   `json:"dev_fee_payment_event"`
	CollectionSlug      string        `json:"collection_slug"`
	CreatedDate         TimeNano      `json:"created_date"`
	ModifiedDate        TimeNano      `json:"modified_date"`
	ContractAddress     Address       `json:"contract_address"`
	LogIndex            interface{}   `json:"log_index"`
	EventType           EventType     `json:"event_type"`
	AuctionType         string        `json:"auction_type"`
	StartingPrice       string        `json:"starting_price"`
	EndingPrice         string        `json:"ending_price"`
	Duration            interface{}   `json:"duration"`
	MinPrice            Number        `json:"min_price"`
	OfferedTo           Number        `json:"offered_to"`
	BidAmount           Number        `json:"bid_amount"`
	TotalPrice          Number        `json:"total_price"`
	CustomEventName     interface{}   `json:"custom_event_name"`
	Quantity            string        `json:"quantity"`
	PayoutAmount        interface{}   `json:"payout_amount"`
	EventTimestamp      string        `json:"event_timestamp"`
	Relayer             string        `json:"relayer"`
	Collection          int64         `json:"collection"`
	PayoutAccount       interface{}   `json:"payout_account"`
	PayoutAssetContract interface{}   `json:"payout_asset_contract"`
	PayoutCollection    interface{}   `json:"payout_collection"`
	BuyOrder            uint64        `json:"buy_order"`
	SellOrder           uint64        `json:"sell_order"`
}

func (e Event) IsBundle() bool {
	return e.AssetBundle != nil
}

type PaymentToken struct {
	Symbol   string  `json:"symbol"`
	Address  Address `json:"address"`
	ImageURL string  `json:"image_url"`
	Name     string  `json:"name"`
	Decimals int64   `json:"decimals"`
	EthPrice string  `json:"eth_price"`
	UsdPrice string  `json:"usd_price"`
}

type Transaction struct {
	ID               int64    `json:"id"`
	FromAccount      Account  `json:"from_account"`
	ToAccount        Account  `json:"to_account"`
	CreatedDate      TimeNano `json:"created_date"`
	ModifiedDate     TimeNano `json:"modified_date"`
	TransactionHash  string   `json:"transaction_hash"`
	TransactionIndex string   `json:"transaction_index"`
	BlockNumber      string   `json:"block_number"`
	BlockHash        string   `json:"block_hash"`
	Timestamp        string   `json:"timestamp"`
}

type AssetBundle struct {
	Maker         *Account       `json:"maker"`
	Slug          string         `json:"slug"`
	Assets        []*Asset       `json:"assets"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	ExternalLink  string         `json:"external_link"`
	AssetContract *AssetContract `json:"asset_contract"`
	Permalink     string         `json:"permalink"`
	SellOrders    interface{}    `json:"sell_orders"`
}

type EventType string

const (
	EventTypeNone               EventType = ""
	EventTypeCreated            EventType = "created"
	EventTypeSuccessful         EventType = "successful"
	EventTypeCancelled          EventType = "cancelled"
	EventTypeBidEntered         EventType = "bid_entered"
	EventTypeBidWithdrawn       EventType = "bid_withdrawn"
	EventTypeTransfer           EventType = "transfer"
	EventTypeApprove            EventType = "approve"
	EventTypeCompositionCreated EventType = "composition_created"
)

type AuctionType string

const (
	AuctionTypeNone     AuctionType = ""
	AuctionTypeEnglish  AuctionType = "english"
	AuctionTypeDutch    AuctionType = "dutch"
	AuctionTypeMinPrice AuctionType = "min-price"
)

type RetrievingEventsParams struct {
	AssetContractAddress Address
	TokenID              int32
	AccountAddress       Address
	EventType            EventType
	OnlyOpensea          bool
	AuctionType          AuctionType
	Offset               int
	Limit                int
	OccurredBefore       int64
	OccurredAfter        int64
}

func NewRetrievingEventsParams() *RetrievingEventsParams {
	return &RetrievingEventsParams{
		AssetContractAddress: NullAddress,
		TokenID:              0,
		AccountAddress:       NullAddress,
		EventType:            EventTypeNone,
		OnlyOpensea:          false,
		AuctionType:          AuctionTypeNone,
		Offset:               0,
		Limit:                100,
		OccurredBefore:       time.Now().Unix(),
		OccurredAfter:        time.Now().Unix() - 3600,
	}
}

func (p *RetrievingEventsParams) SetAssetContractAddress(addr string) (err error) {
	p.AssetContractAddress, err = ParseAddress(addr)
	return
}

func (p *RetrievingEventsParams) SetAccountAddress(addr string) (err error) {
	p.AccountAddress, err = ParseAddress(addr)
	return
}

func (p RetrievingEventsParams) Encode() string {
	q := url.Values{}

	if p.AssetContractAddress != NullAddress {
		q.Set("asset_contract_address", p.AssetContractAddress.String())
	}
	if p.TokenID != 0 {
		q.Set("token_id", fmt.Sprintf("%d", p.TokenID))
	}
	if p.AccountAddress != NullAddress {
		q.Set("account_address", p.AccountAddress.String())
	}
	if p.EventType != EventTypeNone {
		q.Set("event_type", string(p.EventType))
	}
	if p.OnlyOpensea {
		q.Set("only_opensea", "true")
	} else {
		q.Set("only_opensea", "false")
	}
	if p.AuctionType != AuctionTypeNone {
		q.Set("auction_type", string(p.AuctionType))
	}
	q.Set("limit", fmt.Sprintf("%d", p.Limit))
	q.Set("offset", fmt.Sprintf("%d", p.Offset))
	q.Set("occurred_after", fmt.Sprintf("%d", p.OccurredAfter))
	q.Set("occurred_before", fmt.Sprintf("%d", p.OccurredBefore))

	return q.Encode()
}

func (o Opensea) RetrievingEvents(params *RetrievingEventsParams) ([]*Event, error) {
	ctx := context.TODO()
	return o.RetrievingEventsWithContext(ctx, params)
}

func (o Opensea) RetrievingEventsWithContext(ctx context.Context, params *RetrievingEventsParams) (events []*Event, err error) {
	if params == nil {
		params = NewRetrievingEventsParams()
	}

	events = []*Event{}
	for true {
		path := "/api/v1/events/?" + params.Encode()
		b, err := o.getPath(ctx, path)
		if err != nil {
			return nil, err
		}

		out := &struct {
			AssetEvents []*Event `json:"asset_events"`
		}{}

		err = json.Unmarshal(b, out)
		if err != nil {
			return nil, err
		}

		// Filters
		tmp := make([]*Event, len(out.AssetEvents))
		cnt := 0
		for _, e := range out.AssetEvents {
			// remove incorrect asset, the events are bundled collection
			if params.AssetContractAddress != NullAddress {
				if e.Asset != nil && e.Asset.AssetContract.Address != params.AssetContractAddress {
					continue
				}

				if e.AssetBundle != nil {
					ok := false
					for _, a := range e.AssetBundle.Assets {
						if a.AssetContract.Address == params.AssetContractAddress {
							ok = true
							break
						}
					}
					if !ok {
						continue
					}
				}
			}

			tmp[cnt] = e
			cnt++
		}
		events = append(events, tmp[0:cnt]...)

		if len(out.AssetEvents) < params.Limit {
			break
		}
		params.Offset += params.Limit
	}

	return
}
