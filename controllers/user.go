package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	. "github.com/zjw1918/go-web/myconst"
	"github.com/zjw1918/go-web/forms"
	"github.com/zjw1918/go-web/utils"
	"github.com/zjw1918/go-web/model"
)

type UserController struct{}
var userModel = new(model.UserModel)

//getUserID ...
func GetUserID(c *gin.Context) int64 {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID != nil {
		return utils.ConvertToInt64(userID)
	}
	return 0
}

//getSessionUserInfo ...
func GetSessionUserInfo(c *gin.Context) (userSessionInfo model.UserSessionInfo) {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID != nil {
		userSessionInfo.ID = utils.ConvertToInt64(userID)
		userSessionInfo.Username = session.Get("user_name").(string)
		userSessionInfo.Email = session.Get("user_email").(string)
	}
	return userSessionInfo
}

func (ctrl UserController) SigninGet(c *gin.Context) {
	c.HTML(http.StatusOK, "signin.tmpl", nil)
}

func (ctrl UserController) SigninPost(c *gin.Context) {
	var signinForm forms.SigninForm

	if c.Bind(&signinForm) != nil {
		c.HTML(http.StatusNotAcceptable, "signin.tmpl", gin.H{
			"message": "invalid form info",
		})
		return
	}

	user, err := userModel.Signin(signinForm)
	if err == nil {
		session := sessions.Default(c)
		session.Set("user_id", user.ID)
		session.Set("user_email", user.Email)
		session.Set("user_name", user.Username)
		session.Save()

		//c.JSON(http.StatusOK, gin.H{"message": "User signed in", "user": user})
		c.Redirect(http.StatusFound, "/")
	} else {
		//c.JSON(406, gin.H{"message": "Invalid signin details", "error": err.Error()})
		c.HTML(http.StatusNotAcceptable, "signin.tmpl", gin.H{
			"message": err.Error(),
		})
	}
}

func (ctrl UserController) SignupGet(c *gin.Context)  {
	c.HTML(http.StatusOK, "signup.tmpl", nil)
}

func (ctrl UserController) SignupPost(c *gin.Context)  {
	var signupForm forms.SignupForm

	if c.Bind(&signupForm) != nil {
		c.HTML(http.StatusNotAcceptable, "signup.tmpl", gin.H{
			"message": "invalid form info",
		})
		return
	}

	user, err := userModel.Signup(signupForm)
	if err != nil {
		c.HTML(http.StatusNotAcceptable, "signup.tmpl", gin.H{
			"message": err.Error(),
		})
		return
	}

	if user.ID > 0 {
		session := sessions.Default(c)
		session.Set("user_id", user.ID)
		session.Set("user_email", user.Email)
		session.Set("user_name", user.Username)
		session.Save()
		c.JSON(200, gin.H{"message": "Success signup", "user": user})
	} else {
		c.JSON(406, gin.H{"message": "Could not signup this user", "error": err.Error()})
	}
}

func (ctrl UserController) Signout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(200, gin.H{"message": "Signed out..."})
}

func updateSession(c *gin.Context, flag bool)  {
	session := sessions.Default(c)
	session.Set(IsAuthed, flag)
	session.Save()
}

// find all users
func (ctrl UserController) GetAllUsers(c *gin.Context)  {
	users, err := userModel.All()
	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"users": users})
	}
}

