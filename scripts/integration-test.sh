#!/bin/bash

echo "Preparing environments.."
docker compose -f deployments/test/docker-compose.yml up -d
if [ $? -ne 0 ]; then
    echo "Error creating environment"
    docker compose -f deployments/test/docker-compose.yml down
    exit 1
fi
echo "Preparing successful done"

echo "Running tests with empty database..."

cd test/testsuite && python3 run_test.py
if [ $? -ne 0 ]; then
    echo "Integration test failed"
    exit 1
fi
echo "Tests passed"

rm -r __pycache__ && cd ../..
docker compose -f deployments/test/docker-compose.yml down
if [ $? -ne 0 ]; then
    echo "Error delete test environment"
    exit 1
fi
echo "Test enviroment deleted"
