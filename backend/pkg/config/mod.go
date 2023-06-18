package config

import "os"

var Environment string
var Port = "3001"

func env(key, fallback string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return fallback
	}
	return value
}

func Init() {
	initAuth()
	initPostgres()

	Environment = env("ENVIRONMENT", "dev")

	if Environment == "prod" {
		Port = "80"
	}
}
