import { SyncDataModel } from "../models/ui/sync-models.js";
import { CmdModel, tokenModel, Header, GetSyncRequestModel, ApplySyncRequestModel } from "../models/api/server-cmd-model.js";
import { SERVER_SYNC_STATE, SYNC_STORE, SYNC_PENDING, CMD_CATEGORY, CMD_ID } from "../app-constants.js";
import { AppGlobalContext } from "../app-global-context.js";
import * as appSettings from "tns-core-modules/application-settings";
import { ServicePointDatastoreModel } from "../models/db/service-point-models.js";
import { ServerDataStoreDataModel } from "../models/api/server-data-store-data-model.js";
import { ConfDatastoreModel } from "../models/db/conf-model.js";
import { PatientMasterDatastoreModel } from "../models/db/patient-master-model.js";
import { ScheduleDatastoreModel } from "../models/db/schedule-model.js";
import { PatientAdmissionDatastoreModel } from "../models/db/patient-admission-model.js";
import { PatientPersonalDetailsDatastoreModel } from "../models/db/patient-personal-details-model.js";
import { PatientMedicalDetailsDatastoreModel } from "../models/db/patient-medical-details-model.js";
import { ActionTxnDatastoreModel } from "../models/db/action-txn-model.js";
import { AppMessageDbSyncHandler } from "../workers/app-message-dbsync-handler.js";
import { ServerWorkerEventDataModel } from "../models/api/server-worker-event-data-model.js";
import { SyncDb } from "../helpers/sync-db-helper.js";
import { RequestManager } from "./request-manager.js";
import { ServerWorkerContext, SYNC_TYPE } from "./server-worker-context.js";
import { SyncStoreManager } from "./sync-store-manager.js";
import { IDatastoreModel } from "../models/db/idatastore-model.js";
import { DatabaseHelper } from "../helpers/database-helper.js";

export class ServerHelper {

    public static sendToServerCallback: (msg: any) => void;
    static postMessageCallback: (msg: ServerWorkerEventDataModel) => void;
    constructor() {
    }

    public static Init(postMessageFn: (msg: ServerWorkerEventDataModel) => void) {
        ServerHelper.postMessageCallback = postMessageFn;
        console.log("in server helper init")
    }

    public static SyncProcess(syncstate: number) {

        switch (syncstate) {

            case SERVER_SYNC_STATE.SEND_AUTH_CMD:
                // send auth cmd
                // {"header":{"crc":"12","category":1,"commandid":1,"seqid":3},"payload":{"token":"Dev6AD88A481524BABF"}}

                console.log("SERVER_SYNC_STATE.SEND_AUTH_CMD");

                ServerWorkerContext.syncState = SERVER_SYNC_STATE.SEND_AUTH_CMD;

                const authcmd = this.AuthCmd();
                console.log("authcmd:", authcmd);

                ServerHelper.sendToServerCallback(authcmd);

                break;

            case SERVER_SYNC_STATE.AUTHOURIZED:
                // sync state authorized

                console.log("SERVER_SYNC_STATE.AUTHOURIZED");

                ServerWorkerContext.syncState = SERVER_SYNC_STATE.AUTHOURIZED;

                // read syncstore
                SyncStoreManager.ReadSyncStore().then(
                    (val) => {
                        console.log("reading sync store completed..")
                        this.SwitchSyncState();
                    },
                    (err) => {
                        console.log(err);
                    }
                );

                ServerWorkerContext.isSyncInprogress = true;
                ServerWorkerContext.syncType = SYNC_TYPE.FULL

                break;

            case SERVER_SYNC_STATE.DIFFERENTIAL_SYNC_STARTED:
                // sync state differential sync

                console.log("SERVER_SYNC_STATE.DIFFERENTIAL_SYNC_STARTED");

                ServerWorkerContext.syncState = SERVER_SYNC_STATE.DIFFERENTIAL_SYNC_STARTED;

                SyncStoreManager.ReadSyncStore().then(
                    (val) => {
                        console.log("reading sync store completed..")
                        this.SwitchSyncState();
                    },
                    (err) => {
                        console.log(err);
                    }
                );

                ServerWorkerContext.isSyncInprogress = true;
                ServerWorkerContext.syncType = SYNC_TYPE.DIFFERENTIAL

                break;


            case SERVER_SYNC_STATE.SYNC_TO_SERVER:
                //sync to server
                // send apply sync cmd
                // {"header":{"crc":"12","category":3,"commandid":51,"seqid":3},"payload":{"storename":"","storedata":[{"uuid":"PA001","bedno":"A0001"}]}}
                // Read Sync store and getnext store

                console.log("SERVER_SYNC_STATE.SYNC_TO_SERVER");

                //get next store
                var storename = SyncStoreManager.getNextStore(SERVER_SYNC_STATE.SYNC_TO_SERVER)
                console.log("SYNC_TO_SERVER storename:", storename)

                if (storename.currentStoreName != "") {
                    if (storename.currentStoreName == "getNextStore") {
                        this.SwitchSyncState();
                    } else {
                        DatabaseHelper.getSyncPendingDataStore(storename.currentStoreName)
                            .then(
                                (val) => {
                                    const syncCmd = this.ApplySyncCmd(storename.currentStoreName, val);
                                    console.log("apply syncCmd", syncCmd);
                                    ServerHelper.sendToServerCallback(syncCmd);
                                },
                                (err) => {
                                    console.log("getSyncPendingDataStore err:", err);
                                }
                            )
                    }
                } else {
                    ServerWorkerContext.syncState = SERVER_SYNC_STATE.SYNC_TO_SERVER;
                    this.SwitchSyncState();
                }

                break;

            case SERVER_SYNC_STATE.SYNC_TO_SERVER_COMPLETED:
                //sync to server completed

                console.log("SERVER_SYNC_STATE.SYNC_TO_SERVER_COMPLETED");
                ServerWorkerContext.syncState = SERVER_SYNC_STATE.SYNC_TO_SERVER_COMPLETED;
                SyncStoreManager.readSyncComplete = false;
                this.SwitchSyncState();

                break;

            case SERVER_SYNC_STATE.SYNC_FROM_SERVER:
                //sync from server
                // send get sync cmd with store name
                // {"header":{"crc":"12","category":3,"commandid":50,"seqid":3},"payload":{"storename":"","updatedon":"2018-10-30T00:00:00Z"}}

                console.log("SERVER_SYNC_STATE.SYNC_FROM_SERVER");

                //get next store
                storename = SyncStoreManager.getNextStore(SERVER_SYNC_STATE.SYNC_FROM_SERVER)
                console.log("SYNC_from_SERVER storename:", storename)


                if (storename.currentStoreName != "") {

                    if (storename.currentStoreName == "getNextStore") {
                        this.SwitchSyncState();
                    } else {
                        const syncCmd = this.GetSyncCmd(storename.currentStoreName, storename.lastSynched);
                        console.log("get syncCmd", syncCmd);
                        ServerHelper.sendToServerCallback(syncCmd);
                    }

                } else {
                    ServerWorkerContext.syncState = SERVER_SYNC_STATE.SYNC_FROM_SERVER;
                    this.SwitchSyncState();
                }

                break;

            case SERVER_SYNC_STATE.SYNC_FROM_SERVER_COMPLETED:
                //sync from server completed

                console.log("SERVER_SYNC_STATE.SYNC_FROM_SERVER_COMPLETED");

                if (ServerWorkerContext.syncType === SYNC_TYPE.DIFFERENTIAL) {
                    console.log("Differential Sync Completed");
                }

                ServerWorkerContext.isSyncInprogress = false;
                SyncStoreManager.updateSyncStore()

                break;
        }
    }

    public static SwitchSyncState() {

        console.log("SyncState", ServerWorkerContext.syncState);

        switch (ServerWorkerContext.syncState) {
            case SERVER_SYNC_STATE.NONE:
                this.SyncProcess(SERVER_SYNC_STATE.SEND_AUTH_CMD);
                break;

            case SERVER_SYNC_STATE.SEND_AUTH_CMD:
                this.SyncProcess(SERVER_SYNC_STATE.AUTHOURIZED);
                break;

            case SERVER_SYNC_STATE.DIFFERENTIAL_SYNC_INITIALISE:
                this.SyncProcess(SERVER_SYNC_STATE.DIFFERENTIAL_SYNC_STARTED);
                break;

            case SERVER_SYNC_STATE.DIFFERENTIAL_SYNC_STARTED:
                this.SyncProcess(SERVER_SYNC_STATE.SYNC_TO_SERVER);
                break;

            case SERVER_SYNC_STATE.AUTHOURIZED:
                this.SyncProcess(SERVER_SYNC_STATE.SYNC_TO_SERVER);
                break;

            case SERVER_SYNC_STATE.SYNC_TO_SERVER:
                this.SyncProcess(SERVER_SYNC_STATE.SYNC_TO_SERVER_COMPLETED)
                break;

            case SERVER_SYNC_STATE.SYNC_TO_SERVER_COMPLETED:
                this.SyncProcess(SERVER_SYNC_STATE.SYNC_FROM_SERVER)
                break;

            case SERVER_SYNC_STATE.SYNC_FROM_SERVER:
                this.SyncProcess(SERVER_SYNC_STATE.SYNC_FROM_SERVER_COMPLETED)
                break;
        }
    }

    public static AuthCmd() {
        // {"header":{"crc":"12","category":1,"commandid":1,"seqid":3},"payload":{"token":"Dev6AD88A481524BABF"}}

        const authcmd = new CmdModel();
        authcmd.header = new Header();

        authcmd.header.crc = '12';
        authcmd.header.category = CMD_CATEGORY.CMD_CAT_DEV_REGISTRATION;
        authcmd.header.commandid = CMD_ID.CMD_DEV_REGISTRATION;

        const tokenmodel = new tokenModel();
        // tokenmodel.token = AppGlobalContext.Token;
        tokenmodel.token = appSettings.getString("AUTH_TOKEN");

        authcmd.payload = tokenmodel;

        // set sequence number and map sequence no to request packet
        RequestManager.setSequencetNumber(authcmd);

        const cmdstring = JSON.stringify(authcmd);

        return cmdstring;
    }

    public static GetSyncCmd(strname: string, lastSynched: Date) {
        // {"header":{"crc":"12","category":3,"commandid":50,"seqid":3},"payload":{"storename":"","updatedon":"2018-10-30T00:00:00Z"}}

        const getSyncCmd = new CmdModel();
        getSyncCmd.header = new Header();

        getSyncCmd.header.crc = '12';
        getSyncCmd.header.category = CMD_CATEGORY.CMD_CAT_SYNC;
        getSyncCmd.header.commandid = CMD_ID.CMD_GET_STORE_SYNC;

        const getrequest = new GetSyncRequestModel();
        getrequest.storename = strname;

        // for first time sync - last synched empty then send default date
        if (lastSynched) {
            getrequest.updatedon = lastSynched;
        } else {
            getrequest.updatedon = new Date("2018-10-30T00:00:00Z");
        }

        getSyncCmd.payload = getrequest;

        // set sequence number and map sequence no to request packet
        RequestManager.setSequencetNumber(getSyncCmd);

        const cmdstring = JSON.stringify(getSyncCmd);

        return cmdstring
    }

    public static ApplySyncCmd(storename: any, storedata: any[]) {
        // {"header":{"crc":"12","category":3,"commandid":51,"seqid":3},"payload":{"storename":"","storedata":[{"uuid":"PA001","bedno":"A0001"}]}}

        const applySyncCmd = new CmdModel();
        applySyncCmd.header = new Header();

        applySyncCmd.header.crc = '12';
        applySyncCmd.header.category = CMD_CATEGORY.CMD_CAT_SYNC;
        applySyncCmd.header.commandid = CMD_ID.CMD_APPLY_STORE_SYNC;

        const applyReqModel = new ApplySyncRequestModel();
        applyReqModel.storename = storename;
        applyReqModel.storedata = storedata;

        applySyncCmd.payload = applyReqModel;

        // set sequence number and map sequence no to request packet
        RequestManager.setSequencetNumber(applySyncCmd);

        const cmdstring = JSON.stringify(applySyncCmd);

        return cmdstring;

    }

    public static CmdProcessor(respMsg: any) {

        // TODO: check if authorized if yes, set GlobalContext to Authorized
        // then call SwitchSyncState

        console.log(" in CmdProcessor..");

        console.log("respMsg", respMsg);

        let respDataModel = new CmdModel();
        respDataModel.header = new Header();
        respDataModel = JSON.parse(respMsg);

        // get request cmd packet
        const requestCmd = RequestManager.getRequest(respDataModel.header.seqid);
        console.log("requestCmd", requestCmd);

        if (requestCmd) {

            //handle response to cmd send by device
            switch (requestCmd.header.category, requestCmd.header.commandid) {

                case CMD_CATEGORY.CMD_CAT_DEV_REGISTRATION && CMD_ID.CMD_DEV_REGISTRATION:
                    // auth request cmd response
                    if (respDataModel.payload.ack == true) {
                        this.SwitchSyncState();
                    }

                    break;

                case CMD_CATEGORY.CMD_CAT_SYNC && CMD_ID.CMD_GET_STORE_SYNC:
                    // get sync request cmd response

                    this.SwitchSyncState();

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
                        }
                    }

                    break;

                case CMD_CATEGORY.CMD_CAT_SYNC && CMD_ID.CMD_APPLY_STORE_SYNC:
                    // apply sync request cmd response
                    // sync to server response - update individual tbl sync flag 
                    SyncStoreManager.updateTblSyncPending(requestCmd.payload.storename, requestCmd.payload.storedata[0].sync_pending_time);

                    this.SwitchSyncState();
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
        console.log("Service Point tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const servicePointDatastoreModel = new ServicePointDatastoreModel();
            const item = <ServicePointDatastoreModel>tblData[i];
            servicePointDatastoreModel.uuid = item.uuid;
            servicePointDatastoreModel.sp_name = item.sp_name;
            servicePointDatastoreModel.short_desc = item.short_desc;
            servicePointDatastoreModel.sp_state = item.sp_state;
            servicePointDatastoreModel.sp_state_since = item.sp_state_since;
            servicePointDatastoreModel.updated_on = item.updated_on;
            servicePointDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            // console.log("sp store data:", servicePointDatastoreModel);

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.SERVICE_POINT;
            serverDataStoreDataModel.data = servicePointDatastoreModel;

            // console.log("sp server data store model", serverDataStoreDataModel);

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

            // update sync table last synced
            DatabaseHelper.updateSyncStoreLastSynched(data.payload.ackdata.storename, data.payload.ackdata.updatedon);


        }

    }

    public static handleConfResponse(data: CmdModel) {
        console.log("Conf tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const confDatastoreModel = new ConfDatastoreModel();
            const item = <ConfDatastoreModel>tblData[i];
            confDatastoreModel.uuid = item.uuid;
            confDatastoreModel.conf_type_code = item.conf_type_code;
            confDatastoreModel.conf = item.conf;
            confDatastoreModel.updated_on = item.updated_on;
            confDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            // console.log("conf store data:", confDatastoreModel);

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.CONF;
            serverDataStoreDataModel.data = confDatastoreModel;

            // console.log("conf server data store model", serverDataStoreDataModel);

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

            // update sync table last synced
            DatabaseHelper.updateSyncStoreLastSynched(data.payload.ackdata.storename, data.payload.ackdata.updatedon);

        }

    }

    public static handlePatientMasterResponse(data: CmdModel) {
        console.log("Patient master tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const patientMasterDatastoreModel = new PatientMasterDatastoreModel();
            const item = <PatientMasterDatastoreModel>tblData[i];

            patientMasterDatastoreModel.uuid = item.uuid;
            patientMasterDatastoreModel.patient_reg_no = item.patient_reg_no;
            patientMasterDatastoreModel.fname = item.fname;
            patientMasterDatastoreModel.lname = item.lname;
            patientMasterDatastoreModel.mob_no = item.mob_no;
            patientMasterDatastoreModel.age = item.age;
            patientMasterDatastoreModel.blood_grp = item.blood_grp;
            patientMasterDatastoreModel.gender = item.gender;
            patientMasterDatastoreModel.updated_on = item.updated_on;
            patientMasterDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            // console.log("patient master store data:", patientMasterDatastoreModel);

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.PATIENT_MASTER;
            serverDataStoreDataModel.data = patientMasterDatastoreModel;

            // console.log("patient master server data store model", serverDataStoreDataModel);

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

            // update sync table last synced
            DatabaseHelper.updateSyncStoreLastSynched(data.payload.ackdata.storename, data.payload.ackdata.updatedon);

        }

    }

    public static handleScheduleResponse(data: CmdModel) {
        console.log("Patient schedule tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const scheduleDatastoreModel = new ScheduleDatastoreModel();
            const item = <ScheduleDatastoreModel>tblData[i];

            scheduleDatastoreModel.uuid = item.uuid;
            scheduleDatastoreModel.admission_uuid = item.admission_uuid;
            scheduleDatastoreModel.conf_type_code = item.conf_type_code;
            scheduleDatastoreModel.conf = item.conf;
            scheduleDatastoreModel.updated_on = item.updated_on;
            scheduleDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            // console.log("schedule store data:", scheduleDatastoreModel);

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.SCHEDULE;
            serverDataStoreDataModel.data = scheduleDatastoreModel;

            // console.log("schedule server data store model", serverDataStoreDataModel);

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

            // update sync table last synced
            DatabaseHelper.updateSyncStoreLastSynched(data.payload.ackdata.storename, data.payload.ackdata.updatedon);

        }

    }

    public static handlePatientAdmissionResponse(data: CmdModel) {
        console.log("Patient Admsn tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const patientAdmissionDatastoreModel = new PatientAdmissionDatastoreModel();
            const item = <PatientAdmissionDatastoreModel>tblData[i];

            patientAdmissionDatastoreModel.uuid = item.uuid;
            patientAdmissionDatastoreModel.patient_uuid = item.patient_uuid;
            patientAdmissionDatastoreModel.patient_reg_no = item.patient_reg_no;
            patientAdmissionDatastoreModel.bed_no = item.bed_no;
            patientAdmissionDatastoreModel.status = item.status;
            patientAdmissionDatastoreModel.sp_uuid = item.sp_uuid;
            patientAdmissionDatastoreModel.dr_incharge = item.dr_incharge;
            patientAdmissionDatastoreModel.admitted_on = item.admitted_on;
            patientAdmissionDatastoreModel.discharged_on = item.discharged_on;
            patientAdmissionDatastoreModel.updated_on = item.updated_on;
            patientAdmissionDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            // console.log("patient admsn store data:", patientAdmissionDatastoreModel);

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.PATIENT_ADMISSION;
            serverDataStoreDataModel.data = patientAdmissionDatastoreModel;

            // console.log("patient admsn server data store model", serverDataStoreDataModel);

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

            // update sync table last synced
            DatabaseHelper.updateSyncStoreLastSynched(data.payload.ackdata.storename, data.payload.ackdata.updatedon);

        }

    }

    public static handlePersonalDetailsResponse(data: CmdModel) {
        console.log("Personal details tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const patientPersonalDetailsDatastoreModel = new PatientPersonalDetailsDatastoreModel();
            const item = <PatientPersonalDetailsDatastoreModel>tblData[i];

            patientPersonalDetailsDatastoreModel.uuid = item.uuid;
            patientPersonalDetailsDatastoreModel.patient_uuid = item.patient_uuid;
            patientPersonalDetailsDatastoreModel.admission_uuid = item.admission_uuid;
            patientPersonalDetailsDatastoreModel.age = item.age;
            patientPersonalDetailsDatastoreModel.weight = item.weight;
            patientPersonalDetailsDatastoreModel.other_details = item.other_details;
            patientPersonalDetailsDatastoreModel.updated_on = item.updated_on;
            patientPersonalDetailsDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            // console.log("personal details store data:", patientPersonalDetailsDatastoreModel);

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.PERSONAL_DETAILS;
            serverDataStoreDataModel.data = patientPersonalDetailsDatastoreModel;

            // console.log("personal details server data store model", serverDataStoreDataModel);

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

            // update sync table last synced
            DatabaseHelper.updateSyncStoreLastSynched(data.payload.ackdata.storename, data.payload.ackdata.updatedon);

        }

    }

    public static handleMedicalDetailsResponse(data: CmdModel) {
        console.log("Medical details tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const patientMedicalDetailsDatastoreModel = new PatientMedicalDetailsDatastoreModel();
            const item = <PatientMedicalDetailsDatastoreModel>tblData[i];

            patientMedicalDetailsDatastoreModel.uuid = item.uuid;
            patientMedicalDetailsDatastoreModel.patient_uuid = item.patient_uuid;
            patientMedicalDetailsDatastoreModel.admission_uuid = item.admission_uuid;
            patientMedicalDetailsDatastoreModel.reason_for_admission = item.reason_for_admission;
            patientMedicalDetailsDatastoreModel.patient_medical_hist = item.patient_medical_hist;
            patientMedicalDetailsDatastoreModel.treatment_recieved_before = item.treatment_recieved_before;
            patientMedicalDetailsDatastoreModel.family_hist = item.family_hist;
            patientMedicalDetailsDatastoreModel.menstrual_hist = item.menstrual_hist;
            patientMedicalDetailsDatastoreModel.allergies = item.allergies;
            patientMedicalDetailsDatastoreModel.personal_history = item.personal_history;
            patientMedicalDetailsDatastoreModel.general_physical_exam = item.general_physical_exam;
            patientMedicalDetailsDatastoreModel.systematic_exam = item.systematic_exam;
            patientMedicalDetailsDatastoreModel.updated_on = item.updated_on;
            patientMedicalDetailsDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            // console.log("medical details store data:", patientMedicalDetailsDatastoreModel);

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.MEDICAL_DETAILS;
            serverDataStoreDataModel.data = patientMedicalDetailsDatastoreModel;

            // console.log("medical details server data store model", serverDataStoreDataModel);

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

            // update sync table last synced
            DatabaseHelper.updateSyncStoreLastSynched(data.payload.ackdata.storename, data.payload.ackdata.updatedon);

        }

    }

    public static handleActionTxnResponse(data: CmdModel) {
        console.log("Action Txn tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const actionTxnDatastoreModel = new ActionTxnDatastoreModel();
            const item = <ActionTxnDatastoreModel>tblData[i];

            actionTxnDatastoreModel.uuid = item.uuid;
            actionTxnDatastoreModel.schedule_uuid = item.schedule_uuid;
            actionTxnDatastoreModel.txn_data = item.txn_data;
            actionTxnDatastoreModel.txn_date = item.txn_date;
            actionTxnDatastoreModel.txn_state = item.txn_state;
            actionTxnDatastoreModel.conf_type_code = item.conf_type_code;
            actionTxnDatastoreModel.runtime_config_data = item.runtime_config_data;
            actionTxnDatastoreModel.updated_on = item.updated_on;
            actionTxnDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            // console.log("action txn store data:", actionTxnDatastoreModel);

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.ACTION_TXN;
            serverDataStoreDataModel.data = actionTxnDatastoreModel;

            // console.log("action txn server data store model", serverDataStoreDataModel);

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

            // update sync table last synced
            DatabaseHelper.updateSyncStoreLastSynched(data.payload.ackdata.storename, data.payload.ackdata.updatedon);

        }

    }

}


