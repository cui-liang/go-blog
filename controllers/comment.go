package controllers

import (
	"cuiliang/models"
	"cuiliang/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCommentList(c *gin.Context) {
	var comments []models.Comment
	db := utils.Db
	// 查询多条记录
	if err := db.Preload("SubComments").Where("post_id = ? AND parent_id = ? AND status = ?", 0, 0, 1).Find(&comments).Error; err != nil {
		fmt.Println("Failed to query comments")
		return
	}
	params := utils.MergeParams(c, gin.H{
		"title":    utils.MergeTitle([]string{"留言"}),
		"Comments": comments,
	})
	c.HTML(http.StatusOK, "guestbook.html", params)
}

func CreateComment(c *gin.Context) {
	var form models.Comment
	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusBadRequest, "Invalid form data"+err.Error())
		return
	}

	db := utils.Db
	if err := db.Create(&form).Error; err != nil {
		c.String(http.StatusInternalServerError, "Failed to create comment")
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/post/%d", form.PostID))
}

func EditComment(c *gin.Context) {
	var comment models.Comment
	idStr := c.Param("id")
	if err := utils.Db.First(&comment, idStr).Error; err != nil {
		c.String(http.StatusNotFound, "Comment not found")
		return
	}
	params := utils.MergeParams(c, gin.H{
		"title":   "修改评论",
		"Comment": comment,
	})

	c.HTML(http.StatusOK, "comment.html", params)
}

func UpdateComment(c *gin.Context) {
	var form models.Comment
	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusBadRequest, "Invalid form data")
		return
	}

	var comment models.Comment
	if err := utils.Db.First(&comment, form.ID).Error; err != nil {
		c.String(http.StatusNotFound, "Comment not found")
		return
	}

	comment.Content = form.Content

	if err := utils.Db.Save(&comment).Error; err != nil {
		c.String(http.StatusInternalServerError, "Failed to update comment")
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/post/%d", comment.PostID))
}

func DeleteComment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid comment ID")
		return
	}

	var comment models.Comment
	if err := utils.Db.First(&comment, id).Error; err != nil {
		c.String(http.StatusNotFound, "Comment not found")
		return
	}

	if err := utils.Db.Delete(&comment).Error; err != nil {
		c.String(http.StatusInternalServerError, "Failed to delete comment")
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/post/%d", comment.PostID))
}
