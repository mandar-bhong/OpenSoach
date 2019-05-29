import { ActionDataStoreModel } from "../db/action-datastore";
import { DataActionItem } from "./action-model";
import { SchedularConfigData } from "./chart-models";
import { ACTION_STATUS } from "~/app/app-constants";
//import { Observable } from "tns-core-modules/data/observable";

export class ActionItemVMModel {
    actionItemDBModel: ActionDataStoreModel
    dbModel: DataActionItem;
    actionStatus: ACTION_STATUS;
    configData: SchedularConfigData;
    iscompleted: boolean;
    txnData: any;
    isActionActive: boolean;
    hasTxnData: boolean;
    doctorOrderModel: DoctorsOrderItem;
    conf_type_code: string;
}

export class DoctorsOrderItem {
    uuid: string;
    admission_uuid: string
    doctor_id: number
    doctors_orders: string
    comment: string
    ack_by: number
    ack_time: string
    status: number
    order_created_time: string
    order_type: string
    document_uuid: string
    document_name: string
    doctype: string
    updated_by: number
    updated_on: string
    sync_pending: number
    client_updated_at: string
    ack_by_name: string
    order_by_name:string
}
