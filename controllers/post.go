package controllers

import (
	"cuiliang/models"
	"cuiliang/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"net/http"
	"strconv"
)

func GetPostDetail(c *gin.Context) {
	var post models.Post
	var comments []models.Comment
	id := c.Param("id")
	db := utils.Db
	if err := db.Preload("Comments").First(&post, id).Error; err != nil {
		fmt.Println("Record not found!")
	}
	fmt.Println(post)

	// 检查用户是否已经访问过
	if _, err := c.Request.Cookie(fmt.Sprintf("post_%s_visited", id)); errors.Is(err, http.ErrNoCookie) {
		// 如果用户未访问过，则增加浏览量并设置 Cookie
		db.Model(&post).UpdateColumn("views", gorm.Expr("views + 1"))
		// 设置 Cookie，有效期可以根据需求设置
		c.SetCookie(fmt.Sprintf("post_%s_visited", id), "true", 0, "/", "", false, true)
	}

	// 查询多条记录
	if err := utils.Db.Preload("SubComments").Where("post_id = ? AND parent_id=?", id, 0).Find(&comments).Error; err != nil {
		fmt.Println("Record not found!")
	}
	params := utils.MergeParams(c, gin.H{
		"title":    post.Title,
		"Post":     post,
		"Comments": comments,
	})

	c.HTML(http.StatusOK, "post.html", params)
}

func GetPostList(c *gin.Context) {
	var posts []models.Post
	var count int64
	var title []string
	var pagePath string
	pageSize := 10
	pageStr := c.Param("page")
	if pageStr == "" {
		pageStr = "1"
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid page parameter"})
		return
	}

	offset := (page - 1) * pageSize

	// 按分类查询
	param := c.Param("param")
	db := utils.Db
	if param == "technology" {
		pagePath = "/technology"
		title = append(title, "技术")
		db = db.Where("category = ?", 1)
	} else if param == "life" {
		pagePath = "/life"
		title = append(title, "生活")
		db = db.Where("category = ?", 2)
	} else if param == "run" {
		pagePath = "/run"
		title = append(title, "跑步")
		db = db.Where("category = ?", 3)
	}
	tag := c.Param("tag")
	if tag != "" {
		pagePath = "/tag/" + tag
		title = append(title, tag)
		db = db.Joins("JOIN tbl_tag ON tbl_tag.post_id = tbl_post.id").Where("name = ?", tag)
	}

	// 带日期的规则
	year, month, day, _ := utils.GetDateFromParam(c)
	if year != 0 && month != 0 {
		if day != 0 {
			date := fmt.Sprintf("%d-%02d-%02d", year, month, day)
			db = db.Where("DATE_FORMAT(created_at, ?) = ?", "%Y-%m-%d", date)
		} else {
			date := fmt.Sprintf("%d-%02d", year, month)
			db = db.Where("DATE_FORMAT(created_at, ?) = ?", "%Y-%m", date)
		}
	}

	// 搜索
	q := c.Query("q")
	if q != "" {
		db = db.Where("title like ? OR content like ?", "%"+q+"%", "%"+q+"%")
	}

	db.Model(&models.Post{}).Count(&count)
	db = db.Preload("Tags").Preload("Comments")
	db.Limit(pageSize).Offset(offset).Order("id DESC").Find(&posts)
	pageCount := int(math.Ceil(float64(count) / float64(pageSize)))

	// 分页处理
	// TODO: 如果不在1-pageCount，则返回404页面
	//if c.Param("page") == "1" || page < 1 || page > pageCount {
	//	c.Redirect(http.StatusFound, "/")
	//}

	if c.Param("page") != "" {
		title = append(title, fmt.Sprintf("第 %d 页", page))
	}

	var prevPage, nextPage string
	if page > 1 {
		prevPage = fmt.Sprintf("%d", page-1)
	}

	if page <= pageCount {
		nextPage = fmt.Sprintf("%d", page+1)
	}
	var start int

	if page < 5 {
		start = 1
	} else if page > pageCount-3 {
		start = pageCount - 6
	} else {
		start = page - 3
	}

	pages := make([]int, 7)
	for i := range pages {
		pages[i] = start + i
	}

	params := utils.MergeParams(c, gin.H{
		"title":     utils.MergeTitle(title),
		"Posts":     posts,
		"Page":      page,
		"PageCount": pageCount,
		"PrePage":   prevPage,
		"NextPage":  nextPage,
		"Pages":     pages,
		"PagePath":  pagePath + "/page/",
	})
	c.HTML(http.StatusOK, "index.html", params)
}
