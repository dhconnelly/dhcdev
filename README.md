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

1.  Build the tools: `cd build && go build && cd ../serve && go build && cd ..`
2.  Compile the site and start serving: `./build/build && ./serve/serve`
3.  Open/refresh http://localhost:8080 in your browser

While iterating on the pages or posts, only steps 2 and 3 need to be repeated.

Pass `-h` to either the `build` or `serve` commands to see a few options.

### Build and run locally with Docker

    docker-compose up -d --build

### Run integration tests

    ./test.sh

### Deploy to production

The docker image (including the static site and the server) is automatically
built and deployed to Digital Ocean via Github Actions on each commit to the
main branch.

## Implementation

There are two parts:

-   `build`: a static site builder. It copies files from one directory to
    another, rendering any Markdown (`.md`) files into HTML and copying all
    other files without modification.

-   `serve`: serves static files. There's some wrapper logic around the Go
    standard library's `http.FileServer` to log requests and response codes,
    expose basic metrics using `expvar` at `/debug/vars`,
    and add a `Cache-Control` header for Cloudflare CDN.

The `Dockerfile` builds both tools, calls `build` to build the files in `pages`,
and calls `serve` to serve them.

## License

The content in `pages` is licensed under CC BY 4.0. The rest of the files are
licensed under the MIT license. See `LICENSE.md` for more details.
