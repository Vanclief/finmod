package market

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vanclief/ez"
)

func TestOrderBookGetDepth(t *testing.T) {
	sampleOrderBook := getTestOrderBook()

	index, err := sampleOrderBook.GetDepth(5)
	assert.Nil(t, err)
	assert.Equal(t, float64(5), index)
	index, err = sampleOrderBook.GetDepth(11)
	assert.Nil(t, err)
	assert.Equal(t, float64(11), index)
	index, err = sampleOrderBook.GetDepth(14.33)
	assert.Nil(t, err)
	assert.Equal(t, float64(15), index)
	index, err = sampleOrderBook.GetDepth(100)
	assert.NotNil(t, err)
	assert.Equal(t, ez.ENOTFOUND, ez.ErrorCode(err))
	assert.Equal(t, float64(0), index)
	index, err = sampleOrderBook.GetDepth(-2)
	assert.NotNil(t, err)
	assert.Equal(t, ez.EINVALID, ez.ErrorCode(err))
	assert.Equal(t, float64(0), index)
	index, err = sampleOrderBook.GetDepth(9)
	assert.Nil(t, err)
	assert.Equal(t, float64(9), index)
}

func getTestOrderBook() OrderBook {
	return OrderBook{
		Time: time.Now().Unix(),
		Asks: []OrderBookRow{
			{
				Price:       0,
				Volume:      1,
				AccumVolume: 1,
			},
			{
				Price:       1,
				Volume:      1,
				AccumVolume: 2,
			},
			{
				Price:       2,
				Volume:      1,
				AccumVolume: 3,
			},
			{
				Price:       3,
				Volume:      1,
				AccumVolume: 4,
			},
			{
				Price:       4,
				Volume:      1,
				AccumVolume: 5,
			},
			{
				Price:       5,
				Volume:      1,
				AccumVolume: 6,
			},
			{
				Price:       6,
				Volume:      1,
				AccumVolume: 7,
			},
			{
				Price:       7,
				Volume:      1,
				AccumVolume: 8,
			},
			{
				Price:       8,
				Volume:      1,
				AccumVolume: 9,
			},
			{
				Price:       9,
				Volume:      1,
				AccumVolume: 10,
			},
		},
		Bids: []OrderBookRow{
			{
				Price:       20,
				Volume:      1,
				AccumVolume: 1,
			},
			{
				Price:       19,
				Volume:      1,
				AccumVolume: 2,
			},
			{
				Price:       18,
				Volume:      1,
				AccumVolume: 3,
			},
			{
				Price:       17,
				Volume:      1,
				AccumVolume: 4,
			},
			{
				Price:       16,
				Volume:      1,
				AccumVolume: 5,
			},
			{
				Price:       15,
				Volume:      1,
				AccumVolume: 6,
			},
			{
				Price:       14,
				Volume:      1,
				AccumVolume: 7,
			},
			{
				Price:       13,
				Volume:      1,
				AccumVolume: 8,
			},
			{
				Price:       12,
				Volume:      1,
				AccumVolume: 9,
			},
			{
				Price:       11,
				Volume:      1,
				AccumVolume: 10,
			},
		},
	}
}
