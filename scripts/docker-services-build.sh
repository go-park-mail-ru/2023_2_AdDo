#!/bin/bash

echo 'Building musicon backend...' 	
docker compose -f deployments/prod/docker-compose.yml build
echo 'Musicon backend built successful'
