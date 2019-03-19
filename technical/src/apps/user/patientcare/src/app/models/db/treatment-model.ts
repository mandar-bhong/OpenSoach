import { IDatastoreModel } from "./idatastore-model";

export class TreatmentDatastoreModel implements IDatastoreModel {
    uuid: string;
    admission_uuid: string;
    treatment_done: string;
    treatment_performed_time: Date;
    details: string;
    post_observation: string;
    updated_by: number;
    updated_on: Date;
    sync_pending: number;
    client_updated_at: string;
    getModelValues(): any[] {
        return [this.uuid, this.admission_uuid, this.treatment_done, this.treatment_performed_time, this.details, this.post_observation, this.updated_by,
        this.updated_on, this.sync_pending, this.client_updated_at];
    }
}