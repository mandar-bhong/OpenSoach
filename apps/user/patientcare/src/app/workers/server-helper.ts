import { SERVER_SYNC_STATE } from "../app-constants.js";
import { ServerWorkerEventDataModel } from "../models/api/server-worker-event-data-model.js";
import { ServerWorkerContext, SYNC_TYPE } from "./server-worker-context.js";
import { SyncStoreManager } from "./sync-store-manager.js";
import { DatabaseHelper } from "../helpers/database-helper.js";
import { CommandRequestGenerator } from "./command-request-generator.js";
import { DocumentSyncHelper } from "./document-sync-helper.js";

export class ServerHelper {

    public static sendToServerCallback: (msg: any) => void;
    static postMessageCallback: (msg: ServerWorkerEventDataModel) => void;
    constructor() {
    }

    public static init(postMessageFn: (msg: ServerWorkerEventDataModel) => void) {
        ServerHelper.postMessageCallback = postMessageFn;
        console.log("in server helper init")
    }

    public static syncProcess(syncstate: SERVER_SYNC_STATE) {

        ServerWorkerContext.syncState = syncstate;

        switch (syncstate) {

            case SERVER_SYNC_STATE.SEND_AUTH_CMD:
                // send auth cmd
                // {"header":{"crc":"12","category":1,"commandid":1,"seqid":3},"payload":{"token":"Dev6AD88A481524BABF"}}

                console.log("SERVER_SYNC_STATE.SEND_AUTH_CMD");

                const authcmd = CommandRequestGenerator.authCmd();
                console.log("authcmd:", authcmd);

                ServerHelper.sendToServerCallback(authcmd);

                break;

            case SERVER_SYNC_STATE.READ_SYNC_STORE:
                // sync state READ_SYNC_STORE

                console.log("SERVER_SYNC_STATE.READ_SYNC_STORE");

                // read syncstore
                SyncStoreManager.readSyncStore().then(
                    (val) => {
                        console.log("reading sync store completed..")
                        ServerWorkerContext.isSyncInprogress = true;
                        ServerWorkerContext.syncType = SYNC_TYPE.FULL;
                        SyncStoreManager.getFilteredStoreList();
                        ServerWorkerContext.syncState = SERVER_SYNC_STATE.READ_SYNC_STORE_COMPLETED;
                        this.switchSyncState();
                    },
                    (err) => {
                        console.log(err);
                    }
                );

                break;

            case SERVER_SYNC_STATE.DIFFERENTIAL_SYNC_STARTED:
                // sync state differential sync

                console.log("SERVER_SYNC_STATE.DIFFERENTIAL_SYNC_STARTED");

                SyncStoreManager.readSyncStore().then(
                    (val) => {
                        console.log("reading sync store completed..")
                        ServerWorkerContext.isSyncInprogress = true;
                        ServerWorkerContext.syncType = SYNC_TYPE.DIFFERENTIAL;
                        SyncStoreManager.getFilteredStoreList();
                        ServerWorkerContext.syncState = SERVER_SYNC_STATE.READ_SYNC_STORE_COMPLETED;
                        this.switchSyncState();
                    },
                    (err) => {
                        console.log(err);
                    }
                );

                break;


            case SERVER_SYNC_STATE.SYNC_TO_SERVER:
                //sync to server
                // send apply sync cmd
                // {"header":{"crc":"12","category":3,"commandid":51,"seqid":3},"payload":{"storename":"","storedata":[{"uuid":"PA001","bedno":"A0001"}]}}
                // Read Sync store and getnext store

                console.log("SERVER_SYNC_STATE.SYNC_TO_SERVER");

                //get next store
                var storename = SyncStoreManager.getNextStore(SERVER_SYNC_STATE.SYNC_TO_SERVER)
                console.log("SYNC_TO_SERVER storename:", storename)

                if (storename.currentStoreName != "") {
                    DatabaseHelper.getSyncPendingDataStore(storename.currentStoreName)
                        .then(
                            (val) => {
                                const syncCmd = CommandRequestGenerator.applySyncCmd(storename.currentStoreName, val);
                                console.log("apply syncCmd", syncCmd);
                                ServerHelper.sendToServerCallback(syncCmd);
                            },
                            (err) => {
                                console.log("getSyncPendingDataStore err:", err);
                            }
                        )
                } else {
                    ServerWorkerContext.syncState = SERVER_SYNC_STATE.SYNC_TO_SERVER_COMPLETED;
                    this.switchSyncState();
                }

                break;

            case SERVER_SYNC_STATE.SYNC_FROM_SERVER:
                //sync from server
                // send get sync cmd with store name
                // {"header":{"crc":"12","category":3,"commandid":50,"seqid":3},"payload":{"storename":"","updatedon":"2018-10-30T00:00:00Z"}}

                console.log("SERVER_SYNC_STATE.SYNC_FROM_SERVER");

                //get next store
                storename = SyncStoreManager.getNextStore(SERVER_SYNC_STATE.SYNC_FROM_SERVER)
                console.log("SYNC_from_SERVER storename:", storename)


                if (storename.currentStoreName != "") {
                    const syncCmd = CommandRequestGenerator.getSyncCmd(storename.currentStoreName, storename.lastSynched);
                    console.log("get syncCmd", syncCmd);
                    ServerHelper.sendToServerCallback(syncCmd);

                } else {
                    ServerWorkerContext.syncState = SERVER_SYNC_STATE.SYNC_FROM_SERVER_COMPLETED;
                    this.switchSyncState();
                }

                break;

            case SERVER_SYNC_STATE.SYNC_FROM_SERVER_COMPLETED:
                //sync from server completed

                console.log("SERVER_SYNC_STATE.SYNC_FROM_SERVER_COMPLETED");

                DocumentSyncHelper.sync();

                if (ServerWorkerContext.syncType === SYNC_TYPE.DIFFERENTIAL) {
                    console.log("Differential Sync Completed");
                }

                ServerWorkerContext.isSyncInprogress = false;
                SyncStoreManager.updateSyncStore()

                break;
        }
    }

    public static switchSyncState() {

        console.log("SyncState", ServerWorkerContext.syncState);

        switch (ServerWorkerContext.syncState) {
            case SERVER_SYNC_STATE.NONE:
                this.syncProcess(SERVER_SYNC_STATE.SEND_AUTH_CMD);
                break;

            case SERVER_SYNC_STATE.SEND_AUTH_CMD_SUCCESS:
                this.syncProcess(SERVER_SYNC_STATE.READ_SYNC_STORE);
                break;

            case SERVER_SYNC_STATE.READ_SYNC_STORE_COMPLETED:
                this.syncProcess(SERVER_SYNC_STATE.SYNC_TO_SERVER);
                break;

            case SERVER_SYNC_STATE.SYNC_TO_SERVER:
                this.syncProcess(SERVER_SYNC_STATE.SYNC_TO_SERVER)
                break;

            case SERVER_SYNC_STATE.SYNC_TO_SERVER_COMPLETED:
                this.syncProcess(SERVER_SYNC_STATE.SYNC_FROM_SERVER)
                break;

            case SERVER_SYNC_STATE.SYNC_FROM_SERVER:
                this.syncProcess(SERVER_SYNC_STATE.SYNC_FROM_SERVER)
                break;

            case SERVER_SYNC_STATE.SYNC_FROM_SERVER_COMPLETED:
                this.syncProcess(SERVER_SYNC_STATE.SYNC_FROM_SERVER_COMPLETED)
                break;


            case SERVER_SYNC_STATE.DIFFERENTIAL_SYNC_INITIALISE:
                this.syncProcess(SERVER_SYNC_STATE.DIFFERENTIAL_SYNC_STARTED);
                break;

            case SERVER_SYNC_STATE.DIFFERENTIAL_SYNC_STARTED:
                this.syncProcess(SERVER_SYNC_STATE.SYNC_TO_SERVER);
                break;
        }
    }

}


