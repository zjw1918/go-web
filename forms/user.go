package forms

type SigninForm struct {
	Username     string `form:"username" json:"username" binding:"required,max=100"`
	Password string `form:"password" json:"password" binding:"required"`
}

type SignupForm struct {
	Username     string `form:"username" json:"username" binding:"required,max=100"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}
