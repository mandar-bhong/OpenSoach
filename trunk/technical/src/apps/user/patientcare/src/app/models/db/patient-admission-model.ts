import { IDatastoreModel } from "./idatastore-model";

export class PatientAdmissionDatastoreModel implements IDatastoreModel {
    uuid: string;
    patient_uuid: string;
    patient_reg_no: string;
    bed_no: string;
    status: string;
    sp_uuid: string;
    dr_incharge: number;
    admitted_on: Date;
    discharged_on: Date;
    updated_on: Date;
    sync_pending: number;
    sync_pending_time: Date;
    getModelValues(): any[] {
        return [this.uuid, this.patient_uuid, this.patient_reg_no, this.bed_no, this.status, this.sp_uuid,
        this.dr_incharge, this.admitted_on, this.discharged_on, this.updated_on, this.sync_pending, this.sync_pending_time,
        this.sync_pending_time];
    }
}