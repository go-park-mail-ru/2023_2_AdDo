.PHONY: unit-test
		integration-test
		mocks-gen
		mocks-clean
		clean

# Запуск юнит-тестов в файл
unit-test:
	$(CURDIR)/scripts/mocks-clean.sh
	$(CURDIR)/scripts/unit-test.sh

# генерация моков
mocks-gen:
	$(CURDIR)/scripts/mocks-gen.sh

# удаление моков
mocks-clean:
	$(CURDIR)/scripts/mocks-clean.sh

# запуск интеграционных тестов
integration-test:
	$(CURDIR)/scripts/docker-base-build.sh
	$(CURDIR)/scripts/docker-services-build.sh test
	$(CURDIR)/scripts/integration-test.sh

# Удаление всех контейнеров, образов, сетей и вольюмов на хосте!
hard-clean:
	$(CURDIR)/scripts/docker-hard-clean.sh
