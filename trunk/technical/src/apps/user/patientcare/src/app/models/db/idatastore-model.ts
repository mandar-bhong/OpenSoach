export interface IDatastoreModel {
    uuid: string;
    sync_pending: number;
    sync_pending_time: Date;
    getModelValues(): any[]
}