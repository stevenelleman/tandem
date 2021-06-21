GO ?= go

# Go Vendor Commands
vendor-all: vendor-golang vendor-public-api vendor-grpc

vendor-golang:
	cd ./libraries/golang; ${GO} mod tidy -v; ${GO} mod vendor -v;

vendor-public-api:
	cd ./services/public-api; ${GO} mod tidy -v; ${GO} mod vendor -v;

vendor-grpc:
	cd ./services/grpc; ${GO} mod tidy -v; ${GO} mod vendor -v;

# Service Builds
build-all: build-grpc build-public-api build-web-frontend

build-grpc:
	cd ./services/grpc; go build -o ../../builds/main-grpc ./cmd/main.go

build-public-api:
	cd ./services/public-api; go build -o ../../builds/main-public-api ./cmd/main.go

build-web-frontend:
	cd ./services/web-frontend/internal; yarn run build && mv ./build ../../../builds/web-frontend-build