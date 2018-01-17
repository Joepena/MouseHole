package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gobuffalo/envy"
	"github.com/unrolled/secure"

	"github.com/gobuffalo/buffalo/middleware/csrf"
	"github.com/gobuffalo/buffalo/middleware/i18n"
	"github.com/gobuffalo/packr"
	"github.com/joepena/mouse_hole/api"
	"github.com/joepena/mouse_hole/models"
	"github.com/gorilla/sessions"
)

// TODO: remove this from source later
var SERVER_SECRET = []byte("057E3CE6B941756FD9CAB17D93C522F7C3745A78066A278E83999FFF547C5A8F")

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var T *i18n.Translator
var hub = api.NewHub()

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		// init DB
		models.GetDBInstance()

		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_mouse_hole_session",
			SessionStore: sessions.NewCookieStore(SERVER_SECRET),
		})

		// turn context to MMContext
		app.Use(func(next buffalo.Handler) buffalo.Handler {
			return func(c buffalo.Context) error {
				// change the context to MMContext
				return next(MMContext{c})
			}
		})

		// Automatically redirect to SSL
		app.Use(ssl.ForceSSL(secure.Options{
			SSLRedirect:     ENV == "production",
			SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
		}))

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		app.Use(csrf.New)

		// Setup and use translations:
		var err error
		if T, err = i18n.New(packr.NewBox("../locales"), "en-US"); err != nil {
			app.Stop(err)
		}

		app.Use(T.Middleware())

		apiGroup := app.Group("/api")
		apiGroup.Use(authenticateRequest)
		apiGroup.Use(ReadRequestAssigner)
		apiGroup.GET("/", apiHandler)

		// init hub for api package
		go hub.Run()

		app.GET("/events", eventsHandler)
		app.GET("/events_socket", eventSocketHandler)
		app.GET("/user", createUserHandler)
		app.ServeFiles("/assets", assetsBox)
	}

	return app
}
