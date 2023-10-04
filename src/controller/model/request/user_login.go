package request

type UserLogin struct {
	Email    string `JSON:"email" binding:"required,email"`
	Password string `JSON:"password" binding:"required,min=6,containsany=!@#$%*"`
}