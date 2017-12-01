package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

type Gift struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Name      string    `json:"name" db:"name"`
	PriceInt  int64     `json:"price" db:"price"`
	Price     string    `json:"-" db:"-"`
	Url       string    `json:"url" db:"url"`
	PersonID  uuid.UUID `json:"person_id" db:"person_id"`
	EventID   uuid.UUID `json:"event_id" db:"event_id"`
	Status    string    `json:"status" db:"status"`
}

// String is not required by pop and may be deleted
func (g Gift) String() string {
	jg, _ := json.Marshal(g)
	return string(jg)
}

// Gifts is not required by pop and may be deleted
type Gifts []Gift

// String is not required by pop and may be deleted
func (g Gifts) String() string {
	jg, _ := json.Marshal(g)
	return string(jg)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (g *Gift) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var v []validate.Validator
	v = append(v, &validators.StringIsPresent{Field: g.Name, Name: "Name"})
	if g.Url != "" {
		v = append(v, &validators.URLIsPresent{Field: g.Url, Name: "Url"})
	}
	return validate.Validate(
		v...,
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (g *Gift) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (g *Gift) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (u *Gift) AfterFind(tx *pop.Connection) error {
	u.Price = RenderMoney(u.PriceInt)
	return nil
}

func (u *Gift) BeforeCreate(tx *pop.Connection) error {
	var err error
	u.PriceInt, err = ParseMoney(u.Price)
	return errors.WithStack(err)
}

func (u *Gift) BeforeSave(tx *pop.Connection) error {
	var err error
	u.PriceInt, err = ParseMoney(u.Price)
	return errors.WithStack(err)
}
