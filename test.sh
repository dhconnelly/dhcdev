#!/bin/bash

test() {
    docker-compose down && \
    docker-compose up -d --build && \
    sleep 3 && \
    echo "checking service health..." && \
    curl -s -w %{http_code} -O localhost:8080/healthz | grep -q 200 && \
    docker-compose down
}

if test; then
    echo "Success"
else
    echo "Failure" && exit 1
fi
