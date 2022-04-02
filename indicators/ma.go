package indicators

import (
	"github.com/vanclief/ez"
	"github.com/vanclief/finmod/market"
)

func MovingAverage(candles []market.Candle, period int) ([]float64, error) {
	op := "indicators.MovingAverage"

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

func ExponentialMovingAverage(candles []market.Candle, period int) ([]float64, error) {
	const op = "indicators.ExponentialMovingAverage"

	if candles == nil {
		return nil, ez.New(op, ez.EINVALID, "Candle array missing", nil)
	} else if len(candles) < period {
		return nil, ez.New(op, ez.EINVALID, "Period argument is larger than the length of candles", nil)
	} else if period <= 0 {
		return nil, ez.New(op, ez.EINVALID, "Period can't be less than 1", nil)
	}

	var emaArray []float64
	p := 2 / (float64(period) + 1)

	for i := range candles {
		if i < period {
			continue
		}

		if len(emaArray) == 0 {
			emaArray = append(emaArray, average(candles[:period]))
		} else {
			ema := (candles[i].Close * p) + (emaArray[len(emaArray)-1] * (1 - p))
			emaArray = append(emaArray, ema)
		}

	}

	return emaArray, nil
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

	var smmaArray []float64

	for i := range candles {
		if i < period-1 {
			continue
		}

		sum, err := closingSum(candles[:i+1], period)
		if err != nil {
			return nil, ez.Wrap(op, err)
		}

		if len(smmaArray) == 0 {
			smma1 := sum / float64(period)
			smmaArray = append(smmaArray, smma1)
		} else if len(smmaArray) == 1 {
			smma2 := (smmaArray[0]*(float64(period)-1.0) + candles[i].Close) / float64(period)
			smmaArray = append(smmaArray, smma2)
		} else {

			prevSMMA := smmaArray[len(smmaArray)-1]
			prevSum := prevSMMA * float64(period)
			smmai := (prevSum - prevSMMA + candles[i].Close) / float64(period)
			smmaArray = append(smmaArray, smmai)
		}
	}

	return smmaArray, nil
}

func closingSum(candles []market.Candle, period int) (sum float64, err error) {
	const op = "indicators.closingSum"

	if len(candles) < period {
		err = ez.New(op, ez.EINVALID, "Period is larger than the amount of candles", nil)
		return
	}

	for _, candle := range candles[len(candles)-period:] {
		sum = sum + candle.Close
	}

	return sum, nil
}
