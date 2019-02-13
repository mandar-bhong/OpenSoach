import { IDatastoreModel } from "./idatastore-model";

export class PatientMedicalDetailsDatastoreModel implements IDatastoreModel {
    uuid: string;
    patient_uuid: string;
    admission_uuid: string;
    present_complaints: string;
    reason_for_admission: string;
    history_present_illness: string;
    past_history: string;
    treatment_before_admission: string;
    investigation_before_admission:string;
    family_history: string;
    allergies: string;
    personal_history: string;
    updated_by: number;
    updated_on: Date;
    sync_pending: number;
    client_updated_at: Date;
    getModelValues(): any[] {
        return [this.uuid, this.patient_uuid, this.admission_uuid, this.present_complaints,this.reason_for_admission, 
        this.history_present_illness,this.past_history,this.treatment_before_admission, this.investigation_before_admission, 
        this.family_history, this.allergies, this.personal_history, this.updated_by,this.updated_on, this.sync_pending,
        this.client_updated_at];
    }
}