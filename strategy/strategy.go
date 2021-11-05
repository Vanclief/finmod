package strategy

import (
	"time"

	"github.com/vanclief/ez"
	"github.com/vanclief/finmod/market"
)

// Tick - The object that is passed on every tick
type Tick struct {
	Candle          *market.Candle
	PendingActions  *Actions
	OpenOrders      map[string]market.Order
	OpenPositions   map[string]market.Position
	ClosedPositions map[string]market.Position
	Capital         float64
}

// GetLocalTime - Returns the time of the candle in the local timezone
func (t *Tick) GetLocalTime() time.Time {
	return time.Unix(t.Candle.Time, 0)
}

// GetLocationTime - Returns the time of the candle for a certain location/timezone
func (t *Tick) GetLocationTime(location string) (time.Time, error) {
	const op = "Tick.GetLocationTime"

	loc, err := time.LoadLocation(location)
	if err != nil {
		return time.Time{}, ez.Wrap(op, err)
	}

	candleTime := time.Unix(t.Candle.Time, 0)

	return candleTime.In(loc), nil
}

// Strategy is the interface of a strategy implementation
type Strategy interface {
	GetName() string
	GetVersion() string
	PreloadTick(tick *Tick) error
	Tick(tick *Tick) (*Actions, error)
}
