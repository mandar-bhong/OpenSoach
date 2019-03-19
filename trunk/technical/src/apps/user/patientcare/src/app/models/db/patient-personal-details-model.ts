import { IDatastoreModel } from "./idatastore-model";

export class PatientPersonalDetailsDatastoreModel implements IDatastoreModel {
    uuid: string;
    patient_uuid: string;
    admission_uuid: string;
    age: string;
    other_details: string;
    person_accompanying:string;
    updated_by: number;
    updated_on: Date;
    sync_pending: number;
     client_updated_at: string;
    getModelValues(): any[] {
        return [this.uuid, this.patient_uuid, this.admission_uuid, this.age,this.other_details,this.person_accompanying,
        this.updated_by,this.updated_on, this.sync_pending,this.client_updated_at];
    }
}