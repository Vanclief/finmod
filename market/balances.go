package market

import "fmt"

// BalanceSnapshot - An snapshot of the current balance of a user
type BalanceSnapshot struct {
	Balance    float64 `json:"balance"`
	Equity     float64 `json:"equity"`
	Margin     float64 `json:"margin"`
	FreeMargin float64 `json:"free_margin"`
}

func (b *BalanceSnapshot) String() string {
	return fmt.Sprintf("Balance: %f, Equity: %f, Margin: %f, Free Margin: %f\n", b.Balance, b.Equity, b.Margin, b.FreeMargin)
}

// AssetsSnapshot - An snapshot of the total amount of assets
type AssetsSnapshot struct {
	Time   float64   `json:"time"`
	Assets []Balance `json:"assets"`
}

func (a *AssetsSnapshot) String() string {
	return fmt.Sprintf("Time: %f, Assets: %v\n", a.Time, a.Assets)
}

// Balance - An asset and its unitary amount
type Balance struct {
	Asset  *Asset  `json:"asset"`
	Amount float64 `json:"amount"`
}

func (b *Balance) String() string {
	return fmt.Sprintf("Asset: %s, Amount: %f\n", b.Asset.String(), b.Amount)
}
