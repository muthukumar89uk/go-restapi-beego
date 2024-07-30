package repository

import (
	"fmt"
	"testBeego/helpers"
	"testBeego/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection() *gorm.DB {
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s  sslmode=disable", helpers.Host, helpers.Port, helpers.User, helpers.Password, helpers.Dbname)

	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		fmt.Println("Db Connection is failed", err)
		return nil
	}

	fmt.Println("DB Drive Connected")

	return db
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(&models.Employee{}, &models.Address{})
}
