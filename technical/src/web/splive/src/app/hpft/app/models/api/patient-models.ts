import { Time } from "@angular/common";

export class PatientDataAddRequest {
    patientdetails: string;
    medicaldetails: string;
    patientfiletemplate: number;
    patientid: number;
    spid: number;
    status: number;
    spname: string;
    dischargedon: Date;
}
export class PatientDetailAddRequest {
    patientname: string;
    age: number;
    emergencycontactno: string;
    patientregistrationno: string;
    ward: number;
    bedno: number;
    admissiondate: Date;
    dischargedate: Date;
    bloodgroup: string;
    weight: number;
    drinst: string;
    gender: string;
}
export class MedicalDetailAddRequest {
    reasonadmission: string;
    treatmentdone: string;
    allergies: string;
    patientmedicalhistory: string;

}
// export class PatientFileAddRequest {
//     patientfile: string;
//     spcid: number;
// }

export class PatientDetaFilterRequest {
    patientname: string;
    age: number;
    emergencycontactno: string;
    patientregistrationno: string;
    ward: string;
    bedno: number;
    admissiondate: Date;
    dischargedate: Date;
    statusid: number;
}
export class PatientDataListResponse {
    patientdetails: string;
    medicaldetails: string;
    patientfiletemplate: string;
    patientid: number;
}
export class StatusChangeRequest {
    status: number;
    patientid: number;
    // discharge:  Date;
}

export class PatientTxnRequest {
    spid: number;
    // startdate: Date;
    // enddate: Date;
    servinid: number;
}

export class PersonDataAddRequest {
    reportname: string;
    age: number;
    persondate: Date;
    patientregistrationno: string;
    ward: number;
    bedno: number;
    admissiondate: Date;
    dischargedate: Date;
    bloodgroup: string;
    weight: number;
    drinst: string;
    gender: string;
}

//Chandan......
// patient master add
export class PatientAddRequest {
    uuid: string;
    patientregno: string;
    fname: string;
    lname: string;
    mobno: string;
    age: string;
    bloodgrp: string;
    gender: string;
}

//Master List Display
export class PatientRequestFilterDataModel {
    cpmid: number;
    fname: string;
    lname: string;
    mobno: string;
    spid: number;
    bedno: string;
    status: number;
    admittedon: Date;
}

// patient master add
export class AdmissionAddRequest {
    uuid: string;
    patientid: number;
    patientregno: string;
    bedno: string;
    status: number;
    spid: number;
    drincharge: number;
    admittedon: Date;
    dischargedon: Date;
}


//Patient List for Search the list first
export class PatientFilterRequest {
    fname: string;
    lname: string;
    mobno: string;
    spid: number;
    bedno: string;
    status: number;
    patientregno: string;
    admittedon: Date;
}


//Patient List Response for display first list
export class PatientDetaListResponse {
    patientid: number;
    admissionid: number;
    cpmid: number;
    fname: string;
    lname: string;
    mobno: string;
    bedno: string;
    status: number;
    spid: number;
    patientregno: string;
    drincharge: number;
    admittedon: Date;
    dischargedon: Date;
}


//for search  in add patient
export class PatientSearchRequestFilter {

    fname: string;
    lname: string;
    mobno: string;
}

//patient list for add patient
export class PatientSearchResponseFilter {
    patientid: number;
    cpmid: number;
    patientregno: string;
    fname: string;
    lname: string;
    mobno: string;
    age: string;
    bloodgrp: string;
    gender: string;
    createdon: Date;
    updatedon: Date;
}

//Patient Update Request
export class PatientUpdateRequest {
    patientid: number;
    patientregno: string;
    fname: string;
    lname: string;
    mobno: string;
    age: string;
    bloodgrp: string;
    gender: string;
}

//Patient Update Response
export class PatientUpdateResponse {
    patientid: number;
    admissionid: number;
    cpmid: number;
    patientregno: string;
    fname: string;
    lname: string;
    mobno: string;
    age: string;
    bloodgrp: string;
    gender: string;
    createdon: Date;
    updatedon: Date;
}

//update admission request
export class AdmissionUpdateRequest {
    admissionid: number;
    uuid: string;
    patientid: number;
    patientregno: string;
    bedno: string;
    status: number;
    spid: number;
    drincharge: number;
    admittedon: Date;
    dischargedon: Date;
}

// Admission Info Response
export class AdmissionUpdateResponse {
    admissionid: number;
    cpmid: number;
    patientid: number;
    patientregno: string;
    bedno: string;
    status: number;
    spid: number;
    drincharge: number;
    admittedon: Date;
    dischargedon: Date;
    createdon: Date;
    updatedon: Date;
}

//Status Api for patient to show where patient is hospitalized or not
export class PatientStatus {
    admissionid: number;
    status: number;
}

//Define for gloablly
export class JSONBaseDataModel<T> {
    version: number;
    data: T;
}


//person accompanying start from here.

export class PersonAccompanyingInfo {
    name: string;
    pesrsonage: string;
    gender: number;
    personaddress: string;
    relationshipwithpatient: string;
    contact: string;
    alternatecontact: string;
}
// Personal detail request
export class PersonalDetailsRequest {
    personaldetailsid: number;
    patientid: number;
    admissionid: number;
    uuid: string;
    age: string;
    otherdetails: string;
    personaccompanying: string;
}
// Personal detail response
export class PersonalDetailsResponse {
    personaldetailsid: number;
    patientid: number;
    admissionid: number;
    uuid: string;
    age: string;
    otherdetails: string;
    personaccompanying: string;
}



//
export class AdmissionAddResponseModel {
    admissionid: number;
    personaldetailsid: number;
    medicaldetailsid: number;
}

//status patient discharge
export class AdmissionStatusRequest {
    admissionid: number;
    status: number;
    dischargedon: Date;
}

//Medical data start from here
export class JSONInnerData {
    text: string;
    date: Date;
}

export class MedicalDetailsRequest {
    uuid: string;
    patientid: number;
    admissionid: number;
    presentcomplaints: string;
    reasonforadmission: string;
    historypresentillness: string;
    pasthistory: string;
    treatmentbeforeadmission: string;
    investigationbeforeadmission: string;
    familyhistory: string;
    allergies: string;
    personalhistory: string;
    medicaldetialsid: number;
}

export class MedicalDetailsResponse {
    uuid: string;
    patientid: number;
    admissionid: number;
    presentcomplaints: string;
    reasonforadmission: string;
    historypresentillness: string;
    pasthistory: string;
    treatmentbeforeadmission: string;
    investigationbeforeadmission: string;
    familyhistory: string;
    allergies: string;
    personalhistory: string;
    medicaldetialsid: number;
}
export class MedicalDetailsResponseJSON {

    presentcomplaints: JSONBaseDataModel<JSONInnerData>;
    reasonforadmission: JSONBaseDataModel<JSONInnerData>;
    historypresentillness: JSONBaseDataModel<JSONInnerData>;
    pasthistory: JSONBaseDataModel<JSONInnerData>;
    treatmentbeforeadmission: JSONBaseDataModel<JSONInnerData>;
    investigationbeforeadmission: JSONBaseDataModel<JSONInnerData>;
    familyhistory: JSONBaseDataModel<JSONInnerData>;
    allergies: JSONBaseDataModel<JSONInnerData>;
    personalhistory: JSONBaseDataModel<JSONInnerData>;
}

export class PresentComplaint {
    medicaldetialsid: number;
    presentcomplaints: string;
}

export class ReasonForAdmission {

    medicaldetialsid: number;
    reasonforadmission: string;
}

export class HistoryPresentIllness {
    medicaldetialsid: number;
    historypresentillness: string;
}

export class PastHistory {
    medicaldetialsid: number;
    pasthistory: string;
}

export class TreatmentBeforeAdmission {
    medicaldetialsid: number;
    treatmentbeforeadmission: string;
}

export class InvestigationBeforeAdmission {
    medicaldetialsid: number;
    investigationbeforeadmission: string;
}

export class FamilyHistory {
    medicaldetialsid: number;
    familyhistory: string;
}

export class Allergies {
    medicaldetialsid: number;
    allergies: string;
}

//Medical Personal history start from here
export class PersonalHistoryInfo {
    medicaldetialsid: number;
    weight: string;
    weighttendency: string;
    alcohalquantity: string;
    alcohalcomment: string;
    smokingquantity: string;
    smokingcomment: string;
    other: string;
    date: Date;
}

//Medical data start from here
export class WaightTendencyInfo {
    tendency: number;
}

export class PersonalHistory {
    medicaldetialsid: number;
    personalhistory: string;
}



//Pethology start from here
export class PathologyResponse{
    id: number
    uuid: string;
    pathologyid: number;
    admissionid: number;
    testperformed: string;
    treatmentid: number;
    comments: string;
    testresult: string;
    documentid: number;
    documentuuid: string;
    documentname: string;
    testperformedtime: string;
    value: string;
    // txndata: T;
}
export class PathologyFilterRequest
{
    admissionid: number;
}

//Treatment start from here
export class TreatmentResponse{
    id: number
    uuid: string;
    treatmentid: number;
    admissionid: number;
    treatmentdone: string;
    details: string;
    postobservation: string;
    documentid: number;
    documentuuid: string;
    documentname: string;
    treatmentperformedtime: string;
    value: string;
}
export class TreatmentFilterRequest
{
    admissionid: number;
}
 
export class CheckPatientRequest
{
    patientid: number;
}
export class CheckPatientResponse
{
    patientid: number;
    status: number;
}