#!/bin/bash

docker compose -f deployments/prod/docker-compose.yml down
if [ $? -ne 0 ]; then
    echo "Error while execute 'docker-compose -f deployments/prod/docker-compose.yml down'"
    exit 1
fi

# shellcheck disable=SC2046
docker stop $(docker ps -a -q) -f
# shellcheck disable=SC2046
docker rm $(docker ps -a -q) -f
# shellcheck disable=SC2046
docker rmi $(docker images -q) -f
# shellcheck disable=SC2046
docker volume rm $(docker volume ls -q) -f
# shellcheck disable=SC2046
docker network rm $(docker network ls -q) -f
echo "Successful cleaning"
