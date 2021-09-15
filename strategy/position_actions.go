package strategy

import (
	"github.com/rs/xid"
	"github.com/vanclief/finmod/market"
)

type UpdatePositionAction struct {
	ID         string            `json:"id"`
	PositionID string            `json:"position_id"`
	ActionType market.ActionType `json:"action_type"`
	OrderType  market.OrderType  `json:"order_type"`
	Volume     float64           `json:"volume"`
	StopLoss   float64           `json:"stop_loss"`
	TakeProfit float64           `json:"take_profit"`
	Slippage   float64           `json:"slippage"`
}

func (action *UpdatePositionAction) GenerateID() {
	action.ID = xid.New().String()
}

type ClosePositionAction struct {
	ID         string `json:"id"`
	PositionID string `json:"position_id"`
}

func (action *ClosePositionAction) GenerateID() {
	action.ID = xid.New().String()
}
