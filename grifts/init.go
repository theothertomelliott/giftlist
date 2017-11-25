package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/theothertomelliott/giftlist/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
