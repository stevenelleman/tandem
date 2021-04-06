package constants

// TODO: How to more tightly couple these values to k8s.yaml? It would be idea to have one

const (
	MaxConns                  = 40
	PublicAPIPort             = 8000
	LocalWebFrontendHost      = "http://localhost:3000"
	DockerComposeHost         = "http://localhost"
	TiltHost                  = "http://localhost:8000"
	SGServiceAddress          = "grpc:8001"
	APIStoreDriver            = "postgres"
	APIStorePort              = 5432
	APIStoreUser              = "postgres-user"
	APIStorePassword          = "secret"
	APIStoreName              = "postgresdb"
	APIStoreMigraterTableName = "public-api"
	APIStoreMigrationPath     = "./migrations"
)
