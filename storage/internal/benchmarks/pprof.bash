#!/bin/bash

NUM_FILES=$1

# Loop over Numfiles int adding string to var
FILES_TO_PPROF="cpu0"

for (( i=1 ; i<$NUM_FILES ; i++ ))
do
	FILES_TO_PPROF="$FILES_TO_PPROF cpu$i"
done

echo $FILES_TO_PPROF

go tool pprof -http ":8000" $FILES_TO_PPROF
