package models

import "time"

type User struct {
	ID        int
	Username  string
	Password  string
	Email     string
	FullName  string
	CreatedAt time.Time `gorm:"type:DATETIME;autoCreateTime"`
	UpdatedAt time.Time `gorm:"type:DATETIME;autoUpdateTime"`
	IsAdmin   bool
}

type Auth struct {
	ID       int
	UserID   int
	Source   string
	SourceID string
}
