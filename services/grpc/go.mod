module sg/services/sg

go 1.13

require (
	github.com/stretchr/testify v1.6.1
	google.golang.org/grpc v1.32.0

	libraries/golang v0.0.0
)

replace libraries/golang v0.0.0 => ../../libraries/golang
