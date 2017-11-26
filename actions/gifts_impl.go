package actions

import (
	"github.com/markbates/pop"
	"github.com/pkg/errors"
	"github.com/theothertomelliott/giftlist/models"
)

func getGiftWithRelations(tx *pop.Connection, gifts ...models.Gift) ([]GiftWithRelations, error) {
	var giftsWithRelations = make([]GiftWithRelations, 0, len(gifts))
	for _, gift := range gifts {
		// Obtain associated Event
		event := &models.Event{}
		if err := tx.Find(event, gift.EventID); err != nil {
			return nil, errors.WithStack(err)
		}

		// Obtain associated Person
		person := &models.Person{}
		if err := tx.Find(person, gift.PersonID); err != nil {
			return nil, errors.WithStack(err)
		}

		giftsWithRelations = append(
			giftsWithRelations,
			GiftWithRelations{
				Gift:   gift,
				Event:  event,
				Person: person,
			},
		)
	}
	return giftsWithRelations, nil
}
