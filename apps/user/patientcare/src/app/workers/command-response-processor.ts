import { CmdModel } from "../models/api/server-cmd-model.js";
import { RequestManager } from "./request-manager.js";
import { CMD_ID, CMD_CATEGORY, SERVER_SYNC_STATE, SYNC_STORE, SYNC_PENDING, APP_MODE } from "../app-constants.js";
import { ServerWorkerContext } from "./server-worker-context.js";
import { ServerHelper } from "./server-helper.js";
import { SyncStoreManager } from "./sync-store-manager.js";
import { ServicePointDatastoreModel } from "../models/db/service-point-models.js";
import { IDatastoreModel } from "../models/db/idatastore-model.js";
import { ServerDataStoreDataModel } from "../models/api/server-data-store-data-model.js";
import { AppMessageDbSyncHandler } from "./app-message-dbsync-handler.js";
import { DatabaseHelper } from "../helpers/database-helper.js";
import { ConfDatastoreModel } from "../models/db/conf-model.js";
import { PatientMasterDatastoreModel } from "../models/db/patient-master-model.js";
import { ScheduleDatastoreModel } from "../models/db/schedule-model.js";
import { PatientAdmissionDatastoreModel } from "../models/db/patient-admission-model.js";
import { PatientPersonalDetailsDatastoreModel } from "../models/db/patient-personal-details-model.js";
import { PatientMedicalDetailsDatastoreModel } from "../models/db/patient-medical-details-model.js";
import { ActionTxnDatastoreModel } from "../models/db/action-txn-model.js";
import { DoctorsOrdersDatastoreModel } from "../models/db/doctors-orders-model.js";
import { TreatmentDatastoreModel } from "../models/db/treatment-model.js";
import { TreatmentDocDatastoreModel } from "../models/db/treatment-doc-model.js";
import { PathologyRecordDatastoreModel } from "../models/db/pathology-record-model.js";
import { PathologyRecordDocDatastoreModel } from "../models/db/pathology-record-doc-model.js";
import { ActionDataStoreModel } from "../models/db/action-datastore.js";
import { UserDatastoreModel } from "../models/db/user-model.js";
import * as appSettings from "tns-core-modules/application-settings";

export class CommandResponseProcessor {

    public static cmdProcessor(respMsg: any) {

        // TODO: check if authorized if yes, set GlobalContext to Authorized
        // then call SwitchSyncState

        console.log(" in CmdProcessor..");

        console.log("respMsg", respMsg);

        const respDataModel: CmdModel = JSON.parse(respMsg);

        // get request cmd packet
        const requestCmd = RequestManager.getRequest(respDataModel.header.seqid);
        console.log("requestCmd", JSON.stringify(requestCmd));

        if (requestCmd) {

            //handle response to cmd send by device
            switch (requestCmd.header.category, requestCmd.header.commandid) {

                case CMD_CATEGORY.CMD_CAT_DEV_REGISTRATION && CMD_ID.CMD_DEV_REGISTRATION:
                    // auth request cmd response
                    if (respDataModel.payload.ack == true) {
                        ServerWorkerContext.syncState = SERVER_SYNC_STATE.SEND_AUTH_CMD_SUCCESS;
                        ServerHelper.switchSyncState();
                    } else {
                        // TODO : handle for failure
                    }

                    break;

                case CMD_CATEGORY.CMD_CAT_SYNC && CMD_ID.CMD_GET_STORE_SYNC:
                    // get sync request cmd response

                    ServerHelper.switchSyncState();

                    if (respDataModel.payload.ack == true && respDataModel.payload.ackdata != null) {
                        switch (respDataModel.payload.ackdata.storename) {

                            case SYNC_STORE.SERVICE_POINT:
                                this.handleServicePointResponse(respDataModel);
                                break;
                            case SYNC_STORE.CONF:
                                this.handleConfResponse(respDataModel);
                                break;
                            case SYNC_STORE.PATIENT_MASTER:
                                this.handlePatientMasterResponse(respDataModel);
                                break;
                            case SYNC_STORE.SCHEDULE:
                                this.handleScheduleResponse(respDataModel);
                                break;
                            case SYNC_STORE.PATIENT_ADMISSION:
                                this.handlePatientAdmissionResponse(respDataModel);
                                break;
                            case SYNC_STORE.PERSONAL_DETAILS:
                                this.handlePersonalDetailsResponse(respDataModel);
                                break;
                            case SYNC_STORE.MEDICAL_DETAILS:
                                this.handleMedicalDetailsResponse(respDataModel);
                                break;
                            case SYNC_STORE.ACTION_TXN:
                                this.handleActionTxnResponse(respDataModel);
                                break;
                            case SYNC_STORE.DOCTORS_ORDERS:
                                this.handleDoctorsOrdersResponse(respDataModel);
                                break;

                            case SYNC_STORE.TREATMENT:
                                this.handleTreatmentResponse(respDataModel);
                                break;
                            case SYNC_STORE.TREATMENT_DOC:
                                this.handleTreatmentDocResponse(respDataModel);
                                break;
                            case SYNC_STORE.PATHOLOGY_RECORD:
                                this.handlePathologyRecordOrdersResponse(respDataModel);
                                break;
                            case SYNC_STORE.PATHOLOGY_RECORD_DOC:
                                this.handlePathologyRecordDocResponse(respDataModel);
                                break;
                            case SYNC_STORE.ACTION:
                                this.handleActionResponse(respDataModel);
                                break;
                            case SYNC_STORE.USER:
                                this.handleUserResponse(respDataModel);
                                break;
                            case SYNC_STORE.PATIENT_MONITOR_MAPPING:
                                if (appSettings.getNumber("APP_MODE") == APP_MODE.USER_DEVICE) {
                                    this.handlePatientMonitorMappingViewResponse(respDataModel);
                                }
                                break;
                        }

                        if (respDataModel.payload.ackdata.updatedon) {
                            // update sync table last synced
                            DatabaseHelper.updateSyncStoreLastSynched(respDataModel.payload.ackdata.storename, respDataModel.payload.ackdata.updatedon);
                        }
                    }

                    break;

                case CMD_CATEGORY.CMD_CAT_SYNC && CMD_ID.CMD_APPLY_STORE_SYNC:
                    // apply sync request cmd response
                    // sync to server response - update individual tbl sync flag 

                    if (respDataModel.payload.ack == true) {
                        SyncStoreManager.updateTblSyncPending(requestCmd.payload.storename, requestCmd.payload.storedata[0].client_updated_at);
                    }

                    ServerHelper.switchSyncState();
                    break;
            }
        } else {
            //handle server notification
            switch (respDataModel.header.category, respDataModel.header.commandid) {

                case CMD_CATEGORY.CMD_CAT_SERVER_NOTIFICATION && CMD_ID.CMD_APPLY_STORE_SYNC:
                    console.log("notification cat response:", respDataModel);

                    SyncStoreManager.syncFromServerChanged(respDataModel.payload.storename)

                    break;

            }

        }

    }

    public static handleServicePointResponse(data: CmdModel) {
        // console.log("Service Point tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const servicePointDatastoreModel = new ServicePointDatastoreModel();
            const item = <ServicePointDatastoreModel>tblData[i];
            Object.assign(servicePointDatastoreModel, item);
            servicePointDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            // console.log("sp store data:", servicePointDatastoreModel);

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.SERVICE_POINT;
            serverDataStoreDataModel.data = servicePointDatastoreModel;

            // console.log("sp server data store model", serverDataStoreDataModel);

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

        }

    }

    public static handleConfResponse(data: CmdModel) {
        // console.log("Conf tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const confDatastoreModel = new ConfDatastoreModel();
            const item = <ConfDatastoreModel>tblData[i];
            Object.assign(confDatastoreModel, item);
            confDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            // console.log("conf store data:", confDatastoreModel);

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.CONF;
            serverDataStoreDataModel.data = confDatastoreModel;

            // console.log("conf server data store model", serverDataStoreDataModel);

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);
        }

    }

    public static handlePatientMasterResponse(data: CmdModel) {
        // console.log("Patient master tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const patientMasterDatastoreModel = new PatientMasterDatastoreModel();
            const item = <PatientMasterDatastoreModel>tblData[i];

            Object.assign(patientMasterDatastoreModel, item);
            patientMasterDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            // console.log("patient master store data:", patientMasterDatastoreModel);

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.PATIENT_MASTER;
            serverDataStoreDataModel.data = patientMasterDatastoreModel;

            // console.log("patient master server data store model", serverDataStoreDataModel);

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

        }

    }

    public static handleScheduleResponse(data: CmdModel) {
        // console.log("Patient schedule tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const scheduleDatastoreModel = new ScheduleDatastoreModel();
            const item = <ScheduleDatastoreModel>tblData[i];

            Object.assign(scheduleDatastoreModel, item);
            scheduleDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            // console.log("schedule store data:", scheduleDatastoreModel);

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.SCHEDULE;
            serverDataStoreDataModel.data = scheduleDatastoreModel;

            // console.log("schedule server data store model", serverDataStoreDataModel);

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

        }

    }

    public static handlePatientAdmissionResponse(data: CmdModel) {
        // console.log("Patient Admsn tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const patientAdmissionDatastoreModel = new PatientAdmissionDatastoreModel();
            const item = <PatientAdmissionDatastoreModel>tblData[i];

            Object.assign(patientAdmissionDatastoreModel, item);
            patientAdmissionDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            // console.log("patient admsn store data:", patientAdmissionDatastoreModel);

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.PATIENT_ADMISSION;
            serverDataStoreDataModel.data = patientAdmissionDatastoreModel;

            // console.log("patient admsn server data store model", serverDataStoreDataModel);

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

        }

    }

    public static handlePersonalDetailsResponse(data: CmdModel) {
        // console.log("Personal details tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const patientPersonalDetailsDatastoreModel = new PatientPersonalDetailsDatastoreModel();
            const item = <PatientPersonalDetailsDatastoreModel>tblData[i];

            Object.assign(patientPersonalDetailsDatastoreModel, item);
            patientPersonalDetailsDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            // console.log("personal details store data:", patientPersonalDetailsDatastoreModel);

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.PERSONAL_DETAILS;
            serverDataStoreDataModel.data = patientPersonalDetailsDatastoreModel;

            // console.log("personal details server data store model", serverDataStoreDataModel);

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

        }

    }

    public static handleMedicalDetailsResponse(data: CmdModel) {
        // console.log("Medical details tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const patientMedicalDetailsDatastoreModel = new PatientMedicalDetailsDatastoreModel();
            const item = <PatientMedicalDetailsDatastoreModel>tblData[i];

            Object.assign(patientMedicalDetailsDatastoreModel, item);
            patientMedicalDetailsDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            // console.log("medical details store data:", patientMedicalDetailsDatastoreModel);

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.MEDICAL_DETAILS;
            serverDataStoreDataModel.data = patientMedicalDetailsDatastoreModel;

            // console.log("medical details server data store model", serverDataStoreDataModel);

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

        }

    }

    public static handleActionTxnResponse(data: CmdModel) {
        // console.log("Action Txn tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const actionTxnDatastoreModel = new ActionTxnDatastoreModel();
            const item = <ActionTxnDatastoreModel>tblData[i];

            Object.assign(actionTxnDatastoreModel, item);
            actionTxnDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            // console.log("action txn store data:", actionTxnDatastoreModel);

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.ACTION_TXN;
            serverDataStoreDataModel.data = actionTxnDatastoreModel;

            // console.log("action txn server data store model", serverDataStoreDataModel);

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

        }

    }

    public static handleDoctorsOrdersResponse(data: CmdModel) {
        // console.log("Doctors Orders tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const doctorsOrdersDatastoreModel = new DoctorsOrdersDatastoreModel();
            const item = <DoctorsOrdersDatastoreModel>tblData[i];

            Object.assign(doctorsOrdersDatastoreModel, item);
            doctorsOrdersDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            // console.log("action txn store data:", doctorsOrdersDatastoreModel);

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.DOCTORS_ORDERS;
            serverDataStoreDataModel.data = doctorsOrdersDatastoreModel;

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

        }

    }

    public static handleTreatmentResponse(data: CmdModel) {
        // console.log("Treatment tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const treatmentDatastoreModel = new TreatmentDatastoreModel();
            const item = <TreatmentDatastoreModel>tblData[i];

            Object.assign(treatmentDatastoreModel, item);
            treatmentDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.TREATMENT;
            serverDataStoreDataModel.data = treatmentDatastoreModel;

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

        }

    }

    public static handleTreatmentDocResponse(data: CmdModel) {
        // console.log("Treatment doc tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const treatmentDocDatastoreModel = new TreatmentDocDatastoreModel();
            const item = <TreatmentDocDatastoreModel>tblData[i];

            Object.assign(treatmentDocDatastoreModel, item);
            treatmentDocDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.TREATMENT_DOC;
            serverDataStoreDataModel.data = treatmentDocDatastoreModel;

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

        }

    }

    public static handlePathologyRecordOrdersResponse(data: CmdModel) {
        // console.log("Pathology record tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const pathologyRecordDatastoreModel = new PathologyRecordDatastoreModel();
            const item = <PathologyRecordDatastoreModel>tblData[i];

            Object.assign(pathologyRecordDatastoreModel, item);
            pathologyRecordDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.PATHOLOGY_RECORD;
            serverDataStoreDataModel.data = pathologyRecordDatastoreModel;

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

        }

    }

    public static handlePathologyRecordDocResponse(data: CmdModel) {
        // console.log("Pathology record doc tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const pathologyRecordDocDatastoreModel = new PathologyRecordDocDatastoreModel();
            const item = <PathologyRecordDocDatastoreModel>tblData[i];

            Object.assign(pathologyRecordDocDatastoreModel, item);
            pathologyRecordDocDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.PATHOLOGY_RECORD_DOC;
            serverDataStoreDataModel.data = pathologyRecordDocDatastoreModel;

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

        }

    }

    public static handleActionResponse(data: CmdModel) {
        // console.log("action tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const actionDataStoreModel = new ActionDataStoreModel();
            const item = <ActionDataStoreModel>tblData[i];

            Object.assign(actionDataStoreModel, item);
            actionDataStoreModel.sync_pending = SYNC_PENDING.FALSE;

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.ACTION;
            serverDataStoreDataModel.data = actionDataStoreModel;

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

        }

    }

    public static handleUserResponse(data: CmdModel) {
        // console.log("user tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const userDatastoreModel = new UserDatastoreModel();
            const item = <UserDatastoreModel>tblData[i];

            Object.assign(userDatastoreModel, item);
            userDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.USER;
            serverDataStoreDataModel.data = userDatastoreModel;

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

        }

    }

    public static handlePatientMonitorMappingViewResponse(data: CmdModel) {
        // console.log("patient monitor mapping view data", data);

        const tblData = data.payload.ackdata.data

        for (let i = 0; i < tblData.length; i++) {

            DatabaseHelper.selectByID("patientlistbyadmissionuuid", tblData[i].uuid).then(
                (val) => {
                    if (val.length != 0) {
                        const patientAdmissionDatastoreModel = new PatientAdmissionDatastoreModel();
                        const item = <PatientAdmissionDatastoreModel>tblData[i];

                        Object.assign(patientAdmissionDatastoreModel, item);
                        patientAdmissionDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

                        // console.log("patient admsn store data:", patientAdmissionDatastoreModel);

                        const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
                        serverDataStoreDataModel.datastore = SYNC_STORE.PATIENT_ADMISSION;
                        serverDataStoreDataModel.data = patientAdmissionDatastoreModel;

                        // console.log("patient admsn server data store model", serverDataStoreDataModel);

                        new AppMessageDbSyncHandler().handleDeleteMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

                        // delete other store admission related entries
                        var storeName = [];
                        storeName.push("patient_personal_details_tbl", "patient_medical_details_tbl", "schedule_tbl",
                            "action_tbl", "action_txn_tbl", "doctors_orders_tbl", "treatment_tbl", "pathology_record_tbl")

                        storeName.forEach(item => {
                            DatabaseHelper.deleteDataStoreDataByAdmisionUuid(item, patientAdmissionDatastoreModel.uuid)
                        });

                    }
                });
        }

    }

}

