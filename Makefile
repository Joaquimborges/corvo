.PHONY: test
test:
	@go clean -testcache && go test -v -cover -short ./...

.PHONY: format
format:
	@echo "### Formatting the source code ###"
	@go fmt ./...

.PHONY: vet
vet:
	@echo "### Checking for code issues ###"
	@go vet ./...