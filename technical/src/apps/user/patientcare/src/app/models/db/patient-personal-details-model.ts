import { IDatastoreModel } from "./idatastore-model";

export class PatientPersonalDetailsDatastoreModel implements IDatastoreModel {
    uuid: string;
    patient_uuid: string;
    admission_uuid: string;
    age: string;
    weight: string;
    other_details: string;
    updated_on: Date;
    sync_pending: number;
    client_updated_at: Date;
    getModelValues(): any[] {
        return [this.uuid, this.patient_uuid, this.admission_uuid, this.age, this.weight, this.other_details,
        this.updated_on, this.sync_pending,this.client_updated_at];
    }
}