module grpc-service

go 1.13

require (
	github.com/shell/libraries/golang v0.0.0
	github.com/shell/services/grpc v0.0.0
	github.com/stretchr/testify v1.6.1
	google.golang.org/grpc v1.32.0
)

replace (
	github.com/shell/libraries/golang v0.0.0 => ../../libraries/golang
	github.com/shell/services/grpc v0.0.0 => ./
)
