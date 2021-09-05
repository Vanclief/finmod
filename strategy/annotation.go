package strategy

import (
	"github.com/rs/xid"
	"github.com/vanclief/ez"
	"github.com/vanclief/state/interfaces"
)

// Annotation - A note of explanation or comment added to the graph
type Annotation struct {
	ID           string  `json:"id"`
	Time         int64   `json:"time"`
	CreationTime int64   `json:"creation_time"`
	Price        float64 `json:"price"`
	ExecutionID  string  `json:"execution_id"`
	Label        string  `json:"label"`
	Tooltip      string  `json:"tooltip"`
}

// GetSchema returns the database schema for the Annotation model
func (a *Annotation) GetSchema() *interfaces.Schema {
	return &interfaces.Schema{Name: "annotations", PKey: "id"}
}

// GetID returns the ID from the Annotation model
func (a *Annotation) GetID() string {
	return a.ID
}

// Update sets the value of the Annotation instance
func (a *Annotation) Update(i interface{}) error {
	const op = "Annotation.Update"

	model, ok := i.(*Annotation)
	if !ok {
		return ez.New(op, ez.EINVALID, "Provided interface is not of type Annotation", nil)
	}

	*a = *model

	return nil
}

// NewAnnotation
func NewAnnotation(t int64, price float64, label, tooltip string) *Annotation {

	annotation := &Annotation{
		ID:      xid.New().String(),
		Time:    t,
		Price:   price,
		Label:   label,
		Tooltip: tooltip,
	}

	return annotation
}
