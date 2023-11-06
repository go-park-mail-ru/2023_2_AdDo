#!/bin/bash

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
