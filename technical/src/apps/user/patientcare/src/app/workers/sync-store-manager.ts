import { SYNC_STORE } from "../app-constants";

export class SyncStoreManager {
    static syncStore: any; // list of sync tables
    static currentStore: string;
    public static ReadSyncStore() {
        // Read sync store and save in syncStore
        
    }

    public static getNextStore(syncState: string): string {
        // switch on SyncState
        // case ToServer
        // get the store ordered by sync order and whose sync_to_server_pending is pending
        // and whose sync type is syncToServer or supporting both and is after currentstore
        // case FromServer
        // if full sync
        // get the store ordered by sync order and
        // whose sync type is syncFromServer or supporting both and is after currentstore
        // if differentialsync
        // get the store ordered by sync order and sync_from_server_pending is true
        // whose sync type is syncFromServer or supporting both and is after currentstore 
        

        this.currentStore = "schedule_tbl";
        return this.currentStore;
    }

    public static updateSyncStore()
    {
        // read sync table again
        // switch on SyncState
        // case ToServer        
        // check for each store
        // if the previous sync_to_server_pending was false and now it is true, do not update the store
        // if the previous sync_to_server_pending was true and the sync_to_server_pending_time are equal set sync_to_server_pending to false
        // case FromServer
         // if the previous sync_from_server_pending was false and now it is true, do not update the store
        // if the previous sync_from_server_pending was true and the sync_from_server_pending_time are equal set sync_from_server_pending to false

        // if there is conflict in data read and new data, trigger different sync again
        // else stop
    }

    public static syncFromServerChanged(datastore:SYNC_STORE)
    {
        // update sync_table for the datastore.
        // columns to update, sync_from_server_pending=true, sync_from_server_pending_time=now
        // if sync not in progress
        // trigger differential sync, sync_to_server
    }

    public static syncToServerChanged(datastore:SYNC_STORE)
    {
        // update sync_table for the datastore.
        // columns to update, sync_to_server_pending=true, sync_to_server_pending_time=now
         // if sync not in progress
        // trigger differential sync, sync_from_server
    }
}