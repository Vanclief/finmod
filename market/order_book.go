package market

// OrderBook - A record of active buy and sell orders in a single market
type OrderBook struct {
	Time       float64
	SellOrders []OrderBookRow `json:"sell_orders"`
	BuyOrders  []OrderBookRow `json:"buy_orders"`
}

// OrderBookRow - A single order from the order book
type OrderBookRow struct {
	Price       float64 `json:"price"`
	Volume      float64 `json:"volume"`
	TotalVolume float64 `json:"total_volume"`
}
