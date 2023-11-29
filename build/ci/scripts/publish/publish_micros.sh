#!/bin/bash

image_tag=$1
registry=registry.musicon.space
for arch in /tmp/images/*.tar; do
    image_name=$(basename $arch .tar)
    image=$registry/$image_name
    cat $arch | docker import - $image:$image_tag
    cat $arch | docker import - $image:latest
    docker push $image --all-tags
    if [ $? -ne "0" ]; then
        echo "Error while publishing $image"
        exit 1
    fi
done