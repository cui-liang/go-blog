package controllers

import (
	"cuiliang/models"
	"cuiliang/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTagList(c *gin.Context) {
	var tags []models.Tag
	db := utils.Db
	db.Select("name, count(name) as count").Group("name").Order("count DESC").Find(&tags)

	params := utils.MergeParams(c, gin.H{
		"title": utils.MergeTitle([]string{"标签"}),
		"Tags":  tags,
	})
	c.HTML(http.StatusOK, "tags.html", params)
}
