package login

import (
	"fmt"

	ghelper "opensoach.com/core/helper"
	gmodels "opensoach.com/models"
	lmodels "opensoach.com/spl/models"
	repo "opensoach.com/spl/repository"
	"opensoach.com/spl/webserver/login/dbaccess"
)

type Service interface {
	Login(username string, password string) (bool, interface{})
	//GetProducts(exeContext *wmodels.ExecutionContext) (bool, interface{})
}

// Implement service with empty struct
type LoginService struct {
}

// Implement service functions
func (LoginService) Login(username string, password string) (bool, interface{}) {

	dbErr, dbData := dbaccess.ValidateLogin(repo.Instance().Context.Dynamic.DB, username, password)

	if dbErr != nil {
		//logger.Error("DB Error occured while login. Error: %#v", dbErr.Error())
		fmt.Printf("DB Error occured while login. Error: %#v \n", dbErr.Error())
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *dbData

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = lmodels.MOD_ERR_LOGIN_INVALID_USER
		return false, errModel
	}

	isSuccess, token := ghelper.CreateToken()

	if !isSuccess {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_SERVER
		return false, errModel
	}

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = lmodels.MOD_ERR_LOGIN_INVALID_USER
		return false, errModel
	}

	resp := lmodels.LoginResponse{}
	resp.UserID = dbRecord[0].ID
	resp.Category = dbRecord[0].UserCategory
	resp.State = dbRecord[0].UserState
	resp.Token = token

	fmt.Printf("DB Data. : %#v", dbData)

	return true, &resp
}
