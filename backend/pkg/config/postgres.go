package config

var PostgresHost string
var PostgresUser string
var PostgresPassword string
var PostgresDbName string
var PostgresPort string

func initPostgres() {
	PostgresHost = env("POSTGRES_HOST", "localhost")
	PostgresUser = env("POSTGRES_USER", "postgres")
	PostgresPassword = env("POSTGRES_PASSWORD", "secretpassword")
	PostgresDbName = env("POSTGRES_DB_NAME", "postgres")
	PostgresPort = env("POSTGRES_PORT", "5432")
}
