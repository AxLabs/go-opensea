package opensea

import (
	"testing"
)

func TestGetOrders(t *testing.T) {
	is := initializeTest(t)

	ret, err := o.GetOrders(contract)
	is.Nil(err)

	print(*ret[0])
}
