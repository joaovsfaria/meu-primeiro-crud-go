package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/meu-primeiro-crud-go/src/configuration/validation"
	"github.com/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/meu-primeiro-crud-go/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init updateUser Controller",
		zap.String("journey", "UpdateUser"),
	)
	var userRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validade user info", err,
			zap.String("journey", "updateUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value")
		c.JSON(errRest.Code, errRest)
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)

	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		logger.Error(
			"Error trying to call updateUser service",
			err,
			zap.String("journey", "updateUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User updated succesfully",
		zap.String("userID", userId),
		zap.String("journey", "updateUser"))

	c.Status(http.StatusOK)

}
