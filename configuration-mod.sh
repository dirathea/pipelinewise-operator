#!/usr/bin/env bash

CONFIGURATION_PATH=$1
OUTPUT_PATH=$2

mkdir -p $OUTPUT_PATH

for filePath in $CONFIGURATION_PATH/*; do
    key=$(basename $filePath)
    perl -0777 -pe 's/:\s+\|\s+!vault/: !vault/g' $filePath > $OUTPUT_PATH/$key
done