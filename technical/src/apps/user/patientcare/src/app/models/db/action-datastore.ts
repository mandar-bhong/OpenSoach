
import { IDatastoreModel } from './idatastore-model.js';

export class ActionDataStoreModel implements IDatastoreModel {
    uuid: string;
    admission_uuid: string;
    conf_type_code: string;
    schedule_uuid: string;
    exec_time: Date;
    sync_pending: number;
    sync_pending_time: Date;
    getModelValues(): any[] {
        return [this.uuid, this.admission_uuid, this.conf_type_code, this.schedule_uuid, this.exec_time, this.sync_pending,this.sync_pending_time];
    }
}

export class ActionsData {
    actions: ActionDataStoreModel[];
    enddate: Date;
}