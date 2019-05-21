import { IDatastoreModel } from "./idatastore-model";

export class PatientAdmissionDatastoreModel implements IDatastoreModel {
    uuid: string;
    patient_uuid: string;
    patient_reg_no: string;
    bed_no: string;
    status: number;
    sp_uuid: string;
    dr_incharge: number;
    admitted_on: string;
    discharged_on: string;
    updated_by: number;
    updated_on: string;
    sync_pending: number;
    client_updated_at: string;
    getModelValues(): any[] {
        return [this.uuid, this.patient_uuid, this.patient_reg_no, this.bed_no, this.status, this.sp_uuid,
        this.dr_incharge, this.admitted_on, this.discharged_on, this.updated_by, this.updated_on, this.sync_pending, this.client_updated_at];
    }
}