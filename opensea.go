package opensea

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
)

var (
	MainnetAPI = "https://api.opensea.io/api/v1/"
	RinkebyAPI = "https://rinkeby-api.opensea.io/api/v1/"
)

type Opensea struct {
	API    string
	APIKey string
}

func NewOpensea(apiKey string) (*Opensea, error) {
	o := &Opensea{
		API:    MainnetAPI,
		APIKey: apiKey,
	}
	return o, nil
}

func (o Opensea) GetSingleAsset(assetContractAddress string, tokenID *big.Int) (*Asset, error) {
	ctx := context.TODO()
	return o.GetSingleAssetWithContext(ctx, assetContractAddress, tokenID)
}

func (o Opensea) GetSingleAssetWithContext(ctx context.Context, assetContractAddress string, tokenID *big.Int) (*Asset, error) {
	path := fmt.Sprintf("asset/%s/%s", assetContractAddress, tokenID.String())
	b, err := o.getPath(ctx, path)
	if err != nil {
		return nil, err
	}
	ret := new(Asset)
	return ret, json.Unmarshal(b, ret)
}

func (o Opensea) getPath(ctx context.Context, path string) ([]byte, error) {
	return o.getURL(ctx, o.API+path)
}

func (o Opensea) getURL(ctx context.Context, url string) ([]byte, error) {
	client := new(http.Client)
	fmt.Println(url)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	req.Header.Add("X-API-KEY", o.APIKey)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// if resp.StatusCode != http.StatusOK {
	// 	return nil, fmt.Errorf("Backend returns status %d msg: %s", resp.StatusCode, string(body))
	// }

	// // If it returns array, not an error
	// if string(body[0]) != "[" {
	// 	e := new(errorResponse)
	// 	err = json.Unmarshal(body, e)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	if e.Body != nil {
	// 		return nil, e.Error()
	// 	}
	// }

	return body, nil
}
