import { AppMessageHandler } from "./app-message-handler.js";
import { ServerDataStoreDataModel } from "../models/api/server-data-store-data-model.js";
import { ServerWorkerEventDataModel } from "../models/api/server-worker-event-data-model.js";
import { IDatastoreModel } from "../models/db/idatastore-model.js";
import { SyncStoreManager } from "./sync-store-manager.js";
import { ScheduleDatastoreModel } from "../models/db/schedule-model.js";
import { SYNC_STORE } from "../app-constants.js";
import { ScheduleDatastoreMessageHandler } from "./schedule-datastore-message-handler.js";
import { CancelScheduleDatastoreMessageHandler } from "./cancel-schedule-datastore-message-handler.js";

export class AppMessageUIHandler extends AppMessageHandler {

    constructor() {
        super();
    }
    async handleMessage(msg: ServerDataStoreDataModel<IDatastoreModel>, postMessageFn: (msg: ServerWorkerEventDataModel) => void) {
        super.handleMessage(msg, postMessageFn);
        switch (msg.datastore) {
            case SYNC_STORE.SCHEDULE:
                const scheduleDatastoreModel = <ScheduleDatastoreModel>this.dataModel.data
                //0:   scheduleDatastoreModel.status
                if (scheduleDatastoreModel.status == 0) {
                    const scheduleDatastoreMessageHandler = new ScheduleDatastoreMessageHandler();
                    this.postActionContext = scheduleDatastoreMessageHandler.handleMessage(scheduleDatastoreModel);
                    this.postAction = this.schedulePostAction;
                } else if (scheduleDatastoreModel.status == 1) {
                    const cancelScheduleDatastoreMessageHandler = new CancelScheduleDatastoreMessageHandler();                 
                    this.postActionContext = await cancelScheduleDatastoreMessageHandler.handleMessage(scheduleDatastoreModel);                  
                    this.postAction = this.schedulePostAction;                    
                }
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
            const item = new ServerDataStoreDataModel<IDatastoreModel>()
            item.datastore = SYNC_STORE.ACTION
            item.data = element;
            appMessageUIHandler.handleMessage(item, this.postMessageCallback);
        });
    }
}