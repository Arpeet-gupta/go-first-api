package models

import (
	"gorm.io/gorm"
)

//Author `has many` relation with Book
type Author struct {
	gorm.Model
	FullName string `json:"fullname"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	AllBook  []Book `json:"books" gorm:"constraint:OnDelete:CASCADE"`
}

//Book `has a` relation with the Author
type Book struct {
	gorm.Model
	Name      string `json:"name"`
	Publisher string `json:"publisher"`
	AuthorID  uint
}
