package opensea

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"testing"
)

func TestAssetResponse(t *testing.T) {
	inputFile, err := ioutil.ReadFile("test-files/opensea-assets-collectibles.json")
	if err != nil {
		log.Fatalf("file not found...")
	}

	osAsset := AssetResponse{}

	err = json.Unmarshal(inputFile, &osAsset)
	if err != nil {
		log.Fatalf(err.Error())
	}

	assert.NotNil(t, osAsset)
}

func TestStatResponse(t *testing.T) {
	inputFile, err := ioutil.ReadFile("test-files/opensea-stats-doodles-official.json")
	if err != nil {
		log.Fatalf("file not found...")
	}

	osStat := StatResponse{}

	err = json.Unmarshal(inputFile, &osStat)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	assert.NotNil(t, osStat)
}
