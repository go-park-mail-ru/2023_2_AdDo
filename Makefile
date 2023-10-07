.PHONY: docker-service-build unit-test  mocks-clean database-clean docker-service-empty-db-test docker-service-test fill_database docker-service-start clean deploy hard_deploy

#билд контейнера с бэкендом
docker-service-build:
	echo 'Building musicon backend...'
	docker compose build

# Запуск юнит-тестов в файл
unit-test:
	@make mocks-clean
	@echo "Generating mocks..."
	@mockgen -source=internal/pkg/user/domain.go -destination=test/mocks/user/user_mock.go -package=user_mock;
	@mockgen -source=internal/pkg/session/domain.go -destination=test/mocks/session/session_mock.go -package=session_mock;
	@mockgen -source=internal/pkg/track/domain.go -destination=test/mocks/track/track_mock.go -package=track_mock;
	@mockgen -source=internal/pkg/artist/domain.go -destination=test/mocks/artist/artist_mock.go -package=artist_mock;
	@mockgen -source=internal/pkg/album/domain.go -destination=test/mocks/album/album_mock.go -package=album_mock;
	@echo "Running unit tests..."
	@go test ./...
	@go test -coverprofile=all_files ./... -coverpkg=./...
	@cat all_files | grep -v "cmd" | grep -v "test" | grep -v "init" > testing_files
	@go tool cover -func=testing_files

#удаление моков
mocks-clean:
	@echo "Mocks deleting..."
	@sudo rm -r test/mocks

#очистка базы данных
database-clean:
	@echo "Database cleaning..."
	@sudo rm -r ~/db-data

#запуск интеграционных тестов на пустой базе данных
docker-service-empty-db-test:
	@make database-clean
	@make docker-service-start
	@echo "Running tests with empty database..."
	@python3 test/testsuite/test_empty.py

#запуск интеграционных тестов на пустой базе данных
docker-service-filled-db-test:
	@make database-clean
	@make fill_database_mock_data
	@make docker-service-start
	@echo "Running tests with filled database..."
	@python3 test/testsuite/test_filled.py

# запуск интеграционных тестов
docker-service-test:
	@make docker-service-empty-db-test

#генерируем скрипт для заполнения базы данных основными данными
fill_database:
	@cd database && python3 fill_db_script.py;

#генерируем скрипт для заполнения базы данных моковыми данными
fill_database_mock_data:
	@cd database && python3 fill_db_script_mock_data.py;

# запуск сервиса в докере
docker-service-start:
	@make fill_database
	@echo "Starting containers..."
	@docker compose up -d

# Остановка контейнеров, удаление контейнеров, образов и сетей
clean:
	@echo "Stopping containers..."
	@-docker stop $$(docker ps -a -q)
	@echo "Removing containers..."
	@-docker rm $$(docker ps -a -q)
	@echo "Removing images..."
	@-docker rmi $$(docker images -q)
	@echo "Removing networks..."
	@-docker network rm $$(docker network ls -q)

# Деплой без очистки данных - мы не теряем созданных пользователей и их лайки
deploy:
	@make clean
	@make docker-service-start

#Деплой полностью сервиса без пользователей только с музыкой
hard_deploy:
	@make clean
	@make database-clean
	@make fill_database
	@make docker-service-start
