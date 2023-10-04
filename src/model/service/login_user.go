package service

import (
	"go.uber.org/zap"

	"github.com/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/meu-primeiro-crud-go/src/model"
)

func (ud *userDomainService) LoginUserServices(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init loginUser model",
		zap.String("journey", "loginUser"))
	
	userDomain.EncryptPassword()
	//Valida se o email já está cadastrado
	user, err := ud.FindUserByEmailAndPasswordServices(
		userDomain.GetEmail(),
		userDomain.GetPassword(),
	)
	if err != nil {
		return nil
	}
	logger.Info(
		"CreateUser service executed successfully",
		zap.String("userId", user.GetID()),
		zap.String("journey", "loginUser"))
	return userDomainRepository, nil
}
