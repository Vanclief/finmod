package strategy

import (
	"fmt"
	"github.com/rs/xid"
	"github.com/vanclief/ez"
	"github.com/vanclief/state/interfaces"
)

type AddAnnotationsAction struct {
	Annotations []Annotation
}

func (a *AddAnnotationsAction) String() string {
	str := "Annotations:\n"
	for _, v := range a.Annotations {
		str += fmt.Sprintf("%s\n", v.String())
	}
	return str
}

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

func (a *Annotation) String() string {
	return fmt.Sprintf("ID: %s, Time: %d, Creation Time: %d, Price: %f, ExecutionID: %s, Label: %s, Tooltip: %s\n",
		a.ID,
		a.Time,
		a.CreationTime,
		a.Price,
		a.ExecutionID,
		a.Label,
		a.Tooltip,
	)
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

func NewAnnotation(t int64, price float64, label, tooltip string) *Annotation {
	return &Annotation{
		ID:      xid.New().String(),
		Time:    t,
		Price:   price,
		Label:   label,
		Tooltip: tooltip,
	}
}
