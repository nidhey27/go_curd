package models

import "gorm.io/gorm"

type Posts struct {
	gorm.Model
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}
