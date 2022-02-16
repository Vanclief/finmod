package indicators

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vanclief/finmod/market"
)

func TestWilliamsFractal(t *testing.T) {

	candles, _, _, _, _, err := loadCandlesFromFile("./test_dataset/BINANCE_ETHUSD_60.csv")
	assert.Nil(t, err)

	fractals := WilliamFractals(candles[:1])
	assert.Len(t, fractals, 0)
	assert.Nil(t, fractals)

	fractals = WilliamFractals(candles[:2])
	assert.Len(t, fractals, 0)
	assert.Nil(t, fractals)

	fractals = WilliamFractals(candles[:3])
	assert.Len(t, fractals, 0)
	assert.Nil(t, fractals)

	fractals = WilliamFractals(candles[:4])
	assert.Len(t, fractals, 0)
	assert.Nil(t, fractals)

	fractals = WilliamFractals(candles[:5])
	assert.Len(t, fractals, 1)
	assert.NotNil(t, fractals)

	fractals = WilliamFractals(candles)
	assert.NotNil(t, fractals)

	// Noria prints
	for _, v := range fractals {
		if v.Type == "up" {
			fmt.Printf("%v,%v\n", (v.Time-candles[0].Time)/3600, v.Price+3)
		} else {
			fmt.Printf("%v,%v\n", (v.Time-candles[0].Time)/3600, v.Price-3)
		}
	}
}

func TestExample(t *testing.T) {

	candles := []market.Candle{
		{Time: time.Now().Unix(), Open: 10, Close: 10, High: 10, Low: 14},
		{Time: time.Now().Add(1 * time.Second).Unix(), Open: 10, Close: 10, High: 15, Low: 12},
		{Time: time.Now().Add(2 * time.Second).Unix(), Open: 10, Close: 10, High: 20, Low: 10},
		{Time: time.Now().Add(3 * time.Second).Unix(), Open: 10, Close: 10, High: 15, Low: 12},
		{Time: time.Now().Add(4 * time.Second).Unix(), Open: 10, Close: 10, High: 10, Low: 14},
	}

	fractals := WilliamFractals(candles)
	assert.Len(t, fractals, 2)
	assert.NotNil(t, fractals)
}
