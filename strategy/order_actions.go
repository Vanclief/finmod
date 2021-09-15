package strategy

import (
	"github.com/rs/xid"
	"github.com/vanclief/finmod/market"
)

type CreateOrderAction struct {
	ID         string            `json:"id"`
	Symbol     string            `json:"symbol"`
	ActionType market.ActionType `json:"action_type"`
	OrderType  market.OrderType  `json:"order_type"`
	Price      float64           `json:"price"`
	Volume     float64           `json:"volume"`
	StopLoss   float64           `json:"stop_loss"`
	TakeProfit float64           `json:"take_profit"`
	Slippage   float64           `json:"slippage"`
}

func (action *CreateOrderAction) GenerateID() {
	action.ID = xid.New().String()
}

type UpdateOrderAction struct {
	ID         string  `json:"id"`
	OrderID    string  `json:"order_id"`
	Price      float64 `json:"price"`
	Volume     float64 `json:"volume"`
	StopLoss   float64 `json:"stop_loss"`
	TakeProfit float64 `json:"take_profit"`
	Slippage   float64 `json:"slippage"`
}

func (action *UpdateOrderAction) GenerateID() {
	action.ID = xid.New().String()
}

type CancelOrderAction struct {
	ID      string `json:"id"`
	OrderID string `json:"order_id"`
}

func (action *CancelOrderAction) GenerateID() {
	action.ID = xid.New().String()
}
