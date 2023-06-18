package utils

import (
	"letters/pkg/config"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

func NewJwt(id uint, tag string) (string, error) {
	claims := jwt.MapClaims{
		"id":  strconv.Itoa(int(id)),
		"tag": tag,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.Secret))
}
