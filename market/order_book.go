package market

import (
	"fmt"
	"math"

	"github.com/vanclief/ez"
)

// OrderBook - A record of active buy and sell orders in a single market
type OrderBook struct {
	Time int64          `json:"time"`
	Asks []OrderBookRow `json:"asks"` // ordered from lowest to highest
	Bids []OrderBookRow `json:"bids"` // ordered from highest to lowest
}

func (ob *OrderBook) String() string {
	return fmt.Sprintf("Time: %v, Asks: %v, Bids: %v\n", ob.Time, ob.Asks, ob.Bids)
}

// OrderBookRow - A single order from the order book
type OrderBookRow struct {
	Price       float64 `json:"price"`
	Volume      float64 `json:"volume"`
	AccumVolume float64 `json:"accum_volume"`
}

func (obr *OrderBookRow) String() string {
	return fmt.Sprintf("Price: %v, Volume: %v, AccumVolume: %v\n", obr.Price, obr.Volume, obr.AccumVolume)
}

// Display - returns a
func (ob *OrderBook) Display(side string, rowCount int) []OrderBookRow {

	var rows []OrderBookRow

	if side == "asks" {
		for j, row := range ob.Asks {

			if j >= rowCount {
				break
			}

			rows = append(rows, row)
		}
	} else {
		for j, row := range ob.Bids {

			if j >= rowCount {
				break
			}

			rows = append(rows, row)
		}
	}

	return rows
}

// GetDepth - returns the accumulated volume from a determined price
// a price onwards
func (ob *OrderBook) GetDepth(price float64) (float64, error) {
	op := "OrderBook.GetDepth"

	if price <= 0 {
		return 0, ez.New(op, ez.EINVALID, "Price can't be equal or below zero", nil)
	}

	firstAsk, lastAsk := ob.Asks[0], ob.Asks[len(ob.Asks)-1]
	firstBid, lastBid := ob.Bids[0], ob.Bids[len(ob.Bids)-1]
	if firstAsk.Price <= price && price <= lastAsk.Price {
		// price is in asks interval
		for k := range ob.Asks {
			if ob.Asks[k].Price > price {
				index := math.Max(0, float64(k-1))
				return ob.Asks[int(index)].Price, nil
			}
		}
		return lastAsk.Price, nil
	} else if lastBid.Price <= price && price <= firstBid.Price {
		// price is in bids interval
		for k := range ob.Bids {
			if ob.Bids[k].Price < price {
				index := math.Max(0, float64(k-1))
				return ob.Bids[int(index)].Price, nil
			}
		}
		return lastBid.Price, nil
	}
	return 0, ez.New(op, ez.ENOTFOUND, "No depth for the selected price", nil)
}
