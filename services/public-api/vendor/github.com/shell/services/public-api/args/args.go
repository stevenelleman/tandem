package args

import (
	"os"
	"strconv"

	"github.com/shell/libraries/golang/datastore/migrater"
	"github.com/shell/libraries/golang/guts/connection/service/psql_conn"
	"github.com/shell/libraries/golang/guts/connection/service/sg_conn"
	"github.com/shell/services/public-api/mapper"
)

const (
	// Environment Variables
	psqlPort = "PSQL_PORT"
	psqlDriver = "PSQL_DRIVER"
	psqlUser = "PSQL_USER"
	psqlSecret = "PSQL_SECRET"
	psqlName = "PSQL_NAME"
	psqlMigraterTableName = "PSQL_MIGRATER_TABLE_NAME"
	psqlMigrationPath = "PSQL_MIGRATION_PATH"
	grpcAddress = "GPRC_ADDRESS"

	// Constants
	maxConns = 40
)

func MakeArgsFromEnv(store string) (*psql_conn.StoreArgs, *sg_conn.StoreArgs) {
	envPort := os.Getenv(psqlPort)
	port, err := strconv.Atoi(envPort)
	if err != nil {
		panic("port is invalid")
	}

	driver := os.Getenv(psqlDriver)
	if driver == "" {
		panic("driver is invalid")
	}

	user := os.Getenv(psqlUser)
	if user == "" {
		panic("user is invalid")
	}

	secret := os.Getenv(psqlSecret)
	if secret == "" {
		panic("secret is invalid")
	}

	name := os.Getenv(psqlName)
	if name == "" {
		panic("name is invalid")
	}

	migraterTableName := os.Getenv(psqlMigraterTableName)
	if migraterTableName == "" {
		panic("migraterTableName is invalid")
	}

	migrationPath := os.Getenv(psqlMigrationPath)
	if migrationPath == "" {
		panic("migrationPath is invalid")
	}

	address := os.Getenv(grpcAddress)
	if address == "" {
		panic("address is invalid")
	}


	migraterArgs := migrater.MakeArgs(
		driver,
		migraterTableName,
		migrationPath,
	)

	psqlArgs := psql_conn.MakeArgs(
		driver,
		store,
		user,
		secret,
		name,
		port,
		maxConns,
		migraterArgs,
		mapper.MapTables,
	)

	sgArgs := sg_conn.MakeArgs(address)

	return psqlArgs, sgArgs
}
