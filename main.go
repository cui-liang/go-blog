// main.go
package main

import (
	"cuiliang/controllers"
	"cuiliang/middleware"
	"cuiliang/models"
	"cuiliang/utils"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

func main() {
	r := gin.Default()

	r.Use(middleware.LoggerMiddleware).
		Use(middleware.NavigationMiddleware).
		Use(middleware.ArchiveMiddleware).
		Use(middleware.CalendarMiddleware)

	// 设置静态文件路由映射
	r.Static("/css", "./static/css")
	r.Static("/js", "./static/js")
	r.Static("/node", "./node_modules")
	r.Static("/dist", "./static/dist")
	r.Static("/images", "./static/images")
	r.Static("/uploads", "./static/uploads")
	r.StaticFile("/favicon.ico", "./static/favicon.ico")
	r.HTMLRender = loadTemplates("./views")

	// 设置一个用于存储登录状态的 session
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mySession", store))

	front := r.Group("/", middleware.CommonMiddleware)
	{
		// 定义路由
		front.GET("/", controllers.GetPostList)
		front.GET("/page/:page", controllers.GetPostList)
		front.GET("/:param", controllers.GetPostList)
		front.GET("/:param/page/:page", controllers.GetPostList)
		front.GET("/post/:id", controllers.GetPostDetail)
		front.GET("/tag/:tag", controllers.GetPostList)
		front.GET("/tag/:tag/page/:page", controllers.GetPostList)
		front.GET("/tags", controllers.GetTagList)
		front.GET("/guestbook", controllers.GetCommentList)
		front.GET("/rss", controllers.GetRss)
		front.GET("/subscribe", controllers.GetCommentList)
		front.GET("/login", controllers.Login)
		front.POST("/login", controllers.LoginCheck)
		front.GET("/about", func(c *gin.Context) {
			params := utils.MergeParams(c, gin.H{
				"title": utils.MergeTitle([]string{"关于"}),
			})
			c.HTML(http.StatusOK, "about.html", params)
		})
		front.GET("/licence", func(c *gin.Context) {
			params := utils.MergeParams(c, gin.H{
				"title": utils.MergeTitle([]string{"声明"}),
			})
			c.HTML(http.StatusOK, "licence.html", params)
		})
		front.GET("/donate", func(c *gin.Context) {
			params := utils.MergeParams(c, gin.H{
				"title": utils.MergeTitle([]string{"捐赠"}),
			})
			c.HTML(http.StatusOK, "donate.html", params)
		})
	}

	// 启动 Gin 服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our views map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFilesFuncs(filepath.Base(include), template.FuncMap{
			"MarkdownToHTML": utils.MarkdownToHTML,
			"ShortDate": func(date string) string {
				t, err := time.Parse("2006-01-02 15:04:05", date)
				if err != nil {
					return ""
				}
				return t.Format("06-01-02")
			},
			"ShortDateTime": func(date string) string {
				t, err := time.Parse("2006-01-02 15:04:05", date)
				if err != nil {
					return ""
				}
				return t.Format("06-01-02 15:04")
			},
			"CategoryPath": func(category int) string {
				return models.GetCategoryPath(category)
			},
			"CategoryName": func(category int) string {
				return models.GetCategoryName(category)
			},
		}, files...)
	}

	return r
}
