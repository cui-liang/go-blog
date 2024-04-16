package models

import (
	"time"
)

type Post struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	Title     string `form:"title" binding:"required"`
	Content   string `form:"content" binding:"required"`
	Excerpt   string `form:"excerpt" binding:"required"`
	Category  int    `form:"category" binding:"required"`
	Permalink string `form:"permalink"`
	Views     int
	CreatedAt time.Time `gorm:"type:DATETIME;autoCreateTime"`
	UpdatedAt time.Time `gorm:"type:DATETIME;autoUpdateTime"`
	Comments  []Comment // 关联的评论
	Tags      []Tag     `gorm:"foreignKey:PostID"`
}

type Archive struct {
	Post
	Date   string
	Year   string
	Month  string
	Count  int
	Active bool
}

type Calendar struct {
	Post
	Date string
}

type Category struct {
	ID   int
	Name string
	Path string
}

// 初始化文章分类
var categories = map[int]Category{
	1: {ID: 1, Name: "技术", Path: "technology"},
	2: {ID: 2, Name: "生活", Path: "life"},
	3: {ID: 3, Name: "跑步", Path: "run"},
}

func GetCategoryName(categoryID int) string {
	category, exists := categories[categoryID]
	if !exists {
		return "未分类"
	}
	return category.Name
}

func GetCategoryPath(categoryID int) string {
	category, exists := categories[categoryID]
	if !exists {
		return "unclassified"
	}
	return category.Path
}
