package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/joepena/mouse_hole/api"
)

func apiHandler(c buffalo.Context) error {
	api.RegisterClient(hub, c)

	return nil
}
