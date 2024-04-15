package middleware

import "github.com/gin-gonic/gin"

func NavigationMiddleware(c *gin.Context) {
	path := c.Request.URL.Path
	navigation := []map[string]interface{}{
		{"name": "首页", "route": "/", "active": path == "/"},
		{"name": "技术", "route": "/technology", "active": path == "/technology"},
		{"name": "生活", "route": "/life", "active": path == "/life"},
		{"name": "跑步", "route": "/running", "active": path == "/running"},
		{"name": "数码", "route": "/digital", "active": path == "/digital"},
		{"name": "标签", "route": "/tags", "active": path == "/tags"},
		{"name": "留言", "route": "/guestbook", "active": path == "/guestbook"},
		{"name": "关于", "route": "/about", "active": path == "/about"},
		{"name": "RSS", "route": "/rss", "active": path == "/rss", "blank": true},
		//{"name": "订阅", "route": "/subscribe", "active": path == "/subscribe"},
		{"name": "声明", "route": "/licence", "active": path == "/licence"},
		//{"name": "捐赠", "route": "/donate", "active": path == "/donate"},
	}
	c.Set("navigation", navigation)
	c.Next()
}
