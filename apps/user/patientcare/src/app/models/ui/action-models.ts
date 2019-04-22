import { DataActionItem } from "./action-model";

export class ActionListViewModel {
    dbmodel: any;
}

export class ActionDBModel {
    uuid: string;
    admission_uuid: string;
    conf_type_code: string;
    schedule_uuid: string;
    scheduled_time: string;
    sync_pending: number;
}

export class ActionTxnDBModel {
    uuid: string;
    schedule_uuid: string;
    txn_data: string;
    scheduled_time: string;
    txn_state: number;
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
    userid: number;
    user_fname: string;
    user_lname: string;
    email: string;
    pin: number;
}
export class UserCreateFormRequest{
    email: string;
    password: string;
    newpin: string;
    reenterpin: string
}

