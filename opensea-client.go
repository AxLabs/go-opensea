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
	mainnetAPI = "https://api.opensea.io"
	rinkebyAPI = "https://rinkeby-api.opensea.io"
)

type Opensea struct {
	API        string
	APIKey     string
	httpClient *http.Client
}

type errorResponse struct {
	Success bool `json:"success" bson:"success"`
}

func (e errorResponse) Error() string {
	return "Not success"
}

func NewOpensea(apiKey string) (*Opensea, error) {
	o := &Opensea{
		API:        mainnetAPI,
		APIKey:     apiKey,
		httpClient: defaultHttpClient(),
	}
	return o, nil
}

func NewOpenseaRinkeby(apiKey string) (*Opensea, error) {
	o := &Opensea{
		API:        rinkebyAPI,
		APIKey:     apiKey,
		httpClient: defaultHttpClient(),
	}
	return o, nil
}

// TODO
//func (o Opensea) GetAssets(params GetAssetsParams) (*AssetResponse, error) {
//	ctx := context.TODO()
//	return o.GetAssetsWithContext(ctx, params)
//}

// TODO
//func (o Opensea) GetAssetsWithContext(ctx context.Context, params GetAssetsParams) (*AssetResponse, error) {
//	path := fmt.Sprintf("/api/v1/assets")
//	b, err := o.GetPath(ctx, path)
//	if err != nil {
//		return nil, err
//	}
//	ret := new(AssetResponse)
//	return ret, json.Unmarshal(b, ret)
//}

func (o Opensea) GetSingleAsset(assetContractAddress string, tokenID *big.Int) (*Asset, error) {
	ctx := context.TODO()
	return o.GetSingleAssetWithContext(ctx, assetContractAddress, tokenID)
}

func (o Opensea) GetSingleAssetWithContext(ctx context.Context, assetContractAddress string, tokenID *big.Int) (*Asset, error) {
	path := fmt.Sprintf("/api/v1/asset/%s/%s", assetContractAddress, tokenID.String())
	b, err := o.GetPath(ctx, path)
	if err != nil {
		return nil, err
	}
	ret := new(Asset)
	return ret, json.Unmarshal(b, ret)
}

func (o Opensea) GetPath(ctx context.Context, path string) ([]byte, error) {
	return o.getURL(ctx, o.API+path)
}

func (o Opensea) getURL(ctx context.Context, url string) ([]byte, error) {
	client := o.httpClient
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	req.Header.Add("X-API-KEY", o.APIKey)
	req.Header.Add("Accept", "application/json")
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
		e := new(errorResponse)
		err = json.Unmarshal(body, e)
		if err != nil {
			return nil, err
		}
		if !e.Success {
			return nil, e
		}

		return nil, fmt.Errorf("Backend returns status %d msg: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

func (o Opensea) SetHttpClient(httpClient *http.Client) {
	o.httpClient = httpClient
}

func defaultHttpClient() *http.Client {
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
