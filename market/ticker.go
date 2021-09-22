package market

// Ticker - Latest price data for an asset
type Ticker struct {
	Time   int64
	Ask    float64
	Bid    float64
	Last   float64
	Volume float64
	Side   string // buy or sell
}
