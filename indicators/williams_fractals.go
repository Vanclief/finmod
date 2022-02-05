package indicators

import (
	"fmt"
	"github.com/vanclief/finmod/market"
	"math"
)

type WilliamFractal struct {
	Time        int64
	Price       float64
	FractalType int // 5 para la de 5 candles, 6 para la de 6 y así
}

func (wf *WilliamFractal) String() string {
	return fmt.Sprintf("%d %f %d", wf.Time, wf.Price, wf.FractalType)
}

// Taken from https://forexsb.com/wiki/_media/indicators/fractal-types.png

func HasRamp(candles []market.Candle) bool {
	firstThree := candles[:3]
	lastThree := candles[len(candles)-3:]
	return math.Abs(firstThree[0].Close) < math.Abs(firstThree[1].Close) && math.Abs(firstThree[1].Close) < math.Abs(firstThree[2].Close) &&
		math.Abs(lastThree[0].Close) > math.Abs(lastThree[1].Close) && math.Abs(lastThree[1].Close) > math.Abs(lastThree[2].Close)
}

func CandleDifferenceIsWithinDifferencePercentage(candle1, candle2 market.Candle, highLow string) bool {
	percentage := 5 / 100.0
	if highLow == "high" {
		return math.Abs(candle1.High-candle2.High)/candle1.High < percentage
	} else {
		return math.Abs(candle1.Low-candle2.Low)/candle1.Low < percentage
	}
}

func Cleanup(dirtyFractals []WilliamFractal) (cleanedFractals []WilliamFractal) {
	for _, v := range dirtyFractals {
		if v.Time != 0 {
			cleanedFractals = append(cleanedFractals, v)
		}
	}
	return cleanedFractals
}

func WilliamsFractalController(candles []market.Candle) (fractals []WilliamFractal) {
	if len(candles) < 5 {
		return
	}
	i := 0
	j := 5
	for {
		if j >= len(candles) {
			break
		}
		testSubset := candles[i:j]
		fractals = append(fractals, WilliamsFractalLocator(testSubset))
		i++
		j++
	}
	return Cleanup(fractals)
}

func WilliamsFractalLocator(candles []market.Candle) (foundFractal WilliamFractal) {
	// All fractals require at least 5 candles
	if len(candles) < 5 {
		return foundFractal
	}
	if !HasRamp(candles) {
		return foundFractal
	}
	discriminantNumberOfCandles := len(candles) - 4
	switch discriminantNumberOfCandles {
	case 1:
		return WilliamFractal{
			Time:        candles[2].Time,
			Price:       candles[2].Close,
			FractalType: 5,
		}
	case 2:
		if CandleDifferenceIsWithinDifferencePercentage(candles[2], candles[3], "high") {
			return WilliamFractal{
				Time:        candles[3].Time,
				Price:       candles[3].Close,
				FractalType: 5,
			}
		} else {
			return foundFractal
		}
	case 3:
		diff23 := CandleDifferenceIsWithinDifferencePercentage(candles[2], candles[3], "high")
		diff24 := CandleDifferenceIsWithinDifferencePercentage(candles[2], candles[4], "high")
		if diff24 {
			if !diff23 {
				return WilliamFractal{
					Time:        candles[3].Time,
					Price:       candles[3].Close,
					FractalType: 5,
				}
			} else {
				return WilliamFractal{
					Time:        candles[4].Time,
					Price:       candles[4].Close,
					FractalType: 5,
				}
			}
		} else {
			return foundFractal
		}
	case 4:
		diff23 := CandleDifferenceIsWithinDifferencePercentage(candles[2], candles[3], "high")
		diff25 := CandleDifferenceIsWithinDifferencePercentage(candles[2], candles[5], "high")
		diff34 := CandleDifferenceIsWithinDifferencePercentage(candles[3], candles[4], "low")
		diff45 := CandleDifferenceIsWithinDifferencePercentage(candles[4], candles[5], "high")
		if diff25 && diff34 {
			if !diff23 && diff45 {
				return WilliamFractal{
					Time:        candles[4].Time,
					Price:       candles[4].Close,
					FractalType: 5,
				}
			} else if diff23 && !diff45 {
				return WilliamFractal{
					Time:        candles[3].Time,
					Price:       candles[3].Close,
					FractalType: 5,
				}
			} else {
				return foundFractal
			}
		} else {
			return foundFractal
		}

	case 5:
		diff23 := CandleDifferenceIsWithinDifferencePercentage(candles[2], candles[3], "high")
		diff24 := CandleDifferenceIsWithinDifferencePercentage(candles[2], candles[4], "high")
		diff46 := CandleDifferenceIsWithinDifferencePercentage(candles[4], candles[6], "high")
		diff34 := CandleDifferenceIsWithinDifferencePercentage(candles[3], candles[4], "low")
		diff45 := CandleDifferenceIsWithinDifferencePercentage(candles[4], candles[5], "low")
		diff56 := CandleDifferenceIsWithinDifferencePercentage(candles[5], candles[6], "high")
		if !diff23 && !diff34 && !diff45 && diff56 && diff24 && diff46 {
			return WilliamFractal{
				Time:        candles[3].Time,
				Price:       candles[3].Close,
				FractalType: 5,
			}
		} else {
			return foundFractal
		}
	}
	return foundFractal
}
