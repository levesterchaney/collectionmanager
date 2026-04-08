run:
	go run cmd/api/main.go

build:
	go build -o bin/api cmd/api/main.go

deps:
	go mod download

clean:
	rm -rf bin/

.PHONY: run build deps clean