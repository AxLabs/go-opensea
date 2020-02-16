package opensea

import (
	"encoding/json"
)

type Address string

const NullAddress Address = "0x0000000000000000000000000000000000000000"

type Asset struct {
	TokenID          string  `json:"token_id"`
	ImageURL         string  `json:"image_url"`
	ImageOriginalURL string  `json:"image_original_url"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	ExternalLink     string  `json:"external_link"`
	Owner            Account `json:"owner"`
	Permalink        string  `json:"permalink"`
	Traits           []Trait `json:"traits"`
}

type Account struct {
	User          *User   `json:"user"`
	ProfileImgURL string  `json:"profile_img_url"`
	Address       Address `json:"address"`
	Config        string  `json:"config"`
	DiscordID     string  `json:"discord_id"`
}

type User struct {
	Username string `json:"username"`
}

type Trait struct {
	TraitType   string          `json:"trait_type"`
	Value       json.RawMessage `json:"value"`
	DisplayType interface{}     `json:"display_type"`
	MaxValue    interface{}     `json:"max_value"`
	TraitCount  int64           `json:"trait_count"`
	Order       interface{}     `json:"order"`
}

type Value struct {
	Integer *int64
	String  *string
}
