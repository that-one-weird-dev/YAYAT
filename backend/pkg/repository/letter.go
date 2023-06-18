package repository

import "letters/pkg/models"

func GetLast10Letters() []models.Letter {
	var letters []models.Letter
	DB.Order("created_at desc").Limit(10).Preload("User").Find(&letters)

	return letters
}

func GetLetterById(id string) models.Letter {
	var letter models.Letter
	DB.Find(&letter, id)

	return letter
}

func CreateLetter(letter *models.Letter) {
	DB.Create(letter)
}
