#!/bin/bash

test() {
    docker-compose down && \
    docker-compose up -d --build && \
    sleep 5 && curl localhost:8080 && \
    docker-compose down
}

if test; then
    echo "Success"
else
    echo "Failure" && exit 1
fi
