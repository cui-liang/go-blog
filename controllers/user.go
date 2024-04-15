package controllers

import (
	"cuiliang/models"
	"cuiliang/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		// 渲染注册页面
		c.HTML(http.StatusOK, "register.html", nil)
		return
	}

	var form struct {
		Username        string `form:"username" binding:"required"`
		FullName        string `form:"full_name" binding:"required"`
		Email           string `form:"email" binding:"required"`
		Password        string `form:"password" binding:"required"`
		PasswordConfirm string `form:"password_confirm" binding:"required"`
	}

	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusBadRequest, "Invalid form data"+err.Error())
		return
	}

	// 检查密码和确认密码是否匹配
	if form.Password != form.PasswordConfirm {
		c.String(http.StatusBadRequest, "Passwords do not match")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to hash password")
		return
	}

	user := models.User{
		Username: form.Username,
		Password: string(hashedPassword),
		Email:    form.Email,
		FullName: form.FullName,
	}

	if err := utils.Db.Create(&user).Error; err != nil {
		c.String(http.StatusInternalServerError, "Failed to register user")
		return
	}

	c.String(http.StatusOK, "User registered successfully")
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	// 清除session中的username信息
	session.Delete("username")
	session.Delete("fullName")
	session.Delete("isAdmin")
	err := session.Save()
	if err != nil {
		return
	}

	// 重定向到登录页面
	c.Redirect(http.StatusFound, "/")
}

func Login(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		params := utils.MergeParams(c, gin.H{
			"title": utils.MergeTitle([]string{"登录"}),
		})
		c.HTML(http.StatusOK, "login.html", params)
		return
	}
	var form struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusBadRequest, "Invalid form data")
		return
	}

	var user models.User
	if err := utils.Db.Where("username = ?", form.Username).First(&user).Error; err != nil {
		c.String(http.StatusUnauthorized, "Invalid username or password")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
		c.String(http.StatusUnauthorized, "Invalid username or password")
		return
	}

	session := sessions.Default(c)
	session.Set("username", user.Username)
	session.Set("fullName", user.FullName)
	session.Set("isAdmin", user.IsAdmin)
	session.Save()

	// 重定向到登录页面
	c.Redirect(http.StatusFound, "/")
}
