package models_test

import (
	"testing"

	"github.com/theothertomelliott/giftlist/models"
)

func TestValidateGift(t *testing.T) {
	var tests = []struct {
		name                 string
		gift                 models.Gift
		expectedErrorsString string
	}{
		{
			name:                 "Empty gift",
			expectedErrorsString: `{"errors":{"name":["Name can not be blank."]}}`,
		},
		{
			name: "Empty URL",
			gift: models.Gift{
				Name: "Name",
			},
			expectedErrorsString: `{"errors":{}}`,
		},
		{
			name: "Invalid URL",
			gift: models.Gift{
				Name: "Name",
				Url:  "Invalid",
			},
			expectedErrorsString: `{"errors":{"url":["Url does not match url format. Err: parse Invalid: invalid URI for request"]}}`,
		},
		{
			name: "Valid",
			gift: models.Gift{
				Name: "Name",
				Url:  "http://telliott.io",
			},
			expectedErrorsString: `{"errors":{}}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotErrors, err := test.gift.Validate(nil)
			if err != nil {
				t.Error(err)
			}
			if gotErrors.String() != test.expectedErrorsString {
				t.Errorf("Expected %+v, got %+v", test.expectedErrorsString, gotErrors.String())
			}
		})
	}
}
