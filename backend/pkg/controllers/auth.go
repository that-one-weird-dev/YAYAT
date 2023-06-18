package controllers

import (
	"letters/pkg/repository"
	"letters/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type LoginBody struct {
	Tag      string `json:"tag"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func PostLogin(c *fiber.Ctx) error {
	body := new(LoginBody)

	err := c.BodyParser(body)
	if err != nil {
		return err
	}

	user, err := repository.FindUserByTag(body.Tag)
	if err != nil {
		return err
	}

	if !utils.CheckPasswordHash(body.Password, user.Password) {
		c.Status(fiber.StatusUnauthorized)
		return nil
	}

	tokenString, err := utils.NewJwt(user.ID, user.Tag)
	if err != nil {
		return err
	}

	return c.JSON(LoginResponse{
		Token: tokenString,
	})
}
