package response

type UserResponse struct{
	ID 			string `JSON:"id"`
	Email 	 	string `JSON:"email"`
	Name 		string `JSON:"name"`
	Age 		int8   `JSON:"age"`
}