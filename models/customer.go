package models

import (
	"gorm.io/gorm"
)

type Customers struct {
	ID uint `gorm:"primary key;autoIncrement" json:"id"`

	Firstname *string `json:"firstname"`

	Lastname *string `json:"lastname"`

	Birthdate *string `json:"birthdate"`

	Gender *string `json:"gender"`

	Email *string `json:"email"`

	Address *string `json:"address"`
}

func MigrateBooks(db *gorm.DB) error {

	err := db.AutoMigrate(&Customers{})

	return err

}
