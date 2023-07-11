# syntax=docker/dockerfile:1

# thanks to https://docs.docker.com/language/golang/build-images

FROM golang:1.20 AS build-stage

# build the server
WORKDIR /app
COPY go.mod ./
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build

# build the site compiler
WORKDIR /app/build
COPY build/*.go ./
RUN CGO_ENABLED=0 GOOS=linux go build

# compile the site
WORKDIR /app
COPY pages/ ./pages/
RUN ./build/build -srcDir=pages -dstDir=target

# build lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /app
COPY --from=build-stage /app/dhcdev ./
COPY --from=build-stage /app/target/ ./target/

# start
EXPOSE 8080
USER nonroot:nonroot
CMD ["./dhcdev", "-port=8080", "-serveDir=target"]
