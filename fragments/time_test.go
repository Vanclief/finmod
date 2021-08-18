package fragments

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vanclief/finmod/market"
)

func TestWithinTimeWindow(t *testing.T) {

	// Sunday, August 15, 2021 12:08:13 PM
	candle := &market.Candle{Time: 1629047293}

	// Case 1: Is within the time window
	ok, err := WithinHours(candle, "8:30", "20:30")
	assert.Nil(t, err)
	assert.True(t, ok)

	// Case 2: Is within the time window
	ok, err = WithinHours(candle, "12:08", "12:09")
	assert.Nil(t, err)
	assert.True(t, ok)

	// Case 3: Is NOT within the time window
	ok, err = WithinHours(candle, "5:08", "11:09")
	assert.Nil(t, err)
	assert.False(t, ok)

	// Case 4: Is NOT within the time window
	ok, err = WithinHours(candle, "12:01", "12:07")
	assert.Nil(t, err)
	assert.False(t, ok)

	// Case 5: Is NOT within the time window
	ok, err = WithinHours(candle, "12:09", "20:10")
	assert.Nil(t, err)
	assert.False(t, ok)

	// Case 6: Is within the time window
	ok, err = WithinHours(candle, "00:00", "23:59")
	assert.Nil(t, err)
	assert.True(t, ok)

}
