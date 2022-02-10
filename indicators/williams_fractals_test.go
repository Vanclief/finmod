package indicators

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWilliamsFractal(t *testing.T) {
	fractalSize := 5
	candles, _, _, _, _, err := loadCandlesFromFile("./test_dataset/BINANCE_ETHUSD_60.csv")
	assert.Nil(t, err)
	ans := WilliamsFractalController(candles, fractalSize)
	for _, v := range ans {
		if v.Type == "up" {
			fmt.Printf("%v,%v\n", (v.Time-candles[0].Time)/3600, v.Price+3)
		} else {
			fmt.Printf("%v,%v\n", (v.Time-candles[0].Time)/3600, v.Price-3)
		}
	}
}
