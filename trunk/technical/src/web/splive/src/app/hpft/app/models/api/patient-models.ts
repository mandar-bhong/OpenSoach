export class PatientDataAddRequest {
    patientdetails: string;
    medicaldetails: string;
    patientfiletemplate: number;
    patientid: number;
    spid: number;
    status: number;
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
}
