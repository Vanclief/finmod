package market

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vanclief/ez"
)

func SaveCandlesToFile(candles *[]Candle, filepath string) error {
	const op = "market.SaveCandlesToFile"

	f, err := os.Create(filepath)
	if err != nil {
		return ez.Wrap(op, err)
	}

	defer f.Close()
	candleList := *candles

	for _, candle := range candleList {
		fmt.Fprintln(f, candle) // print values to f, one per line
	}

	return nil
}

func LoadCandlesFromFile(filepath string) ([]Candle, error) {
	const op = "market.loadCandlesFromFile"

	var candles []Candle

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

		v, err := strconv.ParseFloat(lineArr[5], 64)
		if err != nil {
			return nil, ez.Wrap(op, err)
		}

		candle := Candle{
			Time:   time,
			Open:   o,
			High:   h,
			Low:    l,
			Close:  c,
			Volume: v,
		}

		candles = append(candles, candle)
	}

	return candles, nil
}

func TestModifyInterval(t *testing.T) {

	// Load the test candles
	candles1Min, err := LoadCandlesFromFile("test_dataset/candles_1_min")
	assert.Nil(t, err)
	candles5Min, err := LoadCandlesFromFile("test_dataset/candles_5_min")
	assert.Nil(t, err)
	candles15Min, err := LoadCandlesFromFile("test_dataset/candles_15_min")
	assert.Nil(t, err)
	candles1H, err := LoadCandlesFromFile("test_dataset/candles_1_h")
	assert.Nil(t, err)

	// Change granularity to 5 min
	candles, err := ModifyInterval(candles1Min, 5)
	assert.Nil(t, err)
	assert.NotNil(t, candles)
	assert.Equal(t, candles5Min, candles)

	// Change granularity to 15 min
	candles, err = ModifyInterval(candles1Min, 15)
	assert.Nil(t, err)
	assert.NotNil(t, candles)
	assert.Equal(t, candles15Min, candles)

	// Change granularity to 1 Hour
	candles, err = ModifyInterval(candles1Min, 60)
	assert.Nil(t, err)
	assert.NotNil(t, candles)
	assert.Equal(t, candles1H, candles)

	// What happens if we only have 1 candle
	candles, err = ModifyInterval(candles1Min[0:1], 1)
	assert.Nil(t, err)
	assert.NotNil(t, candles)

	// What happens if we only have 2 candles
	candles, err = ModifyInterval(candles1Min[0:2], 2)
	assert.Nil(t, err)
	assert.NotNil(t, candles)
	assert.Len(t, candles, 1)

	// Candles with 5 minute time interval cannot be changed to 2 minute interval because 5 % 2 != 0
	candles, err = ModifyInterval(candles5Min, 2)
	assert.Nil(t, candles)
	assert.NotNil(t, err)

	// Candles with 15 minute granularity cannot be changed into 1 minute
	candles, err = ModifyInterval(candles15Min, 1)
	assert.Nil(t, candles)
	assert.NotNil(t, err)

	// Candles with 5 minute granularity can be turned into any multiple of 5 candles
	candles, err = ModifyInterval(candles5Min, 15)
	assert.Nil(t, err)
	assert.NotNil(t, candles)

	// ... but not 16 because 16 % 5 != 0
	candles, err = ModifyInterval(candles5Min, 16)
	assert.Nil(t, candles)
	assert.NotNil(t, err)

	// 1 hour candles can be modified to 2 hour candles
	candles, err = ModifyInterval(candles1H, 120)
	assert.Nil(t, err)
	assert.NotNil(t, candles)
}
