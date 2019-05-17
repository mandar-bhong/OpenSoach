package models

type DBEPSPDataModel struct {
	ID           int64  `db:"sp_id_fk" json:"spid"`
	Name         string `db:"sp_name" json:"spname"`
	CategoryName string `db:"spc_name" json:"spcname"`
}

type DBEPSPServConfDataModel struct {
	ServInId     int64  `db:"id" dbattr:"pri,auto"  json:"servinid"`
	ServConfId   int64  `db:"serv_conf_id_fk" json:"servconfid"`
	ConfTypeCode string `db:"conf_type_code" json:"conftypecode"`
	ServConfName string `db:"serv_conf_name" json:"servconfname"`
	ServConf     string `db:"serv_conf" json:"servconf"`
}

type DBEPSPPatientConfDataModel struct {
	ServInId       int64  `db:"id" dbattr:"pri,auto"  json:"servinid"`
	ServConfId     int64  `db:"serv_conf_id_fk" json:"servconfid"`
	ConfTypeCode   string `db:"conf_type_code" json:"conftypecode"`
	ServConfName   string `db:"serv_conf_name" json:"servconfname"`
	ServConf       string `db:"serv_conf" json:"servconf"`
	PatientDetails string `db:"patient_details" json:"patientdetails"`
	MedicalDetails string `db:"medical_details" json:"medicaldetails"`
}
