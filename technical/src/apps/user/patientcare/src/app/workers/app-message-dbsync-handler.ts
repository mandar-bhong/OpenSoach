import { AppMessageHandler } from "./app-message-handler.js";
import { ServerDataStoreDataModel } from "../models/api/server-data-store-data-model.js";
import { ServerWorkerEventDataModel } from "../models/api/server-worker-event-data-model.js";

export class AppMessageDbSyncHandler extends AppMessageHandler {

    constructor() {
        super();
    }

    handleMessage(msg: ServerDataStoreDataModel, postMessageFn: (msg: ServerWorkerEventDataModel) => void) {
        super.handleMessage(msg, postMessageFn);
        this.saveToDataStore();
        this.notifyUI();
    }
}