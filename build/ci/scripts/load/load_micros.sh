#!/bin/bash

image_tag=$1
registry=registry.musicon.space
for arch in /tmp/images/*.tar; do
    image_name=$(basename $arch .tar)
    cat $arch | docker import - $registry/$image_name:$image_tag
    if [ $? -ne "0" ]; then
        echo "Error while loading $registry/$image_name:$image_tag"
        exit 1
    fi
done
