package login

import (
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

	dbaccess.ValidateLogin(username, password)

	return "test"
}
