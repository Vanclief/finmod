package market

import (
	"fmt"

	"github.com/vanclief/ez"
)

// Asset - A resource with economic value
type Asset struct {
	Name      string `json:"name"`
	Symbol    string `json:"symbol"`
	AltSymbol string `json:"alt_symbol"`
}

func (a *Asset) String() string {
	return fmt.Sprintf("%s: %s\n", a.Name, a.Symbol)
}

// NewAsset creates a new Asset from a name and a symbol
func NewAsset(symbol, name string) (asset Asset, err error) {
	const op = "market.NewAsset"

	if symbol == "" {
		return asset, ez.New(op, ez.EINVALID, "Missing asset symbol", nil)
	} else if name == "" {
		return asset, ez.New(op, ez.EINVALID, "Missing asset name", nil)
	}

	return Asset{Symbol: symbol, Name: name}, nil
}
