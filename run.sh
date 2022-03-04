#!/bin/bash

go build src/server.go 

# export GIN_MODE=release
export GIN_MODE=debug

./server