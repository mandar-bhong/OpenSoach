import { IDatastoreModel } from "./idatastore-model";

export class TreatmentDocDatastoreModel implements IDatastoreModel {
    uuid: string;
    sync_pending: number;
    client_updated_at: string;
    treatment_uuid: string;
    document_uuid: string;
    updated_on: string;
    getModelValues(): any[] {
        return [this.treatment_uuid, this.document_uuid];
    }
}