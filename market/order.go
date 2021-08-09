package market

import (
	"fmt"
	"time"

	"github.com/vanclief/ez"
	"github.com/vanclief/state/interfaces"
)

// ActionType determines if the action is to buy or to sell
type ActionType string

const (
	// BuyAction is a purchase action
	BuyAction ActionType = "buy"
	// SellAction is a selling action
	SellAction ActionType = "sell"
)

// OrderType determines the type of market order
type OrderType string

const (
	// MarketOrder will attempt to buy or sell at whatever the current price is
	MarketOrder OrderType = "market"
	// LimitOrder will attempt to buy or sell at a specific price
	LimitOrder OrderType = "limit"
)

// OrderStatus - determines the status of the order
type OrderStatus string

const (
	//  FulfilledOrder - The entirety of the order was filled
	FulfilledOrder OrderStatus = "fulfilled"
	//  PartialyFilledOrder - A the order has been filled partialy
	PartialyFilledOrder OrderStatus = "partial_fill"
	//  UnfilledOrder - The order has not been filled
	UnfilledOrder OrderStatus = "unfilled"
	//  CanceledOrder - The order has been cancelled
	CanceledOrder OrderStatus = "canceled"
)

// Order - Set of instructions to purchase or sell an asset
type Order struct {
	ID             string
	Action         ActionType
	Type           OrderType
	Pair           *Pair
	Price          float64
	Volume         float64
	ExecutedVolume float64
	Fee            float64
	Cost           float64
	Status         OrderStatus
	OpenTime       time.Time
	CloseTime      time.Time
	Trades         []string
}

// GetSchema returns the database schema for the Order model
func (o *Order) GetSchema() *interfaces.Schema {
	return &interfaces.Schema{Name: "orders", PKey: "id"}
}

// GetID returns the ID from the Order model
func (o *Order) GetID() string {
	return o.ID
}

// Update sets the value of the Order instance
func (o *Order) Update(i interface{}) error {
	const op = "Order.Update"

	order, ok := i.(*Order)
	if !ok {
		return ez.New(op, ez.EINVALID, "Provided interface is not of type Order", nil)
	}

	*o = *order

	return nil
}

func (o *Order) String() string {

	var pairStr string
	if o.Pair != nil {
		pairStr = o.Pair.String()
	}

	return fmt.Sprintf(
		"ID: %s | Action: %s | Type: %s | Pair: %s | Price: %.4f | Volume: %.4f | ExecVolume: %.4f | Fee: $%.4f | Cost: $%.4f | Status: %s | OT: %s | CT: %s",
		o.ID,
		o.Action,
		o.Type,
		pairStr,
		o.Price,
		o.Volume,
		o.ExecutedVolume,
		o.Fee,
		o.Cost,
		o.Status,
		o.OpenTime,
		o.CloseTime,
	)
}
