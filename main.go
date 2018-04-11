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
	userInfo := controllers.GetSessionUserInfo(c)

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title"	: "Mode",
		"time"	: time.Now().Format(time.RFC3339),
		"user"	: userInfo,
	})
}

func main() {
	fmt.Println("hello")

	db.Init()

	// Migrate the schema
	db.GetDB().AutoMigrate(&model.User{})
	defer db.GetDB().Close()

	r := gin.Default()
	store := sessions.NewCookieStore([]byte(myconst.SessionKey))
	r.Use(sessions.Sessions(myconst.CookieName, store))
	r.LoadHTMLGlob("./public/templates/*")
	r.Static("/public", "./public")

	user := new(controllers.UserController)
	authorized := r.Group("/")
	authorized.Use(AuthRequired())
	{
		authorized.GET("/", index)
		authorized.GET("/users", user.GetAllUsers)
	}

	r.GET("/signin", user.SigninGet)
	r.POST("/signin", user.SigninPost)
	r.GET("/signup", user.SignupGet)
	r.POST("/signup", user.SignupPost)
	r.GET("/signout", user.Signout)

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.tmpl", nil)
	})

	r.Run()
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := controllers.GetUserID(c)
		if session == 0 {
			//c.JSON(http.StatusForbidden, gin.H{
			//	"message": "not authed",
			//})
			c.Redirect(http.StatusFound, "/signin")
			c.Abort()
		}
	}
}
