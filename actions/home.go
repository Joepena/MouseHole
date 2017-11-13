package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/joepena/mouse_hole/models"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	readReq := &models.ReadRequest{
		"test_app",
		"event",
		//Event{},
	}
	events := models.GetDBInstance().Read(readReq)

	c.Set("events", events)
	return c.Render(200, r.HTML("index.html"))
}

