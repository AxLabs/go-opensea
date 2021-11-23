package opensea

import (
	"context"
	"encoding/json"
)

type Contract struct {
	Collection                  Collection  `json:"collection" bson:"collection"`
	Address                     Address     `json:"address" bson:"address"`
	AssetContractType           string      `json:"asset_contract_type" bson:"asset_contract_type"`
	CreatedDate                 string      `json:"created_date" bson:"created_date"`
	Name                        string      `json:"name" bson:"name"`
	NFTVersion                  string      `json:"nft_version" bson:"nft_version"`
	OpenseaVersion              interface{} `json:"opensea_version" bson:"opensea_version"`
	Owner                       int64       `json:"owner" bson:"owner"`
	SchemaName                  string      `json:"schema_name" bson:"schema_name"`
	Symbol                      string      `json:"symbol" bson:"symbol"`
	TotalSupply                 interface{} `json:"total_supply" bson:"total_supply"`
	Description                 string      `json:"description" bson:"description"`
	ExternalLink                string      `json:"external_link" bson:"external_link"`
	ImageURL                    string      `json:"image_url" bson:"image_url"`
	DefaultToFiat               bool        `json:"default_to_fiat" bson:"default_to_fiat"`
	DevBuyerFeeBasisPoints      int64       `json:"dev_buyer_fee_basis_points" bson:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int64       `json:"dev_seller_fee_basis_points" bson:"dev_seller_fee_basis_points"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers" bson:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  int64       `json:"opensea_buyer_fee_basis_points" bson:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int64       `json:"opensea_seller_fee_basis_points" bson:"opensea_seller_fee_basis_points"`
	BuyerFeeBasisPoints         int64       `json:"buyer_fee_basis_points" bson:"buyer_fee_basis_points"`
	SellerFeeBasisPoints        int64       `json:"seller_fee_basis_points" bson:"seller_fee_basis_points"`
	PayoutAddress               Address     `json:"payout_address" bson:"payout_address"`
}

func (o Opensea) GetSingleContract(assetContractAddress string) (*Contract, error) {
	ctx := context.TODO()
	return o.GetSingleContractWithContext(ctx, assetContractAddress)
}

func (o Opensea) GetSingleContractWithContext(ctx context.Context, assetContractAddress string) (contract *Contract, err error) {
	path := "/api/v1/asset_contract/" + assetContractAddress
	b, err := o.GetPath(ctx, path)
	if err != nil {
		return nil, err
	}

	contract = &Contract{}
	err = json.Unmarshal(b, contract)
	return
}
