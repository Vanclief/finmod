package market

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "time"
)

func TestOrderBook_ReturnVolumeOfOrderBook(t *testing.T) {
  sampleOrderBook := ReturnOrderBook()

  index, err := sampleOrderBook.ReturnVolumeOfOrderBook(5)
  assert.Nil(t, err)
  assert.Equal(t, float64(5), index)
  index, err = sampleOrderBook.ReturnVolumeOfOrderBook(11)
  assert.Nil(t, err)
  assert.Equal(t, float64(11), index)
  index, err = sampleOrderBook.ReturnVolumeOfOrderBook(14.33)
  assert.Nil(t, err)
  assert.Equal(t, float64(15), index)
  index, err = sampleOrderBook.ReturnVolumeOfOrderBook(100)
  assert.NotNil(t, err)
  assert.Equal(t, float64(-1), index)
  index, err = sampleOrderBook.ReturnVolumeOfOrderBook(-2)
  assert.NotNil(t, err)
  assert.Equal(t, float64(-1), index)
  index, err = sampleOrderBook.ReturnVolumeOfOrderBook(9)
  assert.Nil(t, err)
  assert.Equal(t, float64(9), index)
}

func ReturnOrderBook() OrderBook {
  return OrderBook {
    Time: float64(time.Now().Unix()),
    Asks: []OrderBookRow{
      {
        Price: 0,
        Volume: 1,
        TotalVolume: 1,
      },
      {
        Price: 1,
        Volume: 1,
        TotalVolume: 2,
      },
      {
        Price: 2,
        Volume: 1,
        TotalVolume: 3,
      },
      {
        Price: 3,
        Volume: 1,
        TotalVolume: 4,
      },
      {
        Price: 4,
        Volume: 1,
        TotalVolume: 5,
      },
      {
        Price: 5,
        Volume: 1,
        TotalVolume: 6,
      },
      {
        Price: 6,
        Volume: 1,
        TotalVolume: 7,
      },
      {
        Price: 7,
        Volume: 1,
        TotalVolume: 8,
      },
      {
        Price: 8,
        Volume: 1,
        TotalVolume: 9,
      },
      {
        Price: 9,
        Volume: 1,
        TotalVolume: 10,
      },
    },
    Bids: []OrderBookRow {
      {
        Price: 20,
        Volume: 1,
        TotalVolume: 1,
      },
      {
        Price: 19,
        Volume: 1,
        TotalVolume: 2,
      },
      {
        Price: 18,
        Volume: 1,
        TotalVolume: 3,
      },
      {
        Price: 17,
        Volume: 1,
        TotalVolume: 4,
      },
      {
        Price: 16,
        Volume: 1,
        TotalVolume: 5,
      },
      {
        Price: 15,
        Volume: 1,
        TotalVolume: 6,
      },
      {
        Price: 14,
        Volume: 1,
        TotalVolume: 7,
      },
      {
        Price: 13,
        Volume: 1,
        TotalVolume: 8,
      },
      {
        Price: 12,
        Volume: 1,
        TotalVolume: 9,
      },
      {
        Price: 11,
        Volume: 1,
        TotalVolume: 10,
      },
    },
  }
}
