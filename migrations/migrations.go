package migrations

import (
	"github.com/Arpeet-gupta/go-first-api/v3/database"
	"github.com/Arpeet-gupta/go-first-api/v3/models"
)

//Createtable is to create table schema using gorm models
func Createtable() error {
	// database.Db.SingularTable(true)
	err := database.Db.AutoMigrate(&models.Author{}, &models.Book{}).Error
	if err != nil {
		return err
	}
	return nil
}
