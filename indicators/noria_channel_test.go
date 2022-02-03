package indicators

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoriaChannel(t *testing.T) {
	// Golden ratio
	//fitInsidePercentage := 2 / (1 + math.Sqrt(5))
	fitInsidePercentage := 0.95
	testChannelFitsInChannelFactor := 2.2
	candles, _, _, _, _, err := loadCandlesFromFile("./test_dataset/BINANCE_ETHUSD_60.csv")
	assert.Nil(t, err)
	ans := Iterator(candles, fitInsidePercentage, testChannelFitsInChannelFactor)
	for _, v := range ans {
		v.Print()
	}
}
