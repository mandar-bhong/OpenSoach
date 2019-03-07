package models

import (
	"time"

	pcmodels "opensoach.com/prodcore/models"
)

type DBPatientMasterDataModel struct {
	Uuid         string     `db:"uuid" json:"uuid"`
	PatientRegNo string     `db:"patient_reg_no" json:"patientregno"`
	Fname        string     `db:"fname" json:"fname"`
	Lname        string     `db:"lname" json:"lname"`
	MobNo        string     `db:"mob_no" json:"mobno"`
	DateOfBirth  *time.Time `db:"date_of_birth" json:"dateofbirth"`
	Age          string     `db:"age" json:"age"`
	BloodGrp     string     `db:"blood_grp" json:"bloodgrp"`
	Gender       int        `db:"gender" json:"gender"`
	UpdatedBy    int64      `db:"updated_by" json:"updated_by"`
}

type DBPatientMasterInsertRowModel struct {
	DBPatientMasterDataModel
	pcmodels.CPMIDEntityModel
}

type DBPatientUpdateRowModel struct {
	PatientId int64 `db:"id" dbattr:"pri,auto"  json:"patientid"`
	pcmodels.CPMIDEntityModel
	PatientRegNo string `db:"patient_reg_no" json:"patientregno"`
	Fname        string `db:"fname" json:"fname"`
	Lname        string `db:"lname" json:"lname"`
	MobNo        string `db:"mob_no" json:"mobno"`
	Age          string `db:"age" json:"age"`
	BloodGrp     string `db:"blood_grp" json:"bloodgrp"`
	Gender       int    `db:"gender" json:"gender"`
	UpdatedBy    int64  `db:"updated_by" json:"updated_by"`
}

type DBPatientUpdateStatusRowModel struct {
	AdmissionId int64 `db:"id" dbattr:"pri,auto"  json:"admissionid"`
	pcmodels.CPMIDEntityModel
	Status       int       `db:"status" json:"status"`
	DischargedOn time.Time `db:"discharged_on" json:"dischargedon"`
}

type DBPatientFilterModel struct {
	Fname *string `db:"fname" json:"fname"`
	Lname *string `db:"lname" json:"lname"`
	MobNo *string `db:"mob_no" json:"mobno"`
}

type DBAdmissionTblDataModel struct {
	Uuid         string     `db:"uuid" json:"uuid"`
	PatientId    int64      `db:"patient_id_fk" json:"patientid"`
	PatientRegNo string     `db:"patient_reg_no" json:"patientregno"`
	BedNo        string     `db:"bed_no" json:"bedno"`
	Status       int        `db:"status" json:"status"`
	SpId         int64      `db:"sp_id_fk" json:"spid"`
	DrIncharge   int64      `db:"dr_incharge" json:"drincharge"`
	AdmittedOn   time.Time  `db:"admitted_on" json:"admittedon"`
	DischargedOn *time.Time `db:"discharged_on" json:"dischargedon"`
	UpdatedBy    int64      `db:"updated_by" json:"updated_by"`
}

type DBAdmissionTblInsertRowModel struct {
	DBAdmissionTblDataModel
	pcmodels.CPMIDEntityModel
}

type DBAdmissionTblUpdateRowModel struct {
	AdmissionId int64  `db:"id" dbattr:"pri,auto"  json:"admissionid"`
	Uuid        string `db:"uuid" json:"uuid"`
	pcmodels.CPMIDEntityModel
	PatientId    int64      `db:"patient_id_fk" json:"patientid"`
	PatientRegNo string     `db:"patient_reg_no" json:"patientregno"`
	BedNo        string     `db:"bed_no" json:"bedno"`
	Status       int        `db:"status" json:"status"`
	SpId         int64      `db:"sp_id_fk" json:"spid"`
	DrIncharge   int64      `db:"dr_incharge" json:"drincharge"`
	AdmittedOn   time.Time  `db:"admitted_on" json:"admittedon"`
	DischargedOn *time.Time `db:"discharged_on" json:"dischargedon"`
	UpdatedBy    int64      `db:"updated_by" json:"updated_by"`
}

type DBPersonalDetailsDataModel struct {
	PatientId          int64   `db:"patient_id" json:"patientid"`
	AdmissionId        int64   `db:"admission_id_fk" json:"admissionid"`
	Uuid               string  `db:"uuid" json:"uuid"`
	Age                string  `db:"age" json:"age"`
	OtherDetails       *string `db:"other_details" json:"otherdetails"`
	PersonAccompanying *string `db:"person_accompanying" json:"personaccompanying"`
	UpdatedBy          int64   `db:"updated_by" json:"updated_by"`
}

type DBPersonalDetailsInsertRowModel struct {
	DBPersonalDetailsDataModel
	pcmodels.CPMIDEntityModel
}

type DBPersonalDetailsUpdateRowModel struct {
	PersonalDetailsId int64 `db:"id" dbattr:"pri,auto"  json:"personaldetailsid"`
	pcmodels.CPMIDEntityModel
	DBPersonalDetailsDataModel
}

type DBPersonalDetailsUpdatePersonAccompanyingRowModel struct {
	PersonalDetailsId int64 `db:"id" dbattr:"pri,auto"  json:"personaldetailsid"`
	pcmodels.CPMIDEntityModel
	PersonAccompanying string `db:"person_accompanying" json:"personaccompanying"`
	UpdatedBy          int64  `db:"updated_by" json:"updated_by"`
}

type DBMedicalDetailsDataModel struct {
	Uuid                         string  `db:"uuid" json:"uuid"`
	PatientId                    int64   `db:"patient_id" json:"patientid"`
	AdmissionId                  int64   `db:"admission_id_fk" json:"admissionid"`
	PresentComplaints            *string `db:"present_complaints" json:"presentcomplaints"`
	ReasonForAdmission           *string `db:"reason_for_admission" json:"reasonforadmission"`
	HistoryPresentIllness        *string `db:"history_present_illness" json:"historypresentillness"`
	PastHistory                  *string `db:"past_history" json:"pasthistory"`
	TreatmentBeforeAdmission     *string `db:"treatment_before_admission" json:"treatmentbeforeadmission"`
	InvestigationBeforeAdmission *string `db:"investigation_before_admission" json:"investigationbeforeadmission"`
	FamilyHistory                *string `db:"family_history" json:"familyhistory"`
	Allergies                    *string `db:"allergies" json:"allergies"`
	PersonalHistory              *string `db:"personal_history" json:"personalhistory"`
	UpdatedBy                    int64   `db:"updated_by" json:"updated_by"`
}

type DBMedicalDetailsInsertRowModel struct {
	DBMedicalDetailsDataModel
	pcmodels.CPMIDEntityModel
}

type DBMedicalDetailsUpdateRowModel struct {
	MedicalDetailsId int64 `db:"id" dbattr:"pri,auto"  json:"medicaldetialsid"`
	pcmodels.CPMIDEntityModel
	DBMedicalDetailsDataModel
}

type DBMedicalDetailsUpdatePresentComplaintsRowModel struct {
	MedicalDetailsId int64 `db:"id" dbattr:"pri,auto"  json:"medicaldetialsid"`
	pcmodels.CPMIDEntityModel
	PresentComplaints string `db:"present_complaints" json:"presentcomplaints"`
	UpdatedBy         int64  `db:"updated_by" json:"updated_by"`
}

type DBMedicalDetailsUpdateReasonForAdmissionRowModel struct {
	MedicalDetailsId int64 `db:"id" dbattr:"pri,auto"  json:"medicaldetialsid"`
	pcmodels.CPMIDEntityModel
	ReasonForAdmission string `db:"reason_for_admission" json:"reasonforadmission"`
	UpdatedBy          int64  `db:"updated_by" json:"updated_by"`
}

type DBMedicalDetailsUpdateHistoryPresentIllnessRowModel struct {
	MedicalDetailsId int64 `db:"id" dbattr:"pri,auto"  json:"medicaldetialsid"`
	pcmodels.CPMIDEntityModel
	HistoryPresentIllness string `db:"history_present_illness" json:"historypresentillness"`
	UpdatedBy             int64  `db:"updated_by" json:"updated_by"`
}

type DBMedicalDetailsUpdatePastHistoryRowModel struct {
	MedicalDetailsId int64 `db:"id" dbattr:"pri,auto"  json:"medicaldetialsid"`
	pcmodels.CPMIDEntityModel
	PastHistory string `db:"past_history" json:"pasthistory"`
	UpdatedBy   int64  `db:"updated_by" json:"updated_by"`
}

type DBMedicalDetailsUpdateTreatmentBeforeAdmissionRowModel struct {
	MedicalDetailsId int64 `db:"id" dbattr:"pri,auto"  json:"medicaldetialsid"`
	pcmodels.CPMIDEntityModel
	TreatmentBeforeAdmission string `db:"treatment_before_admission" json:"treatmentbeforeadmission"`
	UpdatedBy                int64  `db:"updated_by" json:"updated_by"`
}

type DBMedicalDetailsUpdateInvestigationBeforeAdmissionRowModel struct {
	MedicalDetailsId int64 `db:"id" dbattr:"pri,auto"  json:"medicaldetialsid"`
	pcmodels.CPMIDEntityModel
	InvestigationBeforeAdmission string `db:"investigation_before_admission" json:"investigationbeforeadmission"`
	UpdatedBy                    int64  `db:"updated_by" json:"updated_by"`
}

type DBMedicalDetailsUpdateFamilyHistoryRowModel struct {
	MedicalDetailsId int64 `db:"id" dbattr:"pri,auto"  json:"medicaldetialsid"`
	pcmodels.CPMIDEntityModel
	FamilyHistory string `db:"family_history" json:"familyhistory"`
	UpdatedBy     int64  `db:"updated_by" json:"updated_by"`
}

type DBMedicalDetailsUpdateAllergiesRowModel struct {
	MedicalDetailsId int64 `db:"id" dbattr:"pri,auto"  json:"medicaldetialsid"`
	pcmodels.CPMIDEntityModel
	Allergies string `db:"allergies" json:"allergies"`
	UpdatedBy int64  `db:"updated_by" json:"updated_by"`
}

type DBMedicalDetailsUpdatePersonalHistoryRowModel struct {
	MedicalDetailsId int64 `db:"id" dbattr:"pri,auto"  json:"medicaldetialsid"`
	pcmodels.CPMIDEntityModel
	PersonalHistory *string `db:"personal_history" json:"personalhistory"`
	UpdatedBy       int64   `db:"updated_by" json:"updated_by"`
}

type DBPatientConfUpdateRowModel struct {
	PatientConfId int64 `db:"id" dbattr:"pri,auto"  json:"patientconfid"`
	pcmodels.CPMIDEntityModel
	ConfTypeCode string  `db:"conf_type_code" json:"conftypecode"`
	Conf         string  `db:"conf" json:"conf"`
	ShortDesc    *string `db:"short_desc" json:"shortdesc"`
}

type DBPatientAdmissionStatusInfoModel struct {
	PatientId int64 `db:"patient_id_fk" json:"patientid"`
	Status    int   `db:"status" json:"status"`
}

type PatientUserInfo struct {
	Firstname string `db:"fname" json:"firstname"`
	LastName  string `db:"lname" json:"lastname"`
}

type DBPatientTreatmentDataModel struct {
	Uuid            string  `db:"uuid" json:"uuid"`
	AdmissionId     int64   `db:"admission_id_fk" json:"admissionid"`
	TreatmentDone   string  `db:"treatment_done" json:"treatmentdone"`
	Details         *string `db:"details" json:"details"`
	PostObservation *string `db:"post_observation" json:"postobservation"`
	UpdatedBy       int64   `db:"updated_by" json:"updatedby"`
}

type DBPatientTreatmentInsertRowModel struct {
	DBPatientTreatmentDataModel
	pcmodels.CPMIDEntityModel
}

type DBPatientTreatmentDocInsertRowModel struct {
	TreatmentId int64 `db:"treatment_id_fk" dbattr:"pri"  json:"treatmentid"`
	DocumentId  int64 `db:"document_id_fk" dbattr:"pri"  json:"documentid"`
}

type DBPatientPathologyRecordDataModel struct {
	Uuid          string  `db:"uuid" json:"uuid"`
	AdmissionId   int64   `db:"admission_id_fk" json:"admissionid"`
	TestPerformed string  `db:"test_performed" json:"testperformed"`
	TestResult    *string `db:"test_result" json:"testresult"`
	Comments      *string `db:"comments" json:"comments"`
	UpdatedBy     int64   `db:"updated_by" json:"updatedby"`
}

type DBPatientPathologyRecordInsertRowModel struct {
	DBPatientPathologyRecordDataModel
	pcmodels.CPMIDEntityModel
}

type DBPatientPathologyRecordDocInsertRowModel struct {
	PathologyId int64 `db:"pathology_id_fk" dbattr:"pri"  json:"pathologyid"`
	DocumentId  int64 `db:"document_id_fk" dbattr:"pri"  json:"documentid"`
}

type DBDocumentTblInfoModel struct {
	DocumentUUID string `db:"uuid" json:"documentuuid"`
	DocumentName string `db:"name" json:"documentname"`
}
