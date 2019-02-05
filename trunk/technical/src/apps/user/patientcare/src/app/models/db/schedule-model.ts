import { IDatastoreModel } from "./idatastore-model.js";

export class ScheduleDatastoreModel implements IDatastoreModel {
    uuid: string;
    admission_uuid: string;
    conf_type_code: string;
    conf: string;
    end_date: Date;
    sync_pending: number;
    sync_pending_time: Date;
    getModelValues(): any[] {
        return [this.uuid, this.admission_uuid, this.conf_type_code, this.conf, this.end_date, this.sync_pending, this.sync_pending_time];
    }
}