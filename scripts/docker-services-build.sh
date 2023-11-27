#!/bin/bash

for dockerfile in build/package/micros/*/Dockerfile; do
    image_name=registry.musicon.space/$(basename $(dirname $dockerfile)):test
    docker build -t $image_name -f $dockerfile .
    if [ $? -ne "0" ]; then
        echo "Error while building $image_name image"
        exit 1
    fi
done
