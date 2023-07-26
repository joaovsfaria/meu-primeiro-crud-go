package request

type UserRequest struct{
	Email 	 	string `JSON:"email" binding:"required,email"`
	Password 	string `JSON:"password" binding:"required,min=6,containsany=!@#$%*"`
	Name 		string `JSON:"name" binding:"required,min=4,max=100"`
	Age 		int8 `JSON:"age" binding:"required,numeric,min=1,max=140"`
}

