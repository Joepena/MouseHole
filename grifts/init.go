package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/joepena/mouse_hole/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
