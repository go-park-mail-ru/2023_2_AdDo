#!/bin/bash

docker build -t registry.musicon.space/base -f build/package/base/Dockerfile .
if [ $? -ne 0 ]; then
    echo "Error while execute 'docker build -t registry.musicon.space/base -f build/package/base/Dockerfile .'"
    exit 1
fi
