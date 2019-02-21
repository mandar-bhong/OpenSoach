import { CmdModel } from "../models/api/server-cmd-model.js";
import { RequestManager } from "./request-manager.js";
import { CMD_ID, CMD_CATEGORY, SERVER_SYNC_STATE, SYNC_STORE, SYNC_PENDING } from "../app-constants.js";
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

export class CommandResponseProcessor{
    
    public static cmdProcessor(respMsg: any) {

        // TODO: check if authorized if yes, set GlobalContext to Authorized
        // then call SwitchSyncState

        console.log(" in CmdProcessor..");

        console.log("respMsg", respMsg);

        const respDataModel: CmdModel = JSON.parse(respMsg);

        // get request cmd packet
        const requestCmd = RequestManager.getRequest(respDataModel.header.seqid);
        console.log("requestCmd", requestCmd);

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
                        }
                    }

                    break;

                case CMD_CATEGORY.CMD_CAT_SYNC && CMD_ID.CMD_APPLY_STORE_SYNC:
                    // apply sync request cmd response
                    // sync to server response - update individual tbl sync flag 
                    SyncStoreManager.updateTblSyncPending(requestCmd.payload.storename, requestCmd.payload.storedata[0].sync_pending_time);

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
            servicePointDatastoreModel.updated_by = item.updated_by;
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
            confDatastoreModel.updated_by = item.updated_by;
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
            patientMasterDatastoreModel.updated_by = item.updated_by;
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
            scheduleDatastoreModel.updated_by = item.updated_by;
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
            patientAdmissionDatastoreModel.updated_by = item.updated_by;
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
            patientPersonalDetailsDatastoreModel.other_details = item.other_details;
            patientPersonalDetailsDatastoreModel.person_accompanying = item.person_accompanying;
            patientPersonalDetailsDatastoreModel.updated_by = item.updated_by;
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
            patientMedicalDetailsDatastoreModel.present_complaints = item.present_complaints;
            patientMedicalDetailsDatastoreModel.reason_for_admission = item.reason_for_admission;
            patientMedicalDetailsDatastoreModel.history_present_illness = item.history_present_illness;
            patientMedicalDetailsDatastoreModel.past_history = item.past_history;
            patientMedicalDetailsDatastoreModel.treatment_before_admission = item.treatment_before_admission;
            patientMedicalDetailsDatastoreModel.investigation_before_admission = item.investigation_before_admission;
            patientMedicalDetailsDatastoreModel.family_history = item.family_history;
            patientMedicalDetailsDatastoreModel.allergies = item.allergies;
            patientMedicalDetailsDatastoreModel.personal_history = item.personal_history;
            patientMedicalDetailsDatastoreModel.updated_by = item.updated_by;
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
            actionTxnDatastoreModel.updated_by = item.updated_by;
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

    public static handleDoctorsOrdersResponse(data: CmdModel) {
        console.log("Doctors Orders tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const doctorsOrdersDatastoreModel = new DoctorsOrdersDatastoreModel();
            const item = <DoctorsOrdersDatastoreModel>tblData[i];

            doctorsOrdersDatastoreModel.uuid = item.uuid;
            doctorsOrdersDatastoreModel.admission_uuid = item.admission_uuid;
            doctorsOrdersDatastoreModel.doctor_id = item.doctor_id;
            doctorsOrdersDatastoreModel.doctors_orders = item.doctors_orders;
            doctorsOrdersDatastoreModel.document_uuid = item.document_uuid;
            doctorsOrdersDatastoreModel.document_name = item.document_name;
            doctorsOrdersDatastoreModel.doctype = item.doctype;
            doctorsOrdersDatastoreModel.updated_by = item.updated_by;
            doctorsOrdersDatastoreModel.updated_on = item.updated_on;
            doctorsOrdersDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            // console.log("action txn store data:", doctorsOrdersDatastoreModel);

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.ACTION_TXN;
            serverDataStoreDataModel.data = doctorsOrdersDatastoreModel;

            // console.log("action txn server data store model", serverDataStoreDataModel);

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

            // update sync table last synced
            DatabaseHelper.updateSyncStoreLastSynched(data.payload.ackdata.storename, data.payload.ackdata.updatedon);

        }

    }

}

