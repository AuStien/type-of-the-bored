# Type of the Bored

Type the quote that shows up, **_GET INSTANT FEEDBACK; NOW IN COLORS!_**

Offered as both a CLI and web app/backend.

**NOTE: This repo is only tested on Arch and Ubuntu. If a problem arises on your OS feel free to create a PR to resolve the issue.**

## About

The web app runs in pure HTML/CSS/JS with the only external import being [htmx](https://htmx.org) to update the DOM with new content from the backend.
HTML templating in Go is done with [templ](https://github.com/a-h/templ).

## Getting started

Requirements:

- Go (`>= v1.21`)

_Other CLI tools are used ([templ](https://github.com/a-h/templ) and [gow](https://github.com/mitranim/gow)), but these are automatically downloaded to a local `bin` folder by the Makefile._

**CLI**

1. `make run-cli`

**Web App**

1. `make run-server`
2. Navigate to `localhost:8080` in your browser

---

For help with what commands are available in the Makefile, run `make help`.

## Development

For the best development experience working on the web app, run `make generate-watch` and `make watch-server` in two different terminals to hot-reload the server when changes are made. _Note that you will need to manually refresh the browser._
