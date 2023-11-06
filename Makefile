.PHONY: unit-test
		integration-test
		mocks-clean
		database-clean
		fill-database
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

# запуск сервисов в докере
docker-services-start:
	$(CURDIR)/scripts/docker-services-start.sh

# Удаление продакшен контейнеров и сетей
clean:
	$(CURDIR)/scripts/docker-clean.sh

# Удаление всех контейнеров, образов, сетей и вольюмов на хосте!
hard-clean:
	$(CURDIR)/scripts/docker-hard-clean.sh

# Деплой без очистки данных - мы не теряем созданных пользователей и их лайки
deploy:
	@-make clean
	@make docker-services-start

# Деплой полностью сервиса без пользователей только с музыкой
hard-deploy:
	@-make database-clean
	@make hard-clean
	@make deploy
