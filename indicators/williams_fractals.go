package indicators

import (
	"fmt"
	"github.com/vanclief/finmod/market"
)

type FractalType string

var (
	FractalUp   = FractalType("up")
	FractalDown = FractalType("down")
)

type WilliamFractal struct {
	Time          int64
	Price         float64
	Type          FractalType
	Candles       []market.Candle
	FractalLength int // 5 para la de 5 candles, 6 para la de 6 y así
}

func (wf *WilliamFractal) String() string {
	candlesStrings := ""
	for i, candle := range wf.Candles {
		candlesStrings += fmt.Sprintf("%v,%v\n", i, candle.High)
	}
	for i, candle := range wf.Candles {
		candlesStrings += fmt.Sprintf("%v,%v\n", i, candle.Low)
	}
	return fmt.Sprintf("%d %f %v %d\n%v", wf.Time, wf.Price, wf.Type, wf.FractalLength, candlesStrings)
}

func AllCandlesHighsLowerThanTestCandle(candles []market.Candle, testCandle market.Candle) bool {
	for _, candle := range candles {
		if candle.Time == testCandle.Time {
			continue
		}
		if candle.High > testCandle.High {
			return false
		}
	}
	return true
}

func AllCandlesLowsHigherThanTestCandle(candles []market.Candle, testCandle market.Candle) bool {
	for _, candle := range candles {
		if candle.Time == testCandle.Time {
			continue
		}
		if candle.Low < testCandle.Low {
			return false
		}
	}
	return true
}

func CandlesContainFractal(candles []market.Candle) *WilliamFractal {
	if len(candles) < 5 {
		return nil
	}
	if len(candles)%2 == 0 {
		return nil
	}
	middleCandle := candles[len(candles)/2]
	if AllCandlesHighsLowerThanTestCandle(candles, middleCandle) {
		return &WilliamFractal{
			Time:          middleCandle.Time,
			Price:         middleCandle.High,
			Type:          FractalUp,
			Candles:       candles,
			FractalLength: len(candles),
		}
	} else if AllCandlesLowsHigherThanTestCandle(candles, middleCandle) {
		return &WilliamFractal{
			Time:          middleCandle.Time,
			Price:         middleCandle.Low,
			Type:          FractalDown,
			Candles:       candles,
			FractalLength: len(candles),
		}
	}
	return nil
}

func WilliamsFractalController(candles []market.Candle, gapSize int) (fractals []WilliamFractal) {
	if len(candles) < gapSize {
		return
	}
	i := 0
	j := gapSize
	for {
		if j >= len(candles) {
			break
		}
		testSubset := candles[i:j]
		fractal := CandlesContainFractal(testSubset)
		if fractal != nil {
			fractals = append(fractals, *fractal)
		}
		i++
		j++
	}
	return fractals
}
