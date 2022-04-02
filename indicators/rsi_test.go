package indicators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRSI(t *testing.T) {
	candles, _, _, _, _, err := loadCandlesFromFile("./test_dataset/BINANCE_BTCUSDT_60.csv")
	assert.Nil(t, err)

	calculatedRSI, err := RSI(candles, 14)
	// for i := 0; i < 10; i++ {
	// fmt.Println(calculatedRSI[i])
	// }
	errNil := assert.Nil(t, err)
	notNilRSI := assert.NotNil(t, calculatedRSI)
	if !errNil && notNilRSI {
		return
	}
	// TODO: trading view uses Wilder, this test uses Cutler
	//for k := range rsi[13:] {
	//  assert.LessOrEqual(t, math.Abs(float64(rsi[13+k]-calculatedRSI[k])), 0.1)
	//}
	//calculatedRSI, err = MovingAverage(candles[:10], 14)
	//assert.NotNil(t, err)
	//assert.Nil(t, calculatedRSI)
}
