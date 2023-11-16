# syntax=docker/dockerfile:1

# build the image
FROM golang:1.21-bullseye AS build-stage
WORKDIR /app
RUN GOBIN=/app go install github.com/dhconnelly/sss@v1.0.0

# build the site
COPY . .
RUN ./sss -serve=false -srcDir=pages -dstDir=target -postTmpl=templates/post-template.html

# package the server and files
FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /app
COPY --from=build-stage /app/target/ ./target/
COPY --from=build-stage /app/sss ./sss

# serve
EXPOSE 8080
USER nonroot:nonroot
CMD ["./sss", "-build=false", "-port=8080", "-serveDir=target"]
