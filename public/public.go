package public

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

//go:embed *.html *.png *.ico
var content embed.FS

// GetContent returns the content file system to use.
// If devMode is true, it returns an os.DirFS pointing to the views directory.
// Otherwise, it returns the compiled-in content from the embed.FS.
func GetContent(devMode bool) fs.FS {
	if devMode {
		return os.DirFS("./public")
	}
	return content
}

func ErrorPageHandler(err error, c echo.Context) {
	ctx := c.Request().Context()

	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	log.Ctx(ctx).Error().Err(err).Int("code", code).Str("target_path", c.Request().URL.Path).Msg("HTTP error")

	errorPage := fmt.Sprintf("%d.html", code)
	data, err := content.ReadFile(errorPage)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("error page not found")

		return
	}

	err = c.HTMLBlob(code, data)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("error page render failed")
	}
}
