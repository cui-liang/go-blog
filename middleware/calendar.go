package middleware

import (
	"cuiliang/models"
	"cuiliang/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func CalendarMiddleware(c *gin.Context) {
	year, month, _, _ := utils.GetDateFromParam(c)
	if year == 0 && month == 0 {
		today := time.Now()
		year = today.Year()
		month = int(today.Month())
	}

	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstDay.AddDate(0, 1, -1)
	lastMonth := firstDay.AddDate(0, -1, 0)
	nextMonth := firstDay.AddDate(0, 1, 0)
	calendars := getEventDates(firstDay, nextMonth)

	type Day struct {
		Date     string
		HasEvent bool
	}
	var calendar [][]Day
	var week []Day
	currentDay := firstDay

	// 计算月份开始的前置空格数量
	leadingSpaces := int(firstDay.Weekday())
	// 如果月份不是从周日开始，添加前置空格
	for i := 0; i < leadingSpaces; i++ {
		week = append(week, Day{})
	}

	for currentDay.Before(lastDay) || currentDay.Equal(lastDay) {
		day := Day{
			Date:     currentDay.Format("02"),
			HasEvent: contains(calendars, currentDay.Format("02")),
		}

		week = append(week, day)
		if len(week) == 7 {
			calendar = append(calendar, append([]Day{}, week...))
			week = week[:0] // 清空切片，准备下一周
		}

		currentDay = currentDay.AddDate(0, 0, 1)
	}

	// 处理最后一周，如果不满一周
	if len(week) > 0 {
		for len(week) < 7 {
			week = append(week, Day{})
		}
		calendar = append(calendar, append([]Day{}, week...))
	}
	c.Set("calendar", calendar)
	c.Set("year", year)
	c.Set("month", fmt.Sprintf("%02d", month))
	c.Set("lastYear", lastMonth.Year())
	c.Set("lastMonth", fmt.Sprintf("%02d", lastMonth.Month()))
	c.Set("nextYear", nextMonth.Year())
	c.Set("nextMonth", fmt.Sprintf("%02d", nextMonth.Month()))
	c.Next()
}
func getEventDates(firstDay, nextMonth time.Time) []models.Calendar {
	var calendars []models.Calendar
	utils.Db.Table("tbl_post").Select("DATE_FORMAT(created_at, '%d') as date").
		Where("created_at BETWEEN ? AND ?", firstDay, nextMonth).
		Group("date").
		Find(&calendars)
	return calendars
}

func contains(arr []models.Calendar, target string) bool {
	for _, v := range arr {
		if v.Date == target {
			return true
		}
	}
	return false
}
