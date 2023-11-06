#!/bin/bash

echo "Preparing environments.."

docker build -t musicon-base -f build/package/base/Dockerfile .
if [ $? -ne 0 ]; then
    echo "Error while building musicon-base image"
    exit 1
fi

for dockerfile in build/package/micros/*/Dockerfile; do
    image_name=$(basename $(dirname $dockerfile))
    docker build -t $image_name -f $dockerfile .
    if [ $? -ne "0" ]; then
        echo "Error while building $image_name image"
        exit 1
    fi
done
echo "Built micros images successfully"

docker compose -f deployments/test/docker-compose.yml up -d
if [ $? -ne 0 ]; then
    echo "Error creating environment"
    docker compose -f deployments/test/docker-compose.yml down
    exit 1
fi
echo "Preparing successful done"

echo "Running tests with empty database..."

python3 test/testsuite/run_test.py
if [ $? -ne 0 ]; then
    echo "Integration test failed"
    exit 1
fi
echo "Tests passed"

rm -r test/testsuite/__pycache__
docker compose -f deployments/test/docker-compose.yml down
if [ $? -ne 0 ]; then
    echo "Error delete test environment"
    exit 1
fi
echo "Test enviroment deleted"
