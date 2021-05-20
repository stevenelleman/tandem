package constants

const (
	MaxConns                  = 40
	PublicAPIPort             = 8000
	LocalWebFrontendHost      = "http://localhost:3000"
	DockerComposeHost         = "http://localhost"
	TiltHost                  = "http://www.grouphouse.io"
	SGServiceAddress          = "grpc:8001"
	APIStoreDriver            = "postgres"
	APIStorePort              = 5432
	APIStoreUser              = "postgres-user"
	APIStorePassword          = "secret"
	APIStoreName              = "postgresdb"
	APIStoreMigraterTableName = "public-api"
	APIStoreMigrationPath     = "./migrations"
)