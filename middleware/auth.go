package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	session := sessions.Default(c)
	fullName := session.Get("fullName")
	// 从session中获取isAdmin信息
	isAdmin := session.Get("isAdmin")

	// 将LoggedInUser变量传递给模板
	c.Set("FullName", fullName)
	c.Set("IsAdmin", isAdmin)

	c.Next() // 继续处理后续的请求
}
