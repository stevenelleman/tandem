module sg/services/public-api

go 1.13

require (
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/google/uuid v1.1.2
	github.com/gorilla/schema v1.1.0
	github.com/karrick/godirwalk v1.16.1 // indirect
	github.com/lib/pq v1.7.0
	github.com/sirupsen/logrus v1.6.0 // indirect
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	golang.org/x/sys v0.0.0-20200905004654-be1d3432aa8f // indirect

	libraries v0.0.0
)

replace libraries v0.0.0 => ../../libraries/golang
