package market

import (
	"fmt"
	"time"

	"github.com/vanclief/ez"
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
	Trades         []Trade
}

func (o *Order) String() string {
	return fmt.Sprintf("ID: %s Action: %s Type: %s Volume: %.4f Cost: $%.4f", o.ID, o.Action, o.Type, o.Volume, o.Cost)
}

// OrderRequest - Request to create an order
type OrderRequest struct {
	Action   ActionType
	Type     OrderType
	Pair     *Pair
	Price    float64
	Quantity float64
	Total    float64
}

// NewOrderRequest - creates a valid new OrderRequest
func NewOrderRequest(pair *Pair, action ActionType, orderType OrderType, quantity, price, total float64) (*OrderRequest, error) {
	const op = "market.NewOrderRequest"

	if pair == nil {
		return nil, ez.New(op, ez.EINVALID, "Order must have a pair defined", nil)
	}

	if action == "" || orderType == "" {
		return nil, ez.New(op, ez.EINVALID, "Order must have an action and a type defined", nil)
	}

	if orderType == LimitOrder && price == 0 {
		return nil, ez.New(op, ez.EINVALID, "LimitOrders must have a priced defined", nil)
	}

	if quantity == 0 && total == 0 {
		return nil, ez.New(op, ez.EINVALID, "Total and quantity can't be both zero", nil)
	}

	request := &OrderRequest{
		Action:   action,
		Type:     orderType,
		Pair:     pair,
		Price:    price,
		Quantity: quantity,
		Total:    total,
	}

	return request, nil
}

func (o *OrderRequest) String() string {
	return fmt.Sprintf("Action: %s | Type: %s | Pair: %s | Quantity: %.4f | Price: $%.4f | Total: $%.4f", o.Action, o.Type, o.Pair.String(), o.Quantity, o.Price, o.Total)
}

func (o *OrderRequest) CalculateFields(currentPrice float64) (*OrderRequest, error) {
	const op = "OrderRequest.CalculateFields"

	// Either the total cost or the quantity must be defined
	if o.Total == 0 && o.Quantity == 0 {
		return nil, ez.New(op, ez.EINVALID, "Total and quantity can't be both zero", nil)
	} else if o.Total != 0 && o.Quantity != 0 {
		return nil, ez.New(op, ez.EINVALID, "Either total or quantity must be zero", nil)
	}

	copy := *o

	// MarketOrder
	if o.Type == MarketOrder {
		if o.Quantity != 0 {
			copy.Total = o.Quantity * currentPrice
		} else {
			copy.Quantity = o.Total / currentPrice
		}
	} else {
		if o.Quantity != 0 {
			copy.Total = o.Quantity * o.Price
		} else {
			copy.Quantity = o.Total / o.Price
		}
	}

	return &copy, nil

}
