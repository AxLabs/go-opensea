package opensea

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net"
	"net/http"
	"time"
)

var (
	mainnetAPI = "https://api.opensea.io/api/v1/"
	rinkebyAPI = "https://rinkeby-api.opensea.io/api/v1/"
)

type Opensea struct {
	API    string
	APIKey string
}

func NewOpensea(apiKey string) (*Opensea, error) {
	o := &Opensea{
		API:    mainnetAPI,
		APIKey: apiKey,
	}
	return o, nil
}

func NewOpenseaRinkeby(apiKey string) (*Opensea, error) {
	o := &Opensea{
		API:    rinkebyAPI,
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
	client := httpClient()
	// fmt.Println(url)
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
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Backend returns status %d msg: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

func httpClient() *http.Client {
	client := new(http.Client)
	var transport http.RoundTripper = &http.Transport{
		Proxy:              http.ProxyFromEnvironment,
		DisableKeepAlives:  false,
		DisableCompression: false,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 300 * time.Second,
			DualStack: true,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	client.Transport = transport
	return client
}
