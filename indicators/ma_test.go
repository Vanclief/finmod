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

	candles, expectedMA, _, _, _, err := loadCandlesFromFile("./test_dataset/BINANCE_ETHUSD_60.csv")
	assert.Nil(t, err)
	assert.NotNil(t, candles)

	ma, err := MovingAverage(candles, length)
	assert.Nil(t, err)

	for k := range expectedMA[length-1:] {
		assert.LessOrEqual(t, math.Abs(float64(expectedMA[length-1+k])-ma[k]), 0.1)
	}

	ma, err = MovingAverage(candles[:length-4], length)
	assert.NotNil(t, err)
	assert.Nil(t, ma)
}

func TestSmoothedMovingAverage(t *testing.T) {
	length := 55

	candles, _, _, _, _, err := loadCandlesFromFile("./test_dataset/BINANCE_ETHUSD_60.csv")
	assert.Nil(t, err)

	smma, err := SmoothedMovingAverage(candles, length)
	assert.Nil(t, err)
	assert.NotNil(t, smma)
	assert.Len(t, smma, len(candles)-55)

	// candles, _, _, _, _, err = loadCandlesFromFile("./test_dataset/BINANCE_ETHUSD_60.csv")
	// assert.Nil(t, err)

	// smma, err = SmoothedMovingAverage(candles[:56], length)
	// assert.Nil(t, err)
	// assert.NotNil(t, smma)
	// assert.Len(t, smma, len(candles)-55)
}

func loadCandlesFromFile(filepath string) ([]market.Candle, []float32, []float32, []float32, []float32, error) {
	const op = "loadCandlesFromFile"

	var candles []market.Candle
	var movingAverage []float32
	var volumeWAP []float32
	var volume []float32
	var RSI []float32

	// Load a candle dataset example
	f, err := os.Open(filepath)
	if err != nil {
		return nil, nil, nil, nil, nil, ez.Wrap(op, err)
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
			return nil, nil, nil, nil, nil, ez.Wrap(op, err)
		}

		o, err := strconv.ParseFloat(lineArr[1], 64)
		if err != nil {
			return nil, nil, nil, nil, nil, ez.Wrap(op, err)
		}

		h, err := strconv.ParseFloat(lineArr[2], 64)
		if err != nil {
			return nil, nil, nil, nil, nil, ez.Wrap(op, err)
		}

		l, err := strconv.ParseFloat(lineArr[3], 64)
		if err != nil {
			return nil, nil, nil, nil, nil, ez.Wrap(op, err)
		}

		c, err := strconv.ParseFloat(lineArr[4], 64)
		if err != nil {
			return nil, nil, nil, nil, nil, ez.Wrap(op, err)
		}

		ma, err := strconv.ParseFloat(lineArr[5], 64)
		if err != nil {
			return nil, nil, nil, nil, nil, ez.Wrap(op, err)
		}

		vwap, err := strconv.ParseFloat(lineArr[6], 64)
		if err != nil {
			return nil, nil, nil, nil, nil, ez.Wrap(op, err)
		}

		vol, err := strconv.ParseFloat(lineArr[7], 64)
		if err != nil {
			return nil, nil, nil, nil, nil, ez.Wrap(op, err)
		}

		rsi, err := strconv.ParseFloat(lineArr[8], 64)
		if err != nil {
			return nil, nil, nil, nil, nil, ez.Wrap(op, err)
		}

		candle := market.Candle{
			Time:  time,
			Open:  o,
			High:  h,
			Low:   l,
			Close: c,
		}

		candles = append(candles, candle)
		movingAverage = append(movingAverage, float32(ma))
		volumeWAP = append(volumeWAP, float32(vwap))
		volume = append(volume, float32(vol))
		RSI = append(RSI, float32(rsi))
	}

	return candles, movingAverage, volumeWAP, volume, RSI, nil
}
