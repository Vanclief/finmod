package market

// BalanceSnapshot - An snapshot of the total amount of assets
type BalanceSnapshot struct {
	Time     float64   `json:"time"`
	Balances []Balance `json:"balances"`
}

// Balance - An asset and its unitary amount
type Balance struct {
	Asset  *Asset  `json:"asset"`
	Amount float64 `json:"amount"`
}
