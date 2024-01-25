package models

type Comment struct {
	ID          int
	ParentID    int // 父评论的 ID
	PostID      string
	Author      string
	Content     string
	CreatedAt   string
	Post        Post
	SubComments []Comment `gorm:"foreignKey:ParentID"`
}
