package indicators

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vanclief/finmod/market"
)

func TestExponentialWeightedVolatility(t *testing.T) {

	candles, err := market.LoadCandlesFromFile("./test_dataset/candles_1_h")
	assert.Nil(t, err)

	ewv := ExponentialWeightedVolatility(candles, 0.5)

	assert.Equal(t, float64(28.882156906724013), ewv)
}
