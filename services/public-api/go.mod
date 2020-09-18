module sg/services/public-api

go 1.13

require (
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	sg/libraries/golang v0.0.0
)

// HERE: Need to make consistent with path relative to project root
replace sg/libraries/golang v0.0.0 => ../../../sg/libraries/golang
