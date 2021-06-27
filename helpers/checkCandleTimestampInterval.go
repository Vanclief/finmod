package helpers

import (
  "github.com/vanclief/finmod/market"
  "math"
)

// CheckCandleTimestampInterval returns an integer that represents the minimum timestamp interval
// in minutes in an array of market.Candle
func CheckCandleTimestampInterval(candles []market.Candle) int {
  minDiff := math.MaxFloat64
  for k := range candles {
    if k == 0 {
      continue
    }
    newDiff := math.Abs(float64(candles[k].Time - candles[k - 1].Time))
    if newDiff < minDiff {
      minDiff = newDiff
    }
  }
  return int(minDiff) / 60
}
