package actions

const CURRENT_USER = "current_user"

//func UsersNew(c buffalo.Context) error {
//	u := models.User{}
//	c.Set("user", u)
//	return c.Render(200, r.HTML("users/new.html"))
//}

// UsersCreate registers a new user with the application.

//func UsersCreate(c buffalo.Context) error {
//	u := &models.User{}
//	if err := c.Bind(u); err != nil {
//		return errors.WithStack(err)
//	}
//
//	tx := c.Value("tx").(*pop.Connection)
//	verrs, err := u.Create(tx)
//	if err != nil {
//		return errors.WithStack(err)
//	}
//
//	if verrs.HasAny() {
//		c.Set("user", u)
//		c.Set("errors", verrs)
//		return c.Render(200, r.HTML("users/new.html"))
//	}
//
//	c.Session().Set("current_user_id", u.ID)
//	c.Flash().Add("success", "Welcome to Buffalo!")
//
//	return c.Redirect(302, "/")
//}
