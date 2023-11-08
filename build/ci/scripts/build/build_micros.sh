#!/bin/bash

if [ ! -d "/tmp/images/" ]; then
    mkdir /tmp/images/
fi

for dockerfile in build/package/micros/*/Dockerfile; do
    image_name=$(basename $(dirname $dockerfile))
    docker build -t $image_name -f $dockerfile --output type=tar,dest=/tmp/images/$image_name.tar .
    if [ $? -ne "0" ]; then
        echo "Error while building $image_name image"
        exit 1
    fi
done
echo "Built micros images successfully"