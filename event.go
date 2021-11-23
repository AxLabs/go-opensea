package opensea

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type AssetEventsResponse struct {
	AssetEvents []Event `json:"asset_events" bson:"asset_events"`
}

type Event struct {
	ID                  uint64              `json:"id" bson:"id"`
	Transaction         *Transaction        `json:"transaction" bson:"transaction"`
	PaymentToken        *PaymentToken       `json:"payment_token" bson:"payment_token"`
	Asset               *Asset              `json:"asset" bson:"asset"`
	AssetBundle         *AssetBundle        `json:"asset_bundle" bson:"asset_bundle"`
	WinnerAccount       *Account            `json:"winner_account" bson:"winner_account"`
	FromAccount         *Account            `json:"from_account" bson:"from_account"`
	ToAccount           *Account            `json:"to_account" bson:"to_account"`
	OwnerAccount        *Account            `json:"owner_account" bson:"owner_account"`
	ApprovedAccount     *Account            `json:"approved_account" bson:"approved_account"`
	Seller              *Account            `json:"seller" bson:"seller"`
	DevFeePaymentEvent  *DevFeePaymentEvent `json:"dev_fee_payment_event" bson:"dev_fee_payment_event"`
	CollectionSlug      string              `json:"collection_slug" bson:"collection_slug"`
	CreatedDate         TimeNano            `json:"created_date" bson:"created_date"`
	ModifiedDate        TimeNano            `json:"modified_date" bson:"modified_date"`
	ContractAddress     Address             `json:"contract_address" bson:"contract_address"`
	LogIndex            interface{}         `json:"log_index" bson:"log_index"`
	EventType           EventType           `json:"event_type" bson:"event_type"`
	AuctionType         string              `json:"auction_type" bson:"auction_type"`
	StartingPrice       string              `json:"starting_price" bson:"starting_price"`
	EndingPrice         string              `json:"ending_price" bson:"ending_price"`
	Duration            interface{}         `json:"duration" bson:"duration"`
	MinPrice            Number              `json:"min_price" bson:"min_price"`
	OfferedTo           Number              `json:"offered_to" bson:"offered_to"`
	BidAmount           Number              `json:"bid_amount" bson:"bid_amount"`
	TotalPrice          Number              `json:"total_price" bson:"total_price"`
	CustomEventName     interface{}         `json:"custom_event_name" bson:"custom_event_name"`
	Quantity            string              `json:"quantity" bson:"quantity"`
	PayoutAmount        interface{}         `json:"payout_amount" bson:"payout_amount"`
	EventTimestamp      TimeNano            `json:"event_timestamp" bson:"event_timestamp"`
	Relayer             string              `json:"relayer" bson:"relayer"`
	Collection          uint64              `json:"collection" bson:"collection"`
	PayoutAccount       interface{}         `json:"payout_account" bson:"payout_account"`
	PayoutAssetContract interface{}         `json:"payout_asset_contract" bson:"payout_asset_contract"`
	PayoutCollection    interface{}         `json:"payout_collection" bson:"payout_collection"`
	BuyOrder            uint64              `json:"buy_order" bson:"buy_order"`
	SellOrder           uint64              `json:"sell_order" bson:"sell_order"`
}

func (e Event) IsBundle() bool {
	return e.AssetBundle != nil
}

type PaymentToken struct {
	Symbol   string      `json:"symbol" bson:"symbol"`
	Address  Address     `json:"address" bson:"address"`
	ImageURL string      `json:"image_url" bson:"image_url"`
	Name     string      `json:"name" bson:"name"`
	Decimals int64       `json:"decimals" bson:"decimals"`
	EthPrice interface{} `json:"eth_price" bson:"eth_price"`
	UsdPrice interface{} `json:"usd_price" bson:"usd_price"`
}

type Transaction struct {
	ID               int64    `json:"id" bson:"id"`
	FromAccount      Account  `json:"from_account" bson:"from_account"`
	ToAccount        Account  `json:"to_account" bson:"to_account"`
	CreatedDate      TimeNano `json:"created_date" bson:"created_date"`
	ModifiedDate     TimeNano `json:"modified_date" bson:"modified_date"`
	TransactionHash  string   `json:"transaction_hash" bson:"transaction_hash"`
	TransactionIndex string   `json:"transaction_index" bson:"transaction_index"`
	BlockNumber      string   `json:"block_number" bson:"block_number"`
	BlockHash        string   `json:"block_hash" bson:"block_hash"`
	Timestamp        string   `json:"timestamp" bson:"timestamp"`
}

type AssetBundle struct {
	Maker         *Account       `json:"maker" bson:"maker"`
	Slug          string         `json:"slug" bson:"slug"`
	Assets        []*Asset       `json:"assets" bson:"assets"`
	Name          string         `json:"name" bson:"name"`
	Description   string         `json:"description" bson:"description"`
	ExternalLink  string         `json:"external_link" bson:"external_link"`
	AssetContract *AssetContract `json:"asset_contract" bson:"asset_contract"`
	Permalink     string         `json:"permalink" bson:"permalink"`
	SellOrders    interface{}    `json:"sell_orders" bson:"sell_orders"`
}

// DevFeePaymentEvent is fee transfer event from OpenSea to Dev, It appears to be running in bulk on a regular basis.
type DevFeePaymentEvent struct {
	EventType      string       `json:"event_type" bson:"event_type"`
	EventTimestamp string       `json:"event_timestamp" bson:"event_timestamp"`
	AuctionType    interface{}  `json:"auction_type" bson:"auction_type"`
	TotalPrice     interface{}  `json:"total_price" bson:"total_price"`
	Transaction    Transaction  `json:"transaction" bson:"transaction"`
	PaymentToken   PaymentToken `json:"payment_token" bson:"payment_token"`
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
		OnlyOpensea:          true,
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
		b, err := o.GetPath(ctx, path)
		if err != nil {
			return nil, err
		}

		eventsResp := &AssetEventsResponse{
			AssetEvents: []Event{},
		}
		err = json.Unmarshal(b, eventsResp)
		if err != nil {
			return nil, err
		}

		// Filters
		tmp := make([]*Event, len(eventsResp.AssetEvents))
		cnt := 0
		for i, e := range eventsResp.AssetEvents {
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

			tmp[cnt] = &eventsResp.AssetEvents[i]
			cnt++
		}
		events = append(events, tmp[0:cnt]...)

		if len(eventsResp.AssetEvents) < params.Limit {
			break
		}
		params.Offset += params.Limit
	}

	return
}
