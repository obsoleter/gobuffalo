package actions

import (
	"net/http"
	"os"

	"github.com/markbates/buffalo"
	"github.com/markbates/buffalo/middleware"
	"github.com/markbates/gobuffalo/models"
	"github.com/markbates/going/defaults"
)

var ENV = defaults.String(os.Getenv("GO_ENV"), "development")

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() http.Handler {
	a := buffalo.Automatic(buffalo.Options{
		Env: ENV,
	})

	a.Use(middleware.PopTransaction(models.DB))
	a.ServeFiles("/assets", assetsPath())
	a.GET("/", HomeHandler)

	return a
}