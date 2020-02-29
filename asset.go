package opensea

type Asset struct {
	TokenID              string         `json:"token_id"`
	NumSales             int64          `json:"num_sales"`
	BackgroundColor      string         `json:"background_color"`
	ImageURL             string         `json:"image_url"`
	ImagePreviewURL      string         `json:"image_preview_url"`
	ImageThumbnailURL    string         `json:"image_thumbnail_url"`
	ImageOriginalURL     string         `json:"image_original_url"`
	AnimationURL         string         `json:"animation_url"`
	AnimationOriginalURL string         `json:"animation_original_url"`
	Name                 string         `json:"name"`
	Description          string         `json:"description"`
	ExternalLink         string         `json:"external_link"`
	AssetContract        *AssetContract `json:"asset_contract"`
	Owner                *Account       `json:"owner"`
	Permalink            string         `json:"permalink"`
	// Collection           *Collection    `json:"collection"`
	Decimals int64 `json:"decimals"`
}

type AssetContract struct {
	Address                     Address     `json:"address"`
	AssetContractType           string      `json:"asset_contract_type"`
	CreatedDate                 string      `json:"created_date"`
	Name                        string      `json:"name"`
	NftVersion                  string      `json:"nft_version"`
	OpenseaVersion              interface{} `json:"opensea_version"`
	Owner                       *Account    `json:"owner"`
	SchemaName                  string      `json:"schema_name"`
	Symbol                      string      `json:"symbol"`
	TotalSupply                 interface{} `json:"total_supply"`
	Description                 string      `json:"description"`
	ExternalLink                string      `json:"external_link"`
	ImageURL                    string      `json:"image_url"`
	DefaultToFiat               bool        `json:"default_to_fiat"`
	DevBuyerFeeBasisPoints      int64       `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int64       `json:"dev_seller_fee_basis_points"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  int64       `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int64       `json:"opensea_seller_fee_basis_points"`
	BuyerFeeBasisPoints         int64       `json:"buyer_fee_basis_points"`
	SellerFeeBasisPoints        int64       `json:"seller_fee_basis_points"`
	PayoutAddress               Address     `json:"payout_address"`
}
