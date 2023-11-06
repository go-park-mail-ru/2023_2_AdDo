#!/bin/bash

REGISTRY=$1
USERNAME=$2
PASSWORD=$3

docker login ${REGISTRY} -u ${USERNAME} -p ${PASSWORD}
make deploy