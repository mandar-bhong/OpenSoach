import { SYNC_STORE, SERVER_SYNC_STATE, SYNC_PENDING, DB_SYNC_TYPE } from "../app-constants.js";
import { SyncDb } from "../helpers/sync-db-helper.js";
import { SyncDataModel } from "../models/ui/sync-models.js";

export class SyncStoreManager {
    static syncStore: SyncDataModel[]; // list of sync tables
    static count: number;
    static currentStore: string;
    static readSyncComplete: boolean;
    public static ReadSyncStore() {
        // Read sync store and save in syncStore

        return new Promise((resolve, reject) => {

            console.log("ReadSyncStore..")
            this.syncStore = [];
            SyncDb.getSyncList().then(
                (val) => {
                    val.forEach(item => {
                        let syncDataItem = new SyncDataModel();
                        syncDataItem = item;
                        this.syncStore.push(syncDataItem);
                    });
                    //store ordered by sync order
                    this.syncStore.sort((a, b) => { return a.sync_order - b.sync_order });
                    this.count = 0;
                    this.readSyncComplete = false;
                    resolve();

                },
                (error) => {
                    console.log("getSyncList error:", error);
                    reject(error);
                }
            );
        });

    }

    public static getNextStore(syncState: number): string {
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

        console.log("count:", this.count);
        console.log("readSyncComplete:", this.readSyncComplete);

        if (this.readSyncComplete == false) {
            const currentstore = this.syncStore[this.count];

            switch (syncState) {

                case SERVER_SYNC_STATE.SYNC_TO_SERVER:

                    console.log("currentstore", currentstore);


                    if (currentstore.sync_to_server_pending == SYNC_PENDING.TRUE
                        && currentstore.sync_type == DB_SYNC_TYPE.SYNC_TO_SERVER || DB_SYNC_TYPE.SYNC_TO_AND_FROM_SERVER) {
                        this.currentStore = currentstore.store_name;
                    }

                    break;

                case SERVER_SYNC_STATE.SYNC_FROM_SERVER:

                    if (currentstore.sync_from_server_pending == SYNC_PENDING.TRUE
                        && currentstore.sync_type == DB_SYNC_TYPE.SYNC_FROM_SERVER || DB_SYNC_TYPE.SYNC_TO_AND_FROM_SERVER) {
                        this.currentStore = currentstore.store_name;
                    }

                    break;
            }


            if (this.count < this.syncStore.length - 1) {
                this.count = this.count + 1;
            } else {
                // this.currentStore = "";                
                console.log("read sync completed setting complete flag");
                this.count = 0;
                this.readSyncComplete = true;
            }

        } else {
            console.log("sync complete set current store empty")
            this.currentStore = "";
        }


        console.log("this.currentStore",this.currentStore);

        return this.currentStore;

        // this.currentStore = "schedule_tbl";

    }

    public static updateSyncStore() {
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

    public static syncFromServerChanged(datastore: SYNC_STORE) {
        // update sync_table for the datastore.
        // columns to update, sync_from_server_pending=true, sync_from_server_pending_time=now
        // if sync not in progress
        // trigger differential sync, sync_to_server
    }

    public static syncToServerChanged(datastore: SYNC_STORE) {
        // update sync_table for the datastore.
        // columns to update, sync_to_server_pending=true, sync_to_server_pending_time=now
        // if sync not in progress
        // trigger differential sync, sync_from_server
    }
}