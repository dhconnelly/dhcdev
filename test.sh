#!/bin/bash

PORT=7000

test() {
    docker-compose down && \
    docker-compose up -d --build && \
    sleep 3 && \
    echo "checking service health at port $PORT..." && \
    curl localhost:$PORT && \
    docker-compose down
}

if test; then
    echo "Success"
else
    echo "Failure" && exit 1
fi
