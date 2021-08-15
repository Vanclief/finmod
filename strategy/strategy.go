package strategy

import "github.com/vanclief/finmod/market"

// Tick - The object that is passed on every tick
type Tick struct {
	Candle       *market.Candle
	Orders       []market.Order
	OpenPosition *market.Position
	LastPosition *market.Position
	Assets       float64
	Capital      float64
}

// Actions - The available actions that can be executed by the strategy
type Actions struct {
	OpenOrders  []market.OrderRequest
	CloseOrders []market.Order
}

// Strategy is the interface of a strategy implementation
type Strategy interface {
	GetName() string
	GetVersion() string
	Tick(tick *Tick) (*Actions, error)
}
