#!/bin/bash

echo "Starting containers..."
docker build -t 2023_2_addo-musicon/user -f build/package/user/Dockerfile .
docker compose -f deployments/prod/docker-compose.yml up -d

if [ $? -ne 0 ]; then
    echo "Error while starting containers"
    exit 1
fi

echo "Containers started successful"
