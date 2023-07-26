package controller

import (
	"fmt"
	"log"
	
	
	"github.com/gin-gonic/gin"
	"github.com/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/meu-primeiro-crud-go/src/configuration/validation"
)


func CreateUser(c *gin.Context) {
	log.Println("Init CreaterUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s\n", err.Error())
		errRest := validation.ValidateUserError(err)


			c.JSON(errRest.Code, errRest)
			return
	}
	fmt.Println(userRequest)

}

