package models

import (
	"time"

	hktmodels "opensoach.com/hkt/models"
)

type APIServiceConfAddRequest struct {
	hktmodels.DBServiceConfDataModel
}

type APIServiceInstanceAddRequest struct {
	hktmodels.DBServiceInstanceDataModel
}

type APIServiceInstnaceTxnRequest struct {
	StartDate time.Time `json:"startdate"`
	EndDate   time.Time `json:"enddate"`
}
