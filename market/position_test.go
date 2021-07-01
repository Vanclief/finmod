package market

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewPosition(t *testing.T) {

	baseAsset, _ := NewCryptoAsset("ETH")
	quoteAsset, _ := NewForexAsset("USD")
	pair, _ := NewPair(baseAsset, quoteAsset)

	// Should be able to create a new long position from a trade
	trade := NewTrade("TO2WZX", time.Now(), BuyAction, MarketOrder, pair, 14.39, 10.4233, 150)

	position := NewPosition(trade)
	assert.NotNil(t, position.ID)
	assert.Equal(t, LongPosition, position.Type)
	assert.Equal(t, pair, position.Pair)
	assert.Equal(t, 14.39, position.OpenPrice)
	assert.Equal(t, 10.4233, position.Quantity)
	assert.Equal(t, true, position.Open)
	assert.Len(t, position.Trades, 1)
	assert.Equal(t, *trade, position.Trades[0])

	// Should be able to create a new short position from a trade
	trade = NewTrade("TO2WZX", time.Now(), SellAction, MarketOrder, pair, 14.39, 10.4233, 150)

	position = NewPosition(trade)
	assert.NotNil(t, position.ID)
	assert.Equal(t, ShortPosition, position.Type)
	assert.Equal(t, pair, position.Pair)
	assert.Equal(t, 14.39, position.OpenPrice)
	assert.Equal(t, 10.4233, position.Quantity)
	assert.Equal(t, true, position.Open)
	assert.Len(t, position.Trades, 1)
	assert.Equal(t, *trade, position.Trades[0])
}

func TestModifyPosition(t *testing.T) {

	baseAsset, _ := NewCryptoAsset("ETH")
	quoteAsset, _ := NewForexAsset("USD")
	pair, _ := NewPair(baseAsset, quoteAsset)

	// Case 1: Should be able to modify a long position with a new long trade that increments the position size
	trade1 := NewTrade("TO2WZX", time.Now(), BuyAction, MarketOrder, pair, 14.39, 10.4233, 150)
	position := NewPosition(trade1)

	trade2 := NewTrade("TI7MTX", time.Now(), BuyAction, MarketOrder, pair, 10.270, 48.68549, 500)
	position.Modify(trade2)
	assert.Equal(t, 10.99652470131769, position.OpenPrice)
	assert.Equal(t, 59.10879, position.Quantity)
	assert.Equal(t, true, position.Open)
	assert.Len(t, position.Trades, 2)
	assert.Equal(t, *trade2, position.Trades[1])

	// Case 2: Should be able to modify a long position with a new short trade that reduces the size
	trade3 := NewTrade("TCS6CS", time.Now(), SellAction, MarketOrder, pair, 10.192, 48.68549, 496.202)
	position.Modify(trade3)
	assert.Equal(t, 14.754324946993759, position.OpenPrice)
	assert.Equal(t, 10.423299999999998, position.Quantity)
	assert.Equal(t, true, position.Open)
	assert.Len(t, position.Trades, 3)
	assert.Equal(t, *trade3, position.Trades[2])

	// Case 3: Should be able to modify a long position with a new short trade that closes the position
	trade4 := NewTrade("TCS6CS", time.Now(), SellAction, MarketOrder, pair, 15.39, 10.4233, 160.4233)
	position.Modify(trade4)
	assert.Equal(t, false, position.Open)
	assert.Equal(t, 14.754324946993759, position.OpenPrice)
	assert.Equal(t, 15.39, position.ClosePrice)
	assert.Equal(t, 10.423299999999998, position.Quantity)
	assert.Equal(t, 6.625831779999961, position.Profit)
	assert.Len(t, position.Trades, 4)
	assert.Equal(t, *trade4, position.Trades[3])

	// Case 4: Should be able to modify a short position with a new short trade that increments the position size
	trade5 := NewTrade("TY6CIM", time.Now(), SellAction, MarketOrder, pair, 14.39, 10.4233, 150)
	position = NewPosition(trade5)

	trade6 := NewTrade("TX7ANX", time.Now(), SellAction, MarketOrder, pair, 10.192, 48.68549, 496.202)
	position.Modify(trade6)
	assert.Equal(t, 10.932279295177587, position.OpenPrice)
	assert.Equal(t, 59.10879, position.Quantity)
	assert.Equal(t, true, position.Open)
	assert.Len(t, position.Trades, 2)
	assert.Equal(t, *trade6, position.Trades[1])

	// Case 5: Should be able to modify a short position with a new long trade that reduces the size
	trade7 := NewTrade("TA78IA", time.Now(), BuyAction, MarketOrder, pair, 15.39, 10.4233, 160.4233)
	position.Modify(trade7)
	assert.Equal(t, true, position.Open)
	assert.Equal(t, 9.97790541042105, position.OpenPrice)
	assert.Equal(t, 48.68549, position.Quantity)
	assert.Len(t, position.Trades, 3)
	assert.Equal(t, *trade7, position.Trades[2])

	// Case 6: Should be able to modify a short position with a new long trade that closes the position
	trade8 := NewTrade("T20AGT", time.Now(), BuyAction, MarketOrder, pair, 9.39, 48.68549, 457.61)
	position.Modify(trade8)
	assert.Equal(t, false, position.Open)
	assert.Equal(t, 9.97790541042105, position.OpenPrice)
	assert.Equal(t, 9.39, position.ClosePrice)
	assert.Equal(t, 48.68549, position.Quantity)
	assert.Equal(t, 28.62246297999989, position.Profit)
	assert.Len(t, position.Trades, 4)
	assert.Equal(t, *trade8, position.Trades[3])
}
