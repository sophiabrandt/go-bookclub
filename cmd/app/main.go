package main

import (
	"context"
	"expvar" // Register the expvar handlers
	"fmt"
	"net/http"
	_ "net/http/pprof" // Register the pprof handlers
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ardanlabs/conf"
	"github.com/pkg/errors"
	"github.com/sophiabrandt/go-bookclub/app/handlers"
	lr "github.com/sophiabrandt/go-bookclub/business/logger"
)

// build is the git version of this program. It is set using build flags in the makefile.
var build = "develop"

func main() {
	logger := lr.Init()

	if err := run(logger); err != nil {
		logger.Error(fmt.Sprintf("main: error: ", err))
		os.Exit(1)
	}
}

func run(logger *lr.Logger) error {
	// =========================================================================
	// Configuration

	var cfg struct {
		conf.Version
		Web struct {
			APIHost         string        `conf:"default:0.0.0.0:8000"`
			DebugHost       string        `conf:"default:0.0.0.0:6060"`
			ReadTimeout     time.Duration `conf:"default:5s"`
			WriteTimeout    time.Duration `conf:"default:5s"`
			ShutdownTimeout time.Duration `conf:"default:5s"`
			IdleTimeout     time.Duration `conf:"default:120s"`
		}
	}

	cfg.Version.SVN = build
	cfg.Version.Desc = "Apache 2.0 License"

	if err := conf.Parse(os.Args[1:], "BOOKCLUB", &cfg); err != nil {
		switch err {
		case conf.ErrHelpWanted:
			usage, err := conf.Usage("BOOKCLUB", &cfg)
			if err != nil {
				return errors.Wrap(err, "generating config usage")
			}
			fmt.Println(usage)
			return nil
		case conf.ErrVersionWanted:
			version, err := conf.VersionString("BOOKCLUB", &cfg)
			if err != nil {
				return errors.Wrap(err, "generating config version")
			}
			fmt.Println(version)
			return nil
		}
		return errors.Wrap(err, "parsing config")
	}
	// =========================================================================
	// App Starting

	// Print the build version for our logs. Also expose it under /debug/vars.
	expvar.NewString("build").Set(build)
	logger.Info(fmt.Sprintf("main : Started : Application initializing : version %q", build))
	defer logger.Info("main: Completed")

	logger.Info(fmt.Sprintf("main: Config: %+v;", cfg))

	// =========================================================================
	// Start Debug Service
	//
	// /debug/pprof - Added to the default mux by importing the net/http/pprof package.
	// /debug/vars - Added to the default mux by importing the expvar package.

	logger.Info("main: Initializing debugging support")

	go func() {
		logger.Info(fmt.Sprintf("main: Debug Listening %s", cfg.Web.DebugHost))
		if err := http.ListenAndServe(cfg.Web.DebugHost, http.DefaultServeMux); err != nil {
			logger.Info(fmt.Sprintf("main: Debug Listener closed : %v", err))
		}
	}()

	// =========================================================================
	// Start API Service

	logger.Info("main: Initializing API support")

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	api := http.Server{
		Addr:         cfg.Web.APIHost,
		Handler:      handlers.API(build, shutdown, logger),
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
		IdleTimeout:  cfg.Web.IdleTimeout,
	}

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	// Start the service listening for requests.
	go func() {
		logger.Info(fmt.Sprintf("main: API listening on %s", api.Addr))
		serverErrors <- api.ListenAndServe()
	}()

	// =========================================================================
	// Shutdown

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return errors.Wrap(err, "server error")

	case sig := <-shutdown:
		logger.Info(fmt.Sprintf("main: %v : Start shutdown", sig))

		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), cfg.Web.ShutdownTimeout)
		defer cancel()

		// Asking listener to shutdown and shed load.
		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return errors.Wrap(err, "could not stop server gracefully")
		}
	}

	return nil
}
