all: clean build test bench coverage

clean:
	go clean -testcache -cache
	rm -f demo/ingest
	rm -f demo/query

build:
	go build
	go build -o demo/ingest cli/ingest/main.go
	go build -o demo/query  cli/query/main.go

test:
	go test ./...

bench:
	go test -run=Benchmark -bench=. -benchmem

coverage:
	go test -cover -coverprofile=coverage.out
	go tool cover -func=coverage.out


