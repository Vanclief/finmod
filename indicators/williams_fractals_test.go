package indicators

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWilliamsFractal(t *testing.T) {
	candles, _, _, _, _, err := loadCandlesFromFile("./test_dataset/BINANCE_ETHUSD_60.csv")
	assert.Nil(t, err)
	ans := WilliamsFractalController(candles)
	for _, v := range ans {
		fmt.Println(v.String())
	}
}
