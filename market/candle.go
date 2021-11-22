package market

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/vanclief/ez"
	"github.com/vanclief/finmod/utils"
)

// Candle represents the high, low, open, and closing prices of an asset or security for a specific period
type Candle struct {
	Time   int64   `json:"id"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Volume float64 `json:"volume"`
}

func (p *Candle) String() string {
	unixTime := time.Unix(p.Time, 0)
	return fmt.Sprintf(
		"Time: %s | Open: %.4f | High: %.4f | Low: %.4f | Close: %.4f | Volume: %.4f",
		unixTime,
		p.Open,
		p.High,
		p.Low,
		p.Close,
		p.Volume,
	)
}

// ModifyInterval - takes an array of candles and the desired interval in minutes and returns an array of
// candles with the market data resampled to fit the new interval
func ModifyInterval(candles []Candle, minuteInterval int) ([]Candle, error) {
	op := "market.ModifyInterval"

	// If we only have a candle, return the candle, there is nothing to do
	if len(candles) == 1 {
		return candles, nil
	}

	// Check that if we are requesting a minute interval that there is actually
	// a minute interval between the candles
	minuteCandleDelta := (candles[1].Time - candles[0].Time) / 60
	if minuteCandleDelta != 1 && minuteInterval == 1 {
		return nil, ez.New(op, ez.ECONFLICT, "No 1 minute interval exists", nil)
	}

	if int64(minuteInterval)%minuteCandleDelta != 0 {
		return nil, ez.New(op, ez.ECONFLICT, fmt.Sprintf(`The retrieved candles cannot be modified to the %d minutes timeframe, the delta between candles is %d minutes`, minuteInterval, minuteCandleDelta), nil)
	}

	start, err := findFirstIndex(candles, minuteInterval)
	if err != nil {
		return nil, ez.Wrap(op, err)
	}

	var modifiedCandles []Candle
	firstCandleIndex := start
	pivot := candles[start].Time
	for k, v := range candles[firstCandleIndex+1:] {
		if (utils.Abs(v.Time-pivot)/60)%minuteInterval == 0 {
			compressedCandle, err := compressCandle(candles, start, k+firstCandleIndex+1)
			if err != nil {
				return nil, err
			}
			modifiedCandles = append(modifiedCandles, compressedCandle)
			start = k + firstCandleIndex + 1
		}
	}

	// compress the rest of the candles
	compressedCandle, err := compressCandle(candles, start, len(candles))
	if err != nil {
		return nil, err
	}

	modifiedCandles = append(modifiedCandles, compressedCandle)

	return modifiedCandles, nil
}

// GetMinInterval - Returns the minimun period interval in minutes from an array of
// candles
func GetMinInterval(candles []Candle) int {
	minDiff := math.MaxFloat64
	for k := range candles {
		if k == 0 {
			continue
		}
		newDiff := math.Abs(float64(candles[k].Time - candles[k-1].Time))
		if newDiff < minDiff {
			minDiff = newDiff
		}
	}
	return int(minDiff) / 60
}

// findFirstIndex iterates over the candle array until it finds the first candle with a timestamp that is a
// multiple of the desired interval in minutes, or is a 0 hour (e.g. 1:00:00, 20:00:00)
func findFirstIndex(candles []Candle, minuteInterval int) (int, error) {
	op := "market.findFirstCandle"

	// We get the mod of the first candle's timestamp and the desired interval
	modMinutes := minuteInterval
	if minuteInterval > 60 {
		modMinutes = minuteInterval % 60
	}
	if modMinutes == 0 {
		modMinutes = 60
	}

	i := 0
	for {
		if i == len(candles) {
			return -1, ez.New(op, ez.EINVALID, "Could not find the initial candle", nil)
		}

		currentMinute := time.Unix(candles[i].Time, 0).Minute()
		if currentMinute%modMinutes == 0 || currentMinute == 0 {
			break
		}

		i++
	}

	return i, nil
}

// compressCandle takes the candle array, a start and end indexes and returns a candle that contains
// the candle data in an accumulated form
func compressCandle(candles []Candle, start, end int) (Candle, error) {
	op := "candle.CompressCandle"

	compressedCandle := Candle{
		Time:   candles[start].Time,
		Open:   candles[start].Open,
		High:   0,
		Low:    0,
		Close:  candles[end-1].Close,
		Volume: 0,
	}

	high := -1e308
	low := 1e308
	accum := 0.0
	for _, v := range candles[start:end] {
		if v.Low < low {
			low = v.Low
		}
		if v.High > high {
			high = v.High
		}
		accum += v.Volume
	}
	floatAccum, err := strconv.ParseFloat(fmt.Sprintf("%.5f", accum), 64)
	if err != nil {
		return Candle{}, ez.Wrap(op, err)
	}
	compressedCandle.High = high
	compressedCandle.Low = low
	compressedCandle.Volume = floatAccum
	return compressedCandle, nil
}
