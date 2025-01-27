all: deps check build

format:
	find . -iname "*.go" -exec gofmt -s -l -w {} \;

check:
	go vet ./...

run:
	go run cmd/surl/*.go

deps:
	go mod tidy -v

build:
	go build -trimpath -ldflags "-s -w" cmd/surl/*.go
	# go build -trimpath -ldflags "-s -w -X main.version=`git tag --sort=-version:refname | head -n 1`" cmd/surl/*.go
