# syntax=docker/dockerfile:1

# BUILD STAGE
FROM amazoncorretto:21 AS build-stage
WORKDIR /app

## initialize gradle
COPY ./gradle ./gradle
COPY ./gradlew .
RUN ./gradlew

## build the fat jar
COPY . .
RUN ./gradlew clean installDist --no-daemon

## pre-build the site
ENV BUILD=TRUE
ENV SERVE=FALSE
ENV SOURCE_DIR=./pages
ENV SERVE_DIR=./out
RUN ./build/install/app/bin/app

# RELEASE STAGE
FROM amazoncorretto:21 AS release-stage
WORKDIR /app

## copy over the pre-built site
COPY --from=build-stage /app/out/ ./out/
COPY --from=build-stage /app/build/install/app ./

## serve
EXPOSE 7070
ENV PORT=7070
ENV BUILD=FALSE
ENV SERVE=TRUE
CMD ["./bin/app"]
