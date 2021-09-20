package indicators

import (
	"github.com/vanclief/finmod/market"
)

func ExponentialWeightedVolatility(candles []market.Candle, beta float64) float64 {

	var v float64

	for _, candle := range candles {
		delta := candle.High - candle.Low
		v = beta*v + (1-beta)*delta
	}

	return v
}
