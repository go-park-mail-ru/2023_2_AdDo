#!/bin/bash

IMAGE_TAG=$1
if [ -z "$IMAGE_TAG"];then
    IMAGE_TAG="latest"
fi

REGISTRY=registry.musicon.space
for image in $(docker images ${REGISTRY}/*:${IMAGE_TAG} --format "{{.Repository}}:{{.Tag}}"); do
    echo "docker rmi $image"
    docker rmi $image
    if [ $? -ne "0" ]; then
        echo "Error while publishing $image image"
        exit 1
    fi
done