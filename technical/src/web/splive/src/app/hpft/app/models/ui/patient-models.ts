import {
    MedicalDetailAddRequest,
    PatientDataAddRequest,
    PatientDetailAddRequest,
    PatientAddRequest,
    PatientDetaListResponse,
    PatientFilterRequest,
    PatientSearchRequestFilter,
    PatientSearchResponseFilter,
    PatientUpdateRequest,
    PatientUpdateResponse,
    AdmissionAddRequest,
    AdmissionUpdateRequest,
    AdmissionUpdateResponse,
    JSONBaseDataModel,
    AdmissionStatusRequest,
    MedicalDetailsRequest,
    MedicalDetailsResponse,
    JSONInnerData,
    PersonalHistoryInfo,
    PersonAccompanyingInfo,
    PersonalDetailsResponse,
    PersonalDetailsRequest,
} from '../api/patient-models';
import { NumberCardModule } from '@swimlane/ngx-charts';
import { JSONBaseModel } from './json-base-model';

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

export class PatientAddModal {
    uuid: string;
    patientid: number;
    patientregno: string;
    fname: string;
    lname: string;
    mobno: string;
    age: string;
    bloodgrp: string;
    gender: string;
    cpmid: number;
    spid: number;
    bedno: string;
    status: number;
    admittedon: Date;
    createdon: Date;
    updatedon: Date;
    admissionid: number;
    // patientRequestFilterDataModel: PatientRequestFilterDataModel;
    copyTo(patientAddRequest: PatientAddRequest) {
        patientAddRequest.uuid = this.uuid;
        patientAddRequest.patientregno = this.patientregno;
        patientAddRequest.fname = this.fname;
        patientAddRequest.lname = this.lname;
        patientAddRequest.mobno = this.mobno;
        patientAddRequest.age = this.age;
        patientAddRequest.bloodgrp = this.bloodgrp;
        patientAddRequest.gender = this.gender;
    }
    CopyToUpdate(patientUpdateRequest: PatientUpdateRequest) {
        patientUpdateRequest.patientid = this.patientid;
        patientUpdateRequest.patientregno = this.patientregno;
        patientUpdateRequest.fname = this.fname;
        patientUpdateRequest.lname = this.lname;
        patientUpdateRequest.mobno = this.mobno;
        patientUpdateRequest.age = this.age;
        patientUpdateRequest.bloodgrp = this.bloodgrp;
        patientUpdateRequest.gender = this.gender;
    }
    CopyFromUpdateResponse(patientUpdateResponse: PatientUpdateResponse) {
        this.admissionid = patientUpdateResponse.admissionid;
        this.patientid = patientUpdateResponse.patientid;
        this.cpmid = patientUpdateResponse.cpmid;
        this.patientregno = patientUpdateResponse.patientregno;
        this.fname = patientUpdateResponse.fname;
        this.lname = patientUpdateResponse.lname;
        this.mobno = patientUpdateResponse.mobno;
        this.age = patientUpdateResponse.age;
        this.bloodgrp = patientUpdateResponse.bloodgrp;
        this.gender = patientUpdateResponse.gender;
        this.createdon = patientUpdateResponse.createdon;
        this.updatedon = patientUpdateResponse.updatedon;
    }

}

export class AdmissionAddModel {
    uuid: string;
    patientid: number;
    patientregno: string;
    bedno: string;
    status: number;
    spid: number;
    drincharge: number;
    admittedon: Date;
    dischargedon: Date;
    admissionid: number;
    cpmid: number;
    createdon: Date;
    updatedon: Date;
    copyTo(admissionAddRequest: AdmissionAddRequest) {
        admissionAddRequest.uuid = this.uuid;
        admissionAddRequest.patientid = this.patientid;
        admissionAddRequest.patientregno = this.patientregno;
        admissionAddRequest.bedno = this.bedno;
        admissionAddRequest.status = this.status;
        admissionAddRequest.spid = this.spid;
        admissionAddRequest.drincharge = this.drincharge;
        admissionAddRequest.admittedon = this.admittedon;
        admissionAddRequest.dischargedon = this.dischargedon;
    }
    copyToUpdate(admissionUpdateRequest: AdmissionUpdateRequest) {
        admissionUpdateRequest.admissionid = this.admissionid;
        admissionUpdateRequest.uuid = this.uuid;
        admissionUpdateRequest.patientid = this.patientid;
        admissionUpdateRequest.patientregno = this.patientregno;
        admissionUpdateRequest.bedno = this.bedno;
        admissionUpdateRequest.status = this.status;
        admissionUpdateRequest.spid = this.spid;
        admissionUpdateRequest.drincharge = this.drincharge;
        admissionUpdateRequest.admittedon = this.admittedon;
        admissionUpdateRequest.dischargedon = this.dischargedon;
    }
    copyFromUpdateResponse(admissionUpdateResponse: AdmissionUpdateResponse) {
        this.admissionid = admissionUpdateResponse.admissionid;
        this.cpmid = admissionUpdateResponse.cpmid;
        this.patientid = admissionUpdateResponse.patientid;
        this.patientregno = admissionUpdateResponse.patientregno;
        this.bedno = admissionUpdateResponse.bedno;
        this.status = admissionUpdateResponse.status;
        this.spid = admissionUpdateResponse.spid;
        this.drincharge = admissionUpdateResponse.drincharge;
        this.admittedon = admissionUpdateResponse.admittedon;
        this.dischargedon = admissionUpdateResponse.dischargedon;
        this.createdon = admissionUpdateResponse.createdon;
        this.updatedon = admissionUpdateResponse.updatedon;
    }


}

export class PatientFilterModel {
    fname: string;
    lname: string;
    mobno: string;
    spid: number;
    bedno: string;
    status: number;
    admittedon: Date;
    patientregno: string;
    copyTo(patientDataListRequest: PatientFilterRequest) {
        patientDataListRequest.fname = this.fname;
        patientDataListRequest.bedno = this.bedno;
        patientDataListRequest.lname = this.lname;
        patientDataListRequest.spid = this.spid;
        patientDataListRequest.mobno = this.mobno;
        patientDataListRequest.patientregno = this.patientregno;
        patientDataListRequest.status = this.status;
        patientDataListRequest.admittedon = this.admittedon;
    }
}
export class PatientListDataModel {
    patientid: number;
    admissionid: number;
    cpmid: number;
    fname: string;
    lname: string;
    mobno: string;
    bedno: string;
    status: number;
    spid: number;
    drincharge: number;
    dischargedon: Date;
    copyFrom(patientDetaListResponse: PatientDetaListResponse) {
        this.patientid = patientDetaListResponse.patientid;
        this.admissionid = patientDetaListResponse.admissionid;
        this.cpmid = patientDetaListResponse.cpmid;
        this.fname = patientDetaListResponse.fname;
        this.lname = patientDetaListResponse.lname;
        this.mobno = patientDetaListResponse.mobno;
        this.bedno = patientDetaListResponse.bedno;
        this.status = patientDetaListResponse.status;
        this.spid = patientDetaListResponse.spid;
        this.drincharge = patientDetaListResponse.drincharge;
        this.dischargedon = patientDetaListResponse.dischargedon;
    }
    copyToStatus(admissionStatusRequest: AdmissionStatusRequest) {
        admissionStatusRequest.admissionid = this.admissionid;
        admissionStatusRequest.status = this.status;
        admissionStatusRequest.dischargedon = this.dischargedon;
    }
}

export class PatientCheckModal {
    fname: string;
    lname: string;
    mobno: string;
    copyTo(patientSearchRequestFilter: PatientSearchRequestFilter) {
        patientSearchRequestFilter.fname = this.fname;
        patientSearchRequestFilter.lname = this.lname;
        patientSearchRequestFilter.mobno = this.mobno;
    }
}

export class PatientCheckListDataModal {
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
    copyFrom(patientSearchResponseFilter: PatientSearchResponseFilter) {
        this.patientid = patientSearchResponseFilter.patientid;
        this.cpmid = patientSearchResponseFilter.cpmid;
        this.patientregno = patientSearchResponseFilter.patientregno;
        this.fname = patientSearchResponseFilter.fname;
        this.lname = patientSearchResponseFilter.lname;
        this.mobno = patientSearchResponseFilter.mobno;
        this.age = patientSearchResponseFilter.age;
        this.bloodgrp = patientSearchResponseFilter.bloodgrp;
        this.gender = patientSearchResponseFilter.gender;
        this.createdon = patientSearchResponseFilter.createdon;
        this.updatedon = patientSearchResponseFilter.updatedon;
    }
}

export class PatientPersonalDetails {
    personaldetailsid: number;
    patientid: number;
    admissionid: number;
    uuid: string;
    otherdetails: string;
    personAccompanyingData: JSONBaseDataModel<PersonAccompanyingInfo[]>;
    copyToAddPerson(personalDetailsRequest: PersonalDetailsRequest) {
        personalDetailsRequest.personaldetailsid = this.personaldetailsid;
        personalDetailsRequest.uuid = this.uuid;
        personalDetailsRequest.patientid = this.patientid;
        personalDetailsRequest.admissionid = this.admissionid;
        personalDetailsRequest.personaccompanying = JSON.stringify(this.personAccompanyingData);
    }
    copyFromPerson(personalDetailsResponse: PersonalDetailsResponse) {
        this.personAccompanyingData = new JSONBaseDataModel<PersonAccompanyingInfo[]>();
        const tempPersonAccompanying = JSON.parse(personalDetailsResponse.personaccompanying);
        this.personAccompanyingData.data = [];
        this.personAccompanyingData.data = tempPersonAccompanying.data || null;
        this.personAccompanyingData.version = tempPersonAccompanying.version;
    }

}

export class MedicalDetailsModel {
    uuid: string;
    patientid: number;
    admissionid: number;
    medicaldetialsid: number;
    presentComplaintsData: JSONBaseDataModel<JSONInnerData[]>;
    reasonForAdmissionData: JSONBaseDataModel<JSONInnerData[]>;
    historyPresentIllnessData: JSONBaseDataModel<JSONInnerData[]>;
    pastHistoryData: JSONBaseDataModel<JSONInnerData[]>;
    treatmentBeforeAdmissionData: JSONBaseDataModel<JSONInnerData[]>;
    investigationBeforeAdmissionData: JSONBaseDataModel<JSONInnerData[]>;
    familyHistoryData: JSONBaseDataModel<JSONInnerData[]>;
    allergiesData: JSONBaseDataModel<JSONInnerData[]>;
    personalHistoryData: JSONBaseDataModel<PersonalHistoryInfo>;
    copyToAdd(medicalDetailsRequest: MedicalDetailsRequest) {
        medicalDetailsRequest.uuid = this.uuid;
        medicalDetailsRequest.patientid = this.patientid;
        medicalDetailsRequest.admissionid = this.admissionid;
        medicalDetailsRequest.presentcomplaints = JSON.stringify(this.presentComplaintsData);
        medicalDetailsRequest.reasonforadmission = JSON.stringify(this.reasonForAdmissionData);
        medicalDetailsRequest.historypresentillness = JSON.stringify(this.historyPresentIllnessData);
        medicalDetailsRequest.pasthistory = JSON.stringify(this.pastHistoryData);
        medicalDetailsRequest.treatmentbeforeadmission = JSON.stringify(this.treatmentBeforeAdmissionData);
        medicalDetailsRequest.investigationbeforeadmission = JSON.stringify(this.investigationBeforeAdmissionData);
        medicalDetailsRequest.familyhistory = JSON.stringify(this.familyHistoryData);
        medicalDetailsRequest.allergies = JSON.stringify(this.allergiesData);
        medicalDetailsRequest.personalhistory = JSON.stringify(this.personalHistoryData);
    }
    copyFrom(medicalDetailsResponse: MedicalDetailsResponse) {

        this.medicaldetialsid = medicalDetailsResponse.medicaldetialsid;
        if (medicalDetailsResponse.presentcomplaints != null) {
            this.presentComplaintsData = new JSONBaseDataModel<JSONInnerData[]>();
            const tempPresentComplaints = JSON.parse(medicalDetailsResponse.presentcomplaints);
            this.presentComplaintsData.data = [];
            this.presentComplaintsData.data = tempPresentComplaints.data || [];
            this.presentComplaintsData.version = tempPresentComplaints.version;
        }

        if (medicalDetailsResponse.reasonforadmission != null) {
            this.reasonForAdmissionData = new JSONBaseDataModel<JSONInnerData[]>();
            const tempReasonForAdmission = JSON.parse(medicalDetailsResponse.reasonforadmission);
            this.reasonForAdmissionData.data = [];
            this.reasonForAdmissionData.data = tempReasonForAdmission.data || [];
            this.reasonForAdmissionData.version = tempReasonForAdmission.version;
        }

        if (medicalDetailsResponse.historypresentillness != null) {
            this.historyPresentIllnessData = new JSONBaseDataModel<JSONInnerData[]>();
            const temphistoryPresentIllness = JSON.parse(medicalDetailsResponse.historypresentillness);
            this.historyPresentIllnessData.data = [];
            this.historyPresentIllnessData.data = temphistoryPresentIllness.data || [];
            this.historyPresentIllnessData.version = temphistoryPresentIllness.version;
        }

        if (medicalDetailsResponse.pasthistory != null) {
            this.pastHistoryData = new JSONBaseDataModel<JSONInnerData[]>();
            const temppastHistory = JSON.parse(medicalDetailsResponse.pasthistory);
            this.pastHistoryData.data = [];
            this.pastHistoryData.data = temppastHistory.data || [];
            this.pastHistoryData.version = temppastHistory.version;
        }

        if (medicalDetailsResponse.treatmentbeforeadmission != null) {
            this.treatmentBeforeAdmissionData = new JSONBaseDataModel<JSONInnerData[]>();
            const tempTreatmentBeforeAdmission = JSON.parse(medicalDetailsResponse.treatmentbeforeadmission);
            this.treatmentBeforeAdmissionData.data = [];
            this.treatmentBeforeAdmissionData.data = tempTreatmentBeforeAdmission.data || [];
            this.treatmentBeforeAdmissionData.version = tempTreatmentBeforeAdmission.version;
        }

        if (medicalDetailsResponse.investigationbeforeadmission != null) {
            this.investigationBeforeAdmissionData = new JSONBaseDataModel<JSONInnerData[]>();
            const tempInvestigationBeforeAdmission = JSON.parse(medicalDetailsResponse.investigationbeforeadmission);
            this.investigationBeforeAdmissionData.data = [];
            this.investigationBeforeAdmissionData.data = tempInvestigationBeforeAdmission.data || [];
            this.investigationBeforeAdmissionData.version = tempInvestigationBeforeAdmission.version;
        }

        if (medicalDetailsResponse.familyhistory != null) {
            this.familyHistoryData = new JSONBaseDataModel<JSONInnerData[]>();
            const tempFamilyHistory = JSON.parse(medicalDetailsResponse.familyhistory);
            this.familyHistoryData.data = [];
            this.familyHistoryData.data = tempFamilyHistory.data || [];
            this.familyHistoryData.version = tempFamilyHistory.version;
        }

        if (medicalDetailsResponse.allergies != null) {
            this.allergiesData = new JSONBaseDataModel<JSONInnerData[]>();
            const tempAllergies = JSON.parse(medicalDetailsResponse.allergies);
            this.allergiesData.data = [];
            this.allergiesData.data = tempAllergies.data || [];
            this.allergiesData.version = tempAllergies.version;
        }

        if (medicalDetailsResponse.personalhistory != null) {
            this.personalHistoryData = new JSONBaseDataModel<PersonalHistoryInfo>();
            const tempPersonalHistory = JSON.parse(medicalDetailsResponse.personalhistory);
            this.personalHistoryData.data = tempPersonalHistory || null;
            this.personalHistoryData.version = tempPersonalHistory.version;
        }
    }
}

export class PathologyModel {
    pathologyid: number;
    admissionid: number;
    testperformed: string;
    testresult: string;
    comments: string;
    id: number
    uuid: string;
    treatmentid: number;
    documentid: number;
    documentuuid: string;
    name: string;
}

export class TreatmentModel {
    treatmentid: number;
    admissionid: number;
    treatmentdone: string;
}