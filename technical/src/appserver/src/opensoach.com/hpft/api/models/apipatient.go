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

type APIAdmissionAddResponse struct {
	AdmissionId       int64 `json:"admissionid"`
	PersonalDetailsId int64 `json:"personaldetailsid"`
	MedicalDetailsId  int64 `json:"medicaldetailsid"`
}

type APIAdmissionDetailsResponse struct {
	PersonalDetails hktmodels.DBSplHpftPatientPersonalDetailsRowModel `json:"personaldetails"`
	MedicalDetails  hktmodels.DBSplHpftPatientMedicalDetailsRowModel  `json:"medicaldetails"`
}
