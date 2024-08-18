GOBIN ?= $$(go env GOPATH)/bin

.PHONY: test
test:
	go clean -testcache && go test -v -cover -short ./...

.PHONY: fmt
fmt:
	@echo "### Formatting the source code ###"
	go fmt ./...

.PHONY: vet
vet:
	@echo "### Checking for code issues ###"
	go vet ./...

.PHONY: install-go-test-coverage
install-go-test-coverage:
	go install github.com/vladopajic/go-test-coverage/v2@latest

.PHONY: check-coverage
check-coverage: install-go-test-coverage
	go test ./... -coverprofile=./cover.out -covermode=atomic -coverpkg=./...
	${GOBIN}/go-test-coverage --config=./.testcoverage.yml