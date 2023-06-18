package config

import "time"

var Secret string

const TokenDuration = time.Hour * 24

func initAuth() {
	Secret = env("SECRET", "3a862028-87f1-4ade-8034-c1cf683f7aa4")
}
