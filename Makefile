.PHONY: unit-test
		integration-test
		mocks-gen
		mocks-clean
		kuber-deploy
		build-images
		remove-images

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

# билдит базовый образ и образы микросервисов с тегом latest
build-images:
	$(CURDIR)/scripts/docker-base-build.sh
	$(CURDIR)/scripts/docker-services-build.sh $(args)

# удаляет все образы репозитория registry.musicon.space с тегом latest
remove-images:
	$(CURDIR)/scripts/docker-remove-images.sh $(args)

dev-up:
	make build-images arg1=dev
	docker compose -f deployments/dev/docker-compose.yml up -d

dev-down:
	docker compose -f deployments/dev/docker-compose.yml down
	make remove-images arg1=dev

# запуск интеграционных тестов
integration-test:
	$(CURDIR)/scripts/docker-base-build.sh
	$(CURDIR)/scripts/docker-services-build.sh test
	$(CURDIR)/scripts/integration-test.sh

# паблиш локальных обарзов и деплой в кубер микросервисов
kuber-deploy:
	$(CURDIR)/scripts/docker-publish.sh
	$(CURDIR)/scripts/kuber-deploy.sh
