package market

import (
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/vanclief/ez"
)

// OrderBook - A record of active buy and sell orders in a single market
type OrderBook struct {
	Time     int64          `json:"time"`
	Asks     []OrderBookRow `json:"asks"` // ordered from lowest to highest
	Bids     []OrderBookRow `json:"bids"` // ordered from highest to lowest
	MaxDepth int            `json:"max_depth"`
}

func NewOrderBook(asks, bids []OrderBookRow, maxDepth int) OrderBook {

	ob := OrderBook{
		Time:     time.Now().Unix(),
		Asks:     asks,
		Bids:     bids,
		MaxDepth: maxDepth,
	}

	ob.sort()

	return ob
}

func (ob *OrderBook) sort() {
	sort.SliceStable(ob.Asks, func(i, j int) bool {
		return ob.Asks[i].Price < ob.Asks[j].Price
	})

	sort.SliceStable(ob.Bids, func(i, j int) bool {
		return ob.Bids[i].Price > ob.Bids[j].Price
	})

	askDepth := math.Min(float64(len(ob.Asks)), float64(ob.MaxDepth))
	bidDepth := math.Min(float64(len(ob.Bids)), float64(ob.MaxDepth))

	ob.Asks = ob.Asks[:int(askDepth)]
	ob.Bids = ob.Bids[:int(bidDepth)]

	for i := range ob.Asks {
		if i == 0 {
			ob.Asks[i].AccumVolume = ob.Asks[i].Volume
		} else {
			ob.Asks[i].AccumVolume = ob.Asks[i].Volume + ob.Asks[i-1].AccumVolume
		}
	}

	for i := range ob.Bids {
		if i == 0 {
			ob.Bids[i].AccumVolume = ob.Bids[i].Volume
		} else {
			ob.Bids[i].AccumVolume = ob.Bids[i].Volume + ob.Bids[i-1].AccumVolume
		}
	}
}

func (ob *OrderBook) String() string {
	return fmt.Sprintf("Time: %v, Asks: %v, Bids: %v\n", ob.Time, ob.Asks, ob.Bids)
}

func (ob *OrderBook) Print() {
	fmt.Println("======= OrderBook =======")
	for i := range ob.Asks {
		i = len(ob.Asks) - 1 - i
		fmt.Println(ob.Asks[i])
	}
	fmt.Println("-------------------------")
	for _, bid := range ob.Bids {
		fmt.Println(bid)
	}
}

func (ob *OrderBook) ApplyUpdate(update OrderBookUpdate) error {
	const op = "OrderBook.ApplyUpdate"

	if update.Side == "ask" {
		if update.Volume != 0 {
			found := false
			for i := range ob.Asks {
				if ob.Asks[i].Price == update.Price {
					ob.Asks[i] = OrderBookRow{Price: update.Price, Volume: update.Volume}
					found = true
					break
				}
			}
			if !found {
				if len(ob.Asks) == 0 {
					ob.Asks = append(ob.Asks, OrderBookRow{Price: update.Price, Volume: update.Volume})
					return nil
				}
				for i := range ob.Asks {
					if ob.Asks[i].Price > update.Price {
						ob.Asks = append(ob.Asks, OrderBookRow{})
						copy(ob.Asks[i+1:], ob.Asks[i:])
						ob.Asks[i] = OrderBookRow{Price: update.Price, Volume: update.Volume}
						break
					}
				}
			}
		} else {
			for i := range ob.Asks {
				if ob.Asks[i].Price == update.Price {
					ob.Asks = removeElement(ob.Asks, i)
					break
				}
			}
		}

	} else if update.Side == "bid" {
		if update.Volume != 0 {
			found := false
			for i := range ob.Bids {
				if ob.Bids[i].Price == update.Price {
					ob.Bids[i] = OrderBookRow{Price: update.Price, Volume: update.Volume}
					found = true
					break
				}
			}
			if !found {
				if len(ob.Bids) == 0 {
					ob.Bids = append(ob.Bids, OrderBookRow{Price: update.Price, Volume: update.Volume})
					return nil
				}
				for i := range ob.Bids {
					if ob.Bids[i].Price < update.Price {
						ob.Bids = append(ob.Bids, OrderBookRow{})
						copy(ob.Bids[i+1:], ob.Bids[i:])
						ob.Bids[i] = OrderBookRow{Price: update.Price, Volume: update.Volume}
						break
					}
				}
			}
		} else {
			for i := range ob.Bids {
				if ob.Bids[i].Price == update.Price {
					ob.Bids = removeElement(ob.Bids, i)
					break
				}
			}
		}
	} else {
		return ez.New(op, ez.EINVALID, "update side must be ask or bid", nil)
	}

	return nil
}

func removeElement(slice []OrderBookRow, s int) []OrderBookRow {
	return append(slice[:s], slice[s+1:]...)
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
