package opensea

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

type Event struct {
	ID                  int64        `json:"id"`
	Transaction         *Transaction `json:"transaction"`
	PaymentToken        PaymentToken `json:"payment_token"`
	Asset               *Asset       `json:"asset"`
	AssetBundle         *AssetBundle `json:"asset_bundle"`
	WinnerAccount       *Account     `json:"winner_account"`
	FromAccount         *Account     `json:"from_account"`
	ToAccount           *Account     `json:"to_account"`
	OwnerAccount        *Account     `json:"owner_account"`
	ApprovedAccount     *Account     `json:"approved_account"`
	Seller              *Account     `json:"seller"`
	DevFeePaymentEvent  interface{}  `json:"dev_fee_payment_event"`
	CollectionSlug      string       `json:"collection_slug"`
	CreatedDate         TimeNano     `json:"created_date"`
	ModifiedDate        TimeNano     `json:"modified_date"`
	ContractAddress     Address      `json:"contract_address"`
	LogIndex            interface{}  `json:"log_index"`
	EventType           EventType    `json:"event_type"`
	AuctionType         string       `json:"auction_type"`
	StartingPrice       string       `json:"starting_price"`
	EndingPrice         string       `json:"ending_price"`
	Duration            interface{}  `json:"duration"`
	MinPrice            interface{}  `json:"min_price"`
	OfferedTo           interface{}  `json:"offered_to"`
	BidAmount           interface{}  `json:"bid_amount"`
	TotalPrice          interface{}  `json:"total_price"`
	CustomEventName     interface{}  `json:"custom_event_name"`
	Quantity            string       `json:"quantity"`
	PayoutAmount        interface{}  `json:"payout_amount"`
	EventTimestamp      string       `json:"event_timestamp"`
	Relayer             string       `json:"relayer"`
	Collection          int64        `json:"collection"`
	PayoutAccount       interface{}  `json:"payout_account"`
	PayoutAssetContract interface{}  `json:"payout_asset_contract"`
	PayoutCollection    interface{}  `json:"payout_collection"`
	BuyOrder            interface{}  `json:"buy_order"`
	SellOrder           interface{}  `json:"sell_order"`
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
	EventTypeCreated            EventType = "created"
	EventTypeSuccessful         EventType = "successful"
	EventTypeCancelled          EventType = "cancelled"
	EventTypeBidEntered         EventType = "bid_entered"
	EventTypeBidWithdrawn       EventType = "bid_withdrawn"
	EventTypeTransfer           EventType = "transfer"
	EventTypeApprove            EventType = "approve"
	EventTypeCompositionCreated EventType = "composition_created"
)

func (o Opensea) GetEvents(assetContractAddress string, occurredAfter int64, occurredBefore int64, eventType *EventType) ([]*Event, error) {
	ctx := context.TODO()
	return o.GetEventsWithContext(ctx, assetContractAddress, occurredAfter, occurredBefore, eventType)
}

func (o Opensea) GetEventsWithContext(ctx context.Context, assetContractAddress string, occurredAfter int64, occurredBefore int64, eventType *EventType) (events []*Event, err error) {

	addr, err := ParseAddress(assetContractAddress)
	if err != nil {
		return
	}

	offset := 0
	limit := 100

	q := url.Values{}
	q.Set("asset_contract_address", addr.String())
	q.Set("occurred_after", fmt.Sprintf("%d", occurredAfter))
	q.Set("occurred_before", fmt.Sprintf("%d", occurredBefore))
	q.Set("limit", fmt.Sprintf("%d", limit))
	q.Set("order_by", "created_date")
	q.Set("order_direction", "asc")

	if eventType != nil {
		q.Set("event_type", string(*eventType))
	}

	events = []*Event{}

	for true {
		q.Set("offset", fmt.Sprintf("%d", offset))
		path := "/api/v1/events/?" + q.Encode()
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

		tmp := make([]*Event, len(out.AssetEvents))
		cnt := 0
		for _, e := range out.AssetEvents {
			// remove incorrect asset, the events are bundled collection
			if e.Asset != nil && e.Asset.AssetContract.Address != addr {
				continue
			}

			if e.AssetBundle != nil {
				ok := false
				for _, a := range e.AssetBundle.Assets {
					if a.AssetContract.Address == addr {
						ok = true
						break
					}
				}
				if !ok {
					continue
				}
			}

			tmp[cnt] = e
			cnt++
		}
		events = append(events, tmp[0:cnt]...)

		if len(out.AssetEvents) < limit {
			break
		}
		offset += limit
	}

	return
}
