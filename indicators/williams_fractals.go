package indicators

import (
	"fmt"
	"sort"

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
	Time  int64
	Price float64
	Type  FractalType
}

func (wf *WilliamFractal) String() string {
	return fmt.Sprintf("%d %f %v", wf.Time, wf.Price, wf.Type)
}

func WilliamFractals(candles []market.Candle) (fractals []WilliamFractal) {
	// We need at least 5 candles to find a fractal

	if len(candles) < minCandles {
		return fractals
	}

	for i := 0; i <= len(candles)-5; i++ {
		// While we have enough candles to find a fractal
		foundFractals := findFractals(candles, i)
		fractals = append(fractals, foundFractals...)
	}
	return fractals
}

func findFractals(candles []market.Candle, start int) (foundFractals []WilliamFractal) {

	upFractal, err := findFractal(candles, start, FractalUp)
	if err == nil {
		foundFractals = append(foundFractals, upFractal)
	}

	downFractal, err := findFractal(candles, start, FractalDown)
	if err == nil {
		foundFractals = append(foundFractals, downFractal)
	}

	sort.SliceStable(foundFractals, func(i, j int) bool {
		return foundFractals[i].Time < foundFractals[j].Time
	})

	return foundFractals
}

func findFractal(candles []market.Candle, start int, fType FractalType) (foundFractal WilliamFractal, err error) {
	const op = "findFractal"

	// Step 1) Create our posible fractal
	thirdCandle := candles[start+2]
	foundFractal.Time = thirdCandle.Time

	if fType == FractalUp {
		foundFractal.Price = thirdCandle.High
		foundFractal.Type = FractalUp
	} else {
		foundFractal.Price = thirdCandle.Low
		foundFractal.Type = FractalDown
	}

	// Step 2) Check that the candles of the left are lower than the third candle
	for i := 0; i < 2; i++ {
		if fType == FractalUp && candles[start+i].High >= thirdCandle.High {
			return foundFractal, ez.New(op, ez.EINVALID, "[FractalUp] Third candle doesn't have two preceding lows", nil)
		}

		if fType == FractalDown && candles[start+i].Low <= thirdCandle.Low {
			return foundFractal, ez.New(op, ez.EINVALID, "[FractalDown] Third candle doesn't have two preceding highs", nil)
		}
	}

	// Step 2) Check that there are two lower candles on the right
	rightCandleCount := 0
	for i := start + 3; i < len(candles)-1; i++ {

		fmt.Println("RightCandleCount", rightCandleCount)

		if rightCandleCount == 2 {
			return foundFractal, nil
		}

		if fType == FractalUp {
			if candles[i].High >= thirdCandle.High {
				return foundFractal, ez.New(op, ez.EINVALID, "[FractalUp] No fractal, as new high found", nil)
			} else if candles[i].High < thirdCandle.High {
				rightCandleCount++
			}
		}

		if fType == FractalDown {
			if candles[i].Low <= thirdCandle.Low {
				return foundFractal, ez.New(op, ez.EINVALID, "[FractalDown] No fractal, as new low found", nil)
			} else if candles[i].Low > thirdCandle.Low {
				rightCandleCount++
			}
		}

	}

	return foundFractal, nil
}
