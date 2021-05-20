module golang

go 1.13

require (
	github.com/Masterminds/squirrel v1.4.0
	github.com/gin-gonic/gin v1.6.3
	github.com/gobuffalo/packr/v2 v2.8.0
	github.com/golang/protobuf v1.4.1
	github.com/google/uuid v1.0.0
	github.com/lib/pq v1.2.0
	github.com/rubenv/sql-migrate v0.0.0-20200616145509-8d140a17f351

	github.com/shell/libraries/golang v0.0.0
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/gorp.v2 v2.2.0
)

// NOTE: there is no need to change directory in path (`shell`) when forked for new project
replace github.com/shell/libraries/golang v0.0.0 => ./
