package controllers

import (
	"cuiliang/models"
	"cuiliang/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/feeds"
	"time"
)

func GetRss(c *gin.Context) {
	now := time.Now()
	domain := "www.cuiliang.com"
	feed := &feeds.Feed{
		Title:       "Wblog",
		Link:        &feeds.Link{Href: domain},
		Description: "Wblog,talk about golang,java and so on.",
		Author:      &feeds.Author{Name: "Wangsongyan", Email: "wangsongyanlove@163.com"},
		Created:     now,
	}

	feed.Items = make([]*feeds.Item, 0)

	var posts []models.Post
	// 查询多条记录
	if err := utils.Db.Order("id DESC").Limit(20).Find(&posts).Error; err != nil {
		fmt.Println("Record not found!")
	}

	for _, post := range posts {
		created := post.CreatedAt
		item := &feeds.Item{
			Id:          fmt.Sprintf("%s/post/%d", domain, post.ID),
			Title:       post.Title,
			Link:        &feeds.Link{Href: fmt.Sprintf("%s/post/%d", domain, post.ID)},
			Description: post.Excerpt,
			Created:     created,
		}
		feed.Items = append(feed.Items, item)
	}
	rss, _ := feed.ToRss()

	c.Writer.WriteString(rss)
}
