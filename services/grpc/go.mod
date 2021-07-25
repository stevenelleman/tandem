// Note: Use full path to ensure GOPATH is correct
module web.microservice.shell/services/grpc

go 1.16

require (
	github.com/stretchr/testify v1.6.1
	google.golang.org/grpc v1.32.0
	web.microservice.shell/libraries/golang v0.0.0
)

replace web.microservice.shell/libraries/golang v0.0.0 => ../../libraries/golang
