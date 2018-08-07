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
	SPID              int        `json:"spid"`
	ServiceInstanceID int64      `json:"servinid"`
	StartDate         *time.Time `json:"startdate"`
	EndDate           *time.Time `json:"enddate"`
}
