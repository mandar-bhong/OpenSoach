import { IDatastoreModel } from "./idatastore-model.js";

export class ScheduleDatastoreModel implements IDatastoreModel {
    uuid: string;
    admission_uuid: string;
    conf_type_code: string;
    conf: string;
    end_date: string;
    updated_on: Date;
    sync_pending: number;
    client_updated_at: Date;
    getModelValues(): any[] {
        return [this.uuid, this.admission_uuid, this.conf_type_code, this.conf, this.end_date, this.updated_on, this.sync_pending, this.client_updated_at];
    }
}