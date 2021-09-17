package market

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairSymbol(t *testing.T) {

	base := &Asset{Symbol: "ETH"}
	quote := &Asset{Symbol: "USD"}

	pair := &Pair{Base: base, Quote: quote}

	baseAlt := &Asset{Symbol: "#US30"}
	pairAlt := &Pair{Base: baseAlt}

	assert.Equal(t, "ETHUSD", pair.Symbol(""))
	assert.Equal(t, "ETH/USD", pair.Symbol("/"))
	assert.Equal(t, "#US30", pairAlt.Symbol(""))
	assert.Equal(t, "#US30", pairAlt.Symbol("/"))
}
