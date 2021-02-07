package migrations

import (
	"github.com/Arpeet-gupta/go-first-api/v4/database"
	"github.com/Arpeet-gupta/go-first-api/v4/models"
)

//Createtable is to create table schema using gorm models
func Createtable() error {
	err := database.Db.AutoMigrate(&models.Author{}, &models.Book{})
	if err != nil {
		return err
	}
	return nil
}
