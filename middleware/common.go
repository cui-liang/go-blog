package middleware

import (
	"cuiliang/models"
	"cuiliang/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func CommonMiddleware(c *gin.Context) {
	db := utils.Db
	var links []models.Link
	var tags []models.Tag
	var comments []models.Comment
	//comments
	db.Where("post_id<>?", 0).Select("title,tbl_comment.created_at,author").Joins("Post").Order("tbl_comment.id DESC").Limit(10).Find(&comments)
	// links
	db.Where("status=?", 1).Find(&links)
	// tags
	db.Group("name").
		Order("COUNT(name) DESC").
		Limit(50).
		Find(&tags)

	// theme
	theme, err := c.Cookie("theme")
	if err != nil {
		theme = "light"
	}
	color := "white"
	if theme == "dark" {
		color = "black"
	}

	config, _ := utils.LoadConfig()
	c.Set("title", config.Title)
	c.Set("tags", tags)
	c.Set("CurrentYear", time.Now().Year())
	c.Set("links", links)
	c.Set("comments", comments)
	c.Set("theme", theme)
	c.Set("color", color)
	c.Set("q", c.Query("q"))
	c.Next()
}
