package market

import (
	"fmt"
	"time"

	"github.com/vanclief/ez"
	"github.com/vanclief/state/interfaces"
)

// Trade represents a market trade
type Trade struct {
	ID            string     `json:"id"`
	ExecutionTime time.Time  `json:"execution_time"`
	Action        ActionType `json:"action"`
	OrderType     OrderType  `json:"order_type"`
	Pair          Pair       `json:"pair"`
	Price         float64    `json:"price"`
	Quantity      float64    `json:"quantity"`
	Fee           float64    `json:"fee"`
	Cost          float64    `json:"cost"`
	ExecutionID   string     `json:"execution_id"`
}

// NewTrade creates a new trade
func NewTrade(id string, executionTime time.Time, actionType ActionType, orderType OrderType, pair Pair, price, quantity, cost float64) *Trade {

	var fee float64

	if actionType == BuyAction {
		fee = cost - (price * quantity)
	} else {
		fee = (price * quantity) - cost
	}

	trade := &Trade{
		ID:            id,
		ExecutionTime: executionTime,
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

// GetSchema returns the database schema for the Trade model
func (t *Trade) GetSchema() *interfaces.Schema {
	return &interfaces.Schema{Name: "trades", PKey: "id"}
}

// GetID returns the ID from the Trade model
func (t *Trade) GetID() string {
	return t.ID
}

// Update sets the value of the Trade instance
func (t *Trade) Update(i interface{}) error {
	const op = "Trade.Update"

	trade, ok := i.(*Trade)
	if !ok {
		return ez.New(op, ez.EINVALID, "Provided interface is not of type Trade", nil)
	}

	*t = *trade

	return nil
}

func (t *Trade) String() string {
	return fmt.Sprintf("ID: %s Execution Time: %s Action: %s Type: %s Pair: %s Price: $%.2f Quantity: %.4f Fee: $%.2f Cost: $%.2f\n",
		t.ID,
		t.ExecutionTime,
		t.Action,
		t.OrderType,
		t.Pair,
		t.Price,
		t.Quantity,
		t.Fee,
		t.Cost,
	)
}
