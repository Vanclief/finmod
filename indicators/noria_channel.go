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

func (c *Coordinate) Print() {
	fmt.Printf("%f, %f\n", c.X, c.Y)
}

type Line struct {
	candleStart int64
	candleEnd   int64
	m           float64
	b           float64
}

func (l *Line) Print() string {
	if l.b < 0 {
		return fmt.Sprintf("%fx %f", l.m, l.b)
	} else {
		return fmt.Sprintf("%fx + %f", l.m, l.b)
	}
}

type Lines struct {
	lowLine  Line
	highLine Line
}

func (t *Lines) Print() {
	fmt.Printf("Lines:\nlowLine: %v\nhighLine: %v\n", t.lowLine.Print(), t.highLine.Print())
}

type ChannelSegment struct {
	start    int
	end      int
	lowLine  Line
	highLine Line
}

func (t *ChannelSegment) Print() {
	fmt.Printf("ChannelSegment:\nstart: %d\nend: %d\nlowLine: %v\nhighLine: %v\n----------------\n", t.start, t.end, t.lowLine.Print(), t.highLine.Print())
}

func CandlesMidpoint(candles []market.Candle) (candleLowHighDifference []Coordinate) {
	for i := 0; i < len(candles); i++ {
		candleLowHighDifference = append(candleLowHighDifference, Coordinate{X: float64(i), Y: (candles[i].Low + candles[i].High) / 2})
	}
	return candleLowHighDifference
}

func GetLowHighAvg(candleMidpoints []market.Candle) (average float64) {
	for _, candle := range candleMidpoints {
		average += candle.High - candle.Low
	}
	return average / float64(len(candleMidpoints))
}

func ConvertCandlesToCoordinates(candles []market.Candle, lowHigh string) (candlesCoordinates []Coordinate) {
	for i := 0; i < len(candles); i++ {
		if lowHigh == "low" {
			candlesCoordinates = append(candlesCoordinates, Coordinate{X: float64(i), Y: candles[i].Low})
		} else {
			candlesCoordinates = append(candlesCoordinates, Coordinate{X: float64(i), Y: candles[i].High})
		}
	}
	return candlesCoordinates
}

func RotatePointsByMatrix(points []Coordinate, matrix [][]float64) (rotatedPoints []Coordinate) {
	for i := 0; i < len(points); i++ {
		rotatedPoints = append(rotatedPoints, Coordinate{
			X: points[i].X*matrix[0][0] + points[i].Y*matrix[0][1],
			Y: points[i].X*matrix[1][0] + points[i].Y*matrix[1][1],
		})
	}
	return rotatedPoints
}

func FindMostHighCoordinate(points []Coordinate) (mostHighCoordinate Coordinate) {
	mostHighCoordinate = Coordinate{X: math.Inf(-1), Y: math.Inf(-1)}
	for i := 0; i < len(points); i++ {
		if points[i].Y > mostHighCoordinate.Y {
			mostHighCoordinate = points[i]
		}
	}
	return mostHighCoordinate
}

func FindMostLowCoordinate(points []Coordinate) (mostLowCoordinate Coordinate) {
	mostLowCoordinate = Coordinate{X: math.Inf(1), Y: math.Inf(1)}
	for i := 0; i < len(points); i++ {
		if points[i].Y < mostLowCoordinate.Y {
			mostLowCoordinate = points[i]
		}
	}
	return mostLowCoordinate
}

func LinearRegression(points []Coordinate, startCandle, endCandle int64) Line {
	if len(points) == 0 {
		return Line{
			candleStart: 0,
			candleEnd:   0,
			m:           0,
			b:           0,
		}
	}
	n := float64(len(points))
	sumX := 0.0
	sumY := 0.0
	sumXY := 0.0
	sumX2 := 0.0
	sumY2 := 0.0
	for _, v := range points {
		sumX += v.X
		sumY += v.Y
		sumXY += v.X * v.Y
		sumX2 += math.Pow(v.X, 2)
		sumY2 += math.Pow(v.Y, 2)
	}
	//fmt.Printf("n: %v, sumX: %v, sumY: %v, sumXY: %v, sumX2: %v, sumY2: %v\n", n, sumX, sumY, sumXY, sumX2, sumY2)
	b := ((sumY * sumX2) - (sumX * sumXY)) / ((n * sumX2) - (sumX * sumX))
	m := ((n * sumXY) - (sumX * sumY)) / ((n * sumX2) - (sumX * sumX))

	//fmt.Printf("m: %v, b: %v\n", m, b)
	return Line{
		candleStart: startCandle,
		candleEnd:   endCandle,
		m:           m,
		b:           b,
	}
}

func CandleFitsInChannel(candle market.Candle, lowLine, highLine Line, index int, volatility float64) bool {
	//fmt.Println("candle: ", candle.String())
	//fmt.Println("lowLine: ", lowLine.Print())
	//fmt.Println("highLine: ", highLine.Print())
	//fmt.Println("index: ", index)
	//fmt.Println("volatility: ", volatility)
	factor := 2.2
	lowLineYValue := lowLine.m*float64(index) + lowLine.b
	highLineYValue := highLine.m*float64(index) + highLine.b
	if candle.High > highLineYValue+volatility*factor {
		return false
	} else if candle.Low < lowLineYValue-volatility*factor {
		return false
	}
	return true
}

func Iterator(candles []market.Candle) (channels []ChannelSegment) {
	if len(candles) < 3 {
		return channels
	}

	channel := NoriaChannel(candles[:2], 0.0)
	volatility := GetLowHighAvg(candles[:2])
	i := 0
	j := 2
	for {
		if j >= len(candles)-1 {
			channels = append(channels, ChannelSegment{
				start:    i,
				end:      j,
				lowLine:  channel.lowLine,
				highLine: channel.highLine,
			})
			break
		}
		// Check if test candle fits within existing channel, if it does, recalculate the channel to accommodate the new candle
		// if it does not, break the channel and start a new one
		testCandle := candles[j]
		if CandleFitsInChannel(testCandle, channel.lowLine, channel.highLine, j, volatility) {
			j++
		} else {
			//fmt.Println("i, j: ", i, j)
			channels = append(channels, ChannelSegment{
				start:    i,
				end:      j,
				lowLine:  channel.lowLine,
				highLine: channel.highLine,
			})
			if len(candles)-j < 3 {
				break
			}
			i = j
			j += 3
		}
		channel = NoriaChannel(candles[i:j], float64(i))
		volatility = GetLowHighAvg(candles[i:j])
	}
	return channels
}

func NoriaChannel(candles []market.Candle, offset float64) (lines Lines) {
	if len(candles) < 2 {
		return
	}

	// Golden ratio
	//RATIO := 2 / (1 + math.Sqrt(5))
	RATIO := 0.95

	// 1. Calculate the difference between High and Low for each candle
	candleLowHighDifference := CandlesMidpoint(candles)
	//volatility := GetLowHighAvg(candleLowHighDifference)
	// 2. Calculate Linear Regression using the midpoint
	midLine := LinearRegression(candleLowHighDifference, candles[0].Time, candles[len(candles)-1].Time)
	angle := math.Atan(midLine.m)
	rotationMatrix := [][]float64{
		{math.Cos(angle), math.Sin(angle)},
		{-math.Sin(angle), math.Cos(angle)},
	}
	// 3. Rotate the points by the angle
	lowPointsToRotate := ConvertCandlesToCoordinates(candles, "low")
	highPointsToRotate := ConvertCandlesToCoordinates(candles, "high")
	rotatedLowPoints := RotatePointsByMatrix(lowPointsToRotate, rotationMatrix)
	rotatedHighPoints := RotatePointsByMatrix(highPointsToRotate, rotationMatrix)
	// 4. Offset points by RATIO percentage
	lowCandlesOffset := FindMostLowCoordinate(rotatedLowPoints).Y + (1-RATIO)*(FindMostHighCoordinate(rotatedLowPoints).Y-FindMostLowCoordinate(rotatedLowPoints).Y)
	highCandlesOffset := FindMostLowCoordinate(rotatedHighPoints).Y + RATIO*(FindMostHighCoordinate(rotatedHighPoints).Y-FindMostLowCoordinate(rotatedHighPoints).Y)
	rotatedLowLine := Line{
		candleStart: candles[0].Time,
		candleEnd:   candles[len(candles)-1].Time,
		m:           midLine.m,
		b:           (lowCandlesOffset)/math.Sin(angle+math.Pi/2) - midLine.m*offset,
	}
	rotatedHighLine := Line{
		candleStart: candles[0].Time,
		candleEnd:   candles[len(candles)-1].Time,
		m:           midLine.m,
		b:           (highCandlesOffset)/math.Sin(angle+math.Pi/2) - midLine.m*offset,
	}
	return Lines{
		lowLine:  rotatedLowLine,
		highLine: rotatedHighLine,
	}
}

//func FindProperty(input []Lines, ran int) (property [][]Lines) {
//	if len(input) == 0 {
//		return [][]Lines{}
//	}
//	if ran < 0 {
//		return [][]Lines{}
//	}
//	var mini []Lines
//	for i := 0; i < len(input); i++ {
//		if math.Abs(input[i].lowLine.m) < 1 && math.Abs(input[i].highLine.m) < 1 {
//			mini = append(mini, input[i])
//		} else {
//			if len(mini) > ran {
//				property = append(property, mini)
//				mini = []Lines{}
//			}
//		}
//	}
//	return property
//}
