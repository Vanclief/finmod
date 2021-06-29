package market

import (
	"math"
)

// OrderBook - A record of active buy and sell orders in a single market
type OrderBook struct {
	Time float64
	Asks []OrderBookRow `json:"asks"` // ordered from lowest to highest
	Bids []OrderBookRow `json:"bids"` // ordered from highest to lowest
}

// OrderBookRow - A single order from the order book
type OrderBookRow struct {
	Price       float64 `json:"price"`
	Volume      float64 `json:"volume"`
	TotalVolume float64 `json:"total_volume"` // maybe should be called AccumulatedVolume
}

// ReturnVolumeOfOrderBook gives the accumulated volume from a price onwards,
// it return -1 if it does not find an index for the selected price
func (ob *OrderBook) ReturnVolumeOfOrderBook(price float64) float64 {
	firstAsk, lastAsk := ob.Asks[0], ob.Asks[len(ob.Asks) - 1]
	firstBid, lastBid := ob.Bids[0], ob.Bids[len(ob.Bids) - 1]
	if firstAsk.Price <= price && price <= lastAsk.Price {
		// it belongs to asks
		for k := range ob.Asks {
			if ob.Asks[k].Price > price {
				index := math.Max(0, float64(k - 1))
				return ob.Asks[int(index)].Price
			}
		}
	} else if  lastBid.Price <= price && price <= firstBid.Price {
		// it belongs to bids
		for k := range ob.Bids {
			if ob.Bids[k].Price < price {
				index := math.Max(0, float64(k - 1))
				return ob.Bids[int(index)].Price
			}
		}
	}
	return -1
}