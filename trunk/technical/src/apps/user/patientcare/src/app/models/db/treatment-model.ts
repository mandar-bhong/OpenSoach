import { IDatastoreModel } from "./idatastore-model";

export class TreatmentDatastoreModel implements IDatastoreModel {
    uuid: string;
    admission_uuid: string;
    treatment_done: string;
    details: string;
    post_observation: string;
    updated_by: number;
    updated_on: Date;
    sync_pending: number;
    client_updated_at: Date;
    getModelValues(): any[] {
        return [this.uuid,this.admission_uuid,this.treatment_done,this.details,this.post_observation,this.updated_by,
            this.updated_on,this.sync_pending,this.client_updated_at];
    }
}