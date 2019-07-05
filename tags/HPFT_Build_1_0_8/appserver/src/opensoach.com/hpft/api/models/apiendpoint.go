package models

type APIUserPatientAsscociationRequest struct {
	UsrId     int64  `json:"usrid"`
	SpId      *int64 `json:"spid"`
	PatientId *int64 `json:"patientid"`
}
