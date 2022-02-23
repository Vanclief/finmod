package indicators

import (
	"github.com/vanclief/ez"
	"github.com/vanclief/finmod/market"
)

func MovingAverage(candles []market.Candle, period int) ([]float64, error) {
	op := "movingAverage"

	if candles == nil {
		return nil, ez.New(op, ez.EINVALID, "Candle array missing", nil)
	} else if len(candles) < period {
		return nil, ez.New(op, ez.EINVALID, "Period argument is larger than the length of candles", nil)
	} else if period <= 0 {
		return nil, ez.New(op, ez.EINVALID, "Period can't be less than 1", nil)
	}

	var movingAverage []float64

	i := 0
	j := period - 1

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

func average(candles []market.Candle) float64 {
	result := 0.0
	for _, candle := range candles {
		result += candle.Close
	}
	return float64(result) / float64(len(candles))
}

func SmoothedMovingAverage(candles []market.Candle, period int) ([]float64, error) {
	const op = "indicators.SmoothedMovingAverage"

	if candles == nil {
		return nil, ez.New(op, ez.EINVALID, "Candle array missing", nil)
	} else if len(candles) < period {
		return nil, ez.New(op, ez.EINVALID, "Period is larger than the amount of candles", nil)
	} else if period <= 0 {
		return nil, ez.New(op, ez.EINVALID, "Period can't be less than 1", nil)
	}

	var sum float64
	smmaArray := []float64{0}

	for i, candle := range candles {

		if i == 0 {
			continue
		}

		sum = sum + candle.Close
		n := float64(i + 1)

		smma := (sum - smmaArray[i-1]) / n
		smmaArray = append(smmaArray, smma)
	}

	return smmaArray[period:], nil
}
