import { ServerDataStoreDataModel } from "../models/api/server-data-store-data-model.js";
import { ServerWorkerEventDataModel } from "../models/api/server-worker-event-data-model.js";
import { SERVER_WORKER_EVENT_MSG_TYPE, SYNC_STORE } from "../app-constants.js";
import { DatabaseHelper } from "../helpers/database-helper.js";
import { ScheduleDatastoreModel } from "../models/db/schedule-model.js";
import { IDatastoreModel } from "../models/db/idatastore-model.js";
import { Schedulardata, SchedularConfigData } from "../models/ui/chart-models.js";
import { MedicineHelper } from "../helpers/actions/medicine-helper.js";
import { medicine } from "../common-constants.js";
import { ActionsData, ActionDataStoreModel } from "../models/db/action-datastore.js";
import { PatientMasterDatastoreModel } from "../models/db/patient-master-model.js";
import { PatientAdmissionDatastoreModel } from "../models/db/patient-admission-model.js";

export interface AppMessageHandlerInterface {
    dataModel: ServerDataStoreDataModel<IDatastoreModel>;
    postMessageCallback: (msg: ServerWorkerEventDataModel) => void;
    handleMessage(msg: ServerDataStoreDataModel<IDatastoreModel>, postMessageFn: (msg: ServerWorkerEventDataModel) => void): void
    saveToDataStore(): void;
    notifyUI(): void;
    updateSyncPending(): void;
    notifySync(): void;
}

export class AppMessageHandler implements AppMessageHandlerInterface {
    dataModel: ServerDataStoreDataModel<IDatastoreModel>;
    //ScheduleDatastoreModel
    postMessageCallback: (msg: ServerWorkerEventDataModel) => void;
    handleMessage(msg: ServerDataStoreDataModel<IDatastoreModel>, postMessageFn: (msg: ServerWorkerEventDataModel) => void) {
        this.dataModel = msg;
        console.log('base message handle executed', this.dataModel);
        this.postMessageCallback = postMessageFn;

        switch (msg.datastore) {
            case SYNC_STORE.SCHEDULE:
                const obj = new ScheduleDatastoreModel();
                Object.assign(obj, this.dataModel.data)
                this.dataModel.data = obj;
                this.handleScheduleMessage(obj);
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
        }
    }

    saveToDataStore() {
        try {
            DatabaseHelper.DataStoreInsertUpdate(this.dataModel.datastore, this.dataModel.data.getModelValues());
        } catch (e) {
            console.log(e.error);
        }
    }

    notifyUI() {
        const workerEvent = new ServerWorkerEventDataModel();
        workerEvent.msgtype = SERVER_WORKER_EVENT_MSG_TYPE.DATA_RECEIVED;
        workerEvent.data = [this.dataModel];
        this.postMessageCallback(workerEvent);
    }

    notifySync() {

    }

    handleScheduleMessage(obj: ScheduleDatastoreModel) {
        console.log('handleScheduleMessage executed 3wdqed');

        // const data = <ScheduleDatastoreModel>this.dataModel.data;
        const schedulardata = new Schedulardata();
        schedulardata.data = obj;
        schedulardata.conf = new SchedularConfigData()
        const parsedConf = <SchedularConfigData>JSON.parse(obj.conf);
        schedulardata.conf = parsedConf;
        console.log('schedulardata.conf ', schedulardata.conf);
        //  if (schedulardata.conf_type_code == medicine) {
        console.log('MedicineHelper creating');
        const medicineHelper = new MedicineHelper();
        console.log('MedicineHelper created');
        try {
            const actiondata = <ActionsData>medicineHelper.createMedicineActions(schedulardata);
            console.log('in try catch block');
            parsedConf.endDate = actiondata.enddate;
            obj.conf = JSON.stringify(parsedConf);
            console.log('action inserting..');
            try {
                // .action_tbl_delete
                //   DatabaseHelper.update('action_tbl_delete', []);              
                actiondata.actions.forEach(element => {
                    const actionsdbdata = new ActionDataStoreModel();
                    Object.assign(actionsdbdata, element);
                    DatabaseHelper.DataStoreInsertUpdate(SYNC_STORE.ACTION, actionsdbdata.getModelValues());
                });

            } catch (e) {
                console.log('action inserting failed....', e.error);
            }


        }
        catch (e) {
            console.error('MedicineHelper', e);
        }

        //  }
        // this.dataModel.data
        // create actions
        // set schedule end date
    }

    updateSyncPending() {

        // Update sync table for the this.dataModel.datastore 'sync_pending' to true and sync_pending_time to current time
    }
}