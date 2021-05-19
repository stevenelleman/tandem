module grpc-service

go 1.13

require (
	github.com/stretchr/testify v1.6.1

	github.com/web-microservice-shell/libraries/golang v0.0.0
	github.com/web-microservice-shell/services/grpc v0.0.0
	google.golang.org/grpc v1.32.0
)

replace (
	github.com/web-microservice-shell/libraries/golang v0.0.0 => ../../libraries/golang
	github.com/web-microservice-shell/services/grpc v0.0.0 => ./
)
