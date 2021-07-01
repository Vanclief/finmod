package market

import (
	"fmt"
	"time"
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
	//  FulfilledStatus - The entirety of the order was filled
	FulfilledStatus OrderStatus = "fulfilled"
	//  PartialFillStatus - A the order has been filled partialy
	PartialFillStatus OrderStatus = "partial_fill"
	//  UnfilledStatus - The order has not been filled
	UnfilledStatus OrderStatus = "unfilled"
	//  CanceledStatus - The order has been cancelled
	CanceledStatus OrderStatus = "canceled"
)

// Order - Set of instructions to purchase or sell an asset
type Order struct {
	ID             string
	Action         ActionType
	Type           OrderType
	Pair           Pair
	Price          float64
	Volume         float64
	ExecutedVolume float64
	Fee            float64
	Cost           float64
	Status         OrderStatus
	OpenTime       time.Time
	CloseTime      time.Time
	Trades         []Trade
}

// OrderRequest - Request to create an order
type OrderRequest struct {
	Action   ActionType
	Type     OrderType
	Pair     Pair
	Price    float64
	Quantity float64
	Cost     float64
}

func (o *Order) String() string {
	return fmt.Sprintf("ID: %s Action: %s Type: %s Volume: %.4f Cost: $%.4f", o.ID, o.Action, o.Type, o.Volume, o.Cost)
}

func (o *OrderRequest) String() string {
	return fmt.Sprintf("Action: %s Type: %s Quantity: %.4f Cost: $%.4f", o.Action, o.Type, o.Quantity, o.Cost)
}
