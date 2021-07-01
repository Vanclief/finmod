package market

import (
	"fmt"
	"time"
)

// Trade represents a market trade
type Trade struct {
	ID            string
	ExecutionDate time.Time
	Action        ActionType
	OrderType     OrderType
	Pair          *Pair
	Price         float64
	Quantity      float64
	Fee           float64
	Cost          float64
}

// NewTrade creates a new trade
func NewTrade(id string, executionDate time.Time, actionType ActionType, orderType OrderType, pair *Pair, price, quantity, cost float64) *Trade {

	var fee float64

	if actionType == BuyAction {
		fee = cost - (price * quantity)
	} else {
		fee = (price * quantity) - cost
	}

	trade := &Trade{
		ID:            id,
		ExecutionDate: executionDate,
		Action:        actionType,
		OrderType:     orderType,
		Pair:          pair,
		Price:         price,
		Quantity:      quantity,
		Fee:           fee,
		Cost:          cost,
	}

	return trade
}

func (t *Trade) String() string {
	return fmt.Sprintf("ID: %s Date: %s Action: %s Type: %s Pair: %s Price: $%.2f Quantity: %.4f Fee: $%.2f Cost: $%.2f", t.ID, t.ExecutionDate, t.Action, t.OrderType, t.Pair, t.Price, t.Quantity, t.Fee, t.Cost)
}
