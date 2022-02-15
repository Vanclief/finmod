package indicators

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
