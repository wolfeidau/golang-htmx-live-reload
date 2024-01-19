package main

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type viewLogger struct {
	logger zerolog.Logger
}

func newViewLogger(logger zerolog.Logger) *viewLogger {
	return &viewLogger{logger: logger}
}

func (vl *viewLogger) DebugCtx(ctx context.Context, msg string, fields map[string]any) {
	log.Ctx(ctx).Debug().Fields(fields).Msg(msg)
}

func (vl *viewLogger) ErrorCtx(ctx context.Context, msg string, err error, fields map[string]any) {
	log.Ctx(ctx).Error().Err(err).Fields(fields).Msg(msg)
}

func (vl *viewLogger) Debug(msg string, fields map[string]any) {
	vl.logger.Debug().Fields(fields).Msg(msg)
}

// configureLogger configures the global logger used by the application.
// It sets the output to stderr and enables caller info if devMode is true.
func configureLogger(devMode bool) zerolog.Logger {
	if devMode {
		log.Logger = zerolog.New(zerolog.NewConsoleWriter()).With().Timestamp().Caller().Logger()
	}

	return log.Logger
}
