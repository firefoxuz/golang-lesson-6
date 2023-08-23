test:
	go test -race -coverprofile="coverage.out" -covermode=atomic ./...
	go tool cover -html="coverage.out"

#
#lint:
#	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
#	golangci-lint run

#lint:
#	@if ! command -v golangci-lint &> /dev/null; then \
#		go get -u github.com/golangci/golangci-lint/cmd/golangci-lint; \
#	fi
#	golangci-lint run

