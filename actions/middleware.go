package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/joepena/mouse_hole/models"
)

const READ_REQUEST = "readRequest"

func ReadRequestAssigner(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		readReq := models.NewReadRequest()

		readReq.SetDB("test_app") // this is a test
		c.Set(READ_REQUEST, readReq)

		err := next(c)

		return err
	}
}
