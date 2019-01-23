import { ServerDataStoreDataModel } from "../models/api/server-data-store-data-model.js";
import { ServerWorkerEventDataModel } from "../models/api/server-worker-event-data-model.js";
import { SERVER_WORKER_EVENT_MSG_TYPE } from "../app-constants.js";
import { DatabaseHelper } from "../helpers/database-helper.js";

export interface AppMessageHandlerInterface {
    dataModel: ServerDataStoreDataModel;
    postMessageCallback: (msg: ServerWorkerEventDataModel) => void;
    handleMessage(msg: ServerDataStoreDataModel, postMessageFn: (msg: ServerWorkerEventDataModel) => void): void
    saveToDataStore(): void;
    notifyUI(): void;
    updateSyncPending(): void;
    notifySync(): void;
}

export class AppMessageHandler implements AppMessageHandlerInterface {
    dataModel: ServerDataStoreDataModel;
    postMessageCallback: (msg: ServerWorkerEventDataModel) => void;

    handleMessage(msg: ServerDataStoreDataModel, postMessageFn: (msg: ServerWorkerEventDataModel) => void) {
        this.dataModel = msg;
        console.log('base message handle', this.dataModel);
        this.postMessageCallback = postMessageFn;
    }

    saveToDataStore() {
        DatabaseHelper.DataStoreInsertUpdate(this.dataModel.datastore, this.dataModel.data.getModelValues());
    }

    notifyUI() {
        const workerEvent = new ServerWorkerEventDataModel();
        workerEvent.msgtype = SERVER_WORKER_EVENT_MSG_TYPE.DATA_RECEIVED;
        workerEvent.data = [this.dataModel];
        this.postMessageCallback(workerEvent);
    }

    notifySync() {

    }

    updateSyncPending() {

        // Update sync table for the this.dataModel.datastore 'sync_pending' to true and sync_pending_time to current time
    }
}