package db

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	FirstName string
	LastName  string
	Age       int
	Email     *string
}
