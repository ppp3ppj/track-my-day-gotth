.PHONY: dev
dev: ## run dev
	go build -o ./tmp/$(APP_NAME) ./cmd/$(APP_NAME)/main.go && air -c .air.dev.toml
