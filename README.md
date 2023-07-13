# dhcdev

Source code and static content for my personal web site.

## Usage

All html and md files in the source directory (by default `pages`) will be
compiled into the target directory (by default `target`). The logic for doing
this is very simple:

Markdown (.md) files should have the form:

    --
    title: $TITLE
    --
    $MARKDOWN

The `$MARKDOWN` content will be rendered from Markdown into HTML and the
page `<title>` will be set from `$TITLE`.

HTML (.html) files are passed through unmodified.

### Build and serve the site locally

1.  Build the tools: `go build ./...`
2.  Compile the site and start serving: `./build/build && ./serve/serve`
3.  Open/refresh http://localhost:8080 in your browser

While iterating on the pages or posts, only steps 2 and 3 need to be repeated.

Pass `-h` to either the `build` or `serve` commands to see a few options.

### Build and run locally with Docker

    docker-compose up -d

### Run integration tests

    ./test.sh

### Deploy to production

The docker image (including the static site and the server) is automatically
built and deployed to Digital Ocean via Github Actions on each commit to the
main branch.

## Implementation

There are two packages:

-   `build`: a static site builder. It transforms a directory of HTML and
    Markdown files into a directory of HTML files by rendering the Markdown
    and passing through the HTML unmodified.

-   `serve`: serves static files. There's some wrapper logic around the Go
    standard library's `http.FileServer` to log requests and response codes
    and to expose basic metrics using `expvar` at `/debug/status`.

## License

The content in `pages` is licensed under CC BY 4.0. The rest of the files are
licensed under the MIT license. See `LICENSE.md` for more details.
