#!/bin/bash

echo "Preparing environments.."
docker build -t 2023_2_addo-musicon/base -f build/package/base/Dockerfile .
docker compose -f deployments/test/docker-compose.yml up -d
if [ $? -ne 0 ]; then
    echo "Error creating environment"
    docker compose -f deployments/test/docker-compose.yml down
    exit 1
fi

echo "Preparing successful done"

echo "Running tests with empty database..."
for test in test/testsuite/*; do
    python3 $test
    if [ $? -ne 0 ]; then
        echo "Error integration tests"
        docker compose -f deployments/test/docker-compose.yml down
        exit 1
    fi
done

echo "Tests passed"

docker compose -f deployments/test/docker-compose.yml down
if [ $? -ne 0 ]; then
    echo "Error delete test environment"
    exit 1
fi
echo "Test enviroment deleted"
