package models

import (
	pcmodels "opensoach.com/prodcore/models"
)

type DBUserInfoModel struct {
	UsrId     int64   `dbattr:"pri,auto" db:"id" json:"usrid"`
	UsrName   string  `db:"usr_name" json:"usrname"`
	UroleCode string  `db:"urole_code" json:"urolecode"`
	UroleName string  `db:"urole_name" json:"urolename"`
	Fname     *string `db:"fname" json:"fname"`
	Lname     *string `db:"lname" json:"lname"`
}

type DBPatientMonitorMappingDataModel struct {
	pcmodels.StoreEntityModel
	UsrId     int64  `db:"usr_id_fk" json:"usrid"`
	SpId      *int64 `db:"sp_id_fk" json:"spid"`
	PatientId *int64 `db:"patient_id_fk" json:"patientid"`
	UpdatedBy int64  `db:"updated_by" json:"updatedby"`
}

type DBPatientMonitorMappingInsertRowModel struct {
	DBPatientMonitorMappingDataModel
	pcmodels.CPMIDEntityModel
}

type DBPatientMonitorMappingDeleteRowModel struct {
	UsrId     *int64 `db:"usr_id_fk" json:"usrid"`
	SpId      *int64 `db:"sp_id_fk" json:"spid"`
	PatientId *int64 `db:"patient_id_fk" json:"patientid"`
	pcmodels.CPMIDEntityModel
}
