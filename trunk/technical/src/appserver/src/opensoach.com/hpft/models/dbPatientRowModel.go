package models

type DBPatientDataModel struct {
	PatientDetails      string `db:"patient_details" json:"patientdetails"`
	MedicalDetails      string `db:"medical_details" json:"medicaldetails"`
	PatientFileTemplate string `db:"patient_file_template" json:"patientfiletemplate"`
}

type DBPatientInsertRowModel struct {
	DBPatientDataModel
	CpmId int64 `db:"cpm_id_fk" json:"cpmid"`
}
