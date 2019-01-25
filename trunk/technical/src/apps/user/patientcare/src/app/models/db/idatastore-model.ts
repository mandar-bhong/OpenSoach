export interface IDatastoreModel {
    uuid: string;
    sync_pending: number;
    getModelValues(): any[]
}