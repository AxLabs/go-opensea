package opensea

import (
	"testing"
	"time"
)

func TestGetOrders(t *testing.T) {
	is := initializeTest(t)

	since := time.Now().Unix() - 86400
	ret, err := o.GetOrders(contract, since)
	is.Nil(err)

	print(len(ret))

	print(*ret[0])
}
