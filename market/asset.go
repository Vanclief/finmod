package market

import "github.com/vanclief/ez"

// Asset - A resource with economic value
type Asset struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

// NewAsset creates a new Asset from a name and a symbol
func NewAsset(symbol, name string) (*Asset, error) {
	const op = "market.NewAsset"

	if symbol == "" {
		return nil, ez.New(op, ez.EINVALID, "Missing asset symbol", nil)
	} else if name == "" {
		return nil, ez.New(op, ez.EINVALID, "Missing asset name", nil)
	}

	return &Asset{Symbol: symbol, Name: name}, nil
}
