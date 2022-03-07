package market

import (
	"fmt"
)

// Pair - Quotation of two different assets or currencies, with the value of one being quoted against the other.
type Pair struct {
	Base  Asset `json:"base"`
	Quote Asset `json:"quote"`
}

// NewPair creates a new Pair from two assets
func NewPair(base, quote Asset) Pair {
	return Pair{Base: base, Quote: quote}
}

// String - Implements Stringer interface
func (p *Pair) String() string {
	return fmt.Sprintf("%s/%s\n", p.Base.Symbol, p.Quote.Symbol)
}

// Symbol - Gets the current symbol
func (p *Pair) Symbol(separator string) string {
	if p.Quote.Symbol == "" {
		return p.Base.Symbol
	} else {
		return fmt.Sprintf("%s%s%s", p.Base.Symbol, separator, p.Quote.Symbol)
	}
}
