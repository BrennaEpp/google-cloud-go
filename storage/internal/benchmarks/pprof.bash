#!/bin/bash

NUM_FILES=$1

# Loop over Numfiles int adding string to var
FILES_TO_PPROF="cpu0"

#go tool pprof -http ":8000" $FILES_TO_PPROF

#finish this file
#test this run with a ocuple samples+read
#vendor mods to change pprfo sample rate for branch too
#create branch and deployment
#send stats to a diff metric
#enabled by default