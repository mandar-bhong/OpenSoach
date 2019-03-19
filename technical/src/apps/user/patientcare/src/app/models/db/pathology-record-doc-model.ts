import { IDatastoreModel } from "./idatastore-model";

export class PathologyRecordDocDatastoreModel implements IDatastoreModel {
    uuid: string;
    sync_pending: number;
    client_updated_at: string;
    pathology_record_uuid:string;
    document_uuid:string;
    getModelValues(): any[] {
        return [this.pathology_record_uuid,this.document_uuid];
    }
}