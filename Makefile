default: help

help:
	@echo 
	@echo "available commands"
	@echo "  - dev #live reload"
	@echo "  - build #outputs the application binary"
	@echo "  - docker-build"
	@echo "  - docker-run"
	@echo "  - docker-push"
	@echo "  - clean #deletes the application binary"
	@echo

BIN=server
GO=CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go
GO_BUILD_COMMAND=$(GO) build -a --ldflags '-X main.VERSION=$(TAG) -w -extldflags "-static"' -tags netgo -o $(BIN) ./

TAG=0.1.0
IMAGE=devbytom/binance-wallet

.PHONY: dev
dev:
	go mod download
	go mod vendor
	air .

.PHONY: build
build:
	$(GO_BUILD_COMMAND)

.PHONY: docker-build
docker-build:
	docker build \
	-t $(IMAGE):$(TAG) \
	--build-arg TAG=$(TAG) \
	.

.PHONY: docker-run
docker-run: docker-build
	docker run -i \
	--network host \
	--env-file .env \
	devbytom/binance-wallet:$(TAG)

.PHONY: docker-push
docker-push: docker-build
	docker push $(IMAGE):$(TAG)
	docker tag $(IMAGE):$(TAG) $(IMAGE):latest
	docker push $(IMAGE):latest

.PHONY: clean
clean:
	rm $(BIN)
