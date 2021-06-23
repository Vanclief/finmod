package market

// OrderBook - A record of active buy and sell orders in a single market
type OrderBook struct {
	Time float64
	Asks []OrderBookRow `json:"asks"`
	Bids []OrderBookRow `json:"bids"`
}

// OrderBookRow - A single order from the order book
type OrderBookRow struct {
	Price       float64 `json:"price"`
	Volume      float64 `json:"volume"`
	TotalVolume float64 `json:"total_volume"`
}
