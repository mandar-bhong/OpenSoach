import { IDatastoreModel } from "./idatastore-model";

export class DoctorsOrdersDatastoreModel implements IDatastoreModel {
    uuid: string;
    admission_uuid: string;
    doctor_id: number;
    doctors_orders: string;
    comment: string;
    ack_by: number;
    ack_time: Date;
    status: number;
    order_created_time: Date;
    order_type: string;
    document_uuid: string;
    document_name: string;
    doctype: string;
    updated_by: number;
    updated_on: string;
    sync_pending: number;
    client_updated_at: string;
    getModelValues(): any[] {
        return [this.uuid, this.admission_uuid, this.doctor_id, this.doctors_orders, this.comment, this.ack_by,
        this.ack_time, this.status, this.order_created_time, this.order_type, this.document_uuid, this.document_name,
        this.doctype, this.updated_by, this.updated_on,
        this.sync_pending, this.client_updated_at];
    }
}