package models

type DBDeviceSerConfigModel struct {
	DeviceId      int64  `db:"dev_id_fk" json:"devid"`
	SpId          int64  `db:"sp_id_fk" json:"spid"`
	SerConfId     int64  `db:"serv_conf_id_fk" json:"servconfid"`
	SerConfInstId int64  `db:"serv_conf_in_id" json:"servconfinstid"`
	ServConfCode  string `db:"conf_type_code" json:"servconfcode"`
	ServConfName  string `db:"serv_conf_name" json:"servconfname"`
	ServiceConfig string `db:"serv_conf" json:"serconf"`
}

type DBDevicePatientConfigModel struct {
	DeviceId       int64  `db:"dev_id_fk" json:"devid"`
	SpId           int64  `db:"sp_id_fk" json:"spid"`
	SerConfId      int64  `db:"serv_conf_id_fk" json:"servconfid"`
	SerConfInstId  int64  `db:"serv_conf_in_id" json:"servconfinstid"`
	ServConfCode   string `db:"conf_type_code" json:"servconfcode"`
	ServConfName   string `db:"serv_conf_name" json:"servconfname"`
	ServiceConfig  string `db:"serv_conf" json:"serconf"`
	PatientDetails string `db:"patient_details" json:"patientdetails"`
	MedicalDetails string `db:"medical_details" json:"medicaldetails"`
}
