#!/bin/bash

docker compose -f deployments/prod/docker-compose.yml down
if [ $? -ne 0 ]; then
    echo "Error while execute 'docker-compose -f deployments/prod/docker-compose.yml down'"
    exit 1
fi
