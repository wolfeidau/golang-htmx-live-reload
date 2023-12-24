package main

import (
	"fmt"
	"html/template"
	"io"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/wolfeidau/echo-htmx/assets"
	"github.com/wolfeidau/echo-htmx/public"
	"github.com/wolfeidau/echo-htmx/views"
	middleware "github.com/wolfeidau/echo-middleware"
	templates "github.com/wolfeidau/echo-views"
	live_templates "github.com/wolfeidau/reflex/templates"
)

// configureRouter configures the Echo router with middleware, static file serving,
// templates, and error handling. It takes the Echo instance, a logger, and
// website config flags as parameters.

// It adds the zerolog middleware for logging requests. It serves static files from
// the public and assets directories. It configures template rendering, including
// live reloading of templates in dev mode. It sets the custom error handler.

// This handles all the core middleware and rendering setup for the application.
func configureRouter(e *echo.Echo, logger zerolog.Logger, cfg websiteFlags) error {
	e.Pre(middleware.ZeroLogWithConfig(middleware.ZeroLogConfig{
		Logger: logger,
	}))

	e.StaticFS("/", public.GetContent(cfg.DevMode))

	e.HideBanner = true
	e.Logger.SetOutput(io.Discard)

	templateFuncs := template.FuncMap{
		"LiveReload": func() template.HTML { return template.HTML("") },
	}

	if cfg.DevMode {
		liveReloadHTML, err := live_templates.InjectedHTML()
		if err != nil {
			return fmt.Errorf("failed to load injected html: %w", err)
		}

		templateFuncs["LiveReload"] = func() template.HTML { return template.HTML(liveReloadHTML) }
	}

	htmlTemplates := templates.New(
		templates.WithFS(views.GetContent(cfg.DevMode)),
		templates.WithAutoReload(cfg.DevMode),
		templates.WithLogger(newViewLogger(logger)),
		templates.WithFuncs(templateFuncs),
	)

	e.StaticFS("/assets", assets.GetContent(cfg.DevMode))

	if cfg.DevMode {
		e.StaticFS("/assets", os.DirFS("./assets"))
	}

	err := htmlTemplates.AddWithLayout("layouts/base.html", "pages/*.html")
	if err != nil {
		return fmt.Errorf("failed to add templates: %w", err)
	}

	err = htmlTemplates.Add("fragments/*.html")
	if err != nil {
		return fmt.Errorf("failed to add fragments: %w", err)
	}

	e.Renderer = htmlTemplates
	e.HTTPErrorHandler = public.ErrorPageHandler

	return nil
}
