import { IDatastoreModel } from "./idatastore-model.js";

export class ScheduleDatastoreModel implements IDatastoreModel {
    uuid: string;
    admission_uuid: string;
    conf_type_code: string;
    conf: string;
    updated_on: Date;
    sync_pending: number;
    getModelValues(): any[] {
        return [this.uuid, this.admission_uuid, this.conf_type_code, this.conf, this.updated_on, this.sync_pending];
    }
}