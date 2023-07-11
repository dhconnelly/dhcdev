# dhc.dev

source code and static files for my personal web site

## build and serve the site locally

1.  build the server: `go build`
2.  build the site complier: `cd build && go build && cd ..`
3.  compile the site: `./build/build`
4.  start serving: `./dhcdev` and open http://localhost:8080 in your browser

while iterating on the pages or posts, only steps 3 and 4 need to be repeated.

## build the image and run a container locally

    docker build --tag dhcdev .
    docker run -d -p 8080:8080 dhcdev

## deploy to production

the docker image (including the static site and the server) is automatically
built and deployed to digital ocean via github actions on each commit to the
main branch.
