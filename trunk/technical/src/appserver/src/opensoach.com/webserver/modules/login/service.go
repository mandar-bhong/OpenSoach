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
	Login(username string, password string) (bool, interface{})
	GetProducts(exeContext *wmodels.ExecutionContext) (bool, interface{})
}

// Implement service with empty struct
type LoginService struct {
}

// Implement service with empty struct
type ProductService struct {
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

	dbRecord := *dbData

	if len(dbRecord) < 1 {
		errModel := wmodels.ResponseError{}
		errModel.Code = helper.MOD_ERR_LOGIN_INVALID_USER
		return false, errModel
	}

	resp := wmodels.LoginResponse{}
	resp.UserID = dbRecord[0].ID
	resp.Category = dbRecord[0].UserCategory
	resp.State = dbRecord[0].UserState

	fmt.Printf("DB Data. : %#v", dbData)

	return true, &resp
}

func (ProductService) GetProducts(pExeContext *wmodels.ExecutionContext) (bool, interface{}) {

	err, data := dbaccess.GetUserProducts(pExeContext.SessionInfo.UserID)

	if err != nil {
		errModel := wmodels.ResponseError{}
		errModel.Code = whelper.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	return true, data
}
