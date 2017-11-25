package actions

import (
	"github.com/markbates/pop"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/theothertomelliott/giftlist/models"
)

// buildPersonMap builds a map of all Person names to the corresponding IDs,
// this can be passed to a template to build a form selector.
func buildPersonMap(tx *pop.Connection) (map[string]uuid.UUID, error) {
	people := &models.Persons{}
	// Retrieve all Persons from the DB
	if err := tx.All(people); err != nil {
		return nil, errors.WithStack(err)
	}
	personMap := make(map[string]uuid.UUID)
	for _, person := range *people {
		personMap[person.Name] = person.ID
	}
	return personMap, nil
}
