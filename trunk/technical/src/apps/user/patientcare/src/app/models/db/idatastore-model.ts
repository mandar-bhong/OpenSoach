export interface IDatastoreModel {
    uuid: string;
    sync_pending: number;
    client_updated_at: string;
    getModelValues(): any[]
}