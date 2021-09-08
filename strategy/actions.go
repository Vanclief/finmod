package strategy

// Actions - The available actions that can be executed by the strategy
type Actions struct {
	CreatePositions []CreatePositionAction
	UpdatePositions []UpdatePositionAction
	ClosePositions  []ClosePositionAction
	CancelActions   []CancelAction
	Annotations     []Annotation
}

func (a *Actions) AddCreatePosition(action ...CreatePositionAction) {
	a.CreatePositions = append(a.CreatePositions, action...)
}

func (a *Actions) AddUpdatePosition(action ...UpdatePositionAction) {
	a.UpdatePositions = append(a.UpdatePositions, action...)
}

func (a *Actions) AddClosePositions(action ...ClosePositionAction) {
	a.ClosePositions = append(a.ClosePositions, action...)
}

func (a *Actions) AddCancelAction(action ...CancelAction) {
	a.CancelActions = append(a.CancelActions, action...)
}

func (a *Actions) AddAnnotations(action ...Annotation) {
	a.Annotations = append(a.Annotations, action...)
}
