package indicators

import (
  "fmt"
  "github.com/stretchr/testify/assert"
  "math"
  "testing"
)

func TestVolumeWeightedAveragePrice(t *testing.T) {
  length := 5
  candles, _, expectedVWAP, volume, _, err := loadCandlesFromFile("./test_dataset/BINANCE_BTCUSDT,_60.csv")

  calculatedVWAP, err := VolumeWeightedAveragePrice(candles, volume, length)
  fmt.Println(calculatedVWAP)
  errNil := assert.Nil(t, err)
  notNilCandles := assert.NotNil(t, candles)
  if !errNil && notNilCandles {
    return
  }
  for k := range expectedVWAP[length-1:] {
    assert.LessOrEqual(t, math.Abs(float64(expectedVWAP[length-1+k]-calculatedVWAP[k])), 0.1)
  }
  calculatedVWAP, err = VolumeWeightedAveragePrice(candles[:length-4], volume, length)
  assert.NotNil(t, err)
  assert.Nil(t, calculatedVWAP)
}