package actions

import (
	"github.com/markbates/pop"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/theothertomelliott/giftlist/models"
)

func loadPeopleForEvent(tx *pop.Connection, currentUser uuid.UUID, event *models.Event) (PeopleForEvent, error) {
	// Retrieve all People from the DB
	people := &models.Persons{}
	// Retrieve all Persons from the DB
	if err := tx.Where("user_id = ?", currentUser).All(people); err != nil {
		return nil, errors.WithStack(err)
	}
	var peopleMap = make(map[uuid.UUID]PersonForEvent)
	for _, person := range *people {
		peopleMap[person.ID] = PersonForEvent{
			Person: person,
		}
	}

	// Retrieve all related Gifts from the DB
	gifts := models.Gifts{}
	if err := tx.Where("event_id = ?", event.ID).All(&gifts); err != nil {
		return nil, errors.WithStack(err)
	}
	giftsWithRelations, err := getGiftWithRelations(tx, gifts...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Retrieve all related Budgets from the DB
	budgets := models.Budgets{}
	if err := tx.Where("event_id = ?", event.ID).All(&budgets); err != nil {
		return nil, errors.WithStack(err)
	}

	for _, gift := range giftsWithRelations {
		person := peopleMap[gift.PersonID]
		person.Gifts = append(person.Gifts, gift.Gift)
		peopleMap[gift.PersonID] = person
	}
	for _, budget := range budgets {
		person := peopleMap[budget.PersonID]
		person.Budget = &budget
		peopleMap[budget.PersonID] = person
	}

	var output = make(PeopleForEvent, 0, len(peopleMap))
	for _, personForEvent := range peopleMap {
		output = append(output, personForEvent)
	}

	return output, nil
}

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
