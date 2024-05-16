# syntax=docker/dockerfile:1

# build the image
FROM golang:1.22-bullseye AS build-stage
WORKDIR /app

# build the site
COPY . .
RUN go build
RUN ./dhc.dev -serve=false

# package the server and files
FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /app
COPY --from=build-stage /app/out/ ./out/
COPY --from=build-stage /app/dhc.dev ./dhc.dev

# serve
EXPOSE 8080
USER nonroot:nonroot
CMD ["./dhc.dev", "-build=false", "-port=8080"]
