#!/bin/bash

echo 'Building musicon backend...'
docker build -t musicon-base -f build/package/base/Dockerfile .
docker compose -f deployments/prod/docker-compose.yml build
echo 'Musicon backend built successful'
