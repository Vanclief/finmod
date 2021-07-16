package market

import (
	"fmt"

	"github.com/vanclief/ez"
)

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

	var pairStr string
	if o.Pair != nil {
		pairStr = o.Pair.String()
	}

	return fmt.Sprintf(
		"Action: %s | Type: %s | Pair: %s | Quantity: %.4f | Price: $%.4f | Total: $%.4f",
		o.Action,
		o.Type,
		pairStr,
		o.Quantity,
		o.Price,
		o.Total,
	)
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
