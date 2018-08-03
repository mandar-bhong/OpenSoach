import { MedicalDetailAddRequest, PatientDataAddRequest, PatientDetailAddRequest, StatusChangeRequest } from '../api/patient-models';

export class PatientDataModel {
    patientdetails: PatientDetailAddRequest;
    medicaldetails: MedicalDetailAddRequest;
    patientfiletemplate: number;
    patientid: number;
    spid: number;
    status: number;
    copyTo(patientDataAddRequest: PatientDataAddRequest) {
        patientDataAddRequest.patientdetails = JSON.stringify(this.patientdetails);
        patientDataAddRequest.medicaldetails = JSON.stringify(this.medicaldetails);
        patientDataAddRequest.spid = this.spid;
        patientDataAddRequest.patientfiletemplate = this.patientfiletemplate;
        patientDataAddRequest.status = this.status;
        // patientDataAddRequest.patientfiletemplate = JSON.stringify(this.patientfiletemplate);
        console.log('patientDataAddRequest', patientDataAddRequest);
    }
    copyFrom(patientDataAddRequest: PatientDataAddRequest) {
        this.patientid = patientDataAddRequest.patientid;
        this.patientdetails = new PatientDetailAddRequest();
        Object.assign(this.patientdetails, JSON.parse(patientDataAddRequest.patientdetails));
        this.medicaldetails = new MedicalDetailAddRequest();
        this.spid = patientDataAddRequest.spid;
        this.patientid = patientDataAddRequest.patientid;
        this.patientfiletemplate = patientDataAddRequest.patientfiletemplate;
        this.status = patientDataAddRequest.status;
        // Object.assign(this.medicaldetails, JSON.parse(patientDataAddRequest.medicaldetails));
        // this.patientfiletemplate = new PatientFileAddRequest();
        // Object.assign(this.patientfiletemplate, JSON.parse(patientDataAddRequest.patientfiletemplate));
    }
}

