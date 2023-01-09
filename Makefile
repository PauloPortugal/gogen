.PHONY: audit
audit:
	go list -m all | nancy sleuth

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: coverage
coverage:
	go clean -testcache && go test -race -covermode=atomic -coverprofile=coverage.out ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test:
	go clean -testcache && go test --cover -v ./...