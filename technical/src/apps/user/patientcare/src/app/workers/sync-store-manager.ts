import { SYNC_STORE, SERVER_SYNC_STATE, SYNC_PENDING, DB_SYNC_TYPE } from "../app-constants.js";
import { SyncDb } from "../helpers/sync-db-helper.js";
import { SyncDataModel } from "../models/ui/sync-models.js";
import { ServerWorkerContext, SYNC_TYPE, CurrentStoreModel } from "./server-worker-context.js";
import { DatabaseHelper } from "../helpers/database-helper.js";
import { ServerHelper } from "./server-helper.js";

export class SyncStoreManager {
    static syncStore: SyncDataModel[]; // list of sync tables
    static updatedSyncStore: SyncDataModel[];
    static count: number;
    static currentStore = new CurrentStoreModel();
    static readSyncComplete: boolean;
    static isDataConflict: boolean;


    public static readSyncStore() {
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

    public static readUpdatedSyncStore() {
        // Read sync store and save in syncStore

        return new Promise((resolve, reject) => {

            console.log("ReadUpdatedSyncStore..")
            this.updatedSyncStore = [];
            SyncDb.getSyncList().then(
                (val) => {
                    val.forEach(item => {
                        let syncDataItem = new SyncDataModel();
                        syncDataItem = item;
                        this.updatedSyncStore.push(syncDataItem);
                    });
                    resolve();

                },
                (error) => {
                    console.log("getSyncList error:", error);
                    reject(error);
                }
            );
        });
    }

    public static getNextStore(syncState: number): CurrentStoreModel {
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
            console.log("currentstore", currentstore);

            switch (syncState) {

                case SERVER_SYNC_STATE.SYNC_TO_SERVER:

                    if (currentstore.sync_to_server_pending === SYNC_PENDING.TRUE
                        && (currentstore.sync_type === DB_SYNC_TYPE.SYNC_TO_SERVER || currentstore.sync_type === DB_SYNC_TYPE.SYNC_TO_AND_FROM_SERVER)) {
                        this.currentStore.currentStoreName = currentstore.store_name;
                        console.log("sync to server,current store name :", this.currentStore)
                    } else {
                        this.currentStore.currentStoreName = "getNextStore";
                    }

                    break;

                case SERVER_SYNC_STATE.SYNC_FROM_SERVER:

                    switch (ServerWorkerContext.syncType) {

                        case SYNC_TYPE.FULL:

                            if (currentstore.sync_type === DB_SYNC_TYPE.SYNC_FROM_SERVER || currentstore.sync_type === DB_SYNC_TYPE.SYNC_TO_AND_FROM_SERVER) {
                                this.currentStore.currentStoreName = currentstore.store_name;
                                this.currentStore.lastSynched = currentstore.last_synced;
                            } else {
                                this.currentStore.currentStoreName = "getNextStore";
                            }

                            break;

                        case SYNC_TYPE.DIFFERENTIAL:

                            if (currentstore.sync_from_server_pending === SYNC_PENDING.TRUE
                                && (currentstore.sync_type === DB_SYNC_TYPE.SYNC_FROM_SERVER || currentstore.sync_type === DB_SYNC_TYPE.SYNC_TO_AND_FROM_SERVER)) {
                                this.currentStore.currentStoreName = currentstore.store_name;
                                this.currentStore.lastSynched = currentstore.last_synced;
                            } else {
                                this.currentStore.currentStoreName = "getNextStore";
                            }

                            break;
                    }

                    break;
            }


            if (this.count < this.syncStore.length - 1) {
                this.count = this.count + 1;
            } else {                
                console.log("read sync completed setting complete flag");
                this.count = 0;
                this.readSyncComplete = true;
            }

        } else {
            console.log("sync complete set current store empty")
            this.currentStore.currentStoreName = "";
        }


        console.log("this.currentStore", this.currentStore);

        return this.currentStore;

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


        // read sync table again
        this.readUpdatedSyncStore().then(
            (val) => {
                console.log("reading updated sync store completed..");

                for (var i = 0; i < this.updatedSyncStore.length; i++) {
                    const prevStore = this.syncStore[i];
                    const updatedStore = this.updatedSyncStore[i];

                    switch (prevStore.sync_type) {

                        case DB_SYNC_TYPE.SYNC_TO_SERVER:
                            if ((prevStore.sync_to_server_pending === SYNC_PENDING.TRUE) && (prevStore.sync_to_server_pending_time == updatedStore.sync_to_server_pending_time)) {
                                DatabaseHelper.updateSyncStoreSyncPending(prevStore.store_name, DB_SYNC_TYPE.SYNC_TO_SERVER, SYNC_PENDING.FALSE);
                            } else if (prevStore.sync_to_server_pending_time !== updatedStore.sync_to_server_pending_time) {
                                this.isDataConflict = true;
                            }

                            break;

                        case DB_SYNC_TYPE.SYNC_FROM_SERVER:
                            if ((prevStore.sync_from_server_pending === SYNC_PENDING.TRUE) && (prevStore.sync_from_server_pending_time == updatedStore.sync_from_server_pending_time)) {
                                DatabaseHelper.updateSyncStoreSyncPending(prevStore.store_name, DB_SYNC_TYPE.SYNC_FROM_SERVER, SYNC_PENDING.FALSE);
                            } else if (prevStore.sync_from_server_pending_time !== updatedStore.sync_from_server_pending_time) {
                                this.isDataConflict = true;
                            }

                            // DatabaseHelper.updateSyncStoreLastSynched(prevStore.store_name);

                            break;

                        case DB_SYNC_TYPE.SYNC_TO_AND_FROM_SERVER:

                            if ((prevStore.sync_to_server_pending === SYNC_PENDING.TRUE) && (prevStore.sync_to_server_pending_time == updatedStore.sync_to_server_pending_time)) {
                                DatabaseHelper.updateSyncStoreSyncPending(prevStore.store_name, DB_SYNC_TYPE.SYNC_TO_SERVER, SYNC_PENDING.FALSE);
                            } else if (prevStore.sync_to_server_pending_time !== updatedStore.sync_to_server_pending_time) {
                                this.isDataConflict = true;
                            }

                            if ((prevStore.sync_from_server_pending === SYNC_PENDING.TRUE) && (prevStore.sync_from_server_pending_time == updatedStore.sync_from_server_pending_time)) {
                                DatabaseHelper.updateSyncStoreSyncPending(prevStore.store_name, DB_SYNC_TYPE.SYNC_FROM_SERVER, SYNC_PENDING.FALSE);
                            } else if (prevStore.sync_from_server_pending_time !== updatedStore.sync_from_server_pending_time) {
                                this.isDataConflict = true;
                            }

                            // DatabaseHelper.updateSyncStoreLastSynched(prevStore.store_name);

                            break;

                    }

                }

                if (this.isDataConflict == true) {
                    this.isDataConflict = false;
                    ServerWorkerContext.syncState = SERVER_SYNC_STATE.DIFFERENTIAL_SYNC_INITIALISE
                    ServerHelper.switchSyncState();
                }

            },
            (err) => {
                console.log(err);
            }
        );
    }

    public static syncFromServerChanged(datastore: SYNC_STORE) {
        // update sync_table for the datastore.
        // columns to update, sync_from_server_pending=true, sync_from_server_pending_time=now
        // if sync not in progress
        // trigger differential sync, sync_to_server

        DatabaseHelper.updateSyncStoreSyncPending(datastore, DB_SYNC_TYPE.SYNC_FROM_SERVER, SYNC_PENDING.TRUE)

        if (ServerWorkerContext.isSyncInprogress !== true) {
            ServerWorkerContext.syncState = SERVER_SYNC_STATE.DIFFERENTIAL_SYNC_INITIALISE
            ServerHelper.switchSyncState();
        }

    }

    public static syncToServerChanged(datastore: SYNC_STORE) {
        // update sync_table for the datastore.
        // columns to update, sync_to_server_pending=true, sync_to_server_pending_time=now
        // if sync not in progress
        // trigger differential sync, sync_from_server

        DatabaseHelper.updateSyncStoreSyncPending(datastore, DB_SYNC_TYPE.SYNC_TO_SERVER, SYNC_PENDING.TRUE);

        if (ServerWorkerContext.isSyncInprogress !== true) {
            ServerWorkerContext.syncState = SERVER_SYNC_STATE.DIFFERENTIAL_SYNC_INITIALISE
            ServerHelper.switchSyncState();
        }
    }

    // update individual tbl sync pending
    public static updateTblSyncPending(datastore: SYNC_STORE, syncPendingTime: Date) {
        DatabaseHelper.updateTableSyncPending(datastore, syncPendingTime);
    }



}