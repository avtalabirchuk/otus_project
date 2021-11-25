test:
	go test -race -v `go list ./...`

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.37.0

lint: install-lint-deps
	golangci-lint run ./... -v

build:
	go build -o .idea/app ./cmd/app/main.go

start:
	.idea/app --config configs/config.yml

docker-run:
	@docker build -t image-previewer:latest --target build .
	@docker run -d -p 8080:8080 --name image-previewer image-previewer:latest
up-server:


down-server:


stop-server:


start-server:
