package controller

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/meu-primeiro-crud-go/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("Init findUserByID controller",
		zap.String("journey", "FindUserByID"),
	)

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validade userId",
			err,
			zap.String("journey", "findUserByID"),
		)
		errorMessage := rest_err.NewBadRequestError(
			"UserID is not a valid id",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIDServices(userId)
	if err != nil {
		logger.Error("Error trying to call findUserByID service",
			err,
			zap.String("journey", "findUserByID"),
		)
		c.JSON(err.Code, err)
		return
	}
	logger.Info("FindUserByID controller executed successfuly",
		zap.String("journey", "findUserByID"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init findUserByEmail controller",
		zap.String("journey", "FindUserByEmail"),
	)

	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error trying to validade userEmail",
			err,
			zap.String("journey", "findUserByEmail"),
		)
		errorMessage := rest_err.NewBadRequestError(
			"userEmail is not a valid Email",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Error("Error trying to call findUserByEmail service",
			err,
			zap.String("journey", "findUserByEmail"),
		)
		c.JSON(err.Code, err)
		return
	}
	logger.Info("FindUserByEmail controller executed successfuly",
		zap.String("journey", "findUserByEmail"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))

}
