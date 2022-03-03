package indicators

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoriaChannel(t *testing.T) {
	// Golden ratio
	//percentageOfCandlesThatMustFitInsideChannel := 2 / (1 + math.Sqrt(5))
	percentageOfCandlesThatMustFitInsideChannel := 0.95
	testCandleFitsInChannelFactor := 2.2
	candles, _, _, _, _, err := loadCandlesFromFile("./test_dataset/BINANCE_ETHUSD_60.csv")
	assert.Nil(t, err)
	ans := Iterator(candles, percentageOfCandlesThatMustFitInsideChannel, testCandleFitsInChannelFactor)
	for _, v := range ans {
		fmt.Println(v.String())
	}
}
