// Note: Use full path to ensure GOPATH is correct
module sg/services/public-api

go 1.16

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/stretchr/testify v1.5.1
	gopkg.in/gorp.v2 v2.2.0
	sg/libraries/golang v0.0.0
)

// TODO: would love to reference local libraries directly rather than vendor changes
replace sg/libraries/golang v0.0.0 => ../../libraries/golang
