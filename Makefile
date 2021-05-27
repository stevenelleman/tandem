GO ?= go

# Go Vendor Commands
vendor-all: vendor-golang vendor-public-api vendor-grpc

vendor-golang:
	cd ./libraries/golang; ${GO} mod tidy -v; ${GO} mod vendor -v;

vendor-public-api:
	cd ./services/public-api; ${GO} mod tidy -v; ${GO} mod vendor -v;

vendor-grpc:
	cd ./services/grpc; ${GO} mod tidy -v; ${GO} mod vendor -v;
