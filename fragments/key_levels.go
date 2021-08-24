package fragments

import (
	"sort"
	"time"

	"github.com/vanclief/ez"
	"github.com/vanclief/finmod/market"
)

type KeyLevel struct {
	CreatedAt time.Time
	Value     float64
	Hits      int
}

func FindKeyLevels(candles []market.Candle) ([]KeyLevel, error) {
	const op = "fragments.FindKeyLevels"

	if len(candles) < 5 {
		return nil, ez.New(op, ez.EINVALID, "Must pass at least 5 candles", nil)
	}

	// Step 1: Find all key levels and sort them by value
	keyLevels, avg := FindInflectionPoints(candles)

	sort.Slice(keyLevels, func(i, j int) bool {
		return keyLevels[i].Value < keyLevels[j].Value
	})

	// Step 2: Group the key levels
	var groupedLevels [][]KeyLevel

	for i, k := range keyLevels {

		lastGroup := len(groupedLevels) - 1

		if i == 0 || k.Value-groupedLevels[lastGroup][0].Value > avg {
			group := []KeyLevel{k}
			groupedLevels = append(groupedLevels, group)
		} else {
			group := groupedLevels[lastGroup]
			group = append(group, k)
			groupedLevels[len(groupedLevels)-1] = group
		}
	}

	// Step 3: Reduce the groups to single key levels and sort them by time
	var reducedLevels []KeyLevel

	for i := range groupedLevels {

		if len(groupedLevels[i]) == 1 {
			reducedLevels = append(reducedLevels, groupedLevels[i][0])
		} else {
			last := len(groupedLevels[i]) - 1

			k := KeyLevel{
				CreatedAt: groupedLevels[i][0].CreatedAt,
				Value:     (groupedLevels[i][0].Value + groupedLevels[i][last].Value) / 2,
				Hits:      len(groupedLevels[i]),
			}

			reducedLevels = append(reducedLevels, k)
		}
	}

	sort.Slice(reducedLevels, func(i, j int) bool {
		return reducedLevels[i].CreatedAt.After(reducedLevels[j].CreatedAt)
	})

	return reducedLevels, nil
}

func FindInflectionPoints(candles []market.Candle) ([]KeyLevel, float64) {

	var keyLevels []KeyLevel
	var sum, avg float64

	for i := range candles {

		if i < 2 || i > len(candles)-3 {
			continue
		}

		if candles[i-2].High < candles[i-1].High && candles[i-1].High < candles[i].High {
			if candles[i].High > candles[i+1].High && candles[i+1].High > candles[i+2].High {
				createdAt := time.Unix(candles[i].Time, 0)
				keyLevel := KeyLevel{CreatedAt: createdAt, Value: candles[i].High}
				keyLevels = append(keyLevels, keyLevel)
			}
		}

		if candles[i-2].Low > candles[i-1].Low && candles[i-1].Low > candles[i].Low {
			if candles[i].Low < candles[i+1].Low && candles[i+1].Low < candles[i+2].Low {
				createdAt := time.Unix(candles[i].Time, 0)
				keyLevel := KeyLevel{CreatedAt: createdAt, Value: candles[i].Low}
				keyLevels = append(keyLevels, keyLevel)
			}
		}

		sum = sum + candles[i].High - candles[i].Low
		avg = sum / float64(i)
	}

	return keyLevels, avg
}
