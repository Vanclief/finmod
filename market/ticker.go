package market

// Ticker - Latest price data for an asset
type Ticker struct {
	Time   int64   `json:"time"`
	Ask    float64 `json:"ask"`
	Bid    float64 `json:"bid"`
	Last   float64 `json:"last"`
	Volume float64 `json:"volume"`
	VWAP   float64 `json:"vwap"`
}
