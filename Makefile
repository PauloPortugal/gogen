.PHONY: audit
audit:
	go list -m all | nancy sleuth

.PHONY: test
test:
	go clean -testcache && go fmt ./... && go vet ./... && go test --cover -v ./...