test-all:
	go test -v ./...

test:
	go test -v ./day$(day)

.PHONY: test-all
