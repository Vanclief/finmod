package indicators

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vanclief/ez"
	"github.com/vanclief/finmod/market"
)

func TestMovingAverage(t *testing.T) {
	length := 20

	candles, expectedMA, err := loadCandlesFromFile("./test_dataset/BINANCE_ETHUSD_60.csv")
	assert.Nil(t, err)

	ma, err := MovingAverage(candles, length)
	errNil := assert.Nil(t, err)
	notNilCandles := assert.NotNil(t, candles)
	if !errNil && notNilCandles {
		return
	}
	for k := range expectedMA[length-1:] {
		assert.LessOrEqual(t, math.Abs(float64(expectedMA[length-1+k]-ma[k])), 0.1)
	}

}

func loadCandlesFromFile(filepath string) ([]market.Candle, []float32, error) {
	const op = "loadCandlesFromFile"

	var candles []market.Candle
	var movingAverage []float32

	// Load a candle dataset example
	f, err := os.Open(filepath)
	if err != nil {
		return nil, nil, ez.Wrap(op, err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()

		// Parse each line
		lineArr := strings.Split(text, ",")
		if lineArr[0] == "time" {
			continue
		}

		// Convert from string to datatypes
		time, err := strconv.ParseInt(lineArr[0], 10, 32)
		if err != nil {
			return nil, nil, ez.Wrap(op, err)
		}

		o, err := strconv.ParseFloat(lineArr[1], 64)
		if err != nil {
			return nil, nil, ez.Wrap(op, err)
		}

		h, err := strconv.ParseFloat(lineArr[2], 64)
		if err != nil {
			return nil, nil, ez.Wrap(op, err)
		}

		l, err := strconv.ParseFloat(lineArr[3], 64)
		if err != nil {
			return nil, nil, ez.Wrap(op, err)
		}

		c, err := strconv.ParseFloat(lineArr[4], 64)
		if err != nil {
			return nil, nil, ez.Wrap(op, err)
		}

		ma, err := strconv.ParseFloat(lineArr[5], 64)

		candle := market.Candle{
			Time:  time,
			Open:  o,
			High:  h,
			Low:   l,
			Close: c,
		}

		candles = append(candles, candle)
		movingAverage = append(movingAverage, float32(ma))
	}

	return candles, movingAverage, nil
}
