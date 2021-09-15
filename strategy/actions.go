package strategy

// Actions - The available actions that can be executed by the strategy
type Actions struct {
	CreateOrders    map[string]CreateOrderAction
	UpdateOrders    map[string]UpdateOrderAction
	CancelOrders    map[string]CancelOrderAction
	UpdatePositions map[string]UpdatePositionAction
	ClosePositions  map[string]ClosePositionAction
	Annotations     []Annotation
}

func (a *Actions) AddCreateOrders(actions ...CreateOrderAction) {
	if a.CreateOrders == nil {
		a.CreateOrders = map[string]CreateOrderAction{}
	}

	for _, action := range actions {
		action.GenerateID()
		a.CreateOrders[action.ID] = action
	}
}

func (a *Actions) AddUpdateOrders(actions ...UpdateOrderAction) {
	if a.UpdateOrders == nil {
		a.UpdateOrders = map[string]UpdateOrderAction{}
	}

	for _, action := range actions {
		action.GenerateID()
		a.UpdateOrders[action.ID] = action
	}
}

func (a *Actions) AddCancelOrders(actions ...CancelOrderAction) {
	if a.CancelOrders == nil {
		a.CancelOrders = map[string]CancelOrderAction{}
	}

	for _, action := range actions {
		action.GenerateID()
		a.CancelOrders[action.ID] = action
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

func (a *Actions) AddAnnotations(action ...Annotation) {
	a.Annotations = append(a.Annotations, action...)
}
