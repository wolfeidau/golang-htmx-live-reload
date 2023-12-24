package main

import "github.com/alecthomas/kong"

type websiteFlags struct {
	Version kong.VersionFlag
	Addr    string `default:"localhost:3333"`
	DevMode bool   `default:"false" env:"DEV_MODE"`
}
