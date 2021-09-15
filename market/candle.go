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
func ModifyInterval(candles []Candle, minutes int) ([]Candle, error) {
	op := "market.ModifyInterval"

	if len(candles) == 1 {
		return candles, nil
	}

	diffInMinutes := (candles[1].Time - candles[0].Time) / 60
	if diffInMinutes != 1 && minutes == 1 {
		return nil, ez.New(op, ez.ECONFLICT, "No 1 minute interval exists", nil)
	}
	if int64(minutes)%diffInMinutes != 0 {
		return nil, ez.New(op, ez.ECONFLICT, "The retrieved candles cannot be modified to the requested timeframe", nil)
	}

	var newCandles []Candle
	modMinutes := minutes
	if minutes > 60 {
		modMinutes = minutes % 60
	}

	if modMinutes == 0 {
		modMinutes = 60
	}

	start, err := findFirstIndex(candles, modMinutes)
	if err != nil {
		return nil, ez.Wrap(op, err)
	}

	indexOfFirstCandle := start
	pivot := candles[start].Time
	for k, v := range candles[indexOfFirstCandle+1:] {
		if (utils.Abs(v.Time-pivot)/60)%minutes == 0 {
			compressedCandle, err := compressCandle(candles, start, k+indexOfFirstCandle+1)
			if err != nil {
				return nil, err
			}
			newCandles = append(newCandles, compressedCandle)
			start = k + indexOfFirstCandle + 1
		}
	}

	// compress the rest of the candles
	compressedCandle, err := compressCandle(candles, start, len(candles))
	if err != nil {
		return nil, err
	}

	newCandles = append(newCandles, compressedCandle)

	return newCandles, nil
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
func findFirstIndex(candles []Candle, modMinutes int) (int, error) {
	op := "market.findFirstCandle"
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
