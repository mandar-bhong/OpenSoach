import { IDatastoreModel } from "./idatastore-model";

export class DocumentUploadDatastore implements IDatastoreModel {

    uuid: string;
    doc_path: string;
    doc_name: string;
    updated_on: Date;
    updated_by: number;
    sync_pending: number;
    client_updated_at: Date;
    getModelValues(): any[] {
        return [this.uuid, this.doc_path, this.doc_name, this.updated_by, this.updated_on, this.sync_pending, this.client_updated_at];
    }
}