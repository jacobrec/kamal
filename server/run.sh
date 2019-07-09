#!/bin/bash

function end() {
    echo "leaning files"
    rm server
}
trap end INT

go build
LOG_LEVEL=INFO CONFIG=../.config PORT=8888 ./server

