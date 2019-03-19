import { IDatastoreModel } from "./idatastore-model.js";

export class ScheduleDatastoreModel implements IDatastoreModel {
    uuid: string;
    admission_uuid: string;
    conf_type_code: string;
    conf: string;
    end_date: string;
    updated_by: number;
    updated_on: string;
    sync_pending: number;
    client_updated_at: string;
    status: number; // 0: ACTIVE, 1: CANCEL
    getModelValues(): any[] {
        return [this.uuid, this.admission_uuid, this.conf_type_code, this.conf, this.end_date, this.status, this.updated_by, this.updated_on, this.sync_pending, this.client_updated_at];
    }
}