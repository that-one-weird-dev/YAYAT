package controllers

import (
	"letters/pkg/models"
	"letters/pkg/repository"
	"letters/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type RegisterUserBody struct {
	DisplayName string `json:"displayName"`
	Tag         string `json:"tag"`
	Password    string `json:"password"`
}

type RegisterUserResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	Tag         string `json:"tag"`
	DisplayName string `json:"displayName"`
}

type UserIdResponse struct {
	Tag         string           `json:"tag"`
	DisplayName string           `json:"displayName"`
	Letters     []LetterResponse `json:"letters"`
}

func PostRegisterUser(c *fiber.Ctx) error {
	body := new(RegisterUserBody)

	err := c.BodyParser(body)
	if err != nil {
		return err
	}

	password, err := utils.HashPassword(body.Password)
	if err != nil {
		return err
	}

	user := models.User{
		DisplayName: body.DisplayName,
		Tag:         body.Tag,
		Password:    password,
	}

	repository.CreateUser(&user)

	tokenString, err := utils.NewJwt(user.ID, user.Tag)
	if err != nil {
		return err
	}

	response := RegisterUserResponse{
		Token: tokenString,
	}

	return c.JSON(response)
}

func GetUser(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	user := repository.GetUserById(userId)

	response := UserResponse{
		Tag:         user.Tag,
		DisplayName: user.DisplayName,
	}

	return c.JSON(response)
}

func GetUserId(c *fiber.Ctx) error {
	id := c.Params("id")

	user := repository.GetUserByTagWithLetters(id)

	responseLetters := make([]LetterResponse, len(user.Letters))

	for i, l := range user.Letters {
		responseLetters[i] = toLetterResponse(&l)
	}

	response := UserIdResponse{
		DisplayName: user.DisplayName,
		Tag:         user.Tag,
		Letters:     responseLetters,
	}

	return c.JSON(response)
}
