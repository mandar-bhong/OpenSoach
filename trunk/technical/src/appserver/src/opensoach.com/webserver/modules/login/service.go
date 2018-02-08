package login

type Service interface {
	Login(username string, password string) string
}

// Implement service with empty struct
type LoginService struct {
}

// Implement service functions
func (LoginService) Login(username string, password string) string {
	return "test"
}
