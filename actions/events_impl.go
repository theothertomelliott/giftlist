package actions

import (
	"github.com/markbates/pop"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/theothertomelliott/giftlist/models"
)

// buildEventMap builds a map of all Event names to the corresponding IDs,
// this can be passed to a template to build a form selector.
func buildEventsMap(tx *pop.Connection) (map[string]uuid.UUID, error) {
	events := &models.Events{}
	// Retrieve all Events from the DB
	if err := tx.All(events); err != nil {
		return nil, errors.WithStack(err)
	}
	eventMap := map[string]uuid.UUID{
		"== Choose an Event ===": uuid.Nil,
	}
	for _, event := range *events {
		eventMap[event.Name] = event.ID
	}
	return eventMap, nil
}
