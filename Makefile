.PHONY: unit-test
		integration-test
		mocks-clean
		database-clean
		fill_database
		docker_service_build
		docker_service_start
		clean
		deploy
		hard_deploy

# Запуск юнит-тестов в файл
unit-test:
	@make mocks-clean
	$(CURDIR)/scripts/unit-test.sh

# удаление моков
mocks-clean:
	$(CURDIR)/scripts/mocks-clean.sh

# очистка базы данных
database-clean:
	$(CURDIR)/scripts/database-clean.sh

# запуск интеграционных тестов
integration-test:
	$(CURDIR)/scripts/integration-test.sh

# генерируем скрипт для заполнения базы данных основными данными
fill_database:
	@cd db && python3 fill_db_script.py;

#генерируем скрипт для заполнения базы данных моковыми данными
fill_database_mock_data:
	@cd database && python3 fill_db_script_mock_data.py;

# билд контейнера с бэкендом
docker_services_build: 
	$(CURDIR)/scripts/docker-services-build.sh

# запуск сервисов в докере
docker_services_start:
	$(CURDIR)/scripts/docker-services-start.sh

# Остановка контейнеров, удаление контейнеров, образов и сетей
clean:
	$(CURDIR)/scripts/docker-clean.sh

# Деплой без очистки данных - мы не теряем созданных пользователей и их лайки
deploy:
	@make clean
	@make fill_database
	@make docker_services_start

# Деплой полностью сервиса без пользователей только с музыкой
hard_deploy:
	@make database-clean
	@make deploy
