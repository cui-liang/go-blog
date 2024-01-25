package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func GetDateFromParam(c *gin.Context) (int, int, int, error) {
	// 带日期的规则
	date := c.Param("param")

	// 尝试解析为 "2006-01"
	parsedDate, err := time.Parse("2006-01", date)
	if err == nil {
		return parsedDate.Year(), int(parsedDate.Month()), 0, nil
	}

	// 尝试解析为 "2006-01-02"
	parsedDate, err = time.Parse("2006-01-02", date)
	if err == nil {
		return parsedDate.Year(), int(parsedDate.Month()), parsedDate.Day(), nil
	}

	return 0, 0, 0, fmt.Errorf("invalid date format")
}
