
GOPATH:=$(shell go env GOPATH)
MODIFY=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto

.PHONY: proto
proto:
    
	protoc --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. proto/ProductWeb/ProductWeb.proto
    

.PHONY: build
build: proto

	go build -o ProductWeb-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	go build -o ProductWeb main.go
	docker build -t ponywilliam/product-web .
	docker tag ponywilliam/product-web ponywilliam/product-web
	docker push ponywilliam/product-web