build:
	go build -tags cache cache.go
	go build -o bin/server ./cmd/

run:
	./bin/server