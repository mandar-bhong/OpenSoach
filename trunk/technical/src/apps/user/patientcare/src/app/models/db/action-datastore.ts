import { IDatastoreModel } from "./idatastore-model.js";

export class ActionDataStoreModel implements IDatastoreModel {
    uuid: string;
    admission_uuid: string;
    conf_type_code: string;
    schedule_uuid: string;
    sync_pending: number;
    exec_time: Date;
    getModelValues() {
        return [this.uuid, this.admission_uuid];
    }
}

export class ActionsData {
    actions: ActionDataStoreModel[];
    enddate: Date;
}