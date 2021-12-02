test-all:
	go test -v ./...

test:
	go test ./$(day)

.PHONY: test-all
