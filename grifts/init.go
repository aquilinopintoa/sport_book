package grifts

import (
	"github.com/aquilinopintoa/sport_book/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}