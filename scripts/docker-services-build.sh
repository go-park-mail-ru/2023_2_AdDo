#!/bin/bash

echo 'Building musicon backend...'
docker build -t 2023_2_addo-musicon/user -f build/package/user/Dockerfile .
docker compose -f deployments/prod/docker-compose.yml build
echo 'Musicon backend built successful'
