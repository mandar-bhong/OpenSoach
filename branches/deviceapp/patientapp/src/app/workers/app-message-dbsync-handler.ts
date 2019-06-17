import { AppMessageHandler } from "./app-message-handler.js";
import { ServerDataStoreDataModel } from "../models/api/server-data-store-data-model.js";
import { ServerWorkerEventDataModel } from "../models/api/server-worker-event-data-model.js";
import { IDatastoreModel } from "../models/db/idatastore-model.js";

export class AppMessageDbSyncHandler extends AppMessageHandler {

    constructor() {
        super();
    }

    handleMessage(msg: ServerDataStoreDataModel<IDatastoreModel>, postMessageFn: (msg: ServerWorkerEventDataModel) => void) {
        super.handleMessage(msg, postMessageFn);
        this.saveToDataStore().then(() => {
            this.notifyUI();
        });
    }

    handleDeleteMessage(msg: ServerDataStoreDataModel<IDatastoreModel>, postMessageFn: (msg: ServerWorkerEventDataModel) => void) {
        super.handleMessage(msg, postMessageFn);
        this.deleteFromDataStore().then(() => {
            this.notifyUI();
        });
    }

}