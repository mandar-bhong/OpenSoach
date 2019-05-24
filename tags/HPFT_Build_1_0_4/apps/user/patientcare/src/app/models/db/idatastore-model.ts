export interface IDatastoreModel {
    uuid: string;
    sync_pending: number;
    client_updated_at: string;
    updated_on: string;
    getModelValues(): any[]
}