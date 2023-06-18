package repository

import (
	"letters/pkg/models"
)

func GetUserById(id string) models.User {
	var user models.User
	DB.Find(&user, id)

	return user
}

func CreateUser(user *models.User) {
	DB.Create(user)
}

func FindUserByTag(tag string) (*models.User, error) {
	var user models.User

	result := DB.Where("Tag", &tag).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func GetUserByTagWithLetters(tag string) models.User {
	var user models.User
	DB.Order("created_at desc").Preload("Letters").Limit(10).Where("Tag", &tag).Find(&user)

	return user
}
