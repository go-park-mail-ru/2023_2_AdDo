#!/bin/bash

image_tag=$1
registry=registry.musicon.space
for arch in /tmp/images/*.tar; do
    image_name=$(basename $arch .tar)
    image=$registry/$image_name:$image_tag
    cat $arch | docker import - $image
    docker push $image
    if [ $? -ne "0" ]; then
        echo "Error while publishing $image"
        exit 1
    fi
done