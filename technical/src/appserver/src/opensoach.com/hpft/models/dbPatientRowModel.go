package models

type DBPatientDataModel struct {
	PatientDetails        string `db:"patient_details" json:"patientdetails"`
	MedicalDetails        string `db:"medical_details" json:"medicaldetails"`
	PatientFileTemplateID int64  `db:"patient_file_template" json:"patientfiletemplate"`
	SpId                  int64  `db:"sp_id_fk" json:"spid"`
	Status                int    `db:"status" json:"status"`
}

type DBPatientInsertRowModel struct {
	DBPatientDataModel
	CpmId    int64 `db:"cpm_id_fk" json:"cpmid"`
	ServInId int64 `db:"serv_in_id_fk" json:"servinid"`
}

type DBPatientListDataModel struct {
	PatientId             int64  `db:"id" dbattr:"pri,auto"  json:"patientid"`
	PatientDetails        string `db:"patient_details" json:"patientdetails"`
	MedicalDetails        string `db:"medical_details" json:"medicaldetails"`
	PatientFileTemplateID int64  `db:"patient_file_template" json:"patientfiletemplate"`
	Status                int    `db:"status" json:"status"`
}

type DBPatientUpdateRowModel struct {
	PatientId             int64  `db:"id" dbattr:"pri,auto"  json:"patientid"`
	CpmId                 int64  `db:"cpm_id_fk" json:"cpmid"`
	PatientDetails        string `db:"patient_details" json:"patientdetails"`
	MedicalDetails        string `db:"medical_details" json:"medicaldetails"`
	PatientFileTemplateID int64  `db:"patient_file_template" json:"patientfiletemplate"`
}

type DBPatientUpdateStatusRowModel struct {
	PatientId int64 `db:"id" dbattr:"pri,auto"  json:"patientid"`
	CpmId     int64 `db:"cpm_id_fk" json:"cpmid"`
	Status    int   `db:"status" json:"status"`
}
