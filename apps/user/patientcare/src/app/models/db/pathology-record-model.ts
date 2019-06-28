import { IDatastoreModel } from "./idatastore-model";

export class PathologyRecordDatastoreModel implements IDatastoreModel {
    uuid: string;
    admission_uuid: string;
    test_performed: string;
    test_performed_time: Date;
    test_result: string;
    comments: string;
    updated_by: number;
    updated_on: string;
    sync_pending: number;
    client_updated_at: string;
    getModelValues(): any[] {
        return [this.uuid, this.admission_uuid, this.test_performed, this.test_performed_time, this.test_result, this.comments, this.updated_by,
        this.updated_on, this.sync_pending, this.client_updated_at];
    }
}