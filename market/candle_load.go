package market

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

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
