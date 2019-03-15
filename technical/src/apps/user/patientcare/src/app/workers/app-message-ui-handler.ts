import { AppMessageHandler } from "./app-message-handler.js";
import { ServerDataStoreDataModel } from "../models/api/server-data-store-data-model.js";
import { ServerWorkerEventDataModel } from "../models/api/server-worker-event-data-model.js";
import { IDatastoreModel } from "../models/db/idatastore-model.js";
import { SyncStoreManager } from "./sync-store-manager.js";
import { ScheduleDatastoreModel } from "../models/db/schedule-model.js";
import { SYNC_STORE } from "../app-constants.js";
import { ScheduleDatastoreMessageHandler } from "./schedule-datastore-message-handler.js";

export class AppMessageUIHandler extends AppMessageHandler {

    constructor() {
        super();
    }

    handleMessage(msg: ServerDataStoreDataModel<IDatastoreModel>, postMessageFn: (msg: ServerWorkerEventDataModel) => void) {
        super.handleMessage(msg, postMessageFn);

        switch (msg.datastore) {
            case SYNC_STORE.SCHEDULE:
                const scheduleDatastoreModel = <ScheduleDatastoreModel>this.dataModel.data
                //0:   scheduleDatastoreModel.status
                const scheduleDatastoreMessageHandler = new ScheduleDatastoreMessageHandler();
                this.postActionContext = scheduleDatastoreMessageHandler.handleMessage(scheduleDatastoreModel);
                this.postAction = this.schedulePostAction;
                // 1: call update meyhod in case of cancel schedule of update above method
                break;
        }

        this.saveToDataStore().then(() => {
            this.notifyUI();
            this.notifySync();
            SyncStoreManager.syncToServerChanged(msg.datastore);
            this.postAction();
        });

    }


    // inserting actions after schedule is created.
    schedulePostAction() {
        this.postActionContext.forEach(element => {
            const appMessageUIHandler = new AppMessageUIHandler();
            const x = new ServerDataStoreDataModel<IDatastoreModel>()
            x.datastore = SYNC_STORE.ACTION
            x.data = element;
            appMessageUIHandler.handleMessage(x, this.postMessageCallback);
        });
    }
}