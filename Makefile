GO111MODULES=on

.PHONY: clean
## clean: cleans the library
clean:
	@echo "Cleaning..."
	@go clean ./...

.PHONY: test
## test: runs go test with default values
test:
	go test -v -count=1 -race ./...

.PHONY: setup
## setup: setup go modules
setup:
	@go mod init \
		&& go mod tidy

.PHONY: help
## help: prints this help message
help:
	@echo "\nUsage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
	@echo "\n"
