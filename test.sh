#!/bin/bash

PORT=7070

test() {
    docker compose down && \
    docker compose up -d --build && \
    sleep 5 && \
    echo "checking service health at port $PORT..." && \
    curl -s -w %{http_code} -O localhost:$PORT/healthz | grep -q 200 &&
    docker compose down
}

if test; then
    echo "Success"
else
    echo "Failure" && exit 1
fi
