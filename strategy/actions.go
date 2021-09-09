package strategy

// Actions - The available actions that can be executed by the strategy
type Actions struct {
	CreatePositions map[string]CreatePositionAction
	UpdatePositions map[string]UpdatePositionAction
	ClosePositions  map[string]ClosePositionAction
	CancelActions   []CancelAction
	Annotations     []Annotation
}

func (a *Actions) AddCreatePositions(actions ...CreatePositionAction) {
	if a.CreatePositions == nil {
		a.CreatePositions = map[string]CreatePositionAction{}
	}

	for _, action := range actions {
		action.GenerateID()
		a.CreatePositions[action.ID] = action
	}
}

func (a *Actions) AddUpdatePositions(actions ...UpdatePositionAction) {
	if a.UpdatePositions == nil {
		a.UpdatePositions = map[string]UpdatePositionAction{}
	}

	for _, action := range actions {
		action.GenerateID()
		a.UpdatePositions[action.ID] = action
	}
}

func (a *Actions) AddClosePositions(actions ...ClosePositionAction) {
	if a.ClosePositions == nil {
		a.ClosePositions = map[string]ClosePositionAction{}
	}

	for _, action := range actions {
		action.GenerateID()
		a.ClosePositions[action.ID] = action
	}
}

func (a *Actions) AddCancelActions(actions ...CancelAction) {
	a.CancelActions = append(a.CancelActions, actions...)
}

func (a *Actions) AddAnnotations(action ...Annotation) {
	a.Annotations = append(a.Annotations, action...)
}
