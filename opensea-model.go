package opensea

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"
	"strconv"
	"strings"
	"time"
)

type Number string

func (n Number) Big() *big.Int {
	s := strings.Split(string(n), ".")
	r, _ := new(big.Int).SetString(s[0], 10)
	return r
}

type Address string

const NullAddress Address = "0x0000000000000000000000000000000000000000"

func IsHexAddress(s string) bool {
	if s == "0x0" {
		return true
	}
	if s[0:2] != "0x" {
		return false
	}
	addressLength := 2 + 40
	if len(s) != addressLength {
		return false
	}
	if !isHex(s[2:]) {
		return false
	}
	return true
}

func ParseAddress(address string) (Address, error) {
	if !IsHexAddress(address) {
		return "", errors.New("Invalid address: " + address)
	}
	return Address(strings.ToLower(address)), nil
}

func (a Address) String() string {
	return string(a)
}

func (a Address) IsNullAddress() bool {
	if a.String() == NullAddress.String() {
		return true
	}
	return false
}

func (a *Address) UnmarshalJSON(b []byte) error {
	var s string
	var err error
	if string(b) == "null" {
		s = NullAddress.String()
	} else {
		s, err = strconv.Unquote(string(b))
	}
	if err != nil {
		return err
	}
	*a, err = ParseAddress(s)
	return err
}

func (a Address) MarshalJSON() ([]byte, error) {
	s := strconv.Quote(a.String())
	return []byte(s), nil
}

type Bytes []byte

func (by Bytes) Bytes32() [32]byte {
	var ret [32]byte
	copy(ret[:], by[0:32])
	return ret
}

func (by *Bytes) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	if len(s) == 0 {
		*by = []byte{}
		return nil
	}
	s = s[2:]
	*by, err = hex.DecodeString(s)
	return err
}

func (by Bytes) MarshalJSON() ([]byte, error) {
	s := hex.EncodeToString([]byte(by))
	s = "0x" + s
	s = strconv.Quote(s)
	return []byte(s), nil
}

type TimeNano time.Time

func (t TimeNano) Time() time.Time {
	return time.Time(t)
}

func (t *TimeNano) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	tt := time.Time{}
	tt, err = time.Parse("2006-01-02T15:04:05.999999", s)
	// if strings.Contains(s, ".") {
	//      tt, err = time.Parse("2006-01-02T15:04:05.999999", s)
	// } else {
	//      tt, err = time.Parse("2006-01-02T15:04:05", s)
	// }
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

type Account struct {
	User          User    `json:"user" bson:"user"`
	ProfileImgURL string  `json:"profile_img_url" bson:"profile_img_url"`
	Address       Address `json:"address" bson:"address"`
	Config        string  `json:"config" bson:"config"`
	DiscordID     string  `json:"discord_id" bson:"discord_id"`
}

type User struct {
	Username string `json:"username" bson:"username"`
}

type Trait struct {
	TraitType   string          `json:"trait_type" bson:"trait_type"`
	Value       json.RawMessage `json:"value" bson:"value"`
	DisplayType interface{}     `json:"display_type" bson:"display_type"`
	MaxValue    interface{}     `json:"max_value" bson:"max_value"`
	TraitCount  int64           `json:"trait_count" bson:"trait_count"`
	Order       interface{}     `json:"order" bson:"order"`
}

type Value struct {
	Integer *int64
	String  *string
}

// isHexCharacter returns bool of c being a valid hexadecimal.
func isHexCharacter(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
}

// isHex validates whether each byte is valid hexadecimal string.
func isHex(str string) bool {
	if len(str)%2 != 0 {
		return false
	}
	for _, c := range []byte(str) {
		if !isHexCharacter(c) {
			return false
		}
	}
	return true
}

type AssetResponse struct {
	Assets []Asset `json:"assets" bson:"assets"`
}

type Asset struct {
	ID                   int64          `json:"id" bson:"id"`
	TokenID              string         `json:"token_id" bson:"token_id"`
	NumSales             int64          `json:"num_sales" bson:"num_sales"`
	BackgroundColor      string         `json:"background_color" bson:"background_color"`
	ImageURL             string         `json:"image_url" bson:"image_url"`
	ImagePreviewURL      string         `json:"image_preview_url" bson:"image_preview_url"`
	ImageThumbnailURL    string         `json:"image_thumbnail_url" bson:"image_thumbnail_url"`
	ImageOriginalURL     string         `json:"image_original_url" bson:"image_original_url"`
	AnimationURL         string         `json:"animation_url" bson:"animation_url"`
	AnimationOriginalURL string         `json:"animation_original_url" bson:"animation_original_url"`
	Name                 string         `json:"name" bson:"name"`
	Description          string         `json:"description" bson:"description"`
	ExternalLink         string         `json:"external_link" bson:"external_link"`
	AssetContract        *AssetContract `json:"asset_contract" bson:"asset_contract"`
	Owner                *Account       `json:"owner" bson:"owner"`
	Permalink            string         `json:"permalink" bson:"permalink"`
	Collection           *Collection    `json:"collection" bson:"collection"`
	Decimals             int64          `json:"decimals" bson:"decimals"`
	TokenMetadata        string         `json:"token_metadata" bson:"token_metadata"`
	Traits               interface{}    `json:"traits" bson:"traits"`
}

type AssetContract struct {
	Address                     Address     `json:"address" bson:"address"`
	AssetContractType           string      `json:"asset_contract_type" bson:"asset_contract_type"`
	CreatedDate                 string      `json:"created_date" bson:"created_date"`
	Name                        string      `json:"name" bson:"name"`
	NftVersion                  string      `json:"nft_version" bson:"nft_version"`
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

type Collection struct {
	BannerImageUrl              string      `json:"banner_image_url" bson:"banner_image_url"`
	ChatUrl                     string      `json:"chat_url" bson:"chat_url"`
	CreatedDate                 string      `json:"created_date" bson:"created_date"`
	DefaultToFiat               bool        `json:"default_to_fiat" bson:"default_to_fiat"`
	Description                 string      `json:"description" bson:"description"`
	DevBuyerFeeBasisPoints      string      `json:"dev_buyer_fee_basis_points" bson:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     string      `json:"dev_seller_fee_basis_points" bson:"dev_seller_fee_basis_points"`
	DiscordUrl                  string      `json:"discord_url" bson:"discord_url"`
	DisplayData                 interface{} `json:"display_data" bson:"display_data"`
	ExternalUrl                 string      `json:"external_url" bson:"external_url"`
	Featured                    bool        `json:"featured" bson:"featured"`
	FeaturedImageUrl            string      `json:"featured_image_url" bson:"featured_image_url"`
	Hidden                      bool        `json:"hidden" bson:"hidden"`
	SafelistRequestStatus       string      `json:"safelist_request_status" bson:"safelist_request_status"`
	ImageUrl                    string      `json:"image_url" bson:"image_url"`
	IsSubjectToWhitelist        bool        `json:"is_subject_to_whitelist" bson:"is_subject_to_whitelist"`
	LargeImageUrl               string      `json:"large_image_url" bson:"large_image_url"`
	MediumUsername              string      `json:"medium_username" bson:"medium_username"`
	Name                        string      `json:"name" bson:"name"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers" bson:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  string      `json:"opensea_buyer_fee_basis_points" bson:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints string      `json:"opensea_seller_fee_basis_points" bson:"opensea_seller_fee_basis_points"`
	PayoutAddress               string      `json:"payout_address" bson:"payout_address"`
	RequireEmail                bool        `json:"require_email" bson:"require_email"`
	ShortDescription            string      `json:"short_description" bson:"short_description"`
	Slug                        string      `json:"slug" bson:"slug"`
	TelegramUrl                 string      `json:"telegram_url" bson:"telegram_url"`
	TwitterUsername             string      `json:"twitter_username" bson:"twitter_username"`
	InstagramUsername           string      `json:"instagram_username" bson:"instagram_username"`
	WikiUrl                     string      `json:"wiki_url" bson:"wiki_url"`
}

type StatResponse struct {
	Stats Stat `json:"stats" bson:"stats"`
}

type Stat struct {
	OneDayVolume          float64 `json:"one_day_volume" bson:"one_day_volume"`
	OneDayChange          float64 `json:"one_day_change" bson:"one_day_change"`
	OneDaySales           float64 `json:"one_day_sales" bson:"one_day_sales"`
	OneDayAveragePrice    float64 `json:"one_day_average_price" bson:"one_day_average_price"`
	SevenDayVolume        float64 `json:"seven_day_volume" bson:"seven_day_volume"`
	SevenDayChange        float64 `json:"seven_day_change" bson:"seven_day_change"`
	SevenDaySales         float64 `json:"seven_day_sales" bson:"seven_day_sales"`
	SevenDayAveragePrice  float64 `json:"seven_day_average_price" bson:"seven_day_average_price"`
	ThirtyDayVolume       float64 `json:"thirty_day_volume" bson:"thirty_day_volume"`
	ThirtyDayChange       float64 `json:"thirty_day_change" bson:"thirty_day_change"`
	ThirtyDaySales        float64 `json:"thirty_day_sales" bson:"thirty_day_sales"`
	ThirtyDayAveragePrice float64 `json:"thirty_day_average_price" bson:"thirty_day_average_price"`
	TotalVolume           float64 `json:"total_volume" bson:"total_volume"`
	TotalSales            float64 `json:"total_sales" bson:"total_sales"`
	TotalSupply           float64 `json:"total_supply" bson:"total_supply"`
	Count                 float64 `json:"count" bson:"count"`
	NumOwners             float64 `json:"num_owners" bson:"num_owners"`
	AveragePrice          float64 `json:"average_price" bson:"average_price"`
	NumReports            float64 `json:"num_reports" bson:"num_reports"`
	MarketCap             float64 `json:"market_cap" bson:"market_cap"`
	FloorPrice            float64 `json:"floor_price" bson:"floor_price"`
}

type CollectionSingleResponse struct {
	Collection CollectionSingle `json:"collection" bson:"collection"`
}

type CollectionSingle struct {
	Editors               []Address      `json:"editors" bson:"editors"`
	PaymentTokens         []PaymentToken `json:"payment_tokens" bson:"payment_tokens"`
	PrimaryAssetContracts []Contract     `json:"primary_asset_contracts" bson:"primary_asset_contracts"`
	Traits                interface{}    `json:"traits" bson:"traits"`
	Collection
}
