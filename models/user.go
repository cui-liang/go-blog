package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID      int
	Name    string
	Website string
	Email   string
}

type Auth struct {
	ID       int
	UserID   int
	Source   string
	SourceID string
}
