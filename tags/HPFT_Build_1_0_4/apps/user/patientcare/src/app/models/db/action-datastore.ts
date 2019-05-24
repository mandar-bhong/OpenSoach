
import { IDatastoreModel } from './idatastore-model.js';

export class ActionDataStoreModel implements IDatastoreModel {
    uuid: string;
    admission_uuid: string;
    conf_type_code: string;
    schedule_uuid: string;
    scheduled_time: string;
    is_deleted: number;
    updated_by:number;
    updated_on:string;
    sync_pending: number;
    client_updated_at: string;
    getModelValues(): any[] {
        return [this.uuid, this.admission_uuid, this.conf_type_code, this.schedule_uuid, this.scheduled_time, this.is_deleted,this.updated_by,this.updated_on, this.sync_pending,this.client_updated_at];
    }
}

export class ActionsData {
    actions: ActionDataStoreModel[];
    enddate: string;
}