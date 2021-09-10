package market

// BalanceSnapshot - An snapshot of the current balance of a user
type BalanceSnapshot struct {
	Balance    float64 `json:"balance"`
	Equity     float64 `json:"equity"`
	Margin     float64 `json:"margin"`
	FreeMargin float64 `json:"free_margin"`
}

// AssetsSnapshot - An snapshot of the total amount of assets
type AssetsSnashot struct {
	Time   float64   `json:"time"`
	Assets []Balance `json:"assets"`
}

// Balance - An asset and its unitary amount
type Balance struct {
	Asset  *Asset  `json:"asset"`
	Amount float64 `json:"amount"`
}
