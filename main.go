package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/meu-primeiro-crud-go/src/configuration/database/mongodb"
	"github.com/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/meu-primeiro-crud-go/src/controller"
	"github.com/meu-primeiro-crud-go/src/controller/routes"
	"github.com/meu-primeiro-crud-go/src/model/repository"
	"github.com/meu-primeiro-crud-go/src/model/service"
)

func main() {
	logger.Info("About to start aplication")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnetion(context.Background())
	if err != nil {
		log.Fatalf("error trying to connect to database, error=%s \n", err.Error())
		return
	}

	//init dependences
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
