package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	NikeName string `json:"nike_name"`
	Content string `json:"content"`
}

