package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gobuffalo/envy"
	"github.com/unrolled/secure"

	"github.com/gobuffalo/x/sessions"
	"github.com/rs/cors"
	"github.com/touslesmemes/api/models"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:          ENV,
			SessionStore: sessions.Null{},
			PreWares: []buffalo.PreWare{
				cors.Default().Handler,
			},
			SessionName: "_api_session",
		})
		// Automatically redirect to SSL
		app.Use(ssl.ForceSSL(secure.Options{
			SSLRedirect:     ENV == "production",
			SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
		}))

		// Set the request content type to JSON
		app.Use(middleware.SetContentType("application/json"))

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.PopTransaction)
		// Remove to disable this.
		app.Use(middleware.PopTransaction(models.DB))

		app.GET("/", HomeHandler)

		api := app.Group("/v1")
		api.POST("/auth/login", UsersLogin)

		users := api.Group("/users")
		channels := api.Group("/channels")
		comments := api.Group("/comments")
		posts := api.Group("/posts")

		// restrict access to authenticated users
		users.Use(RestrictedHandlerMiddleware)

		ur := &UsersResource{}
		cr := &ChannelsResource{}
		cr7 := &CommentsResource{} //cr7 stand for cristiano ronaldo
		pr := &PostsResource{}

		users.Middleware.Skip(RestrictedHandlerMiddleware, ur.New, ur.Create)

		channels.Use(RestrictedHandlerMiddleware)
		comments.Use(RestrictedHandlerMiddleware)
		posts.Use(RestrictedHandlerMiddleware)

		users.Resource("", ur)
		channels.Resource("", cr)
		comments.Resource("", cr7)
		posts.Resource("", pr)

		// restrict access to authenticated users

	}

	return app
}
