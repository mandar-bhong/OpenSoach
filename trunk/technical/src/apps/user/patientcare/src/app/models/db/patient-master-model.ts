import { IDatastoreModel } from "./idatastore-model";

export class PatientMasterDatastoreModel implements IDatastoreModel {
    uuid: string;
    patient_reg_no: string;
    fname: string;
    lname: string;
    mob_no: string;
    age: string;
    blood_grp: string;
    gender: string;
    updated_on: Date;
    sync_pending: number;
    client_updated_at: Date;
    getModelValues(): any[] {
        return [this.uuid, this.patient_reg_no, this.fname, this.lname, this.mob_no, this.age, this.blood_grp, this.gender, this.updated_on, this.sync_pending,this.client_updated_at];
    }
}