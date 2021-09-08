package strategy

import "github.com/rs/xid"

type ClosePositionAction struct {
	ID         string `json:"id"`
	PositionID string `json:"position_id"`
}

func (action *ClosePositionAction) GenerateID() {
	action.ID = xid.New().String()
}
