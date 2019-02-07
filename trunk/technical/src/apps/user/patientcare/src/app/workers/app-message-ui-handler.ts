import { AppMessageHandler } from "./app-message-handler.js";
import { ServerDataStoreDataModel } from "../models/api/server-data-store-data-model.js";
import { ServerWorkerEventDataModel } from "../models/api/server-worker-event-data-model.js";
import { IDatastoreModel } from "../models/db/idatastore-model.js";
import { SyncStoreManager } from "./sync-store-manager.js";

export class AppMessageUIHandler extends AppMessageHandler {

    constructor() {
        super();
    }

    handleMessage(msg: ServerDataStoreDataModel<IDatastoreModel>, postMessageFn: (msg: ServerWorkerEventDataModel) => void) {
        super.handleMessage(msg, postMessageFn);
        this.saveToDataStore();
        this.notifyUI();
        this.notifySync();
        SyncStoreManager.syncToServerChanged(msg.datastore);
    }
}