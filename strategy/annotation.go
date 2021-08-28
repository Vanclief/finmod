package strategy

import (
	"fmt"
	"time"

	"github.com/vanclief/ez"
	"github.com/vanclief/go-crypto/sha"
	"github.com/vanclief/go-crypto/utils"
	"github.com/vanclief/state/interfaces"
)

// Annotation - A note of explanation or comment added to the graph
type Annotation struct {
	ID    string    `json:"id"`
	Time  time.Time `json:"time"`
	Price float64   `json:"price"`
	Tag   string    `json:"tag"`
	Type  string    `json:"type"`
	Note  string    `json:"note"`
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
func NewAnnotation(t time.Time, price float64, kind, note string) *Annotation {

	msg := fmt.Sprintf("%s, %f, %s, %s", t, price, kind, note)
	hash := sha.RandomSHA256(msg)

	annotation := &Annotation{
		ID:    utils.BytesToHex(hash.Value),
		Time:  t,
		Price: price,
		Type:  kind,
		Note:  note,
	}

	return annotation
}
