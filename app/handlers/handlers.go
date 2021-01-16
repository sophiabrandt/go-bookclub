// packag handlers contains all the route handlers for the application
package handlers

import (
	"net/http"
	"os"

	"github.com/sophiabrandt/go-bookclub/business/logger"
	"github.com/sophiabrandt/go-bookclub/business/mid"
	"github.com/sophiabrandt/go-bookclub/foundation/web"
)

// API constructs an http.Handler with all application routes defined.
func API(build string, shutdown chan os.Signal, logger *logger.Logger) http.Handler {
	app := web.NewApp(shutdown, mid.Logger(logger), mid.Errors(logger), mid.Metrics(), mid.Panics(logger))

	check := checkGroup{
		logger: logger,
	}

	app.Handle(http.MethodGet, "/readiness", check.readiness)

	return app
}
