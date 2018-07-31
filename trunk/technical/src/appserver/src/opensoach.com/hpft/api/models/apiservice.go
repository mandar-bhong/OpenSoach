package models

import (
	"time"

	hktmodels "opensoach.com/hpft/models"
)

type APIServiceConfAddRequest struct {
	hktmodels.DBServiceConfDataModel
}

type APIServiceInstanceAddRequest struct {
	hktmodels.DBServiceInstanceDataModel
}

type APIServiceInstnaceTxnRequest struct {
	SPID      int       `json:"spid"`
	StartDate time.Time `json:"startdate"`
	EndDate   time.Time `json:"enddate"`
}
