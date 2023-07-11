# syntax=docker/dockerfile:1

FROM golang:1.20

# build the server
WORKDIR /app
COPY go.mod ./
COPY *.go ./
RUN go build

# build the site compiler
WORKDIR /app/build
COPY build/*.go ./
RUN go build

# compile the site
WORKDIR /app
COPY pages/ ./pages/
RUN ./build/build -srcDir=pages -dstDir=target

# start
EXPOSE 8080
CMD ["./dhcdev", "-port=8080", "-serveDir=target"]
