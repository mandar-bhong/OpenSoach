import { IDatastoreModel } from "./idatastore-model";

export class ServicePointDatastoreModel implements IDatastoreModel {
    uuid: string;
    sp_name: string;
    short_desc: string;
    sp_state: number;
    sp_state_since: string;
    updated_by: number;
    updated_on: string;
    sync_pending: number;
    client_updated_at: string;
    getModelValues(): any[] {
        return [this.uuid, this.sp_name, this.short_desc, this.sp_state, this.sp_state_since, this.updated_by, this.updated_on, this.sync_pending,this.client_updated_at];
    }
}