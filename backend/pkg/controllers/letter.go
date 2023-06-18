package controllers

import (
	"letters/pkg/models"
	"letters/pkg/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type LetterBody struct {
	Text string `json:"text"`
}

type LetterResponse struct {
	Text            string `json:"text"`
	PostedAt        string `json:"postedAt"`
	UserDisplayName string `json:"userDisplayName"`
	UserTag         string `json:"userTag"`
}

func toLetterResponse(letter *models.Letter) LetterResponse {
	return LetterResponse{
		Text:            letter.Text,
		UserDisplayName: letter.User.DisplayName,
		UserTag:         letter.User.Tag,
		PostedAt:        letter.CreatedAt.Format("2006-01-02 15:04"),
	}
}

func GetLetters(c *fiber.Ctx) error {
	letters := repository.GetLast10Letters()

	response := make([]LetterResponse, len(letters))

	for i, l := range letters {
		response[i] = toLetterResponse(&l)
	}

	return c.JSON(&response)
}

func GetLetter(c *fiber.Ctx) error {
	id := c.Params("id")
	letter := repository.GetLetterById(id)

	return c.JSON(toLetterResponse(&letter))
}

func PostLetter(c *fiber.Ctx) error {
	body := new(LetterBody)

	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	user := repository.GetUserById(userId)

	letter := models.Letter{
		Text:   body.Text,
		UserId: user.ID,
		User:   user,
	}

	repository.CreateLetter(&letter)
	return c.JSON(toLetterResponse(&letter))
}
