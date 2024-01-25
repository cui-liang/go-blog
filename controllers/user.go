package controllers

import (
	"cuiliang/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// User struct represents a basic user model.
type User struct {
	Username string
	Password string
}

var users = []User{
	{"user1", "password1"},
	{"user2", "password2"},
}

func Login(c *gin.Context) {

	params := utils.MergeParams(c, gin.H{
		"title": utils.MergeTitle([]string{"登录"}),
	})
	c.HTML(http.StatusOK, "login.html", params)
}
func isValidUser(username, password string) bool {
	for _, user := range users {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}

func LoginCheck(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if isValidUser(username, password) {
		session := sessions.Default(c)
		session.Set("username", username)
		session.Save()

		c.Redirect(http.StatusFound, "/dashboard")
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"Error": "Invalid credentials"})
	}
}
