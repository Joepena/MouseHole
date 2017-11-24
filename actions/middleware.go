package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/joepena/mouse_hole/models"
	"github.com/markbates/pop"
)

const READ_REQUEST = "readRequest"

func ReadRequestAssigner(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		readReq := models.NewReadRequest()

		// assign the db to the readrequest and then add it to the context to be passed on
		usr := c.Data()[CURRENT_USER].(models.User)
		readReq.SetDB(usr.DBName)
		c.Set(READ_REQUEST, readReq)

		err := next(c)

		return err
	}
}

// SetCurrentUser attempts to find a user based on the current_user_id
// in the session. If one is found it is set on the context.
func SetCurrentUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid != nil {
			u := &models.User{}
			tx := c.Value("tx").(*pop.Connection)
			err := tx.Find(u, uid)
			if err != nil {
				return err
			}
			c.Set(CURRENT_USER, u)
		}
		return next(c)
	}
}

// Authorize require a user be logged in before accessing a route
func Authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid == nil {
			c.Flash().Add("danger", "You must be authorized to see that page")
			return c.Redirect(302, "/")
		}
		return next(c)
	}
}
