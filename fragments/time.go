package fragments

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/vanclief/ez"
)

// WithinHours - Returns true or false depending if a Candle is within
// the HH:MM time window with inclusive limits.
func WithinHours(candleTime time.Time, start, end string) (bool, error) {
	const op = "fragments.WithinHours"

	// Split the strings
	splitStart := strings.Split(start, ":")
	splitEnd := strings.Split(end, ":")

	if len(splitStart) != 2 {
		msg := fmt.Sprintf("Invalid argument for start, should be HH:SS is %s", start)
		return false, ez.New(op, ez.EINVALID, msg, nil)
	} else if len(splitEnd) != 2 {
		msg := fmt.Sprintf("Invalid argument for end, should be HH:SS is %s", start)
		return false, ez.New(op, ez.EINVALID, msg, nil)
	}

	// Parse
	startHour, err := strconv.Atoi(splitStart[0])
	if err != nil {
		ez.Wrap(op, err)
	}

	startMinutes, err := strconv.Atoi(splitStart[1])
	if err != nil {
		ez.Wrap(op, err)
	}

	endHour, err := strconv.Atoi(splitEnd[0])
	if err != nil {
		ez.Wrap(op, err)
	}

	endMinutes, err := strconv.Atoi(splitEnd[1])
	if err != nil {
		ez.Wrap(op, err)
	}

	// Condition
	if startHour > candleTime.Hour() || candleTime.Hour() > endHour {
		return false, nil
	}

	if startHour == candleTime.Hour() && startMinutes > candleTime.Minute() {
		return false, nil
	} else if endHour == candleTime.Hour() && candleTime.Minute() > endMinutes {
		return false, nil
	}

	return true, nil
}
