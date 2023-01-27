package opensea

import (
	"github.com/cheekybits/is"
	"math/big"
	"os"
	"testing"
)

var (
	o = &Opensea{}
)

func TestGetSingleAsset(t *testing.T) {
	assertIs := initializeTest(t)
	contractAddress := "0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb"
	tokenID := big.NewInt(1)
	singleAsset, err := o.GetSingleAsset(contractAddress, tokenID)

	assertIs.Nil(err)

	assertIs.Equal(singleAsset.Name, "CryptoPunk #1")
	assertIs.Equal(singleAsset.ExternalLink, "https://www.larvalabs.com/cryptopunks/details/1")
	assertIs.Equal(singleAsset.AssetContract.AssetContractType, "non-fungible")
}

func TestGetSingleAssetContract(t *testing.T) {
	assertIs := initializeTest(t)
	contractAddress := "0x06012c8cf97bead5deae237070f9587f8e7a266d"

	cryptoKittiesContract, err := o.GetSingleContract(contractAddress)

	assertIs.Nil(err)

	assertIs.Equal(cryptoKittiesContract.Name, "CryptoKitties")
	assertIs.Equal(cryptoKittiesContract.SchemaName, "ERC721")
	assertIs.Equal(cryptoKittiesContract.Collection.ExternalUrl, "https://www.cryptokitties.co/")
	assertIs.Equal(cryptoKittiesContract.Collection.DiscordUrl, "https://discord.gg/cryptokitties")
}

func initializeTest(t *testing.T) is.I {
	assertIs := is.New(t)
	var err error

	if o != nil {
		o, err = NewOpensea(os.Getenv("API_KEY"))
	}
	assertIs.Nil(err)
	return assertIs
}
