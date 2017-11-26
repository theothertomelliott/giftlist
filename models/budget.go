package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/satori/go.uuid"
)

type Budget struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Maximum   int64     `json:"maximum" db:"maximum"`
	PersonID  uuid.UUID `json:"person_id" db:"person_id"`
	EventID   uuid.UUID `json:"event_id" db:"event_id"`
}

// String is not required by pop and may be deleted
func (b Budget) String() string {
	jb, _ := json.Marshal(b)
	return string(jb)
}

// Budgets is not required by pop and may be deleted
type Budgets []Budget

// String is not required by pop and may be deleted
func (b Budgets) String() string {
	jb, _ := json.Marshal(b)
	return string(jb)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (b *Budget) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (b *Budget) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (b *Budget) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
