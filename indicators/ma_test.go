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

func TestExponentialMovingAverage(t *testing.T) {
	candles := []market.Candle{
		{Close: 34440.1},
		{Close: 34402.6},
		{Close: 34365.1},
		{Close: 34385.6},
		{Close: 34369.6},
		{Close: 34402.6},
		{Close: 34424.1},
		{Close: 34397.6},
		{Close: 34425.1},
		{Close: 34438.6},
		{Close: 34469.6},
		{Close: 34469.6},
		{Close: 34479.6},
		{Close: 34479.6},
		{Close: 34468.6},
		{Close: 34494.6},
		{Close: 34534.6},
		{Close: 34532.6},
		{Close: 34526.6},
		{Close: 34534.6},
		{Close: 34528.6},
		{Close: 34539.6},
		{Close: 34547.6},
		{Close: 34523.6},
		{Close: 34517.6},
	}

	expectedEMA := []float64{
		34483.65,
		34488.97,
		34494.56,
		34497.32,
		34499.25,
	}

	// candles, _, _, _, _, err = loadCandlesFromFile("./test_dataset/BINANCE_ETHUSD_60.csv")
	// assert.Nil(t, err)

	ema, err := ExponentialMovingAverage(candles, 20)
	assert.Nil(t, err)
	assert.NotNil(t, ema)
	assert.Len(t, ema, len(expectedEMA))

	for i := range ema {
		assert.Less(t, expectedEMA[i]-ema[i], 0.4, i)
	}

	candles, _, _, _, _, err = loadCandlesFromFile("./test_dataset/BINANCE_ETHUSD_60.csv")
	assert.Nil(t, err)

	ema, err = ExponentialMovingAverage(candles, 55)
	assert.Nil(t, err)
	assert.NotNil(t, ema)
}

func TestSmoothedMovingAverage(t *testing.T) {
	candles := []market.Candle{
		{Close: 35425},
		{Close: 35456},
		{Close: 35447}, // Inicial 35444.061
		{Close: 35434},
		{Close: 35412},
		{Close: 35412},
		{Close: 35406},
		{Close: 35403},
		{Close: 35403},
		{Close: 35406},
		{Close: 35409},
		{Close: 35405},
	}

	expectedSMMA := []float64{
		35440.5,
		35443.75,
		35439.030,
		35425.515,
		35418.758,
		35412.379,
		35407.689,
		35405.345,
		35405.672,
		35407.336,
		35406.168,
	}

	// candles, _, _, _, _, err = loadCandlesFromFile("./test_dataset/BINANCE_ETHUSD_60.csv")
	// assert.Nil(t, err)

	smma, err := SmoothedMovingAverage(candles, 2)
	assert.Nil(t, err)
	assert.NotNil(t, smma)
	assert.Len(t, smma, len(expectedSMMA))

	for i := range smma {
		assert.Less(t, expectedSMMA[i]-smma[i], 0.4, i)
	}

	candles, _, _, _, _, err = loadCandlesFromFile("./test_dataset/BINANCE_ETHUSD_60.csv")
	assert.Nil(t, err)

	smma, err = SmoothedMovingAverage(candles, 55)
	assert.Nil(t, err)
	assert.NotNil(t, smma)
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
