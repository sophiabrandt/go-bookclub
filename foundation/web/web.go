// package web implements a small web framework extension for the application.
package web

import (
	"context"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/matoous/go-nanoid/v2"
	"github.com/vmihailenco/treemux"
)

// ctxKey represents the type of value for the context key.
type ctxKey int

// KeyValues is how request values are stored/retrieved.
const KeyValues ctxKey = 1

// Values represent state for each request.
type Values struct {
	TraceID    string
	Now        time.Time
	StatusCode int
}

// A Handler is a type that handles an http request within our own little mini
// framework.
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers.
type App struct {
	mux      *treemux.TreeMux
	shutdown chan os.Signal
}

// NewApp creates an App value that handle a set of routes for the application.
func NewApp(shutdown chan os.Signal) *App {
	mux := treemux.New()

	return &App{
		mux:      mux,
		shutdown: shutdown,
	}
}

// SignalShutdown is used to gracefully shutdown the app when an integrity
// issue is identified.
func (a *App) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}

// ServeHTTP implements the http.Handler interface. It's the entry point for
// all http traffic.
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

// Handle sets a handler function for a given HTTP method and path pair
// to the application server mux.
func (a *App) Handle(method string, path string, handler Handler) {
	h := treemux.HTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// Set the context with the required values to
		// process the request.
		id, err := gonanoid.New()
		if err != nil {
			a.SignalShutdown()
			return
		}
		v := Values{
			TraceID: id,
			Now:     time.Now(),
		}
		ctx = context.WithValue(ctx, KeyValues, &v)

		if err := handler(ctx, w, r); err != nil {
			a.SignalShutdown()
			return
		}
	})
	a.mux.Handle(method, path, h)
}
