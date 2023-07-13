# dhc.dev

source code and static files for my personal web site

## build and serve the site locally

1.  build everything: `go build ./...`
2.  compile the site and start serving: `./build/build && ./serve/serve`
3.  open/refresh http://localhost:8080 in your browser

while iterating on the pages or posts, only steps 2 and 3 need to be repeated.

## start locally with docker

    docker-compose up -d

## run integration tests

    ./test.sh

## deploy to production

the docker image (including the static site and the server) is automatically
built and deployed to digital ocean via github actions on each commit to the
main branch.
