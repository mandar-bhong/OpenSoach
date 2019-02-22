import { IDatastoreModel } from "./idatastore-model";

export class PathologyRecordDatastoreModel implements IDatastoreModel {
    uuid: string;
    admission_uuid:string;
    test_performed:string;
    test_result:string;
    comments:string;
    updated_by: number;
    updated_on: Date;
    sync_pending: number;
    client_updated_at: Date;
    getModelValues(): any[] {
        return [this.uuid,this.admission_uuid,this.test_performed,this.test_result,this.comments,this.updated_by,
        this.updated_on,this.sync_pending,this.client_updated_at];
    }
}