test:


lint:

build:
	go build -o .idea/app ./cmd/app/main.go
start:
	.idea/app --config configs/config.yml
build-docker:
	@docker build -t image-previewer:latest --target build .

up-server:


down-server:


stop-server:


start-server:
