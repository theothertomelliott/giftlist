package models_test

import (
	"testing"

	"github.com/theothertomelliott/giftlist/models"
)

func TestRenderMoney(t *testing.T) {
	var tests = []struct {
		in  int64
		out string
	}{
		{
			in:  100,
			out: "1.00",
		},
		{
			in:  2,
			out: "0.02",
		},
		{
			in:  999999999,
			out: "9999999.99",
		},
		{
			in:  299,
			out: "2.99",
		},
		{
			in:  -299,
			out: "-2.99",
		},
		{
			in:  0,
			out: "0.00",
		},
		{
			in:  -1,
			out: "-0.01",
		},
	}
	for _, test := range tests {
		t.Run(test.out, func(t *testing.T) {
			got := models.RenderMoney(test.in)
			if got != test.out {
				t.Errorf("Expected %s, got %s", test.out, got)
			}
		})
	}
}

func TestParseMoney(t *testing.T) {
	var tests = []struct {
		out int64
		in  string
		err bool
	}{
		{
			out: 100,
			in:  "1.00",
		},
		{
			out: 2,
			in:  "0.02",
		},
		{
			out: 999999999,
			in:  "9999999.99",
		},
		{
			out: 299,
			in:  "2.99",
		},
		{
			out: -299,
			in:  "-2.99",
		},
		{
			out: 0,
			in:  "0.00",
		},
		{
			out: -1,
			in:  "-0.01",
		},
		{
			in:  "invalid",
			err: true,
		},
		{
			in:  "1.2.3",
			err: true,
		},
	}
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			got, err := models.ParseMoney(test.in)
			if (err != nil) != test.err {
				t.Errorf("Error output wasn't as expected, got %v", err)
			}
			if got != test.out {
				t.Errorf("Expected %d, got %d", test.out, got)
			}
		})
	}
}
