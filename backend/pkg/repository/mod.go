package repository

import (
	"fmt"
	"letters/pkg/config"
	"letters/pkg/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConnection() {
	var err error

	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", config.PostgresHost, config.PostgresUser, config.PostgresPassword, config.PostgresDbName, config.PostgresPort)

	println("Connecting to db with: " + connString)

	DB, err = newDbConnection(connString, 10)
	if err != nil {
		panic(err.Error())
	}

	models.AutoMigrate(DB)
}

func newDbConnection(connString string, retries int) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	db, err = gorm.Open(postgres.Open(connString), &gorm.Config{})

	for err != nil {
		println("Database connection failed retrying...")

		retries -= 1
		if retries <= 0 {
			return nil, err
		}

		db, err = gorm.Open(postgres.Open(connString), &gorm.Config{})

		time.Sleep(time.Second * 5)
	}

	return db, nil
}
