package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/zjw1918/go-demo/model"
	"time"

	"github.com/zjw1918/go-demo/utils"
)

const COOKIE_NAME = "cookie-name"

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
	db *gorm.DB
)

func index(c *gin.Context)  {
	user := model.User{Username:"zjw", Email:"a@a.com"}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title"	: "Mode",
		"time"	: time.Now().Format(time.RFC3339),
		"user" 	: user,
	})
}

func loginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", nil)
}

func loginPost(c *gin.Context) {
	//session, _ := store.Get(c.Request, COOKIE_NAME)

	username := c.PostForm("username")
	password := c.PostForm("password")

	fmt.Println(username + ", " + password)
	hash, _ := utils.HashPassword(password)
	fmt.Println(utils.CheckPasswordHash("111", hash))


	//session.Values["authed"] = false
	//session.Save(c.Request, c.Writer)

	c.Redirect(http.StatusTemporaryRedirect, "/index")
}

func logout(c *gin.Context) {
	session, _ := store.Get(c.Request, COOKIE_NAME)

	session.Values["authed"] = false
	session.Save(c.Request, c.Writer)
}

func main() {
	var err error
	fmt.Println("hello")

	db, err = gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&model.User{})

	//db.Create(&Student{Name:"Tom", Age: 19, Address: "莲花路1978号"})

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	authorized := r.Group("/")
	authorized.Use(AuthRequired())
	{
		authorized.GET("/", index)
	}

	r.GET("/login", loginGet)
	r.POST("/login", loginPost)
	r.GET("/logout", logout)

	r.Run()
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := store.Get(c.Request, COOKIE_NAME)
		if auth, ok := session.Values["authed"].(bool); !auth || !ok {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "not authed",
			})
			c.Abort()
		}
	}
}


