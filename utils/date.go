package utils

import (
	"time"

	"github.com/vanclief/ez"
)

func RawDateToRFC3339(date string) (time.Time, error) {
	const op = "RawDateToRFC3339"

	parsedDate, err := time.Parse(time.RFC3339, date+"T00:00:00Z") // readed hardcoded HH:MM:SS, makes our lives easier as during a fetch or backtest we don't really want to define the hour - Franco
	if err != nil {
		return parsedDate, ez.Wrap(op, err)
	}

	return parsedDate, nil
}
