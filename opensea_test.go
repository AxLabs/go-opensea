package opensea

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"reflect"
	"testing"

	"github.com/cheekybits/is"
)

var (
	o        = &Opensea{}
	owner    = "0xd868711BD9a2C6F1548F5f4737f71DA67d821090"
	contract = "0xdceaf1652a131f32a821468dc03a92df0edd86ea"
	tokenID  = big.NewInt(50010001)
)

func TestGetSingleAsset(t *testing.T) {
	is := initializeTest(t)

	ret, err := o.GetSingleAsset(contract, tokenID)
	is.Nil(err)

	print(*ret)
}

func initializeTest(t *testing.T) is.I {
	is := is.New(t)
	var err error

	o, err = NewOpensea(os.Getenv("API_KEY"))
	is.Nil(err)
	return is
}

func print(in interface{}) {
	if reflect.TypeOf(in).Kind() == reflect.Struct {
		in, _ = json.Marshal(in)
		in = string(in.([]byte))
	}
	fmt.Println(in)
}
