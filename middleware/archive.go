package middleware

import (
	"cuiliang/models"
	"cuiliang/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func ArchiveMiddleware(c *gin.Context) {
	db := utils.Db
	// 当前日期
	year, month, _, _ := utils.GetDateFromParam(c)
	if year == 0 && month == 0 {
		currentTime := time.Now()
		year = currentTime.Year()
		month = int(currentTime.Month())
	}

	// 格式化时间为 "2006-01-02"
	date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02")

	// 查询下一个月的文章数量
	var next []models.Archive
	db.Table("tbl_post").Select("DATE_FORMAT(created_at, '%Y 年 %m 月') as date, DATE_FORMAT(created_at, '%Y') as year, DATE_FORMAT(created_at, '%m') as month, COUNT(id) as count").
		Where("created_at BETWEEN ? AND ?", date, time.Now().Format("2006-01-02 15:04:05")).
		Group("date").
		Order("date ASC").
		Limit(5).
		Find(&next)

	// 查询之前的文章数量
	var prev []models.Archive
	db.Table("tbl_post").Select("DATE_FORMAT(created_at, '%Y 年 %m 月') as date, DATE_FORMAT(created_at, '%Y') as year, DATE_FORMAT(created_at, '%m') as month, COUNT(id) as count").
		Where("created_at BETWEEN ? AND ?", "2009-09-26", date).
		Group("date").
		Order("date DESC").
		Limit(10 - len(next)).
		Find(&prev)

	// 如果 prev 数量不足 5，重新查询下一个月的文章数量
	if len(prev) < 5 {
		db.Select("DATE_FORMAT(created_at, '%Y 年 %m 月') as date, DATE_FORMAT(created_at, '%Y') as year, DATE_FORMAT(created_at, '%m') as month, COUNT(id) as count").
			Where("created_at BETWEEN ? AND ?", date, time.Now().Format("2006-01-02 15:04:05")).
			Group("date").
			Order("date ASC").
			Limit(10 - len(prev)).
			Find(&next)
	}
	// 反转切片
	reversed := reverse(next)
	c.Set("archive", merge(c, reversed, prev))
	c.Next()
}

// 反转切片的函数
func reverse(arr []models.Archive) []models.Archive {
	length := len(arr)
	reversed := make([]models.Archive, length)
	for i, v := range arr {
		reversed[length-i-1] = v
	}
	return reversed
}

// 合并切片的函数
func merge(c *gin.Context, slices ...[]models.Archive) []models.Archive {
	var merged []models.Archive

	for _, slice := range slices {
		for _, post := range slice {
			// 判断 active 字段是否为 true，如果是则将 name 设置为 "default"
			year, month, _, _ := utils.GetDateFromParam(c)
			y, _ := strconv.Atoi(post.Year)
			m, _ := strconv.Atoi(post.Month)
			if y == year && m == month {
				post.Active = true
			}
			merged = append(merged, post)
		}
	}

	return merged
}
