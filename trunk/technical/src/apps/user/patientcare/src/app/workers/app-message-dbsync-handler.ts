import { AppMessageHandler } from "./app-message-handler.js";
import { ServerDataStoreDataModel } from "../models/api/server-data-store-data-model.js";

export class AppMessageDbSyncHandler extends AppMessageHandler {

    constructor() {
        super();
    }

    handleMessage(msg: ServerDataStoreDataModel, postMessageFn: (msg: any) => void) {
        super.handleMessage(msg, postMessageFn);
        this.notifyUI();
    }
}