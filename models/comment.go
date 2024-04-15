package models

import "time"

type Comment struct {
	ID          int
	ParentID    int // 父评论的 ID
	PostID      int
	Author      string
	Content     string `form:"content" binding:"required"`
	CreatedAt   time.Time
	SubComments []Comment `gorm:"foreignKey:ParentID"`
}
