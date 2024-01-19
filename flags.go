package main

import "github.com/alecthomas/kong"

type websiteFlags struct {
	Version kong.VersionFlag
	Addr    string `default:"localhost:3333" env:"ADDR"`
	DevMode bool   `env:"DEV_MODE"`
}
