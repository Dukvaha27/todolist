package config

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	godotenv.Load()

	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  os.Getenv("DB_URI"),
			PreferSimpleProtocol: true,
		}), &gorm.Config{})
	if err != nil {
		panic("не удалось подключиться к БД: " + err.Error())
	}
	return db
}
