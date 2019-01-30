import { IDatastoreModel } from "./idatastore-model";

export class ActionTxnDatastoreModel implements IDatastoreModel {
    uuid: string;
    schedule_uuid: string;
    admission_uuid: string;
    txn_data: string;
    txn_date: Date;
    txn_state: number;
    conf_type_code: string;
    runtime_config_data: string;
    updated_on: Date;
    sync_pending: number;
    getModelValues(): any[] {
         return [this.uuid, this.admission_uuid, this.schedule_uuid, this.txn_data, this.txn_date, this.txn_state, this.conf_type_code,
        this.updated_on,this.runtime_config_data, this.sync_pending];        
    }
}

