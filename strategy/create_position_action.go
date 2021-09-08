package strategy

import (
	"github.com/rs/xid"
	"github.com/vanclief/finmod/market"
)

type CreatePositionAction struct {
	ID         string            `json:"id"`
	Symbol     string            `json:"symbol"`
	ActionType market.ActionType `json:"action_type"`
	OrderType  market.OrderType  `json:"order_type"`
	Volume     float64           `json:"volume"`
	StopLoss   float64           `json:"stop_loss"`
	TakeProfit float64           `json:"take_profit"`
	Slippage   float64           `json:"slippage"`
}

func (action *CreatePositionAction) GenerateID() {
	action.ID = xid.New().String()
}
