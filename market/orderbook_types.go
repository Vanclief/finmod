package market

import "fmt"

// OrderBookRow - A single order from the order book
type OrderBookRow struct {
	Price       float64 `json:"price"`
	Volume      float64 `json:"volume"`
	AccumVolume float64 `json:"accum_volume"`
}

func (obr *OrderBookRow) String() string {
	return fmt.Sprintf("Price: %v, Volume: %v, AccumVolume: %v\n", obr.Price, obr.Volume, obr.AccumVolume)
}

type OrderBookUpdate struct {
	Price  float64 `json:"price"`
	Volume float64 `json:"volume"`
	Side   string  `json:"side"`
}
