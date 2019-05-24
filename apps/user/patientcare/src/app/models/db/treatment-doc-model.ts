import { IDatastoreModel } from "./idatastore-model";

export class TreatmentDocDatastoreModel implements IDatastoreModel {
    uuid: string;
    treatment_uuid: string;
    document_uuid: string;
    document_name: string;
    doctype: string;
    updated_on: string;
    sync_pending: number;
    client_updated_at: string;
    getModelValues(): any[] {
        return [this.treatment_uuid, this.document_uuid,this.document_name,this.doctype,this.updated_on,this.sync_pending,this.client_updated_at];
    }
}