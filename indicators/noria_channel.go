package indicators

import (
	"fmt"
	"github.com/vanclief/finmod/market"
	"math"
)

type Coordinate struct {
	X float64
	Y float64
}

type Line struct {
	candleStart int64
	candleEnd   int64
	m           float64
	b           float64
}

func (l *Line) Print() string {
	return fmt.Sprintf("start: %d, end: %d, slope: %f, intercept: %f\n", l.candleStart, l.candleEnd, l.m, l.b)
}

type Lines struct {
	lowLine  Line
	highLine Line
}

func (t *Lines) Print() {
	fmt.Printf("Lines:\nlowLine: %vhighLine: %v\n", t.lowLine.Print(), t.highLine.Print())
}

func FindMostHighCandle(candles []market.Candle) Coordinate {
	var maxCandle market.Candle
	var maxCandleIndex int
	for i, candle := range candles {
		if i == 0 {
			maxCandle = candle
			maxCandleIndex = i
		} else if candle.High > maxCandle.High {
			maxCandle = candle
			maxCandleIndex = i
		}
	}
	return Coordinate{float64(maxCandleIndex), maxCandle.High}
}

func FindMostLowCandle(candles []market.Candle) Coordinate {
	var minCandle market.Candle
	var minCandleIndex int
	for i, candle := range candles {
		if i == 0 {
			minCandle = candle
			minCandleIndex = i
		} else if candle.Low < minCandle.Low {
			minCandle = candle
			minCandleIndex = i
		}
	}
	return Coordinate{float64(minCandleIndex), minCandle.Low}
}

func LinearRegression(candles []market.Candle, startCandle, endCandle int64, minmax string) Line {
	if len(candles) == 0 {
		return Line{
			candleStart: 0,
			candleEnd:   0,
			m:           0,
			b:           0,
		}
	}
	var points []float64
	if minmax == "low" {
		for i := 0; i < len(candles); i++ {
			points = append(points, candles[i].Low)
		}
	} else {
		for i := 0; i < len(candles); i++ {
			points = append(points, candles[i].High)
		}
	}

	lowestCandle := FindMostLowCandle(candles)
	highestCandle := FindMostHighCandle(candles)
	n := len(candles)
	sumX := 0.0
	sumY := 0.0
	sumXY := 0.0
	sumX2 := 0.0
	sumY2 := 0.0
	for i := 0; i < n; i++ {
		point := points[i]
		sumX += float64(i)
		sumY += point
		sumXY += float64(i) * point
		sumX2 += float64(i) * float64(i)
		sumY2 += point * point
	}
	//fmt.Printf("n: %v, sumX: %v, sumY: %v, sumXY: %v, sumX2: %v, sumY2: %v\n", n, sumX, sumY, sumXY, sumX2, sumY2)
	b := ((sumY * sumX2) - (sumX * sumXY)) / ((float64(n) * sumX2) - (sumX * sumX))
	m := ((float64(n) * sumXY) - (sumX * sumY)) / ((float64(n) * sumX2) - (sumX * sumX))
	if minmax == "low" {
		b = lowestCandle.Y - lowestCandle.X*m
	} else {
		b = highestCandle.Y - highestCandle.X*m
	}
	fmt.Printf("m: %v, b: %v\n", m, b)
	return Line{
		candleStart: startCandle,
		candleEnd:   endCandle,
		m:           m,
		b:           b,
	}
}

func NoriaChannel(candles []market.Candle, length int) (lines []Lines) {
	if length <= 0 {
		return
	}
	if len(candles) < length || len(candles) == 0 || length == 0 {
		return []Lines{
			{
				lowLine:  Line{},
				highLine: Line{},
			},
		}
	}
	lower := 0
	upper := length
	for upper < len(candles) {
		subarray := candles[lower:upper]
		lines = append(lines, Lines{
			lowLine:  LinearRegression(subarray, subarray[0].Time, subarray[len(subarray)-1].Time, "low"),
			highLine: LinearRegression(subarray, subarray[0].Time, subarray[len(subarray)-1].Time, "high"),
		})
		lower++
		upper++
	}
	return lines
}

func FindProperty(input []Lines, ran int) (property [][]Lines) {
	if len(input) == 0 {
		return [][]Lines{}
	}
	if ran < 0 {
		return [][]Lines{}
	}
	var mini []Lines
	for i := 0; i < len(input); i++ {
		if math.Abs(input[i].lowLine.m) < 1 && math.Abs(input[i].highLine.m) < 1 {
			mini = append(mini, input[i])
		} else {
			if len(mini) > ran {
				property = append(property, mini)
				mini = []Lines{}
			}
		}
	}
	return property
}
