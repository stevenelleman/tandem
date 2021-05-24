// Note: Use full path to ensure GOPATH is correct
module web.microservice.shell/services/public-api

go 1.13

require (
	github.com/gin-gonic/gin v1.6.3

	gopkg.in/gorp.v2 v2.2.0
	web.microservice.shell/libraries/golang v0.0.0
)

// TODO: would love to reference local libraries directly rather than vendor changes
replace web.microservice.shell/libraries/golang v0.0.0 => ../../libraries/golang
