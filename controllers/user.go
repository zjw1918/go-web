package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/zjw1918/go-web/utils"
	"github.com/gin-contrib/sessions"
	. "github.com/zjw1918/go-web/myconst"
	"github.com/zjw1918/go-web/db"
	"github.com/zjw1918/go-web/model"
	"log"
)



func SigninGet(c *gin.Context) {
	c.HTML(http.StatusOK, "signin.tmpl", nil)
}

func SigninPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	log.Println(fmt.Sprintf("user: %s, %s", username, password))
	//hash, _ := utils.HashPassword(password)

	var preUser model.User
	//db.GetDB().First(&preUser, "username = ?", username)
	db.GetDB().First(&preUser, "username = ?", username)
	log.Println(preUser)

	if utils.CheckPasswordHash(preUser.Password, password) {
		updateSession(c, true)
		c.Redirect(http.StatusFound, "/")
	} else {
		c.HTML(http.StatusOK, "signin.tmpl", gin.H{
			"message": "Username or password Error.",
		})
	}
}

func SignupGet(c *gin.Context)  {
	c.HTML(http.StatusOK, "signup.tmpl", nil)
}

func SignupPost(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	fmt.Println(fmt.Sprintf("user: %s, %s, %s", username, password, email))

	var preUser *model.User
	db.GetDB().Where("username = ?", username).First(preUser)

	if preUser != nil {
		c.HTML(http.StatusOK, "/signup", gin.H{
			"message": "User account already existed.",
		})
		return
	}

	hash, _ := utils.HashPassword(password)
	user := &model.User{
		Username: username,
		Password: hash,
		Email: email,
	}

	db.GetDB().Create(user)

	updateSession(c, true)
	c.Redirect(http.StatusFound, "/")
}

func Logout(c *gin.Context) {
	updateSession(c, false)
	c.Redirect(http.StatusFound, "/Signin")
}

func updateSession(c *gin.Context, flag bool)  {
	session := sessions.Default(c)
	session.Set(IsAuthed, flag)
	session.Save()
}



