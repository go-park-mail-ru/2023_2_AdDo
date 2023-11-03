.PHONY: unit-test
		integration-test
		mocks-clean
		database-clean
		fill-database
		docker-service-build
		docker-service-start
		clean
		deploy
		hard-deploy

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
fill-database:
	@cd db && python3 fill_db_script.py;

#генерируем скрипт для заполнения базы данных моковыми данными
fill-database-mock-data:
	@cd database && python3 fill-db-script-mock-data.py;

# билд контейнера с бэкендом
docker-services-build:
	$(CURDIR)/scripts/docker-services-build.sh

# запуск сервисов в докере
docker-services-start:
	$(CURDIR)/scripts/docker-services-start.sh

# Остановка контейнеров, удаление контейнеров, образов и сетей
clean:
	$(CURDIR)/scripts/docker-clean.sh

# Деплой без очистки данных - мы не теряем созданных пользователей и их лайки
deploy:
	@make clean
	@make docker-services-build
	@make docker-services-start

# Деплой полностью сервиса без пользователей только с музыкой
hard-deploy:
	@-make database-clean
	@make deploy

registry-start:
	@docker run -d -p 5000:5000 --restart=always --name musicon-registry registry:2
	@DOCKER_OPTS="$DOCKER_OPTS --insecure-registry musicon-registry:5000"
	@service docker restart
