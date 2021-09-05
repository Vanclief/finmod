package market

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTrade(t *testing.T) {

	baseAsset, _ := NewCryptoAsset("FLOW")
	quoteAsset, _ := NewForexAsset("USD")
	pair, _ := NewPair(baseAsset, quoteAsset)

	// Should be able to create a new trade
	trade := NewTrade("TO2WZX", time.Now(), BuyAction, MarketOrder, pair, 14.39, 10.4233, 150)
	assert.Equal(t, "TO2WZX", trade.ID)
	assert.NotNil(t, trade.ExecutionTime)
	assert.Equal(t, BuyAction, trade.Action)
	assert.Equal(t, MarketOrder, trade.OrderType)
	assert.Equal(t, pair, trade.Pair)
	assert.Equal(t, 14.39, trade.Price)
	assert.Equal(t, 10.4233, trade.Quantity)
	assert.Equal(t, 0.008713000000000193, trade.Fee)
	assert.Equal(t, float64(150), trade.Cost)
}
