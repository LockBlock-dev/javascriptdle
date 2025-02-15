# JavaScriptDle

Prototype of a **dle** (daily game) about [JavaScript](https://developer.mozilla.org/en-US/docs/Web/JavaScript) methods.

This project uses the GoTH stack:

-   [Go](https://go.dev/)
-   [Templ](https://templ.guide/)
-   [HTMX](https://htmx.org/)

## Table of content

-   [**Installation**](#installation)
-   [**How to use**](#how-to-use)
-   [**Copyright**](#copyright)

## Installation

1. Install the dependencies:

    - Download [Go](https://go.dev/dl/) (Go 1.23 required).
    - Download [Templ](https://templ.guide/quick-start/installation).
    - Download or clone the project.

2. Acquire external assets (not committed to avoid unnecessary bloat):

    - Download [HTMX](https://htmx.org/docs/#download-a-copy) and place it in [`public/js/`](./public/js/).
    - Download [OpenSans](https://fonts.google.com/specimen/Open+Sans) and place it in [`public/fonts/`](./public/fonts/).

3. Build the project:

    - Run `make` in the `javascriptdle` directory.

## How to use

-   Run the `javascriptdle` binary inside the [`bin`](./bin) folder.
-   The server listens on port `8080` by default.

## Copyright

See the [license](./LICENSE).
