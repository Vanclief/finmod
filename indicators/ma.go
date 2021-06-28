package indicators

import (
	"github.com/vanclief/finmod/market"
)

func MovingAverage(candles []market.Candle, length int) ([]float32, error) {
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
