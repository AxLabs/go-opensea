package opensea

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestGetSingleContract(t *testing.T) {
	contractAddress := "0xdceaf1652a131f32a821468dc03a92df0edd86ea"
	assertIs := initializeTest(t)

	singleContract, err := o.GetSingleContract(contractAddress)

	assertIs.Nil(err)

	assertIs.Equal(singleContract.Address, contractAddress)
	assertIs.Equal(singleContract.Name, "MyCryptoHeroes:Extension")
	assertIs.Equal(singleContract.AssetContractType, "non-fungible")
	assertIs.Equal(singleContract.Collection.Slug, "mycryptoheroes")
}

func print(in interface{}) {
	if reflect.TypeOf(in).Kind() == reflect.Struct {
		in, _ = json.Marshal(in)
		in = string(in.([]byte))
	}
	fmt.Println(in)
}
