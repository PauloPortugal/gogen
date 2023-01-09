.PHONY: audit
audit:
	go list -m all | nancy sleuth

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: coverage
coverage:
	echo "mode: count" > coverage.out \
	go test -v ./... -covermode=count -coverprofile=profile.out $$d > tmp.out; \
	cat tmp.out; \
	if grep -q "^--- FAIL" tmp.out; then \
		rm tmp.out; \
		exit 1; \
	elif grep -q "build failed" tmp.out; then \
		rm tmp.out; \
		exit 1; \
	elif grep -q "setup failed" tmp.out; then \
		rm tmp.out; \
		exit 1; \
	fi; \
	if [ -f profile.out ]; then \
		cat profile.out | grep -v "mode:" >> coverage.out; \
		rm profile.out; \
	fi;

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test:
	go clean -testcache && go test --cover -v ./...