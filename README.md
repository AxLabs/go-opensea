# go-opensea

Golang library for OpenSea APIs (https://docs.opensea.io/reference).

## Get Started

Get it:
```
go get github.com/gsmachado/go-opensea
```

Use it:
```go
TBD
```

## API Support

This SDK supports the following:

- ✅ [https://api.opensea.io/api/v1/assets](https://docs.opensea.io/reference/getting-assets)
- 🛠 [https://api.opensea.io/api/v1/events](https://docs.opensea.io/reference/retrieving-asset-events)
- 🛠 [https://api.opensea.io/api/v1/collections](https://docs.opensea.io/reference/retrieving-collections)
- 🛠 [https://api.opensea.io/api/v1/bundles](https://docs.opensea.io/reference/retrieving-bundles)
- 🛠 [https://api.opensea.io/api/v1/asset/{asset_contract_address}/{token_id}](https://docs.opensea.io/reference/retrieving-a-single-asset)
- 🛠 [https://api.opensea.io/api/v1/asset_contract/{asset_contract_address}](https://docs.opensea.io/reference/retrieving-a-single-contract)
- 🛠 [https://api.opensea.io/api/v1/collection/{collection_slug}](https://docs.opensea.io/reference/retrieving-a-single-collection)
- 🛠 [https://api.opensea.io/api/v1/collection/{collection_slug}/stats](https://docs.opensea.io/reference/retrieving-collection-stats)

## Development

TBD.

## TODOS

- [ ] reasonable integration tests (running nightly)
- [ ] expose the APIs as methods in the `Opensea` type

## Acknowledgement

This library was forked and reworked from https://github.com/rmanzoku/go-opensea, at 
[this commit](https://github.com/rmanzoku/go-opensea/tree/e0722c7d22bbe26cbf222b9503552d05b44af289). Thanks for it! :tada:
