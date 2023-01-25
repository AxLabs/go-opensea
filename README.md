# go-opensea

This is a Golang library for [OpenSea APIs](https://docs.opensea.io/reference).

## Get Started

Execute the command below to download the library.

```
go get github.com/gsmachado/go-opensea
```

Instantiate an Opensea instance and then call one of its provided methods to interact with the Opensea API.

```go
func getAssetContract() {
    o, err := NewOpensea(os.Getenv("API_KEY"))
    if err != nil {
         panic("could not instantiate Opensea instance")
    }
    singleContract, err := o.GetSingleContract("0xdceaf1652a131f32a821468dc03a92df0edd86ea")
    if err != nil {
        panic("could not get asset contract")
    }
    fmt.Println(singleContract)
}
```

## API Support

This SDK supports the following endpoints:

- 🛠 [https://api.opensea.io/api/v1/assets](https://docs.opensea.io/reference/getting-assets)
- 🛠 [https://api.opensea.io/api/v1/events](https://docs.opensea.io/reference/retrieving-asset-events)
- 🛠 [https://api.opensea.io/api/v1/collections](https://docs.opensea.io/reference/retrieving-collections)
- 🛠 [https://api.opensea.io/api/v1/bundles](https://docs.opensea.io/reference/retrieving-bundles)
- 🛠 [https://api.opensea.io/api/v1/asset/{asset_contract_address}/{token_id}](https://docs.opensea.io/reference/retrieving-a-single-asset)
- ✅ [https://api.opensea.io/api/v1/asset_contract/{asset_contract_address}](https://docs.opensea.io/reference/retrieving-a-single-contract)
- 🛠 [https://api.opensea.io/api/v1/collection/{collection_slug}](https://docs.opensea.io/reference/retrieving-a-single-collection)
- 🛠 [https://api.opensea.io/api/v1/collection/{collection_slug}/stats](https://docs.opensea.io/reference/retrieving-collection-stats)

## Development

Do not hesitate to fork and open Pull Requests that would resolve any of the opened issues, fix potential bugs or support additional features. Take a look at the existing issues in this repository that specify some of the desired features for this library.

## Roadmap (timeline tbd)

#### Support APIs

- [ ] Add support for `/assets`
- [ ] Add support for `/events`
- [ ] Add support for `/collections`
- [ ] Add support for `/bundles`
- [ ] Add support for `/asset/{asset_contract_address}/{token_id}`
- [ ] Add support for `/collection/{collection_slug}`
- [ ] Add support for `/collection/{collection_slug}/stats`

#### Testing

- [ ] Write UTs for all above endpoints
- [ ] Include integration tests for all endpoints

## Acknowledgement

This library was forked and reworked from https://github.com/rmanzoku/go-opensea at 
[this commit](https://github.com/rmanzoku/go-opensea/tree/e0722c7d22bbe26cbf222b9503552d05b44af289). Thanks for it! :tada:

## Sponsoring

Support is highly appreciated, if you like what we've built please consider [sponsoring](https://github.com/sponsors/AxLabs) us via GitHub. 

## Contact

Reach out to us in our [Discord server](https://discord.axlabs.com/) in the `#go-opensea` channel. If you have feature requests or have identified bugs, please don't hesitate to directly open an issue about it.
