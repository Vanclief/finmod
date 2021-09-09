package strategy

import "github.com/vanclief/finmod/market"

// Tick - The object that is passed on every tick
type Tick struct {
	Candle          *market.Candle
	PendingActions  *Actions
	OpenPositions   map[string]market.Position // Current Open Positions
	ClosedPositions map[string]market.Position // Recently Closed Positions
	Capital         float64
}

// Strategy is the interface of a strategy implementation
type Strategy interface {
	GetName() string
	GetVersion() string
	PreloadTick(tick *Tick) error
	Tick(tick *Tick) (*Actions, error)
}
