module sg/services/sg

go 1.13

require (
	github.com/google/uuid v1.0.0
	github.com/stretchr/testify v1.6.1
	google.golang.org/grpc v1.32.0

	sg/libraries/golang v0.0.0
)

// HERE: Need to make consistent with path relative to project root
replace sg/libraries/golang v0.0.0 => ../../../sg/libraries/golang
