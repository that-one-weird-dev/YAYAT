package config

import "os"

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
}
