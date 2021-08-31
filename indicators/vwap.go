package indicators

import (
	"github.com/vanclief/ez"
	"github.com/vanclief/finmod/market"
)

func weightedSum(candles []market.Candle, volumes []float32) float32 {
	num := 0.0
	den := 0.0
	for k, v := range candles {
		num += (v.High + v.Close + v.Low)/3 * float64(volumes[k])
		den += float64(volumes[k])
	}
	if den == 0 {
		den = 1
	}
	return float32(num / den)
}

func VolumeWeightedAveragePrice(candles []market.Candle, volume []float32, length int) ([]float32, error) {
	op := "vwap"

	if candles == nil {
		return nil, ez.New(op, ez.EINVALID, "Candle array missing", nil)
	} else if len(candles) < length {
		return nil, ez.New(op, ez.EINVALID, "Length argument is larger than the length of candles", nil)
	} else if length <= 0 {
		return nil, ez.New(op, ez.EINVALID, "Length can't be less than 1", nil)
	}

	var vwap []float32

	i := 0
	j := length - 1

	for {
		if j == len(candles) {
			break
		}

		vwap = append(vwap, weightedSum(candles[i:j+1], volume[i:j+1]))
		i++
		j++
	}

	return vwap, nil
}
