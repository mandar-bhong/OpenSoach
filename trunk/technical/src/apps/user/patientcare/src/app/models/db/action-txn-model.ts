import { IDatastoreModel } from "./idatastore-model";

export class ActionTxnDatastoreModel implements IDatastoreModel {
    uuid: string;
    schedule_uuid: string;
    admission_uuid: string;
    txn_data: string;
    scheduled_time: Date;
    txn_state: number;
    conf_type_code: string;
    runtime_config_data: string;
    updated_by: number;
    updated_on: string;
    sync_pending: number;
    client_updated_at: string;
    getModelValues(): any[] {
        return [this.uuid, this.admission_uuid, this.schedule_uuid, this.txn_data, this.scheduled_time, this.txn_state, this.conf_type_code,
        this.updated_by,this.updated_on, this.runtime_config_data, this.sync_pending, this.client_updated_at];
    }
}

