package models

import (
	hktmodels "opensoach.com/hkt/models"
)

type APISpCategoryAddRequest struct {
	hktmodels.DBSpCategoryDataModel
}

type APIFopSpAddRequest struct {
	hktmodels.DBFopSpDataModel
}

type APIFopSpDeleteRequest struct {
	hktmodels.DBFopSpDataModel
}
