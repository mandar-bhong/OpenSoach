package models

import (
	"time"

	pcmodels "opensoach.com/prodcore/models"
)

type DBPatientMasterDataModel struct {
	Uuid         string `db:"uuid" json:"uuid"`
	PatientRegNo string `db:"patient_reg_no" json:"patientregno"`
	Fname        string `db:"fname" json:"fname"`
	Lname        string `db:"lname" json:"lname"`
	MobNo        string `db:"mob_no" json:"mobno"`
	Age          string `db:"age" json:"age"`
	BloodGrp     string `db:"blood_grp" json:"bloodgrp"`
	Gender       int    `db:"gender" json:"gender"`
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
}

type DBPatientUpdateStatusRowModel struct {
	AdmissionId  int64     `db:"id" dbattr:"pri,auto"  json:"admissionid"`
	CpmId        int64     `db:"cpm_id_fk" json:"cpmid"`
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
}

type DBPersonalDetailsDataModel struct {
	PatientId    int64  `db:"patient_id" json:"patientid"`
	AdmissionId  int64  `db:"admission_id_fk" json:"admissionid"`
	Uuid         string `db:"uuid" json:"uuid"`
	Age          string `db:"age" json:"age"`
	Weight       string `db:"weight" json:"weight"`
	OtherDetails string `db:"other_details" json:"otherdetails"`
}

type DBPersonalDetailsInsertRowModel struct {
	DBPersonalDetailsDataModel
	pcmodels.CPMIDEntityModel
}

type DBPersonalDetailsUpdateRowModel struct {
	PersonalDetailsId int64 `db:"id" dbattr:"pri,auto"  json:"personaldetailsid"`
	pcmodels.CPMIDEntityModel
	PatientId    int64  `db:"patient_id" json:"patientid"`
	AdmissionId  int64  `db:"admission_id_fk" json:"admissionid"`
	Uuid         string `db:"uuid" json:"uuid"`
	Age          string `db:"age" json:"age"`
	Weight       string `db:"weight" json:"weight"`
	OtherDetails string `db:"other_details" json:"otherdetails"`
}

type DBMedicalDetailsDataModel struct {
	Uuid                    string  `db:"uuid" json:"uuid"`
	PatientId               int64   `db:"patient_id" json:"patientid"`
	AdmissionId             int64   `db:"admission_id_fk" json:"admissionid"`
	ReasonForAdmission      string  `db:"reason_for_admission" json:"reasonforadmission"`
	PatientMedicalHist      string  `db:"patient_medical_hist" json:"patientmedicalhist"`
	TreatmentRecievedBefore string  `db:"treatment_recieved_before" json:"treatmentrecievedbefore"`
	FamilyHist              string  `db:"family_hist" json:"familyhist"`
	MenstrualHist           *string `db:"menstrual_hist" json:"menstrualhist"`
	Allergies               string  `db:"allergies" json:"allergies"`
	PersonalHistory         string  `db:"personal_history" json:"personalhistory"`
	GeneralPhysicalExam     string  `db:"general_physical_exam" json:"generalphysicalexam"`
	SystematicExam          string  `db:"systematic_exam" json:"systematicexam"`
}

type DBMedicalDetailsInsertRowModel struct {
	DBMedicalDetailsDataModel
	pcmodels.CPMIDEntityModel
}

type DBMedicalDetailsUpdateRowModel struct {
	MedicalDetailsId int64  `db:"id" dbattr:"pri,auto"  json:"medicaldetialsid"`
	Uuid             string `db:"uuid" json:"uuid"`
	pcmodels.CPMIDEntityModel
	PatientId               int64   `db:"patient_id" json:"patientid"`
	AdmissionId             int64   `db:"admission_id_fk" json:"admissionid"`
	ReasonForAdmission      string  `db:"reason_for_admission" json:"reasonforadmission"`
	PatientMedicalHist      string  `db:"patient_medical_hist" json:"patientmedicalhist"`
	TreatmentRecievedBefore string  `db:"treatment_recieved_before" json:"treatmentrecievedbefore"`
	FamilyHist              string  `db:"family_hist" json:"familyhist"`
	MenstrualHist           *string `db:"menstrual_hist" json:"menstrualhist"`
	Allergies               string  `db:"allergies" json:"allergies"`
	PersonalHistory         string  `db:"personal_history" json:"personalhistory"`
	GeneralPhysicalExam     string  `db:"general_physical_exam" json:"generalphysicalexam"`
	SystematicExam          string  `db:"systematic_exam" json:"systematicexam"`
}

type DBPatientConfUpdateRowModel struct {
	PatientConfId int64 `db:"id" dbattr:"pri,auto"  json:"patientconfid"`
	pcmodels.CPMIDEntityModel
	ConfTypeCode string  `db:"conf_type_code" json:"conftypecode"`
	Conf         string  `db:"conf" json:"conf"`
	ShortDesc    *string `db:"short_desc" json:"shortdesc"`
}
