export class ActionListViewModel {
    dbmodel: any;
}

export class ActionDBModel {
    uuid: string;
    admission_uuid: string;
    conf_type_code: string;
    schedule_uuid: string;
    exec_time: Date;
    sync_pending: number;
}

export class ActionTxnDBModel {
    uuid: string;
    schedule_uuid: string;
    txn_data: string;
    txn_date: Date;
    txn_state: string;
    conf_type_code: string;
    runtime_config_data: string;
    status: number;
    admission_uuid: string;
}
export class ActionDataDBRequest {
    comment: string;
    value: string;
}
export class UserAuthDBRequest{
    user_fname: string;
    user_lname: string;
    email: string;
    pin: number;
}