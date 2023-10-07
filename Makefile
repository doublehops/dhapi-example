
run:
	go run cmd/server/run.go -config ./config.json

gofmt:
	gofumpt -l -w .

lint:
	golangci-lint --config ./ci/.golangci-lint.yml run

test:
	go test ./... -cover
