export class HospitalFilterRequest {
    hospid: number;
    hospname: string;
}
export class HospitalFilterResponse {

}
//Users List display for hospital request method
export class PatientFilterRequest {
    prodcode: string;
    cpmid: number;
    status: number;
    admissionid: number;
    custname: string;
}
//Users List display for hospital response method
export class PatientFilterResponse {
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

export class LaboratoryReportAddRequest {
    uuid: string;
    admissionid: number;
    testperformed: string;
    testresult: string;
    comments: string;
    testperformedtime: string;
    documentuuidlist: string[];
}


export class HospitalSearchRequest {
    prodcode: string;
}

export class HospitalListResponse {
    custname: string;
    cpmid: number;
}