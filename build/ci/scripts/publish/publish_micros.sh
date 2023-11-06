#!/bin/bash

for arch in /tmp/images/*.tar; do
    image_name=registry.musicon.space/$(basename $arch .tar):latest
    cat $arch | docker import - $image_name
    docker push $image_name
    if [ $? -ne "0" ]; then
        echo "Error while publishing $image_name"
        exit 1
    fi
done