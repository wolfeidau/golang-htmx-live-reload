# golang-htmx-live-reload

This is project provides scaffolding for a [Go](https://go.dev) based web site with live reloading to speed up development.

# Features

This project includes the following:

* Flat package structure for easy development.
* Live reloading using my fork of [reflex](https://github.com/wolfeidau/reflex) which has websockets support.
* Templating using [html/template](https://golang.org/pkg/html/template/).
* [echo-views](https://github.com/wolfeidau/echo-views) package for rendering views.
* [echo-middleware](https://github.com/wolfeidau/echo-middleware) I built which provides a set of useful middleware for logging and caching.
* Error pages located in [public](public) which are rendered based on error code, these can be modified in [public](public).
* When `DEV_MODE` is enabled all assets are live reloaded from the filesystem making updates to templates and assets fast and simple.
* When `DEV_MODE` is disabled all the assets are built into the binary ,which makes shipping the application to production simple.
* Building of [assets](assets) via [esbuild](https://github.com/evanw/esbuild) is configured via reflex, this provides fast and efficient bundling of javascript.
* Releases using [goreleaser](https://goreleaser.com/) which builds the application for multiple platforms.
* [TailwindCSS](https://tailwindcss.com/) for styling.

# Usage

To get started just run `make` and open the site and start editing the templates in the [views](views) folder.

After retrieving dependencies you should see the following output.
```
Watching for changes...
go run github.com/wolfeidau/reflex -c reflex.conf
[02] Starting service
[02] <nil> DBG logger.go:26 > register template includes= layout=layouts/base.html name=pages/index.html
[02] <nil> INF main.go:31 > Serving on http://localhost:3333
```

# References

* [emojifavicon](https://emojifavicon.dev/) for the favicon.

# License

This application is released under Apache 2.0 license and is copyright [Mark Wolfe](https://www.wolfe.id.au/?utm_source=golang-htmx-live-reload).
