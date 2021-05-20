module public-api

go 1.13

require (
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/shell/libraries/golang v0.0.0

	github.com/shell/services/public-api v0.0.0
	gopkg.in/gorp.v2 v2.2.0
)

replace (
    // TODO: would love to reference local libraries directly rather than vendor changes
	github.com/shell/libraries/golang v0.0.0 => ../../libraries/golang
	github.com/shell/services/public-api v0.0.0 => ./
)
