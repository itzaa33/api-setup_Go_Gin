package handler

import "gorm.io/gorm"

type CustomerHandler struct {
	Database *gorm.DB
}
