package actions

import (
	"github.com/gobuffalo/buffalo"
)

func eventsHandler(c buffalo.Context) error {
	// Render the html
	return c.Render(200, r.HTML("index.html"))
}
