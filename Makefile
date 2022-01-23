.PHONY: all run clean build test

all: assemble

### Targets for developers
lint:
	@echo "\nApplying golint\n"
	@golint ./...

integration-test:
	@echo "\nRunning integration tests\n"
	@go test -cover -run Integration ./...

unit-test:
	@echo "\nRunning unit tests\n"
	@go test -cover -short ./...

test: unit-test integration-test
	@echo "\nRunning tests\n"

### Target for Docker container
build:
	@echo "\nBuilding application"
	@go build -o application cmd/main.go

### Targets for users
assemble: clean
	@echo "\nBuilding docker image"
	@docker build --tag recipe-aggregator .

install:
	@echo "\nCreating executable"
	@$(shell scripts/link_script.sh)
	@echo "\nGiven execution permission to executable file"
	@chmod +x recipe-aggregator
	@echo "\nAll setup\nRecipe-aggregator is ready\n ./recipe-aggregator to use"

clean:
	@echo "\nRemove recipe-aggregator executable"
	-@rm recipe-aggregator 2>/dev/null || echo "\nExecutable file recipe-aggregator not found to remove"
	@echo "\nRemove recipe-aggregator image"
	-@docker rmi recipe-aggregator 2>/dev/null || echo "\nDocker image recipe-aggregator not found to remove"
