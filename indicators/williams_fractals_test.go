package indicators

import (
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

	// for _, v := range fractals {
	// 	if v.Type == "up" {
	// 		fmt.Printf("%v,%v\n", (v.Time-candles[0].Time)/3600, v.Price+3)
	// 	} else {
	// 		fmt.Printf("%v,%v\n", (v.Time-candles[0].Time)/3600, v.Price-3)
	// 	}
	// }
}

func TestExample(t *testing.T) {

	candles := []market.Candle{
		{Time: time.Now().Unix(), High: 35268.5, Low: 35267},
		{Time: time.Now().Add(1 * time.Second).Unix(), High: 35269.5, Low: 35266},
		{Time: time.Now().Add(2 * time.Second).Unix(), High: 35270.5, Low: 35263},
		{Time: time.Now().Add(3 * time.Second).Unix(), High: 35267, Low: 35263},
		{Time: time.Now().Add(4 * time.Second).Unix(), High: 35268, Low: 35265},
		{Time: time.Now().Add(5 * time.Second).Unix(), High: 35266, Low: 35263},
		{Time: time.Now().Add(6 * time.Second).Unix(), High: 35266.5, Low: 35264},
	}

	fractals := WilliamFractals(candles)
	assert.Len(t, fractals, 2)
	assert.NotNil(t, fractals)
}
