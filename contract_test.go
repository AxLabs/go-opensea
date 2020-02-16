package opensea

import (
	"testing"
)

func TestGetSingleContract(t *testing.T) {
	is := initializeTest(t)

	ret, err := o.GetSingleContract(contract)
	is.Nil(err)

	print(*ret)
}
