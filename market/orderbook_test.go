package market

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewOrderBook(t *testing.T) {
	bids := []OrderBookRow{
		{
			Price:  5,
			Volume: 1.3,
		},
		{
			Price:  2,
			Volume: 1,
		},
		{
			Price:  3,
			Volume: 1,
		},
		{
			Price:  4,
			Volume: 1,
		},
		{
			Price:  7,
			Volume: 1,
		},
		{
			Price:  6,
			Volume: 1,
		},
		{
			Price:  1,
			Volume: 1,
		},
		{
			Price:  8,
			Volume: 1,
		},
		{
			Price:  9,
			Volume: 1,
		},
	}

	asks := []OrderBookRow{
		{
			Price:  20,
			Volume: 1,
		},
		{
			Price:  19,
			Volume: 1,
		},
		{
			Price:  18,
			Volume: 1,
		},
		{
			Price:  17,
			Volume: 1,
		},
		{
			Price:  16,
			Volume: 1,
		},
		{
			Price:  15,
			Volume: 1,
		},
		{
			Price:  14,
			Volume: 1,
		},
		{
			Price:  13,
			Volume: 1,
		},
		{
			Price:  12,
			Volume: 1,
		},
		{
			Price:  11,
			Volume: 1,
		},
	}

	ob := NewOrderBook(asks, bids, 5)

	ob.Print()

	ob = NewOrderBook([]OrderBookRow{}, []OrderBookRow{}, 5)
}

func TestApplyUpdate(t *testing.T) {
	bids := []OrderBookRow{
		{
			Price:  1,
			Volume: 1.3,
		},
		{
			Price:  2,
			Volume: 1,
		},
		{
			Price:  3,
			Volume: 1,
		},
		{
			Price:  4,
			Volume: 1,
		},
		{
			Price:  5,
			Volume: 1,
		},
	}

	asks := []OrderBookRow{
		{
			Price:  12,
			Volume: 1,
		},
		{
			Price:  11,
			Volume: 1,
		},
		{
			Price:  10,
			Volume: 1,
		},
		{
			Price:  9,
			Volume: 1,
		},
		{
			Price:  8,
			Volume: 1,
		},
		{
			Price:  7,
			Volume: 1,
		},
	}

	ob := NewOrderBook(asks, bids, 5)

	update := OrderBookUpdate{13, 1, "bid"}
	err := ob.ApplyUpdate(update)
	assert.Nil(t, err)

	//update = OrderBookUpdate{3.5, 1.4, "bid"}
	//err = ob.ApplyUpdate(update)
	//assert.Nil(t, err)
	//
	//update = OrderBookUpdate{7.5, 2, "ask"}
	//err = ob.ApplyUpdate(update)
	//assert.Nil(t, err)
	//
	//update = OrderBookUpdate{8, 0, "ask"}
	//err = ob.ApplyUpdate(update)
	//assert.Nil(t, err)
	//
	//update = OrderBookUpdate{7, 3, "ask"}
	//err = ob.ApplyUpdate(update)
	//assert.Nil(t, err)

	ob.Print()
}

func BenchmarkApplyUpdate(t *testing.B) {
	bids := []OrderBookRow{
		{
			Price:  1,
			Volume: 1.3,
		},
		{
			Price:  2,
			Volume: 1,
		},
		{
			Price:  3,
			Volume: 1,
		},
		{
			Price:  4,
			Volume: 1,
		},
		{
			Price:  5,
			Volume: 1,
		},
	}

	asks := []OrderBookRow{
		{
			Price:  12,
			Volume: 1,
		},
		{
			Price:  11,
			Volume: 1,
		},
		{
			Price:  10,
			Volume: 1,
		},
		{
			Price:  9,
			Volume: 1,
		},
		{
			Price:  8,
			Volume: 1,
		},
		{
			Price:  7,
			Volume: 1,
		},
	}

	ob := NewOrderBook(asks, bids, 5)

	update := OrderBookUpdate{1.5, 1, "bid"}
	err := ob.ApplyUpdate(update)
	assert.Nil(t, err)

	update = OrderBookUpdate{3.5, 1.4, "bid"}
	err = ob.ApplyUpdate(update)
	assert.Nil(t, err)

	update = OrderBookUpdate{7.5, 2, "ask"}
	err = ob.ApplyUpdate(update)
	assert.Nil(t, err)

	update = OrderBookUpdate{8, 0, "ask"}
	err = ob.ApplyUpdate(update)
	assert.Nil(t, err)

	update = OrderBookUpdate{7, 3, "ask"}
	err = ob.ApplyUpdate(update)
	assert.Nil(t, err)

	//ob.Print()
}

func TestFillOrderBook(t *testing.T) {

	ob := NewOrderBook([]OrderBookRow{}, []OrderBookRow{}, 5)

	update := OrderBookUpdate{7.5, 2, "ask"}
	err := ob.ApplyUpdate(update)
	assert.Nil(t, err)

	update = OrderBookUpdate{6, 2, "bid"}
	err = ob.ApplyUpdate(update)
	assert.Nil(t, err)

	ob.Print()
}

func TestOrderBookGetDepth(t *testing.T) {
	var firstOB, secondOB OrderBook
	firstOB = firstOrderBook()
	secondOB = secondOrderBook()
	result, err := CalculateOverlap(firstOB, secondOB)
	assert.Nil(t, err)
	assert.Equal(t, 3.0, result)

	firstOB = secondOrderBook()
	secondOB = firstOrderBook()
	result, err = CalculateOverlap(firstOB, secondOB)
	assert.Nil(t, err)
	assert.Equal(t, 3.0, result)
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

func firstOrderBook() OrderBook {
	asks := []OrderBookRow{
		{
			Price:  11,
			Volume: 1,
		},
		{
			Price:  12,
			Volume: 1.5,
		},
		{
			Price:  13,
			Volume: 2.5,
		},
		{
			Price:  14,
			Volume: 3,
		},
		{
			Price:  15,
			Volume: 4.5,
		},
		{
			Price:  16,
			Volume: 5.5,
		},
	}

	bids := []OrderBookRow{
		{
			Price:  9,
			Volume: 0.5,
		},
		{
			Price:  8,
			Volume: 1.5,
		},
		{
			Price:  7,
			Volume: 2,
		},
		{
			Price:  6,
			Volume: 2.5,
		},
		{
			Price:  5,
			Volume: 3.5,
		},
		{
			Price:  4,
			Volume: 5,
		},
	}
	return OrderBook{
		Time:     time.Now().Unix(),
		Asks:     asks,
		Bids:     bids,
		MaxDepth: 6,
	}
}

func secondOrderBook() OrderBook {
	asks := []OrderBookRow{
		{
			Price:  15,
			Volume: 1,
		},
		{
			Price:  16,
			Volume: 1,
		},
		{
			Price:  17,
			Volume: 1,
		},

		{
			Price:  18,
			Volume: 1,
		},
		{
			Price:  19,
			Volume: 1,
		},
		{
			Price:  20,
			Volume: 1,
		},
	}

	bids := []OrderBookRow{
		{
			Price:  13,
			Volume: 0.7,
		},
		{
			Price:  12,
			Volume: 2.1,
		},
		{
			Price:  11,
			Volume: 3.4,
		},
		{
			Price:  10,
			Volume: 4.2,
		},
		{
			Price:  9,
			Volume: 5.1,
		},
		{
			Price:  8,
			Volume: 6,
		},
	}

	return OrderBook{
		Time:     time.Now().Unix(),
		Asks:     asks,
		Bids:     bids,
		MaxDepth: 6,
	}
}
