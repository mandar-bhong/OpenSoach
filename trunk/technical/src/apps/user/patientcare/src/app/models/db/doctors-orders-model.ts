import { IDatastoreModel } from "./idatastore-model";

export class DoctorsOrdersDatastoreModel implements IDatastoreModel {
    uuid: string;
    admission_uuid: string;
    doctor_id: number;
    doctors_orders: string;
    document_uuid: string;
    updated_on: Date;
    sync_pending: number;
    client_updated_at: Date;
    getModelValues(): any[] {
        return [this.uuid,this.admission_uuid,this.doctor_id,this.doctors_orders,this.document_uuid,this.updated_on,
        this.sync_pending,this.client_updated_at];
    }
}