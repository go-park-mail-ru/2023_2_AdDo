#!/bin/bash

docker compose -f deployments/test/docker-compose.yml up -d
python3 test/testsuite/run_test.py
if [ $? -ne "0" ];then
    echo echo "Integration tests failed!"
    exit 1
fi
echo "Integration tests passed!"
docker compose -f deployments/test/docker-compose.yml down