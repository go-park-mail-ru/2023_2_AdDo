#!/bin/bash

echo "Login to registry.musicon.space.."
docker login registry.musicon.space
if [ $? -ne "0" ]; then
    echo "Error login to registry.musicon.space"
    exit 1
fi
echo "Successful login to registry.musicon.space"

IMAGE_TAG=$1
if [ -z "$IMAGE_TAG"];then
    IMAGE_TAG="latest"
fi

REGISTRY=registry.musicon.space
for image in $(docker images ${REGISTRY}/*:${IMAGE_TAG} --format "{{.Repository}}:{{.Tag}}"); do
    echo "docker push $image"
    docker push $image
    if [ $? -ne "0" ]; then
        echo "Error while publishing $image image"
        exit 1
    fi
done
