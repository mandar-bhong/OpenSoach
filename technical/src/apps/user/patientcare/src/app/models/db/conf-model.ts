import { IDatastoreModel } from "./idatastore-model";

export class ConfDatastoreModel implements IDatastoreModel {
    uuid: string;
    conf_type_code: string;
    conf: string;
    updated_on: Date;
    updated_by: number;
    sync_pending: number;
    client_updated_at: Date;
    getModelValues(): any[] {
        return [this.uuid, this.conf_type_code, this.conf, this.updated_by, this.updated_on, this.sync_pending,this.client_updated_at];
    }
}