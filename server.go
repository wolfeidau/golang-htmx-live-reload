package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Website struct{}

func newWebsite() *Website {
	return &Website{}
}

func (ws *Website) Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]any{
		"Title": "Home",
	})
}

func (ws *Website) DateTime(c echo.Context) error {
	return c.Render(http.StatusOK, "datetime.html", map[string]any{
		"DateTime": time.Now().Format(time.RFC3339),
	})
}

func (ws *Website) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"status": "ok",
	})
}
