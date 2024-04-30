package db

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name       string
	Price      uint
	CreateById uint
	CreateBy   Admin
}
