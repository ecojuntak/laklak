generate:
	rm -rf gen/**
	cd proto; buf dep update
	buf build
	buf generate

install-dependencies:
	go install \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
        google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc

build:
	go build -o bin/laklak main.go

run: build
	./bin/laklak serve

migrate: build
	./bin/laklak migrate

compose-up:
	docker compose up -d
