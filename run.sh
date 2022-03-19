#!/bin/bash

# delete 'server' file if exists
if [ -f server ] ; then
 rm server
fi

# export GIN_MODE=release
export GIN_MODE=debug

go build src/server.go 

./server