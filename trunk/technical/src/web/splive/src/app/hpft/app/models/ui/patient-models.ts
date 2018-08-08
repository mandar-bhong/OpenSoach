import { MedicalDetailAddRequest, PatientDataAddRequest, PatientDetailAddRequest, StatusChangeRequest } from '../api/patient-models';

export class PatientDataModel {
    patientdetails: PatientDetailAddRequest;
    medicaldetails: MedicalDetailAddRequest;
    patientfiletemplate: number;
    patientid: number;
    spid: number;
    spname: string;
    status: number;
    servinid: number;
    dischargedon: Date;
    copyTo(patientDataAddRequest: PatientDataAddRequest) {
        patientDataAddRequest.patientdetails = JSON.stringify(this.patientdetails);
        patientDataAddRequest.medicaldetails = JSON.stringify(this.medicaldetails);
        patientDataAddRequest.spid = this.spid;
        patientDataAddRequest.patientfiletemplate = this.patientfiletemplate;
        patientDataAddRequest.status = this.status;
        console.log('patientDataAddRequest', patientDataAddRequest);
    }
    copyFrom(patientDataAddRequest: PatientDataAddRequest) {
        this.patientid = patientDataAddRequest.patientid;
        this.patientdetails = new PatientDetailAddRequest();
        Object.assign(this.patientdetails, JSON.parse(patientDataAddRequest.patientdetails));
        this.medicaldetails = new MedicalDetailAddRequest();
        Object.assign(this.medicaldetails, JSON.parse(patientDataAddRequest.medicaldetails));
        this.spid = patientDataAddRequest.spid;
        this.spname = patientDataAddRequest.spname;
        this.patientid = patientDataAddRequest.patientid;
        this.dischargedon = patientDataAddRequest.dischargedon;
        this.patientfiletemplate = patientDataAddRequest.patientfiletemplate;
        this.status = patientDataAddRequest.status;
    }
}

