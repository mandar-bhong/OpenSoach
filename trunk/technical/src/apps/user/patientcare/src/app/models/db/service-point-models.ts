import { IDatastoreModel } from "./idatastore-model";

export class ServicePointDatastoreModel implements IDatastoreModel {
    uuid: string;
    sp_name: string;
    short_desc: string;
    sp_state: number;
    sp_state_since: Date;
    updated_on: Date;
    sync_pending: number;
    sync_pending_time: Date;
    getModelValues(): any[] {
        return [this.uuid, this.sp_name, this.short_desc, this.sp_state, this.sp_state_since, this.updated_on, this.sync_pending,this.sync_pending_time];
    }
}