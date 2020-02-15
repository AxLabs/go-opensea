package opensea

import (
	"context"
	"encoding/json"
	"fmt"
)

type Order struct {
	ID    int64 `json:"id"`
	Asset Asset `json:"asset"`
	// AssetBundle          interface{}          `json:"asset_bundle"`
	// CreatedDate          string               `json:"created_date"`
	// ClosingDate          *string              `json:"closing_date"`
	// ClosingExtendable    bool                 `json:"closing_extendable"`
	// ExpirationTime       int64                `json:"expiration_time"`
	// ListingTime          int64                `json:"listing_time"`
	// OrderHash            string               `json:"order_hash"`
	// Metadata Metadata `json:"metadata"`
	// Exchange             Exchange             `json:"exchange"`
	// Maker                FeeRecipient         `json:"maker"`
	// Taker                FeeRecipient         `json:"taker"`
	CurrentPrice string `json:"current_price"`
	// CurrentBounty        string               `json:"current_bounty"`
	// BountyMultiple       string               `json:"bounty_multiple"`
	// MakerRelayerFee      string               `json:"maker_relayer_fee"`
	// TakerRelayerFee      string               `json:"taker_relayer_fee"`
	// MakerProtocolFee     string               `json:"maker_protocol_fee"`
	// TakerProtocolFee     string               `json:"taker_protocol_fee"`
	// MakerReferrerFee     string               `json:"maker_referrer_fee"`
	// FeeRecipient         FeeRecipient         `json:"fee_recipient"`
	// FeeMethod            int64                `json:"fee_method"`
	// Side                 int64                `json:"side"`
	// SaleKind             int64                `json:"sale_kind"`
	// Target               Target               `json:"target"`
	// HowToCall            int64                `json:"how_to_call"`
	// Calldata             string               `json:"calldata"`
	// ReplacementPattern   string               `json:"replacement_pattern"`
	// StaticTarget         PaymentToken         `json:"static_target"`
	// StaticExtradata      StaticExtradata      `json:"static_extradata"`
	PaymentToken string `json:"payment_token"`
	// PaymentTokenContract PaymentTokenContract `json:"payment_token_contract"`
	// BasePrice            string               `json:"base_price"`
	// Extra                string               `json:"extra"`
	// Quantity             string               `json:"quantity"`
	// Salt                 string               `json:"salt"`
	// V                    int64                `json:"v"`
	// R                    string               `json:"r"`
	// S                    string               `json:"s"`
	// ApprovedOnChain      bool                 `json:"approved_on_chain"`
	// Cancelled            bool                 `json:"cancelled"`
	// Finalized            bool                 `json:"finalized"`
	// MarkedInvalid        bool                 `json:"marked_invalid"`
	// PrefixedHash         string               `json:"prefixed_hash"`
}

// type Metadata struct {
// 	Asset  MetadataAsset `json:"asset"`
// 	Schema string        `json:"schema"`
// }

// type MetadataAsset struct {
// 	ID      string `json:"id"`
// 	Address string `json:"address"`
// }

func (o Opensea) GetOrders(assetContractAddress string) ([]*Order, error) {
	ctx := context.TODO()
	return o.GetOrdersWithContext(ctx, assetContractAddress)
}

func (o Opensea) GetOrdersWithContext(ctx context.Context, assetContractAddress string) (orders []*Order, err error) {
	path := fmt.Sprintf("/wyvern/v1/orders?asset_contract_address=%s", assetContractAddress)
	b, err := o.getPath(ctx, path)
	if err != nil {
		return
	}

	out := &struct {
		Count  int64    `json:"count"`
		Orders []*Order `json:"orders"`
	}{}

	err = json.Unmarshal(b, out)
	if err != nil {
		return
	}

	orders = make([]*Order, 0, len(out.Orders))
	for _, v := range out.Orders {
		orders = append(orders, v)
	}

	return
}
