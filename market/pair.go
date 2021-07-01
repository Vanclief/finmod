package market

import (
	"fmt"

	"github.com/vanclief/ez"
)

// Pair - Quotation of two different assets or currencies, with the value of one being quoted against the other.
type Pair struct {
	Base  *Asset
	Quote *Asset
}

// NewPair creates a new Pair from two assets
func NewPair(base, quote *Asset) (*Pair, error) {
	const op = "market.NewPair"

	if base == nil {
		return nil, ez.New(op, ez.EINVALID, "Missing base asset", nil)
	} else if quote == nil {
		return nil, ez.New(op, ez.EINVALID, "Missing quote asset", nil)
	}

	return &Pair{Base: base, Quote: quote}, nil
}

// String - Implementes Stringer interface
func (p *Pair) String() string {
	return fmt.Sprintf("%s/%s", p.Base.Symbol, p.Quote.Symbol)
}
