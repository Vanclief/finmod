package market

import (
	"fmt"
)

// Pair - Quotation of two different assets or currencies, with the value of one being quoted against the other.
type Pair struct {
	Base      *Asset `json:"base"`
	Quote     *Asset `json:"quote"`
	AltSymbol string `json:"-"` // For internal usage
}

// NewPair creates a new Pair from two assets
func NewPair(base, quote *Asset) *Pair {
	return &Pair{Base: base, Quote: quote}
}

// String - Implementes Stringer interface
func (p *Pair) String() string {
	return fmt.Sprintf("%s/%s", p.Base.Symbol, p.Quote.Symbol)
}

// Symbol - Gets the current symbol
func (p *Pair) Symbol() string {
	if p.AltSymbol != "" {
		return p.AltSymbol
	} else {
		return p.String()
	}
}
