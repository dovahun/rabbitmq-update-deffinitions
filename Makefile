.PHONY: help
help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: run
run: ## запуск скрипта по генерации файла с дифинишинами
	./.scripts/create_definitions.sh -v values-live.yaml -d ci/ -f definitions.json

.PHONY: view
view: ##отображение файла с дифинишинами
	@cat definitions.json

all: run view ##Создание файла с дифинишинами и ввывод на экран содержимого файла