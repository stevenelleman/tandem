GO ?= go

# NOTE: `*** missing separator.  Stop.` may be an indication of spaces being used, rather than tabs

# Script to make branch against Tandem
checkout-against-tandem:
	./scripts/checkout-against-tandem.sh

# Go Vendor Commands
deps: deps-ts-services deps-go-services

# Run ts dependencies to account for local ts cross-dependencies where order matters
deps-ts-services: deps-ts-libraries deps-ts-libraries deps-web-frontend

deps-go-services: deps-golang deps-public-api deps-grpc

deps-ts-libraries:
	for pkg in ./libraries/typescript/*; do \
		if [ -d $$pkg ]; then \
			cd $$pkg; yarn install --force; yarn build; cd ../../../; \
		fi; \
	done;

deps-golang:
	cd ./libraries/golang; ${GO} mod tidy -v; ${GO} mod vendor -v;

deps-public-api:
	cd ./services/public-api; ${GO} mod tidy -v; ${GO} mod vendor -v;

deps-grpc:
	cd ./services/grpc; ${GO} mod tidy -v; ${GO} mod vendor -v;

deps-web-frontend:
	@echo
	@echo "Beware: typescript dependency errors most likely related to cross-dependency(ies) in local ts libraries"
	@echo "Order matters -- local ts libraries without local dependencies should be initialized before libraries that do"
	@echo
	cd ./services/web-frontend/internal; yarn install --force;

# Service Builds
build-all: build-grpc build-public-api build-web-frontend

build-grpc:
	cd ./services/grpc; go build -o ../../builds/main-grpc ./cmd/main.go

build-public-api:
	cd ./services/public-api; go build -o ../../builds/main-public-api ./cmd/main.go

build-web-frontend:
	cd ./services/web-frontend/internal; yarn run build && mv ./build ../../../builds/web-frontend-build