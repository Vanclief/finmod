package helpers

import (
  "github.com/vanclief/finmod/market"
  "math"
)

func CheckCandleTimestampInterval(candles []market.Candle) int {
  minDiff := -1.0
  for k := range candles {
    if k == 0 {
      continue
    }
    newDiff := math.Abs(float64(candles[k].Time - candles[k - 1].Time))
    if newDiff < minDiff {
      minDiff = newDiff
    }
  }
  return int(minDiff)
}
