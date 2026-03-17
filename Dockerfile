# syntax=docker/dockerfile:1

# BUILD STAGE
FROM rust:1-bookworm AS build
WORKDIR /app
COPY Cargo.toml Cargo.lock ./
COPY src/ ./src/
RUN cargo build --release

# Pre-build site
COPY pages/ ./pages/
ENV BUILD=true SERVE=false
RUN ./target/release/dhcdev

# RELEASE STAGE
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=build /app/out/ ./out/
COPY --from=build /app/target/release/dhcdev ./
EXPOSE 7070
ENV PORT=7070 BUILD=false SERVE=true
CMD ["./dhcdev"]
