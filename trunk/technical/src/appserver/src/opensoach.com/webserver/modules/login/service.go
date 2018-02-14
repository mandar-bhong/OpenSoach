package login

import (
	"fmt"
	"opensoach.com/webserver/modules/login/dbaccess"
)

type Service interface {
	Login(username string, password string) string
}

// Implement service with empty struct
type LoginService struct {
}

// Implement service functions
func (LoginService) Login(username string, password string) string {

	dbErr, dbData := dbaccess.ValidateLogin(username, password)

	if(dbErr != nil){
		fmt.Printf("DB Error occured while login. Error: %#v",dbErr.Error())
	}

	fmt.Printf("DB Data. : %#v",dbData)

	return "test"
}
