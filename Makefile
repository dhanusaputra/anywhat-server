help: 
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

proto: ## generate proto
	./third_party/protoc-gen.sh

mock: ## generate mocks from mockery
	mockery --all

gqlgen: ## generate gqlgen
	go run github.com/99designs/gqlgen generate

.DEFAULT_GOAL := help