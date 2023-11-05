#!/bin/bash

echo "Preparing environments.."
docker build -t musicon-base -f build/package/base/Dockerfile . \
    && docker build -t user -f build/package/user/Dockerfile . \
    && docker build -t album -f build/package/album/Dockerfile . \
    && docker build -t artist -f build/package/artist/Dockerfile . \
    && docker build -t entrypoint -f build/package/entrypoint/Dockerfile . \
    && docker build -t images -f build/package/images/Dockerfile . \
    && docker build -t playlist -f build/package/playlist/Dockerfile . \
    && docker build -t session -f build/package/session/Dockerfile . \
    && docker build -t track -f build/package/track/Dockerfile . \

if [ $? -ne 0 ]; then
    echo "Error build images"
    exit 1
fi

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
