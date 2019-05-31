import { HospitalFilterRequest, PatientFilterRequest } from "../api/hospital-models";

export class HospitalDataModel {
    hospid: number;
    hospname: string;
    copyTo(hospitalFilterRequest:HospitalFilterRequest){
        hospitalFilterRequest.hospid = this.hospid;
        hospitalFilterRequest.hospname = this.hospname;
    }
}


export class HospitalFilterModel {
    prodcode: string;
    cpmid: number;
    copyTo(patientDataListRequest: PatientFilterRequest) {
        patientDataListRequest.cpmid = this.cpmid;
        patientDataListRequest.prodcode = this.prodcode;
    }
}