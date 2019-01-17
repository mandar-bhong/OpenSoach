export class ActionListViewModel {
    dbmodel: any;
}

export class ActionDBModel {
    uuid: string;
    admission_uuid: string;
    conf_type_code: string;
    schedule_uuid: string;
    exec_time: Date;
}

export class ActionTxnDBModel {
    uuid: string;
    schedule_uuid: string;
    txn_data: string;
    txn_date: Date;
    txn_state: string;
    conf_type_code: string;
    runtime_config_data: string;
}