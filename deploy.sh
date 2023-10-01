# shellcheck disable=SC2046
docker stop $(docker ps -a -q);
docker rm $(docker ps -a -q);
docker rmi $(docker images -q);
docker network rm $(docker network ls -q);
cd database && python3 fill_db_script.py;
cd .. && docker compose up -d;
