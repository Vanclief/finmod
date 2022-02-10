package indicators

import (
	"fmt"

	"github.com/vanclief/ez"
	"github.com/vanclief/finmod/market"
)

type FractalType string

const minCandles = 5

var (
	FractalUp   = FractalType("up")
	FractalDown = FractalType("down")
)

type WilliamFractal struct {
	Time          int64
	Price         float64
	Type          FractalType
	FractalLength int // 5 para la de 5 candles, 6 para la de 6 y así
}

func (wf *WilliamFractal) String() string {
	return fmt.Sprintf("%d %f %v %d", wf.Time, wf.Price, wf.Type, wf.FractalLength)
}

func WilliamFractals(candles []market.Candle) (fractals []WilliamFractal) {

	if len(candles) < minCandles {
		return fractals
	}

	i := 0
	j := minCandles

	for {
		if j >= len(candles) {
			break
		}

		candleSubset := candles[i:j]
		foundFractal, err := findFractal(candleSubset)
		if err == nil {
			fractals = append(fractals, foundFractal)
		}

		i++
		j++
	}
	return fractals
}

func findFractal(candles []market.Candle) (foundFractal WilliamFractal, err error) {
	const op = "FindFractal"

	if len(candles) < 5 {
		return foundFractal, ez.New(op, ez.EINVALID, "Not enough candles to find a fractal", nil)
	}

	if len(candles)%2 == 0 {
		return foundFractal, ez.New(op, ez.EINVALID, "Must have odd number of candles", nil)
	}

	middleCandle := candles[len(candles)/2]

	if candleIsHighest(candles, middleCandle) {
		foundFractal.Time = middleCandle.Time
		foundFractal.Price = middleCandle.High
		foundFractal.Type = FractalUp
		foundFractal.FractalLength = len(candles)

		return foundFractal, nil

	} else if candleIsLowest(candles, middleCandle) {
		foundFractal.Time = middleCandle.Time
		foundFractal.Price = middleCandle.Low
		foundFractal.Type = FractalDown
		foundFractal.FractalLength = len(candles)

		return foundFractal, nil
	}

	return foundFractal, ez.New(op, ez.EINVALID, "No fractal found", nil)
}

func candleIsHighest(candles []market.Candle, middleCandle market.Candle) bool {
	for _, candle := range candles {
		if candle.Time == middleCandle.Time {
			continue
		}
		if candle.High > middleCandle.High {
			return false
		}
	}
	return true
}

func candleIsLowest(candles []market.Candle, middleCandle market.Candle) bool {
	for _, candle := range candles {
		if candle.Time == middleCandle.Time {
			continue
		}

		if candle.Low < middleCandle.Low {
			return false
		}
	}
	return true
}
