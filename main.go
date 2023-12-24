package main

import (
	"errors"
	"net/http"

	"github.com/alecthomas/kong"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

var (
	version = "dev"
	cfg     websiteFlags
)

func main() {
	kong.Parse(&cfg, kong.Vars{"version": version})
	e := echo.New()

	// setup logger with neatly formatted output for dev mode
	// and structured JSON for production
	logger := configureLogger(cfg.DevMode)

	// setup middleware, logging and other integrations here
	err := configureRouter(e, logger, cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to setup router")
	}

	// setup the core of your application here, routes, handlers, etc
	setupRoutes(e, cfg)

	log.Info().Msgf("Serving on http://%s", cfg.Addr)
	if err = e.Start(cfg.Addr); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal().Err(err).Msg("failed to start listener")
	}
}
