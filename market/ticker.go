package market

import "fmt"

// Ticker - The Latest price data for an asset
type Ticker struct {
	Time   int64   `json:"time"`
	Ask    float64 `json:"ask"`
	Bid    float64 `json:"bid"`
	Last   float64 `json:"last"`
	Volume float64 `json:"volume"`
	VWAP   float64 `json:"vwap"`
}

func (t *Ticker) String() string {
	return fmt.Sprintf("Time: %d, Ask: %f, Bid: %f, Last: %f, Volume: %f, VWAP: %f\n", t.Time, t.Ask, t.Bid, t.Last, t.Volume, t.VWAP)
}
