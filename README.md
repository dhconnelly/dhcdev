# dhcdev

Source code and static content for my personal web site.

## Usage

All files in the source directory (by default `pages`) will be compiled into
the target directory (by default `target`). The logic for doing this is very
simple:

-   Markdown (.md) files will be rendered as HTML (.html) files into the
    target directory. They should have the form:

        === $TITLE ===
        $CONTENT

    All Markdown files will be rendered using a single Go template (by default
    `templates/post-template.html`). During rendering, $TITLE and $CONTENT data
    will be accessible in the `{{.Title}}` and `{{.Content}}` variables.

-   All other files are copied to the target directory without modification.

### Build and serve the site locally

1.  Build the tool: `go build`
2.  Compile the site and start serving: `./dhcdev`
3.  Open/refresh http://localhost:8080 in your browser

To automatically rebuild and serve the site while writing:

    gow -e=go,md run .

### Run integration tests

    ./test.sh

### Deploy to production

The site is served on GCP Cloud Run. It will automatically pick up new pushes
to the main branch and deploy them after merge.

## Implementation

There are three parts:

-   `package build`: implements a simple static site builder. It copies files
    from one directory to another, rendering any Markdown (`.md`) files into
    HTML and copying all other files without modification.

-   `package serve`: implements an http handler for static files. There's some
    wrapper logic around the Go standard library's `http.FileServer` to log
    requests and response codes, expose basic metrics using `expvar` at
    `/debug/vars`, and add a `Cache-Control` header for Cloudflare CDN.

-   `package main` wraps up `build` and `serve`. The `Dockerfile` uses this
    to build and serve the site, but separates the build and serving stages in
    order to deploy the static assets in the Docker image.

## License

The content in `pages` is licensed under CC BY 4.0. The rest of the files are
licensed under the MIT license. See `LICENSE.md` for more details.
