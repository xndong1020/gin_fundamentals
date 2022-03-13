#!/bin/bash

# delete 'server' file if exists
if [ -f server ] ; then
 rm server
fi

go build src/server.go 

# export GIN_MODE=release
export GIN_MODE=debug

./server