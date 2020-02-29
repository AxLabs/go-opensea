package opensea

import (
	"testing"
	"time"
)

func TestGetEvents(t *testing.T) {
	is := initializeTest(t)

	after := time.Now().Unix() - 86400
	before := time.Now().Unix()
	typ := EventTypeSuccessful
	ret, err := o.GetEvents(contract, after, before, &typ)
	is.Nil(err)
	print(len(ret))
	print(*ret[0])
}
