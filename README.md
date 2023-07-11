# dhc.dev

source code and static files for my personal web site

## basic usage

1.  build the server: `go build`
2.  build the site complier: `cd build && go build && cd ..`
3.  compile the site: `./build/build`
4.  start serving: `./dhcdev` and open http://localhost:8080 in your browser

while iterating on the pages or posts, only steps 3 and 4 need to be repeated.

## build and run the container

    docker build --tag dhcdev .
    docker run -d -p 8080:8080 dhcdev

## deploy the site

TODO
