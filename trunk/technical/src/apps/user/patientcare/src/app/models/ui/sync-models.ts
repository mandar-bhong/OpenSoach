export class SyncListViewModel {
    dbmodel: any;
}

export class SyncDataModel {
    store_name: string;
    sync_order: number;
    last_synced: Date;
    sync_type: number;
    sync_to_server_pending:number;
    sync_to_server_pending_time:Date;
    sync_from_server_pending:number;
    sync_from_server_pending_time:Date;
}