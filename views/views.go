package views

import (
	"embed"
	"io/fs"
	"os"
)

//go:embed pages/* layouts/* fragments/*
var content embed.FS

// GetContent returns the content file system to use.
// If devMode is true, it returns an os.DirFS pointing to the views directory.
// Otherwise, it returns the compiled-in content from the embed.FS.
func GetContent(devMode bool) fs.FS {
	if devMode {
		return os.DirFS("./views")
	}
	return content
}
