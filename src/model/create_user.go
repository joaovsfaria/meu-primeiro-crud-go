package model

import(
	"fmt"
	"go.uber.org/zap"

	"github.com/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/meu-primeiro-crud-go/src/configuration/logger"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr{
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	fmt.Println(ud)

	return nil
}