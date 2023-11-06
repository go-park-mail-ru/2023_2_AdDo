#!/bin/bash

for arch in /tmp/images/*.tar; do
    cat $arch | docker import - $(basename $arch .tar):latest
    if [ $? -ne "0" ]; then
        echo "Error while loading $arch"
        exit 1
    fi
done
