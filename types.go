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
	s, err := strconv.Unquote(string(b))
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
