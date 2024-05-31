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
        google.golang.org/grpc/cmd/protoc-gen-go-grpc \
        github.com/vektra/mockery/v2

generate-mock:
	rm -rf mocks
	mockery

build:
	go build -o bin/laklak main.go

run: build
	OTEL_EXPORTER_OTLP_INSECURE=true ./bin/laklak serve

migrate: build
	./bin/laklak migrate

compose-up:
	JAEGER_VERSION=1.52 docker compose up -d

test:
	go clean -testcache
	go test -cover ./...
