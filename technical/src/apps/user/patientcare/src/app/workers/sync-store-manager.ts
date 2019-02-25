import { SYNC_STORE, SERVER_SYNC_STATE, SYNC_PENDING, DB_SYNC_TYPE } from "../app-constants.js";
import { SyncDb } from "../helpers/sync-db-helper.js";
import { SyncDataModel } from "../models/ui/sync-models.js";
import { ServerWorkerContext, SYNC_TYPE, CurrentStoreModel } from "./server-worker-context.js";
import { DatabaseHelper } from "../helpers/database-helper.js";
import { ServerHelper } from "./server-helper.js";

export class SyncStoreManager {
    static syncStore: SyncDataModel[]; // list of sync tables
    static updatedSyncStore: SyncDataModel[];
    static currentStore = new CurrentStoreModel();
    static isDataConflict: boolean;

    static firststore = true;
    static syncToServerStoreList: CurrentStoreModel[];
    static syncFromServerStoreList: CurrentStoreModel[];


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


    public static getFilteredStoreList() {

        console.log("in getFilteredStoreList..")

        this.syncToServerStoreList = [];
        this.syncFromServerStoreList = [];

        this.syncStore.forEach(item => {

            switch (item.sync_type) {

                case DB_SYNC_TYPE.SYNC_TO_SERVER:

                    if (item.sync_to_server_pending === SYNC_PENDING.TRUE) {
                        let currentStoreModel = new CurrentStoreModel()
                        currentStoreModel.currentStoreName = item.store_name;
                        this.syncToServerStoreList.push(currentStoreModel);
                    }

                    break;

                case DB_SYNC_TYPE.SYNC_FROM_SERVER:

                    switch (ServerWorkerContext.syncType) {

                        case SYNC_TYPE.FULL:

                            let currentStoreModel = new CurrentStoreModel()
                            currentStoreModel.currentStoreName = item.store_name;
                            currentStoreModel.lastSynched = item.last_synced;
                            this.syncFromServerStoreList.push(currentStoreModel);

                            break;

                        case SYNC_TYPE.DIFFERENTIAL:

                            if (item.sync_to_server_pending === SYNC_PENDING.TRUE) {
                                let currentStoreModel = new CurrentStoreModel()
                                currentStoreModel.currentStoreName = item.store_name;
                                currentStoreModel.lastSynched = item.last_synced;
                                this.syncFromServerStoreList.push(currentStoreModel);
                            }

                            break;
                    }

                    break;

                case DB_SYNC_TYPE.SYNC_TO_AND_FROM_SERVER:

                    if (item.sync_to_server_pending === SYNC_PENDING.TRUE) {
                        let currentStoreModel = new CurrentStoreModel()
                        currentStoreModel.currentStoreName = item.store_name;
                        this.syncToServerStoreList.push(currentStoreModel);
                    }

                    switch (ServerWorkerContext.syncType) {

                        case SYNC_TYPE.FULL:

                            let currentStoreModel = new CurrentStoreModel()
                            currentStoreModel.currentStoreName = item.store_name;
                            currentStoreModel.lastSynched = item.last_synced;
                            this.syncFromServerStoreList.push(currentStoreModel);

                            break;

                        case SYNC_TYPE.DIFFERENTIAL:

                            if (item.sync_to_server_pending === SYNC_PENDING.TRUE) {
                                let currentStoreModel = new CurrentStoreModel()
                                currentStoreModel.currentStoreName = item.store_name;
                                currentStoreModel.lastSynched = item.last_synced;
                                this.syncFromServerStoreList.push(currentStoreModel);
                            }

                            break;
                    }
                    break;
            }
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

        switch (syncState) {

            case SERVER_SYNC_STATE.SYNC_TO_SERVER:

                if (this.syncToServerStoreList.length == 0) {
                    this.currentStore.currentStoreName = "";
                    this.firststore = true;
                } else {
                    if (this.firststore == true) {
                        this.currentStore = this.syncToServerStoreList[0];
                        this.firststore = false;
                    } else {
                        const index = this.syncToServerStoreList.indexOf(this.currentStore);
                        if (index >= 0 && index < this.syncToServerStoreList.length - 1) {
                            this.currentStore = this.syncToServerStoreList[index + 1]
                        } else {
                            this.currentStore.currentStoreName = "";
                            this.firststore = true;
                        }
                    }
                }

                break;

            case SERVER_SYNC_STATE.SYNC_FROM_SERVER:

                if (this.syncFromServerStoreList.length == 0) {
                    this.currentStore.currentStoreName = "";
                    this.firststore = true;
                } else {
                    if (this.firststore == true) {
                        this.currentStore = this.syncFromServerStoreList[0];
                        this.firststore = false;
                    } else {
                        const index = this.syncFromServerStoreList.indexOf(this.currentStore);
                        if (index >= 0 && index < this.syncFromServerStoreList.length - 1) {
                            this.currentStore = this.syncFromServerStoreList[index + 1]
                        } else {
                            this.currentStore.currentStoreName = "";
                            this.firststore = true;
                        }
                    }
                }

                break;
        }

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