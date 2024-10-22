run:
	@ go run cmd/api/main.go

migrate:
	@ go run cmd/migration/main.go

test:
	@ go test -v ./test/...

test-coverage:
	@ go test -coverprofile=coverage.out ./...
	@ go tool cover -html=coverage.out

test-controller:
	@ go test ./test/controller-test -v

.PHONY: run migrate test