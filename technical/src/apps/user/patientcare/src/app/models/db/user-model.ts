import { IDatastoreModel } from "./idatastore-model";

export class UserDatastoreModel implements IDatastoreModel {
    uuid: string;
    usr_id:number;
    usr_name:string;
    urole_name:string;
    fname:string;
    lname:string;
    updated_on: Date;
    sync_pending: number;
    client_updated_at: string;
    getModelValues(): any[] {
        return [this.usr_id,this.usr_name,this.urole_name,this.fname,this.lname,this.updated_on,this.sync_pending,this.client_updated_at];
    }
}