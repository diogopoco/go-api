package entity

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title       string
	Descriptiom string
	CreatedBy   uint
}