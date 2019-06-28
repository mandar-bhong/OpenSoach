package models

import (
	"time"

	pcmodels "opensoach.com/prodcore/models"
)

type DBSearchFieldOperatorRequestFilterDataModel struct {
	Fopcode  *string `db:"fopcode" json:"fopcode"`
	FopName  *string `db:"fop_name" json:"fopname"`
	MobileNo *string `db:"mobile_no" json:"mobileno"`
	EmailId  *string `db:"email_id" json:"emailid"`
	CpmId    int64   `db:"cpm_id_fk" json:"cpmid"`
	FopState *int    `db:"fop_state" json:"fopstate"`
}

type DBSearchFieldOperatorResponseFilterDataModel struct {
	FopId    int64   `db:"id" dbattr:"pri,auto"  json:"fopid"`
	Fopcode  string  `db:"fopcode" json:"fopcode"`
	FopName  *string `db:"fop_name" json:"fopname"`
	MobileNo string  `db:"mobile_no" json:"mobileno"`
	EmailId  *string `db:"email_id" json:"emailid"`
	FopState int     `db:"fop_state" json:"fopstate"`
	FopArea  int     `db:"fop_area" json:"foparea"`
}

type DBSearchComplaintRequestFilterDataModel struct {
	ComplaintTitle *string `db:"complaint_title" json:"complainttitle"`
	ComplaintState *int    `db:"complaint_state" json:"complaintstate"`
	SpId           *int64  `db:"sp.sp_id_fk" json:"spid"`
	CpmId          int64   `db:"sp.cpm_id_fk" json:"cpmid"`
}

type DBSearchComplaintResponseFilterDataModel struct {
	ComplaintId    int64      `db:"id" dbattr:"pri,auto"  json:"complaintid"`
	SpId           int64      `db:"sp_id_fk"  json:"spid"`
	SpName         string     `db:"sp_name" json:"spname"`
	ComplaintTitle string     `db:"complaint_title" json:"complainttitle"`
	Description    *string    `db:"description" json:"description"`
	ComplaintBy    string     `db:"complaint_by" json:"complaintby"`
	Severity       *int       `db:"severity" json:"severity"`
	RaisedOn       time.Time  `db:"raised_on" json:"raisedon"`
	ComplaintState int        `db:"complaint_state" json:"complaintstate"`
	ClosedOn       *time.Time `db:"closed_on" json:"closedon"`
}

type DBSearchServiceConfRequestFilterModel struct {
	ConfTypeCode *string `db:"conf_type_code" json:"conftypecode"`
	ServConfName *string `db:"serv_conf_name" json:"servconfname"`
	CpmId        int64   `db:"cpm_id_fk" json:"cpmid"`
}

type DBSearchServiceConfResponseFilterModel struct {
	ServConfId   int64     `db:"id" dbattr:"pri,auto"  json:"servconfid"`
	CpmId        int64     `db:"cpm_id_fk" json:"cpmid"`
	SpcId        int64     `db:"spc_id_fk" json:"spcid"`
	ConfTypeCode string    `db:"conf_type_code" json:"conftypecode"`
	ServConfName string    `db:"serv_conf_name" json:"servconfname"`
	ShortDesc    *string   `db:"short_desc" json:"shortdesc"`
	CreatedOn    time.Time `db:"created_on" json:"createdon"`
	UpdatedOn    time.Time `db:"updated_on" json:"updatedon"`
}

type DBSearchServiceInstanceRequestFilterModel struct {
	SpcName *string `db:"spc_name" json:"spcname"`
	CpmId   int64   `db:"serv_conf_in.cpm_id_fk" json:"cpmid"`
}

type DBSearchServiceInstanceResponseFilterModel struct {
	SpId         int64  `db:"sp_id_fk" json:"spid"`
	ServInId     int64  `db:"id" dbattr:"pri,auto"  json:"servinid"`
	ConfTypeCode string `db:"conf_type_code" json:"conftypecode"`
	ServConfName string `db:"serv_conf_name" json:"servconfname"`
}

type DBSearchServicePointRequestFilterDataModel struct {
	SpName  *string `db:"sp_name" json:"spname"`
	SpcId   *int64  `db:"spc_id_fk" json:"spcid"`
	SpState *int    `db:"sp_state" json:"spstate"`
	DevId   *int64  `db:"dev_id_fk" dbattr:"pri"  json:"devid"`
	CpmId   *int64  `db:"sp.cpm_id_fk" json:"cpmid"`
}

type DBSearchServicePointResponseFilterDataModel struct {
	SpId         int64                    `db:"sp_id_fk" json:"spid"`
	SpName       string                   `db:"sp_name" json:"spname"`
	SpcId        int64                    `db:"spc_id_fk" json:"spcid"`
	SpcName      string                   `db:"spc_name" json:"spcname"`
	DeviceData   []DBDeviceShortDataModel `json:"devicedata"`
	ServConfId   *int64                   `db:"serv_conf_id_fk" json:"servconfid"`
	SpState      int                      `db:"sp_state" json:"spstate"`
	SpStateSince time.Time                `db:"sp_state_since" json:"spstatesince"`
}

type DBSearchDeviceRequestFilterDataModel struct {
	Serialno        *string `db:"serialno" json:"serialno"`
	DevName         *string `db:"dev_name" json:"devname"`
	ConnectionState *int    `db:"connection_state" json:"connectionstate"`
	CpmId           int64   `db:"cpm_id_fk" json:"cpmid"`
}

type DBSearchDeviceResponseFilterDataModel struct {
	DevId                int64      `db:"dev_id_fk" dbattr:"pri"  json:"devid"`
	DevName              *string    `db:"dev_name" json:"devname"`
	Serialno             string     `db:"serialno" json:"serialno"`
	CreatedOn            time.Time  `db:"created_on" json:"createdon"`
	UpdatedOn            time.Time  `db:"updated_on" json:"updatedon"`
	ConnectionState      *int       `db:"connection_state" json:"connectionstate"`
	ConnectionStateSince *time.Time `db:"connection_state_since" json:"connectionstatesince"`
	SyncState            *int       `db:"sync_state" json:"syncstate"`
	SyncStateSince       *time.Time `db:"sync_state_since" json:"syncstatesince"`
	BatteryLevel         *int       `db:"battery_level" json:"batterylevel"`
	BatteryLevelSince    *time.Time `db:"battery_level_since" json:"batterylevelsince"`
}

type DBSearchFeedbackRequestFilterDataModel struct {
	SpId     *int64 `db:"sp_id_fk" json:"spid"`
	Feedback *int   `db:"feedback" json:"feedback"`
	CpmId    int64  `db:"cpm_id_fk" json:"cpmid"`
}

type DBSearchFeedbackResponseFilterDataModel struct {
	FeedbackId      int64   `db:"id" dbattr:"pri,auto"  json:"feedbackid"`
	Feedback        int     `db:"feedback" json:"feedback"`
	FeedbackComment *string `db:"feedback_comment" json:"feedbackcomment"`
}

type DBSearchPatientRequestFilterDataModel struct {
	CpmId        int64      `db:"patient.cpm_id_fk" json:"cpmid"`
	Fname        *string    `db:"fname" json:"fname"`
	Lname        *string    `db:"lname" json:"lname"`
	PatientRegNo *string    `db:"padmsn.patient_reg_no" json:"patientregno"`
	MobNo        *string    `db:"mob_no" json:"mobno"`
	SpId         *int64     `db:"sp_id_fk" json:"spid"`
	BedNo        *string    `db:"bed_no" json:"bedno"`
	Status       *int       `db:"status" json:"status"`
	AdmittedOn   *time.Time `db:"admitted_on" json:"admittedon"`
}

type DBSearchPatientResponseFilterDataModel struct {
	PatientId          int64      `db:"patient_id_fk" json:"patientid"`
	PatientRegNo       *string    `db:"patient_reg_no" json:"patientregno"`
	AdmissionId        *int64     `db:"id" dbattr:"pri,auto"  json:"admissionid"`
	CpmId              *int64     `db:"cpm_id_fk" json:"cpmid"`
	Fname              string     `db:"fname" json:"fname"`
	Lname              string     `db:"lname" json:"lname"`
	MobNo              string     `db:"mob_no" json:"mobno"`
	EmergencyContactNo *string    `db:"emergency_contact_no" json:"emergencycontactno"`
	BedNo              *string    `db:"bed_no" json:"bedno"`
	Status             *int       `db:"status" json:"status"`
	SpId               *int64     `db:"sp_id_fk" json:"spid"`
	DrIncharge         *int64     `db:"dr_incharge" json:"drincharge"`
	AdmittedOn         *time.Time `db:"admitted_on" json:"admittedon"`
	DischargedOn       *time.Time `db:"discharged_on" json:"dischargedon"`
}

type DBSearchPatientMasterRequestFilterDataModel struct {
	CpmId int64   `db:"cpm_id_fk" json:"cpmid"`
	Fname *string `db:"fname" json:"fname"`
	Lname *string `db:"lname" json:"lname"`
	MobNo *string `db:"mob_no" json:"mobno"`
}

type DBSearchPatientMasterResponseFilterDataModel struct {
	pcmodels.StoreEntityModel
	PatientId    int64     `db:"id" dbattr:"pri,auto"  json:"patientid"`
	CpmId        int64     `db:"cpm_id_fk" json:"cpmid"`
	PatientRegNo string    `db:"patient_reg_no" json:"patientregno"`
	Fname        string    `db:"fname" json:"fname"`
	Lname        string    `db:"lname" json:"lname"`
	MobNo        string    `db:"mob_no" json:"mobno"`
	Age          string    `db:"age" json:"age"`
	BloodGrp     string    `db:"blood_grp" json:"bloodgrp"`
	Gender       int       `db:"gender" json:"gender"`
	CreatedOn    time.Time `db:"created_on" json:"createdon"`
	UpdatedOn    time.Time `db:"updated_on" json:"updatedon"`
}

type DBSearchPatientActionTxnRequestFilterDataModel struct {
	CpmId        int64   `db:"actn_txn.cpm_id_fk" json:"cpmid"`
	AdmissionId  int64   `db:"actn_txn.admission_id_fk" json:"admissionid"`
	ConfTypeCode *string `db:"actn_txn.conf_type_code" json:"conftypecode"`
}

type DBSearchPatientActionTxnResponseFilterDataModel struct {
	AdmissionId   int64     `db:"admission_id_fk" json:"admissionid"`
	PatientConfId int64     `db:"patient_conf_id_fk" json:"patientconfid"`
	ScheduledTime time.Time `db:"scheduled_time" json:"scheduledtime"`
	TxnData       string    `db:"txn_data" json:"txndata"`
	TxnState      int64     `db:"txn_state" json:"txnstate"`
	ConfTypeCode  string    `db:"conf_type_code" json:"conftypecode"`
	ActionName    string    `db:"action_name" json:"actionname"`
	UpdatedBy     int64     `db:"updated_by" json:"updated_by"`
	FirstName     string    `db:"fname" json:"firstname"`
	LastName      string    `db:"lname" json:"lastname"`
}

type DBSearchPatientDoctorOrdersRequestFilterDataModel struct {
	CpmId       int64 `db:"dorders.cpm_id_fk" json:"cpmid"`
	AdmissionId int64 `db:"admission_id_fk" json:"admissionid"`
}

type DBSearchPatientDoctorOrdersResponseFilterDataModel struct {
	DoctorsOrdersId  int64      `db:"id" dbattr:"pri,auto"  json:"id"`
	Uuid             string     `db:"uuid" json:"uuid"`
	AdmissionId      int64      `db:"admission_id_fk" json:"admissionid"`
	DoctorId         int64      `db:"doctor_id_fk" json:"doctorid"`
	DoctorFirstName  *string    `db:"doctor_fname" json:"doctorfname"`
	DoctorLastName   *string    `db:"doctor_lname" json:"doctorlname"`
	DoctorsOrders    string     `db:"doctors_orders" json:"doctorsorders"`
	Comment          *string    `db:"comment" json:"comment"`
	AckBy            *int64     `db:"ack_by" json:"ackby"`
	AckByFirstName   *string    `db:"ack_by_fname" json:"ackbyfname"`
	AckByLastName    *string    `db:"ack_by_lname" json:"ackbylname"`
	AckTime          *time.Time `db:"ack_time" json:"acktime"`
	Status           *int       `db:"status" json:"status"`
	OrderCreatedTime *time.Time `db:"order_created_time" json:"ordercreatedtime"`
	OrderType        *string    `db:"order_type" json:"ordertype"`
	DocumentId       *int64     `db:"document_id_fk" json:"documentid"`
	DocumentUUID     *string    `db:"document_uuid" json:"documentuuid"`
	Name             *string    `db:"name" json:"name"`
}

type DBSearchPatientTreatmentRequestFilterDataModel struct {
	CpmId       int64 `db:"cpm_id_fk" json:"cpmid"`
	AdmissionId int64 `db:"admission_id_fk" json:"admissionid"`
}

type DBSearchPatientTreatmentResponseFilterDataModel struct {
	TreatmentId            int64                    `db:"id" dbattr:"pri,auto"  json:"treatmentid"`
	AdmissionId            int64                    `db:"admission_id_fk" json:"admissionid"`
	TreatmentDone          string                   `db:"treatment_done" json:"treatmentdone"`
	TreatmentPerformedTime *time.Time               `db:"treatment_performed_time" json:"treatmentperformedtime"`
	Details                *string                  `db:"details" json:"details"`
	PostObservation        *string                  `db:"post_observation" json:"postobservation"`
	DocumentList           []DBDocumentTblInfoModel `db:"document_id" json:"documentlist"`
	CreatedOn              time.Time                `db:"created_on" json:"createdon"`
}

type DBSearchPatientPathologyRecordRequestFilterDataModel struct {
	CpmId       int64 `db:"cpm_id_fk" json:"cpmid"`
	AdmissionId int64 `db:"admission_id_fk" json:"admissionid"`
}

type DBSearchPatientPathologyRecordResponseFilterDataModel struct {
	PathologyId       int64                    `db:"id" dbattr:"pri,auto"  json:"pathologyid"`
	AdmissionId       int64                    `db:"admission_id_fk" json:"admissionid"`
	TestPerformed     string                   `db:"test_performed" json:"testperformed"`
	TestPerformedTime *time.Time               `db:"test_performed_time" json:"testperformedtime"`
	TestResult        *string                  `db:"test_result" json:"testresult"`
	Comments          *string                  `db:"comments" json:"comments"`
	DocumentList      []DBDocumentTblInfoModel `db:"document_id" json:"documentlist"`
	CreatedOn         time.Time                `db:"created_on" json:"createdon"`
}

type DBSearchPatientConfRequestFilterDataModel struct {
	CpmId        int64  `db:"cpm_id_fk" json:"cpmid"`
	AdmissionId  int64  `db:"admission_id_fk" json:"admissionid"`
	ConfTypeCode string `db:"conf_type_code" json:"conftypecode"`
}

type DBSearchPatientConfResponseFilterDataModel struct {
	PatientConfId int64     `db:"id" dbattr:"pri,auto"  json:"patientconfid"`
	AdmissionId   int64     `db:"admission_id_fk" json:"admissionid"`
	ConfTypeCode  string    `db:"conf_type_code" json:"conftypecode"`
	Conf          string    `db:"conf" json:"conf"`
	StartDate     time.Time `db:"start_date" json:"startdate"`
	EndDate       time.Time `db:"end_date" json:"enddate"`
	Status        int       `db:"status" json:"status"`
}

type DBDeviceSearchPatientRequestFilterDataModel struct {
	CpmId        int64   `db:"cpm_id_fk" json:"cpmid"`
	Fname        *string `db:"fname" json:"fname"`
	Lname        *string `db:"lname" json:"lname"`
	PatientRegNo *string `db:"patient_reg_no" json:"patientregno"`
	MobNo        *string `db:"mob_no" json:"mobno"`
	SpId         *int64  `db:"sp_id_fk" json:"spid"`
	BedNo        *string `db:"bed_no" json:"bedno"`
	SpName       *string `db:"sp_name" json:"spname"`
}

type DBDeviceSearchPatientResponseFilterDataModel struct {
	PatientId     int64  `db:"patient_id_fk" json:"patientid"`
	AdmissionId   int64  `db:"admission_id"  json:"admissionid"`
	CpmId         int64  `db:"cpm_id_fk" json:"cpmid"`
	UserId        int64  `db:"usr_id_fk"  json:"usrid"`
	UpmmId        int64  `db:"upmmid"  json:"upmmid"`
	PatientRegNo  string `db:"patient_reg_no" json:"patientregno"`
	BedNo         string `db:"bed_no" json:"bedno"`
	Fname         string `db:"fname" json:"fname"`
	Lname         string `db:"lname" json:"lname"`
	SpId          int64  `db:"sp_id_fk" json:"spid"`
	UpmmPatientId *int64  `db:"upmm_patient_id_fk" json:"upmmidpatientid"`
	UpmmSpId      *int64  `db:"upmm_sp_id_fk" json:"upmmidspid"`
	SpName        string `db:"sp_name" json:"spname"`
	Monitored     int    `db:"monitored" json:"monitored"`
}
