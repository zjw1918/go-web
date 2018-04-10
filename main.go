package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"

	"github.com/zjw1918/go-web/myconst"
	"github.com/zjw1918/go-web/db"
	"github.com/zjw1918/go-web/controllers"
	"github.com/zjw1918/go-web/model"
)


func index(c *gin.Context)  {
	user := model.User{Username:"jim", Email:"a@a.com"}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title"	: "Mode",
		"time"	: time.Now().Format(time.RFC3339),
		"user" 	: &user,
	})
}

func main() {
	fmt.Println("hello")

	db.Init()
	defer db.GetDB().Close()

	r := gin.Default()
	store := sessions.NewCookieStore([]byte(myconst.SessionKey))
	r.Use(sessions.Sessions(myconst.CookieName, store))
	r.LoadHTMLGlob("./public/templates/*")
	r.Static("/public", "./public")

	authorized := r.Group("/")
	authorized.Use(AuthRequired())
	{
		authorized.GET("/", index)
	}

	r.GET("/Signin", controllers.SigninGet)
	r.POST("/Signin", controllers.SigninPost)
	r.GET("/signup", controllers.SignupGet)
	r.POST("/signup", controllers.SignupPost)
	r.GET("/logout", controllers.Logout)

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.tmpl", nil)
	})

	r.Run()
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if auth, ok := session.Get(myconst.IsAuthed).(bool); !auth || !ok {
			//c.JSON(http.StatusForbidden, gin.H{
			//	"message": "not authed",
			//})
			c.Redirect(http.StatusFound, "/Signin")
			c.Abort()
		}
	}
}


