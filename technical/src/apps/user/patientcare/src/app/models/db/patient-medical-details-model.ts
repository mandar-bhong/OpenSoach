import { IDatastoreModel } from "./idatastore-model";

export class PatientMedicalDetailsDatastoreModel implements IDatastoreModel {
    uuid: string;
    patient_uuid: string;
    admission_uuid: string;
    reason_for_admission: string;
    patient_medical_hist: string;
    treatment_recieved_before: string;
    family_hist: string;
    menstrual_hist: string;
    allergies: string;
    personal_hist: string;
    general_physical_exam: string;
    systematic_exam: string;
    updated_on: Date;
    sync_pending: number;
    getModelValues(): any[] {
        return [this.uuid, this.patient_uuid, this.admission_uuid, this.reason_for_admission, this.patient_medical_hist,
        this.treatment_recieved_before, this.family_hist, this.menstrual_hist, this.allergies, this.personal_hist,
        this.general_physical_exam, this.systematic_exam, this.updated_on, this.sync_pending];
    }
}