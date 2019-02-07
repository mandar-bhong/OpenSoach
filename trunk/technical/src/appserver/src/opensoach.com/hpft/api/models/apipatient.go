package models

import (
	hktmodels "opensoach.com/hpft/models"
)

type APIPatientAddRequest struct {
	hktmodels.DBPatientMasterDataModel
}

type APIAdmissionAddRequest struct {
	hktmodels.DBAdmissionTblDataModel
}

type APIPersonalDetailsAddRequest struct {
	hktmodels.DBPersonalDetailsDataModel
}

type APIMedicalDetailsAddRequest struct {
	hktmodels.DBMedicalDetailsDataModel
}
