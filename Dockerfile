# syntax=docker/dockerfile:1

# thanks to https://docs.docker.com/language/golang/build-images

FROM golang:1.20 AS build-stage

# copy over files and build the tool
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build

# build the site
RUN ./dhcdev -serve=false -srcDir=pages -dstDir=target -postTmpl=templates/post-template.html

# build lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /app
COPY --from=build-stage /app/dhcdev ./
COPY --from=build-stage /app/target/ ./target/

# start the server
EXPOSE 8080
USER nonroot:nonroot
CMD ["./dhcdev", "-build=false", "-port=8080", "-serveDir=target"]
