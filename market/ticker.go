package market

// Ticker - Latest price data for an asset
type Ticker struct {
	Time   int64
	Candle *Candle
	Ask    *OrderBookRow
	Bid    *OrderBookRow
}
