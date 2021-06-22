package indicators

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vanclief/ez"
	"github.com/vanclief/finmod/market"
)

func loadCandlesFromFile(filepath string) ([]market.Candle, error) {
	const op = "loadCandlesFromFile"

	var candles []market.Candle

	// Load a candle dataset example
	f, err := os.Open(filepath)
	if err != nil {
		return nil, ez.Wrap(op, err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()

		// Parse each line
		line := text[1 : len(text)-1]
		lineArr := strings.Split(line, " ")

		// Convert from string to datatypes
		time, err := strconv.ParseInt(lineArr[0], 10, 32)
		if err != nil {
			return nil, ez.Wrap(op, err)
		}

		o, err := strconv.ParseFloat(lineArr[1], 64)
		if err != nil {
			return nil, ez.Wrap(op, err)
		}

		h, err := strconv.ParseFloat(lineArr[2], 64)
		if err != nil {
			return nil, ez.Wrap(op, err)
		}

		l, err := strconv.ParseFloat(lineArr[3], 64)
		if err != nil {
			return nil, ez.Wrap(op, err)
		}

		c, err := strconv.ParseFloat(lineArr[4], 64)
		if err != nil {
			return nil, ez.Wrap(op, err)
		}

		candle := market.Candle{
			Time:  time,
			Open:  o,
			High:  h,
			Low:   l,
			Close: c,
		}

		candles = append(candles, candle)
	}

	return candles, nil
}

func loadMAFromFile(filepath string) ([]float32, error) {
	const op = "loadMAFromFile"

	var arrayMA []float32

	// Load a candle dataset example
	f, err := os.Open(filepath)
	if err != nil {
		return nil, ez.Wrap(op, err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()

		// Parse each line
		line := text[1 : len(text)-1]
		lineArr := strings.Split(line, " ")

		// Convert from string to datatypes
		ma, err := strconv.ParseFloat(lineArr[2], 64)
		if err != nil {
			return nil, ez.Wrap(op, err)
		}

		arrayMA = append(arrayMA, float32(ma))
	}

	return arrayMA, nil
}

func TestMovingAverage(t *testing.T) {

	candles, err := loadCandlesFromFile("test_dataset/BINANCE_ETHUSD_60.csv")
	assert.Nil(t, err)

	expectedMA, err := loadMAFromFile("test_dataset/BINANCE_ETHUSD_60.csv")
	assert.Nil(t, err)

	ma, err := MovingAverage(candles, 20)
	assert.Nil(t, err)
	assert.NotNil(t, candles)
	assert.Equal(t, expectedMA, ma)
}
