package login

import (
	"fmt"

	"opensoach.com/utility/logger"
	"opensoach.com/webserver/modules/login/dbaccess"
	"opensoach.com/webserver/modules/login/helper"
	whelper "opensoach.com/webserver/webhelper"
	wmodels "opensoach.com/webserver/webmodels"
)

type Service interface {
	Login(username string, password string) string
}

// Implement service with empty struct
type LoginService struct {
}

// Implement service functions
func (LoginService) Login(username string, password string) (bool, interface{}) {

	dbErr, dbData := dbaccess.ValidateLogin(username, password)

	if dbErr != nil {
		logger.Error("DB Error occured while login. Error: %#v", dbErr.Error())
		errModel := wmodels.ResponseError{}
		errModel.Code = whelper.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	if len(*dbData) < 1 {
		errModel := wmodels.ResponseError{}
		errModel.Code = helper.MOD_ERR_LOGIN_INVALID_USER
		return false, errModel
	}

	fmt.Printf("DB Data. : %#v", dbData)

	return true, dbData
}
