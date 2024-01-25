package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func MergeParams(c *gin.Context, h gin.H) gin.H {
	if c.Keys == nil {
		return h
	}

	if h == nil || len(h) == 0 {
		return c.Keys
	}

	mergedParams := make(gin.H)

	for key, val := range c.Keys {
		mergedParams[key] = val
	}

	for key, val := range h {
		mergedParams[key] = val
	}

	return mergedParams
}

func MergeTitle(parts []string) string {
	config, _ := LoadConfig()
	breadcrumb := config.Title

	// 如果有传入的部分，则构建面包屑
	if len(parts) > 0 {
		// 将部分添加到面包屑中，按照指定的顺序
		for _, part := range parts {
			breadcrumb = fmt.Sprintf("%s - %s", part, breadcrumb)
		}
	}

	return breadcrumb
}
