test-all:
	go test -v ./...

test:
	go test -failfast -v ./day$(day)

.PHONY: test-all
