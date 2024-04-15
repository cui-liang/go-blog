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

	// 设置一个用于存储登录状态的 session
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mySession", store))

	r.Use(middleware.LoggerMiddleware).
		Use(middleware.NavigationMiddleware).
		Use(middleware.ArchiveMiddleware).
		Use(middleware.CalendarMiddleware).
		Use(middleware.AuthMiddleware)

	// 设置静态文件路由映射
	r.Static("/dist", "./static/dist")
	r.Static("/images", "./static/images")
	r.Static("/uploads", "./static/uploads")
	r.StaticFile("/favicon.ico", "./static/favicon.ico")
	r.HTMLRender = loadTemplates("./views")

	front := r.Group("/", middleware.CommonMiddleware)
	{
		// 定义路由
		front.GET("/", controllers.GetPostList)
		front.GET("/page/:page", controllers.GetPostList)
		front.GET("/:param", controllers.GetPostList)
		front.GET("/:param/page/:page", controllers.GetPostList)
		front.GET("/post/create", controllers.CreatePost)
		front.GET("/post/update/:id", controllers.UpdatePost)
		front.POST("/post/save", controllers.SavePost)
		front.GET("/post/:id", controllers.GetPostDetail)
		front.GET("/tag/:tag", controllers.GetPostList)
		front.GET("/tag/:tag/page/:page", controllers.GetPostList)
		front.GET("/tags", controllers.GetTagList)
		front.GET("/guestbook", controllers.GetCommentList)
		front.POST("/comment/create", controllers.CreateComment)
		front.GET("/comment/edit/:id", controllers.EditComment)
		front.POST("/comment/update", controllers.UpdateComment)
		front.GET("/comment/delete/:id", controllers.DeleteComment)
		front.GET("/rss", controllers.GetRss)
		front.GET("/login", controllers.Login)
		front.POST("/login", controllers.Login)
		front.GET("/register", controllers.Register)
		front.POST("/register", controllers.Register)
		front.GET("/logout", controllers.Logout)
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

	includes, err := filepath.Glob(templatesDir + "/*.html")
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
			"ShortDate": func(date time.Time) string {
				return date.Format("06-01-02")
			},
			"ShortDateTime": func(date time.Time) string {
				return date.Format("06-01-02 15:04")
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
