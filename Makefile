.SILENT:

.PHONY:test
test:
	go test -cover -bench=.

.PHONY:lint
lint:
	golangci-lint run -c ./.golangci.yml > lint.txt
