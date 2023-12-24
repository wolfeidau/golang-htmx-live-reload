package main

import "github.com/labstack/echo/v4"

// setupRoutes configures the routes for the Echo server instance.
// It takes an Echo instance and a config struct as arguments,
// initializes a website handler, registers a route for the index page,
// and returns any error.
func setupRoutes(e *echo.Echo, cfg websiteFlags) error {

	ws := newWebsite()

	e.GET("/", ws.Index)
	e.GET("/datetime", ws.DateTime)

	return nil
}
