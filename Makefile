.PHONY: docker-service-build unit-test  mocks-clean database-clean docker-service-empty-db-test docker-service-test fill_database docker-service-start clean deploy hard_deploy

#билд контейнера с бэкендом
docker-service-build:
	echo 'Building musicon backend...'
	docker compose build

# Запуск юнит-тестов в файл
unit-test:
	@-make mocks-clean
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
	@-sudo rm -r ~/db-data

# запуск интеграционных тестов
docker-service-test:
	#@make hard_deploy
	@echo "Running tests with empty database..."
	@python3 test/testsuite/sign_up_test.py
	@python3 test/testsuite/music_test.py
	@python3 test/testsuite/me_test.py
	@python3 test/testsuite/logout_test.py
	@python3 test/testsuite/login_test.py
	@python3 test/testsuite/listen_test.py
	@python3 test/testsuite/auth_test.py

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
	@echo "Removing volumes..."
	@-docker  compose down --volumes

# Деплой без очистки данных - мы не теряем созданных пользователей и их лайки
deploy:
	@make clean
	@make docker-service-start

#Деплой полностью сервиса без пользователей только с музыкой
hard_deploy:
	@make database-clean
	@make clean
	@make fill_database
	@make docker-service-start
