package request

type UserRequest struct{
	Email 	 	string `JSON:"email"`
	Password 	string `JSON:"password"`
	Name 		string `JSON:"name"`
	Age 		int8 `JSON:"age"`
}

