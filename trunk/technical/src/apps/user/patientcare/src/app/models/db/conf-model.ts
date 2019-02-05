import { IDatastoreModel } from "./idatastore-model";

export class ConfDatastoreModel implements IDatastoreModel {
    uuid: string;
    conf_type_code: string;
    conf: string;
    updated_on: Date;
    sync_pending: number;
    sync_pending_time: Date;
    getModelValues(): any[] {
        return [this.uuid, this.conf_type_code, this.conf, this.updated_on, this.sync_pending,this.sync_pending_time];
    }
}