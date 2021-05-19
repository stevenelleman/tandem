module public-api

go 1.13

require (
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/web-microservice-shell/libraries/golang v0.0.0

	github.com/web-microservice-shell/services/public-api v0.0.0
	gopkg.in/gorp.v2 v2.2.0
)

replace (
	github.com/web-microservice-shell/libraries/golang v0.0.0 => ../../libraries/golang
	github.com/web-microservice-shell/services/public-api v0.0.0 => ./
)
