package indicators

import (
  "github.com/vanclief/ez"
  "github.com/vanclief/finmod/market"
)

func calculateRSI(candles []market.Candle) float32 {
  avgUpward := 0.0
  avgDownward := 0.0
  for i := 1; i < len(candles); i++ {
    if candles[i].Close > candles[i - 1].Close {
      avgUpward += candles[i].Close - candles[i - 1].Close
    } else if candles[i].Close < candles[i - 1].Close {
      avgDownward += candles[i - 1].Close - candles[i].Close
    }
  }
  rs := avgUpward / avgDownward
  //fmt.Println("info:", avgUpward/14, avgDownward/14, rs)
  return float32(100 - 100/(1+rs))
}

func RSI(candles []market.Candle, period int) ([]float32, error) {
  op := "movingAverage"

  if candles == nil {
    return nil, ez.New(op, ez.EINVALID, "Candle array missing", nil)
  } else if len(candles) < 14 {
    return nil, ez.New(op, ez.EINVALID, "Cannot calculate RSI with less than 14 candles", nil)
  }

  var rsiArray []float32

  i := 0
  j := period - 1

  for {
    if j == len(candles) {
      break
    }

    rsiArray = append(rsiArray, calculateRSI(candles[i:j+2]))

    i++
    j++
  }
  return rsiArray, nil
}