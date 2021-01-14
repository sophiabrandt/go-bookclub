// packag handlers contains all the route handlers for the application
package handlers

import (
	"net/http"
	"os"

	"github.com/sophiabrandt/go-bookclub/business/logger"
	"github.com/vmihailenco/treemux"
)

// API constructs an http.Handler with all application routes defined.
func API(build string, shutdown chan os.Signal, logger *logger.Logger) *treemux.TreeMux {
	mux := treemux.New()

	check := checkGroup{
		logger: logger,
	}

	mux.Handle(http.MethodGet, "/test", treemux.HTTPHandlerFunc(check.readiness))

	return mux
}
