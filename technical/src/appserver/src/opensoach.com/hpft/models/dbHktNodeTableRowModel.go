package models

import (
	"time"

	pcmodels "opensoach.com/prodcore/models"
)

type DBSplNodeSpComplaintTableRowModel struct {
	ComplaintId    int64      `db:"id" dbattr:"pri,auto"  json:"complaintid"`
	CpmId          int64      `db:"cpm_id_fk" json:"cpmid"`
	SpId           int64      `db:"sp_id_fk" json:"spid"`
	ComplaintTitle string     `db:"complaint_title" json:"complainttitle"`
	Description    *string    `db:"description" json:"description"`
	ComplaintBy    string     `db:"complaint_by" json:"complaintby"`
	MobileNo       *string    `db:"mobile_no" json:"mobileno"`
	EmailId        *string    `db:"email_id" json:"emailid"`
	EmployeeId     *string    `db:"employee_id" json:"employeeid"`
	Severity       *int       `db:"severity" json:"severity"`
	RaisedOn       time.Time  `db:"raised_on" json:"raisedon"`
	ComplaintState int        `db:"complaint_state" json:"complaintstate"`
	ClosedOn       *time.Time `db:"closed_on" json:"closedon"`
	Remarks        *string    `db:"remarks" json:"remarks"`
	CreatedOn      time.Time  `db:"created_on" json:"createdon"`
	UpdatedOn      time.Time  `db:"updated_on" json:"updatedon"`
}

type DBSplNodeTaskLibTableRowModel struct {
	TaskId    int64     `db:"id" dbattr:"pri,auto"  json:"taskid"`
	CpmId     int64     `db:"cpm_id_fk" json:"cpmid"`
	SpcId     int64     `db:"spc_id_fk" json:"spcid"`
	TaskName  string    `db:"task_name" json:"taskname"`
	ShortDesc *string   `db:"short_desc" json:"shortdesc"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplNodeCpmTableRowModel struct {
	CpmId int64 `db:"cpm_id_fk" dbattr:"pri,auto"  json:"cpmid"`
}

type DBSplNodeDevSpMappingRowModel struct {
	DevId     int64     `db:"dev_id_fk" dbattr:"pri"  json:"devid"`
	SpId      int64     `db:"sp_id_fk" dbattr:"pri"  json:"spid"`
	CpmId     int64     `db:"cpm_id_fk" json:"cpmid"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplNodeDevTableRowModel struct {
	DevId     int64     `db:"dev_id_fk" dbattr:"pri"  json:"devid"`
	DevName   string    `db:"dev_name" json:"devname"`
	CpmId     int64     `db:"cpm_id_fk" json:"cpmid"`
	Serialno  string    `db:"serialno" json:"serialno"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplNodeDevStatusTableRowModel struct {
	DevId                int64     `dbattr:"pri" db:"dev_id_fk" json:"devid"`
	ConnectionState      int       `db:"connection_state" json:"connectionstate"`
	ConnectionStateSince time.Time `db:"connection_state_since" json:"connectionstatesince"`
	SyncState            int       `db:"sync_state" json:"syncstate"`
	SyncStateSince       time.Time `db:"sync_state_since" json:"syncstatesince"`
	BatteryLevel         int       `db:"battery_level" json:"batterylevel"`
	BatteryLevelSince    time.Time `db:"battery_level_since" json:"batterylevelsince"`
	CreatedOn            time.Time `db:"created_on" json:"createdon"`
	UpdatedOn            time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplNodeFieldOperatorTableRowModel struct {
	FopId     int64     `db:"id" dbattr:"pri,auto"  json:"fopid"`
	CpmId     int64     `db:"cpm_id_fk" json:"cpmid"`
	Fopcode   string    `db:"fopcode" json:"fopcode"`
	FopName   *string   `db:"fop_name" json:"fopname"`
	MobileNo  string    `db:"mobile_no" json:"mobileno"`
	EmailId   *string   `db:"email_id" json:"emailid"`
	ShortDesc *string   `db:"short_desc" json:"shortdesc"`
	FopState  int       `db:"fop_state" json:"fopstate"`
	FopArea   int       `db:"fop_area" json:"foparea"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplNodeFopSpTableRowModel struct {
	FopId     int64     `db:"fop_id_fk" dbattr:"pri"  json:"fopid"`
	SpId      int64     `db:"sp_id_fk" dbattr:"pri"  json:"spid"`
	CpmId     int64     `db:"cpm_id_fk" json:"cpmid"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplNodeServiceConfTableRowModel struct {
	ServConfId   int64     `db:"id" dbattr:"pri,auto"  json:"servconfid"`
	CpmId        int64     `db:"cpm_id_fk" json:"cpmid"`
	SpcId        int64     `db:"spc_id_fk" json:"spcid"`
	ConfTypeCode string    `db:"conf_type_code" json:"conftypecode"`
	ServConfName string    `db:"serv_conf_name" json:"servconfname"`
	ShortDesc    *string   `db:"short_desc" json:"shortdesc"`
	ServConf     string    `db:"serv_conf" json:"servconf"`
	CreatedOn    time.Time `db:"created_on" json:"createdon"`
	UpdatedOn    time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplNodeServiceInTxnTableRowModel struct {
	ServInTxnId int64     `db:"id" dbattr:"pri,auto"  json:"servintxnid"`
	CpmId       int64     `db:"cpm_id_fk" json:"cpmid"`
	ServInId    int64     `db:"serv_in_id_fk" json:"servinid"`
	TxnData     string    `db:"txn_data" json:"txndata"`
	TxnDate     time.Time `db:"txn_date" json:"txndate"`
	CreatedOn   time.Time `db:"created_on" json:"createdon"`
	UpdatedOn   time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplNodeServiceInstanceTableRowModel struct {
	ServInId   int64     `db:"id" dbattr:"pri,auto"  json:"servinid"`
	CpmId      int64     `db:"cpm_id_fk" json:"cpmid"`
	ServConfId int64     `db:"serv_conf_id_fk" json:"servconfid"`
	SpId       int64     `db:"sp_id_fk" json:"spid"`
	CreatedOn  time.Time `db:"created_on" json:"createdon"`
	UpdatedOn  time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplNodeSpCategoryTableRowModel struct {
	SpcId     int64     `db:"id" dbattr:"pri,auto"  json:"spcid"`
	CpmId     int64     `db:"cpm_id_fk" json:"cpmid"`
	SpcName   string    `db:"spc_name" json:"spcname"`
	ShortDesc *string   `db:"short_desc" json:"shortdesc"`
	CreatedOn time.Time `db:"created_on" json:"createdon"`
	UpdatedOn time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplNodeSpTableRowModel struct {
	pcmodels.StoreEntityModel
	SpId         int64     `db:"sp_id_fk" dbattr:"pri"  json:"spid"`
	CpmId        int64     `db:"cpm_id_fk" json:"cpmid"`
	SpcId        int64     `db:"spc_id_fk" json:"spcid"`
	SpName       string    `db:"sp_name" json:"spname"`
	ShortDesc    *string   `db:"short_desc" json:"shortdesc"`
	SpState      int       `db:"sp_state" json:"spstate"`
	SpStateSince time.Time `db:"sp_state_since" json:"spstatesince"`
	UpdatedBy    int64     `db:"updated_by" json:"updatedby"`
	CreatedOn    time.Time `db:"created_on" json:"createdon"`
	UpdatedOn    time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplNodeFeedbackTableRowModel struct {
	FeedbackId      int64     `db:"id" dbattr:"pri,auto"  json:"feedbackid"`
	CpmId           int64     `db:"cpm_id_fk" json:"cpmid"`
	SpId            int64     `db:"sp_id_fk" json:"spid"`
	Feedback        int       `db:"feedback" json:"feedback"`
	FeedbackComment *string   `db:"feedback_comment" json:"feedbackcomment"`
	RaisedOn        time.Time `db:"raised_on" json:"raisedon"`
	CreatedOn       time.Time `db:"created_on" json:"createdon"`
}

type DBSplNodeReportTemplateTableRowModel struct {
	ReportId          int64  `db:"id" dbattr:"pri,auto"  json:"reportid"`
	ReportCode        string `db:"report_code" json:"reportcode"`
	ReportDesc        string `db:"report_desc" json:"reportdesc"`
	ReportHeader      string `db:"report_header" json:"reportheader"`
	ReportQueryParams string `db:"report_query_params" json:"reportqueryparams"`
	ReportQuery       string `db:"report_query" json:"reportquery"`
}

type DBSplHpftPatientMasterTableRowModel struct {
	pcmodels.StoreEntityModel
	PatientId       int64      `db:"id" dbattr:"pri,auto"  json:"patientid"`
	CpmId           int64      `db:"cpm_id_fk" json:"cpmid"`
	PatientRegNo    string     `db:"patient_reg_no" json:"patientregno"`
	Fname           string     `db:"fname" json:"fname"`
	Lname           string     `db:"lname" json:"lname"`
	MobNo           string     `db:"mob_no" json:"mobno"`
	DateOfBirth     *time.Time `db:"date_of_birth" json:"dateofbirth"`
	Age             string     `db:"age" json:"age"`
	BloodGrp        string     `db:"blood_grp" json:"bloodgrp"`
	Gender          int        `db:"gender" json:"gender"`
	ClientUpdatedAt *time.Time `db:"client_updated_at" json:"clientupdatedat"`
	CreatedOn       time.Time  `db:"created_on" json:"createdon"`
	UpdatedOn       time.Time  `db:"updated_on" json:"updatedon"`
	UpdatedBy       int64      `db:"updated_by" json:"updated_by"`
}

type DBSplHpftPatientAdmissionTableRowModel struct {
	pcmodels.StoreEntityModel
	AdmissionId     int64      `db:"id" dbattr:"pri,auto"  json:"admissionid"`
	CpmId           int64      `db:"cpm_id_fk" json:"cpmid"`
	PatientId       int64      `db:"patient_id_fk" json:"patientid"`
	PatientRegNo    string     `db:"patient_reg_no" json:"patientregno"`
	BedNo           string     `db:"bed_no" json:"bedno"`
	Status          int        `db:"status" json:"status"`
	SpId            int64      `db:"sp_id_fk" json:"spid"`
	DrIncharge      int64      `db:"dr_incharge" json:"drincharge"`
	AdmittedOn      time.Time  `db:"admitted_on" json:"admittedon"`
	DischargedOn    *time.Time `db:"discharged_on" json:"dischargedon"`
	ClientUpdatedAt *time.Time `db:"client_updated_at" json:"clientupdatedat"`
	CreatedOn       time.Time  `db:"created_on" json:"createdon"`
	UpdatedOn       time.Time  `db:"updated_on" json:"updatedon"`
	UpdatedBy       int64      `db:"updated_by" json:"updated_by"`
}

type DBSplHpftPatientPersonalDetailsRowModel struct {
	pcmodels.StoreEntityModel
	PersonalDetailsId  int64      `db:"id" dbattr:"pri,auto"  json:"personaldetailsid"`
	CpmId              int64      `db:"cpm_id_fk" json:"cpmid"`
	PatientId          int64      `db:"patient_id" json:"patientid"`
	AdmissionId        int64      `db:"admission_id_fk" json:"admissionid"`
	Age                string     `db:"age" json:"age"`
	OtherDetails       *string    `db:"other_details" json:"otherdetails"`
	PersonAccompanying *string    `db:"person_accompanying" json:"personaccompanying"`
	ClientUpdatedAt    *time.Time `db:"client_updated_at" json:"clientupdatedat"`
	CreatedOn          time.Time  `db:"created_on" json:"createdon"`
	UpdatedOn          time.Time  `db:"updated_on" json:"updatedon"`
	UpdatedBy          int64      `db:"updated_by" json:"updated_by"`
}

type DBSplHpftPatientMedicalDetailsRowModel struct {
	pcmodels.StoreEntityModel
	MedicalDetailsId             int64      `db:"id" dbattr:"pri,auto"  json:"medicaldetialsid"`
	CpmId                        int64      `db:"cpm_id_fk" json:"cpmid"`
	PatientId                    int64      `db:"patient_id" json:"patientid"`
	AdmissionId                  int64      `db:"admission_id_fk" json:"admissionid"`
	PresentComplaints            *string    `db:"present_complaints" json:"presentcomplaints"`
	ReasonForAdmission           *string    `db:"reason_for_admission" json:"reasonforadmission"`
	HistoryPresentIllness        *string    `db:"history_present_illness" json:"historypresentillness"`
	PastHistory                  *string    `db:"past_history" json:"pasthistory"`
	TreatmentBeforeAdmission     *string    `db:"treatment_before_admission" json:"treatmentbeforeadmission"`
	InvestigationBeforeAdmission *string    `db:"investigation_before_admission" json:"investigationbeforeadmission"`
	FamilyHistory                *string    `db:"family_history" json:"familyhistory"`
	Allergies                    *string    `db:"allergies" json:"allergies"`
	PersonalHistory              *string    `db:"personal_history" json:"personalhistory"`
	ClientUpdatedAt              *time.Time `db:"client_updated_at" json:"clientupdatedat"`
	CreatedOn                    time.Time  `db:"created_on" json:"createdon"`
	UpdatedOn                    time.Time  `db:"updated_on" json:"updatedon"`
	UpdatedBy                    int64      `db:"updated_by" json:"updated_by"`
}

type DBSplHpftPatientConfTableRowModel struct {
	pcmodels.StoreEntityModel
	PatientConfId   int64      `db:"id" dbattr:"pri,auto"  json:"patientconfid"`
	CpmId           int64      `db:"cpm_id_fk" json:"cpmid"`
	AdmissionId     int64      `db:"admission_id_fk" json:"admissionid"`
	ConfTypeCode    string     `db:"conf_type_code" json:"conftypecode"`
	Conf            string     `db:"conf" json:"conf"`
	EndDate         time.Time  `db:"end_date" json:"enddate"`
	StartDate       *time.Time `db:"start_date" json:"startdate"`
	Status          int        `db:"status" json:"status"`
	ClientUpdatedAt *time.Time `db:"client_updated_at" json:"clientupdatedat"`
	CreatedOn       time.Time  `db:"created_on" json:"createdon"`
	UpdatedOn       time.Time  `db:"updated_on" json:"updatedon"`
	UpdatedBy       int64      `db:"updated_by" json:"updated_by"`
}

type DBSplHpftActionTxnTableRowModel struct {
	pcmodels.StoreEntityModel
	ActionTxnId       int64      `db:"id" dbattr:"pri,auto"  json:"actiontxnid"`
	CpmId             int64      `db:"cpm_id_fk" json:"cpmid"`
	PatientConfId     int64      `db:"patient_conf_id_fk" json:"patientconfid"`
	AdmissionId       int64      `db:"admission_id_fk" json:"admissionid"`
	TxnData           string     `db:"txn_data" json:"txndata"`
	RuntimeConfigData string     `db:"runtime_config_data" json:"runtimeconfigdata"`
	TxnDate           time.Time  `db:"txn_date" json:"txndate"`
	TxnState          int64      `db:"txn_state" json:"txnstate"`
	ConfTypeCode      string     `db:"conf_type_code" json:"conftypecode"`
	ClientUpdatedAt   *time.Time `db:"client_updated_at" json:"clientupdatedat"`
	CreatedOn         time.Time  `db:"created_on" json:"createdon"`
	UpdatedOn         time.Time  `db:"updated_on" json:"updatedon"`
	UpdatedBy         int64      `db:"updated_by" json:"updated_by"`
}

type DBSplHpftConfTableRowModel struct {
	pcmodels.StoreEntityModel
	ConfId          int64      `db:"id" dbattr:"pri,auto"  json:"confid"`
	CpmId           int64      `db:"cpm_id_fk" json:"cpmid"`
	ConfTypeCode    string     `db:"conf_type_code" json:"conftypecode"`
	Conf            string     `db:"conf" json:"conf"`
	ShortDesc       *string    `db:"short_desc" json:"shortdesc"`
	ClientUpdatedAt *time.Time `db:"client_updated_at" json:"clientupdatedat"`
	CreatedOn       time.Time  `db:"created_on" json:"createdon"`
	UpdatedOn       time.Time  `db:"updated_on" json:"updatedon"`
	UpdatedBy       int64      `db:"updated_by" json:"updated_by"`
}

type DBSplHpftDocumentTableRowModel struct {
	pcmodels.StoreEntityModel
	DocId           int64      `db:"id" dbattr:"pri,auto"  json:"docid"`
	CpmId           int64      `db:"cpm_id_fk" json:"cpmid"`
	Name            *string    `db:"name" json:"name"`
	DocType         *string    `db:"doctype" json:"doctype"`
	StoreName       *string    `db:"store_name" json:"storename"`
	Location        *string    `db:"location" json:"location"`
	LocationType    *int       `db:"location_type" json:"location_type"`
	Persisted       int        `db:"persisted" json:"persisted"`
	Updated_by      int64      `db:"updated_by" json:"updatedby"`
	ClientUpdatedAt *time.Time `db:"client_updated_at" json:"clientupdatedat"`
	CreatedOn       time.Time  `db:"created_on" json:"createdon"`
	UpdatedOn       time.Time  `db:"updated_on" json:"updatedon"`
}

type DBSplHpftDoctorsOrdersTableRowModel struct {
	pcmodels.StoreEntityModel
	DoctorsOrdersId  int64      `db:"id" dbattr:"pri,auto"  json:"id"`
	CpmId            int64      `db:"cpm_id_fk" json:"cpmid"`
	AdmissionId      int64      `db:"admission_id_fk" json:"admissionid"`
	DoctorId         int64      `db:"doctor_id_fk" json:"doctorid"`
	DoctorsOrders    string     `db:"doctors_orders" json:"doctorsorders"`
	Comment          *string    `db:"comment" json:"comment"`
	AckBy            *int64     `db:"ack_by" json:"ackby"`
	AckTime          *time.Time `db:"ack_time" json:"acktime"`
	Status           *int       `db:"status" json:"status"`
	OrderCreatedTime *time.Time `db:"order_created_time" json:"ordercreatedtime"`
	OrderType        *string    `db:"order_type" json:"ordertype"`
	DocumentId       *int64     `db:"document_id_fk" json:"documentid"`
	ClientUpdatedAt  *time.Time `db:"client_updated_at" json:"clientupdatedat"`
	CreatedOn        time.Time  `db:"created_on" json:"createdon"`
	UpdatedOn        time.Time  `db:"updated_on" json:"updatedon"`
	UpdatedBy        int64      `db:"updated_by" json:"updatedby"`
}

type DBSplHpftPathologyRecordDocTableRowModel struct {
	PathologyId int64     `db:"pathology_id_fk" dbattr:"pri"  json:"pathologyid"`
	DocumentId  int64     `db:"document_id_fk" dbattr:"pri"  json:"documentid"`
	CreatedOn   time.Time `db:"created_on" json:"createdon"`
	UpdatedOn   time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplHpftPathologyRecordTableRowModel struct {
	pcmodels.StoreEntityModel
	PathologyId       int64      `db:"id" dbattr:"pri,auto"  json:"pathologyid"`
	CpmId             int64      `db:"cpm_id_fk" json:"cpmid"`
	AdmissionId       int64      `db:"admission_id_fk" json:"admissionid"`
	TestPerformed     string     `db:"test_performed" json:"testperformed"`
	TestPerformedTime *time.Time `db:"test_performed_time" json:"testperformedtime"`
	TestResult        *string    `db:"test_result" json:"testresult"`
	Comments          *string    `db:"comments" json:"comments"`
	UpdatedBy         int64      `db:"updated_by" json:"updatedby"`
	ClientUpdatedAt   *time.Time `db:"client_updated_at" json:"clientupdatedat"`
	CreatedOn         time.Time  `db:"created_on" json:"createdon"`
	UpdatedOn         time.Time  `db:"updated_on" json:"updatedon"`
}

type DBSplHpftTreatmentDocTableRowModel struct {
	TreatmentId int64     `db:"treatment_id_fk" dbattr:"pri"  json:"treatmentid"`
	DocumentId  int64     `db:"document_id_fk" dbattr:"pri"  json:"documentid"`
	CreatedOn   time.Time `db:"created_on" json:"createdon"`
	UpdatedOn   time.Time `db:"updated_on" json:"updatedon"`
}

type DBSplHpftTreatmentTableRowModel struct {
	pcmodels.StoreEntityModel
	TreatmentId            int64      `db:"id" dbattr:"pri,auto"  json:"treatmentid"`
	CpmId                  int64      `db:"cpm_id_fk" json:"cpmid"`
	AdmissionId            int64      `db:"admission_id_fk" json:"admissionid"`
	TreatmentDone          string     `db:"treatment_done" json:"treatmentdone"`
	TreatmentPerformedTime *time.Time `db:"treatment_performed_time" json:"treatmentperformedtime"`
	Details                *string    `db:"details" json:"details"`
	PostObservation        *string    `db:"post_observation" json:"postobservation"`
	UpdatedBy              int64      `db:"updated_by" json:"updatedby"`
	ClientUpdatedAt        *time.Time `db:"client_updated_at" json:"clientupdatedat"`
	CreatedOn              *time.Time `db:"created_on" json:"createdon"`
	UpdatedOn              *time.Time `db:"updated_on" json:"updatedon"`
}
