package market

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairSymbol(t *testing.T) {

	base := &Asset{Symbol: "ETH"}
	quote := &Asset{Symbol: "USD"}

	pair := &Pair{Base: base, Quote: quote}
	pairAlt := &Pair{}
	pairAlt.AltSymbol = "#US30"

	assert.Equal(t, "ETHUSD", pair.Symbol(""))
	assert.Equal(t, "ETH/USD", pair.Symbol("/"))
	assert.Equal(t, "#US30", pairAlt.Symbol(""))
	assert.Equal(t, "#US30", pairAlt.Symbol("/"))
}
