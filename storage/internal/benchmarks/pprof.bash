#!/bin/bash

NUM_FILES=$1
START_AT=$(($2+0))

# Loop over Numfiles int adding string to var
FILES_TO_PPROF="cpu$START_AT"

for (( i=$START_AT+1 ; i<$NUM_FILES+$START_AT ; i++ ))
do
	FILES_TO_PPROF="$FILES_TO_PPROF cpu$i"
done

echo $FILES_TO_PPROF

go tool pprof -http ":8000" $FILES_TO_PPROF
