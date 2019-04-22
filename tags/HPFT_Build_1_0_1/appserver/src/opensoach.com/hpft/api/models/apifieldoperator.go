package models

import (
	hktmodels "opensoach.com/hpft/models"
)

type APIFieldOperatorAddRequest struct {
	hktmodels.DBFieldOperatorDataModel
}

type APIFopSpAddRequest struct {
	hktmodels.DBFopSpDataModel
}

type APIFopSpDeleteRequest struct {
	hktmodels.DBFopSpDataModel
}
