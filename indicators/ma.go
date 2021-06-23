package indicators

import (
	"github.com/vanclief/finmod/market"
)

func Average(input []market.Candle) float32 {
	result := 0.0
	for _, v := range input {
		result += v.Close
	}
	return float32(result)/float32(len(input))
}

func MovingAverage(candles []market.Candle, length int) ([]float32, error) {
	var movingAverageArray []float32
	i := 0
	j := length - 1
	for {
		if j == len(candles) {
			break
		}
		movingAverageArray = append(movingAverageArray, Average(candles[i : j + 1]))
		i++
		j++
	}

	return movingAverageArray, nil
}
