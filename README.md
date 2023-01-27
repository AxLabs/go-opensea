# go-opensea

This is a Golang library for [OpenSea API v1](https://docs.opensea.io/reference).

## Get Started

Execute the command below to download the library.

```
go get github.com/AxLabs/go-opensea
```

Instantiate an Opensea instance and then call one of its provided methods to interact with the Opensea API.

```go
// Create the OpenSea Client
opensea, err := NewOpensea(os.Getenv("API_KEY"))
if err != nil {
	panic("could not instantiate Opensea instance")
}
// Get the NFT contract of, e.g., "CryptoKitties"
cryptoKittiesContract, err := opensea.GetSingleContract("0x06012c8cf97bead5deae237070f9587f8e7a266d")
if err != nil {
  panic("could not get contract")
}
fmt.Println("Contract name: " + cryptoKittiesContract.Name)
fmt.Println("Schema Name:   " + cryptoKittiesContract.SchemaName)
fmt.Println("External URL:  " + cryptoKittiesContract.Collection.ExternalUrl)
fmt.Println("Discord URL:   " + cryptoKittiesContract.Collection.DiscordUrl)
```

## API Support

This SDK supports the following endpoints:

- 🛠 Retrieve Assets ([/api/v1/assets](https://docs.opensea.io/reference/getting-assets))
- ✅ Retrieve an Asset ([/api/v1/asset_contract/{asset_contract_address}](https://docs.opensea.io/reference/retrieving-a-single-contract))
- 🛠 Validate an Asset ([/api/v1/asset/{asset_contract_address}/{token_id}/validate](https://docs.opensea.io/reference/validate-assets))
- 🛠 Retrieve Events ([/api/v1/events](https://docs.opensea.io/reference/retrieving-asset-events))
- 🛠 Retrieve Collections ([/api/v1/collections](https://docs.opensea.io/reference/retrieving-collections))
- 🛠 Retrieve Bundles ([/api/v1/bundles](https://docs.opensea.io/reference/retrieving-bundles))
- ✅ Retrieve a Contract ([/api/v1/asset/{asset_contract_address}/{token_id}](https://docs.opensea.io/reference/retrieving-a-single-asset))
- 🛠 Retrieve a Collection ([/api/v1/collection/{collection_slug}](https://docs.opensea.io/reference/retrieving-a-single-collection))
- 🛠 Retrieve Collection Stats ([/api/v1/collection/{collection_slug}/stats](https://docs.opensea.io/reference/retrieving-collection-stats))
- 🛠 Retrieve Owners ([/api/v1/asset/{asset_contract_address}/{token_id}/{owners}](https://docs.opensea.io/reference/retrieve-owners))

## Development

Do not hesitate to fork and open Pull Requests that would resolve any of the opened issues, fix potential bugs or support additional
features. Take a look at the existing issues in this repository that specify some of the desired features for this library.

## Roadmap (timeline tbd)

#### Support APIs including its Query options

- [ ] Add support for `/assets`
- [ ] Add support for `/asset_contract/{asset_contract_address}`
- [ ] Add support for `/asset/{asset_contract_address}/{token_id}/validate`
- [ ] Add support for `/events`
- [ ] Add support for `/collections`
- [ ] Add support for `/bundles`
- [x] Add support for `/asset/{asset_contract_address}/{token_id}`
- [ ] Add support for `/collection/{collection_slug}`
- [ ] Add support for `/collection/{collection_slug}/stats`
- [ ] Add support for `/asset/{asset_contract_address}/{token_id}/{owners}`

#### Testing

- [ ] Write UTs for all above endpoints
- [ ] Include integration tests for all endpoints

## Acknowledgement

This library was reworked from https://github.com/rmanzoku/go-opensea at
[this commit](https://github.com/rmanzoku/go-opensea/tree/e0722c7d22bbe26cbf222b9503552d05b44af289). Thanks for it! :tada:

## Sponsoring

Financial support is highly appreciated!

If you like what we've built please consider [sponsoring us](https://github.com/sponsors/AxLabs) via GitHub.

## Contact

Reach out to us in our [Discord server](https://discord.axlabs.com) in the `#go-opensea` channel. If you have feature requests or have
identified bugs, please don't hesitate to directly open an issue about it.
