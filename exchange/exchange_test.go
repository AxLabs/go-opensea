package exchange

import (
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/cheekybits/is"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var (
	e                = &Exchange{}
	opts             = &bind.TransactOpts{}
	exchangeContract = "0x7be8076f4ea4a4ad08075c2508e481d6c946d12b"
	assetContract    = "0xc03844f07f86ad1d90a1c4a2a8204dcf00f3a991"
	tokenID          = big.NewInt(50010001)
	privateKeyHex    = os.Getenv("PRIVATE_KEY")
	rpc              = os.Getenv("RPC")
)

func TestBuy(t *testing.T) {
	is := initializeTest(t)
	tx, err := e.Buy(opts, assetContract, big.NewInt(20140144))
	is.Nil(err)
	fmt.Println(tx)
}

func initializeTest(t *testing.T) is.I {
	is := is.New(t)
	var err error

	e, err = NewExchange(rpc, exchangeContract)
	is.Nil(err)
	return is
}
