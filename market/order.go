package market

import (
	"fmt"

	"github.com/rs/xid"
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

// Order is a transaction instruction to be executed
type Order struct {
	ID       xid.ID
	Action   ActionType
	Type     OrderType
	Quantity float64
	Total    float64
}

// NewOrder creates and returns a new Order
func NewOrder(action ActionType, orderType OrderType, quantity, total float64) *Order {

	order := &Order{
		ID:       xid.New(),
		Action:   action,
		Type:     orderType,
		Quantity: quantity,
		Total:    total,
	}

	return order
}

func (o *Order) String() string {
	return fmt.Sprintf("ID: %s Action: %s Type: %s Quantity: %.4f Total: $%.4f", o.ID.String()[0:8], o.Action, o.Type, o.Quantity, o.Total)
}
