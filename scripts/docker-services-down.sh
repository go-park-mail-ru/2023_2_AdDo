#!/bin/bash

docker compose -f deployments/test/docker-compose.yml down
if [ $? -ne 0 ]; then
    echo "Error execute "docker compose -f deployments/test/docker-compose.yml down""
    exit 1
fi
