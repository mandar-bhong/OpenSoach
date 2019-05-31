import { ServerDataStoreDataModel } from "../models/api/server-data-store-data-model.js";
import { ServerWorkerEventDataModel } from "../models/api/server-worker-event-data-model.js";
import { SERVER_WORKER_EVENT_MSG_TYPE, SYNC_STORE, ConfigCodeType } from "../app-constants.js";
import { DatabaseHelper } from "../helpers/database-helper.js";
import { ScheduleDatastoreModel } from "../models/db/schedule-model.js";
import { IDatastoreModel } from "../models/db/idatastore-model.js";
import { Schedulardata, SchedularConfigData } from "../models/ui/chart-models.js";
import { MedicineHelper } from "../helpers/actions/medicine-helper.js";
import { ActionsData, ActionDataStoreModel } from "../models/db/action-datastore.js";
import { PatientMasterDatastoreModel } from "../models/db/patient-master-model.js";
import { PatientAdmissionDatastoreModel } from "../models/db/patient-admission-model.js";
import { IntakeHelper } from "../helpers/actions/intake-helper.js";
import { MonitorHelper } from "../helpers/actions/monitor-helper.js";
import { ActionTxnDatastoreModel } from "../models/db/action-txn-model.js";
import { ScheduleDatastoreMessageHandler } from "./schedule-datastore-message-handler.js";
import { DoctorsOrdersDatastoreModel } from "../models/db/doctors-orders-model.js";
import { DocumentUploadDatastore } from "../models/db/document-upload-datastore.js";

export interface AppMessageHandlerInterface {
    dataModel: ServerDataStoreDataModel<IDatastoreModel>;
    postActionContext: any;
    postMessageCallback: (msg: ServerWorkerEventDataModel) => void;
    handleMessage(msg: ServerDataStoreDataModel<IDatastoreModel>, postMessageFn: (msg: ServerWorkerEventDataModel) => void): void
    saveToDataStore(): Promise<{}>;
    notifyUI(): void;
    notifySync(): void;
    postAction: () => void;
}

export class AppMessageHandler implements AppMessageHandlerInterface {
    dataModel: ServerDataStoreDataModel<IDatastoreModel>;
    postActionContext: any;
    postAction: () => void;
    //ScheduleDatastoreModel
    postMessageCallback: (msg: ServerWorkerEventDataModel) => void;
    handleMessage(msg: ServerDataStoreDataModel<IDatastoreModel>, postMessageFn: (msg: ServerWorkerEventDataModel) => void) {
        this.dataModel = msg;
        console.log('base message handle executed');
        this.postMessageCallback = postMessageFn;

        switch (msg.datastore) {
            case SYNC_STORE.SCHEDULE:
                const scheduleDatastoreModel = new ScheduleDatastoreModel();
                Object.assign(scheduleDatastoreModel, this.dataModel.data)
                this.dataModel.data = scheduleDatastoreModel;
                // const scheduleDatastoreMessageHandler = new ScheduleDatastoreMessageHandler();
                // scheduleDatastoreMessageHandler.handleMessage(scheduleDatastoreModel);
                break;
            case SYNC_STORE.ACTION:
                const actionDatastoreModel = new ActionDataStoreModel();
                Object.assign(actionDatastoreModel, this.dataModel.data)
                this.dataModel.data = actionDatastoreModel;
                // const scheduleDatastoreMessageHandler = new ScheduleDatastoreMessageHandler();
                // scheduleDatastoreMessageHandler.handleMessage(scheduleDatastoreModel);
                break;
            case SYNC_STORE.PATIENT_MASTER:
                const patientMasterDatastoreModel = new PatientMasterDatastoreModel();
                Object.assign(patientMasterDatastoreModel, this.dataModel.data)
                this.dataModel.data = patientMasterDatastoreModel;
                break;
            case SYNC_STORE.PATIENT_ADMISSION:
                const patientAdmissionDatastoreModel = new PatientAdmissionDatastoreModel();
                Object.assign(patientAdmissionDatastoreModel, this.dataModel.data)
                this.dataModel.data = patientAdmissionDatastoreModel;
                break;
            case SYNC_STORE.ACTION_TXN:
                const actionTxnDatastoreModel = new ActionTxnDatastoreModel();
                Object.assign(actionTxnDatastoreModel, this.dataModel.data)
                this.dataModel.data = actionTxnDatastoreModel;
                break;
            case SYNC_STORE.DOCTORS_ORDERS:
                const doctorsOrdersDatastoreModel = new DoctorsOrdersDatastoreModel();
                Object.assign(doctorsOrdersDatastoreModel, this.dataModel.data);
                this.dataModel.data = doctorsOrdersDatastoreModel;
                break;
            case SYNC_STORE.DOCUMENT:
                const docUploadDatastoreModel = new DocumentUploadDatastore();
                Object.assign(docUploadDatastoreModel, this.dataModel.data);
                this.dataModel.data = docUploadDatastoreModel;
                break;
            default:
                break;
        }
    }

    saveToDataStore() {
        return new Promise((resolve, reject) => {
            console.log('save to datastore model');
            try {

                setTimeout(() => {
                    DatabaseHelper.DataStoreInsertUpdate(this.dataModel.datastore, this.dataModel.data.getModelValues())
                        .then(() => { resolve() }).catch(e => {
                            reject(e);
                        });
                }, 1);

            } catch (e) {
                console.log(e.error);
                reject(e);
            }
        });
    }

    deleteFromDataStore() {
        return new Promise((resolve, reject) => {
            console.log('delete from datastore..');
            try {
                DatabaseHelper.dataStoreDelete(this.dataModel.datastore, this.dataModel.data.getModelValues())
                    .then(() => { resolve() }).catch(e => {
                        reject(e);
                    });
            } catch (e) {
                console.log(e.error);
                reject(e);
            }
        });
    }

    notifyUI() {
        const workerEvent = new ServerWorkerEventDataModel();
        workerEvent.msgtype = SERVER_WORKER_EVENT_MSG_TYPE.DATA_RECEIVED;
        workerEvent.data = [this.dataModel];
        this.postMessageCallback(workerEvent);
    }

    notifySync() {

    }

}