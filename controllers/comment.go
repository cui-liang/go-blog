package controllers

import (
	"cuiliang/models"
	"cuiliang/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCommentList(c *gin.Context) {
	var comments []models.Comment
	db := utils.Db
	// 查询多条记录
	if err := db.Preload("SubComments").Where("post_id = ? AND parent_id = ?", 0, 0).Find(&comments).Error; err != nil {
		fmt.Println("Failed to query comments")
		return
	}
	params := utils.MergeParams(c, gin.H{
		"title":    utils.MergeTitle([]string{"留言"}),
		"Comments": comments,
	})
	c.HTML(http.StatusOK, "guestbook.html", params)
}
