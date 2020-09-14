module sg/services/sourcegraph

go 1.13

require (
	github.com/stretchr/testify v1.5.1
	google.golang.org/grpc v1.32.0

	libraries v0.0.0
)

replace libraries v0.0.0 => ../../libraries/golang
