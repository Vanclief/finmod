package indicators

import (
	"github.com/vanclief/ez"
	"github.com/vanclief/finmod/market"
)

func MovingAverage(candles []market.Candle, length int) ([]float32, error) {
	op := "movingAverage"
	if candles == nil {
		return nil, ez.New(op, ez.EINVALID, "candle array missing", nil)
	}
	if len(candles) < length {
		return nil, ez.New(op, ez.EINVALID, "length argument is bigger than market.Candle length", nil)
	}
	if length <= 0 {
		return nil, ez.New(op, ez.EINVALID, "length cannot be negative", nil)
	}
	var movingAverage []float32

	i := 0
	j := length - 1

	for {
		if j == len(candles) {
			break
		}

		movingAverage = append(movingAverage, average(candles[i:j+1]))

		i++
		j++
	}

	return movingAverage, nil
}

func average(candles []market.Candle) float32 {
	result := 0.0
	for _, candle := range candles {
		result += candle.Close
	}
	return float32(result) / float32(len(candles))
}
