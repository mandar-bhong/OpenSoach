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
import { TreatmentDatastoreModel } from "../models/db/treatment-model.js";
import { TreatmentDocDatastoreModel } from "../models/db/treatment-doc-model.js";
import { PathologyRecordDatastoreModel } from "../models/db/pathology-record-model.js";
import { PathologyRecordDocDatastoreModel } from "../models/db/pathology-record-doc-model.js";
import { ActionDataStoreModel } from "../models/db/action-datastore.js";
import { UserDatastoreModel } from "../models/db/user-model.js";

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
        // console.log("Conf tbl data", data);

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
        // console.log("Patient master tbl data", data);

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
        // console.log("Patient schedule tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const scheduleDatastoreModel = new ScheduleDatastoreModel();
            const item = <ScheduleDatastoreModel>tblData[i];

            scheduleDatastoreModel.uuid = item.uuid;
            scheduleDatastoreModel.admission_uuid = item.admission_uuid;
            scheduleDatastoreModel.conf_type_code = item.conf_type_code;
            scheduleDatastoreModel.conf = item.conf;
            scheduleDatastoreModel.end_date = item.end_date;
            scheduleDatastoreModel.status = item.status;
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
        // console.log("Patient Admsn tbl data", data);

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
        // console.log("Personal details tbl data", data);

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
        // console.log("Medical details tbl data", data);

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
        // console.log("Action Txn tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const actionTxnDatastoreModel = new ActionTxnDatastoreModel();
            const item = <ActionTxnDatastoreModel>tblData[i];

            actionTxnDatastoreModel.uuid = item.uuid;
            actionTxnDatastoreModel.schedule_uuid = item.schedule_uuid;
            actionTxnDatastoreModel.txn_data = item.txn_data;
            actionTxnDatastoreModel.scheduled_time = item.scheduled_time;
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
        // console.log("Doctors Orders tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const doctorsOrdersDatastoreModel = new DoctorsOrdersDatastoreModel();
            const item = <DoctorsOrdersDatastoreModel>tblData[i];

            doctorsOrdersDatastoreModel.uuid = item.uuid;
            doctorsOrdersDatastoreModel.admission_uuid = item.admission_uuid;
            doctorsOrdersDatastoreModel.doctor_id = item.doctor_id;
            doctorsOrdersDatastoreModel.doctors_orders = item.doctors_orders;
            doctorsOrdersDatastoreModel.comment = item.comment;
            doctorsOrdersDatastoreModel.ack_by = item.ack_by;
            doctorsOrdersDatastoreModel.ack_time = item.ack_time;
            doctorsOrdersDatastoreModel.status = item.status;
            doctorsOrdersDatastoreModel.order_created_time = item.order_created_time;
            doctorsOrdersDatastoreModel.order_type = item.order_type;
            doctorsOrdersDatastoreModel.document_uuid = item.document_uuid;
            doctorsOrdersDatastoreModel.document_name = item.document_name;
            doctorsOrdersDatastoreModel.doctype = item.doctype;
            doctorsOrdersDatastoreModel.updated_by = item.updated_by;
            doctorsOrdersDatastoreModel.updated_on = item.updated_on;
            doctorsOrdersDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            // console.log("action txn store data:", doctorsOrdersDatastoreModel);

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.DOCTORS_ORDERS;
            serverDataStoreDataModel.data = doctorsOrdersDatastoreModel;

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

            // update sync table last synced
            DatabaseHelper.updateSyncStoreLastSynched(data.payload.ackdata.storename, data.payload.ackdata.updatedon);

        }

    }

    public static handleTreatmentResponse(data: CmdModel) {
        // console.log("Treatment tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const treatmentDatastoreModel = new TreatmentDatastoreModel();
            const item = <TreatmentDatastoreModel>tblData[i];

            treatmentDatastoreModel.uuid = item.uuid;
            treatmentDatastoreModel.admission_uuid = item.admission_uuid;
            treatmentDatastoreModel.treatment_done = item.treatment_done;
            treatmentDatastoreModel.treatment_performed_time = item.treatment_performed_time;
            treatmentDatastoreModel.details = item.details;
            treatmentDatastoreModel.post_observation = item.post_observation;
            treatmentDatastoreModel.updated_by = item.updated_by;
            treatmentDatastoreModel.updated_on = item.updated_on;
            treatmentDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.TREATMENT;
            serverDataStoreDataModel.data = treatmentDatastoreModel;

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

            // update sync table last synced
            DatabaseHelper.updateSyncStoreLastSynched(data.payload.ackdata.storename, data.payload.ackdata.updatedon);

        }

    }

    public static handleTreatmentDocResponse(data: CmdModel) {
        // console.log("Treatment doc tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const treatmentDocDatastoreModel = new TreatmentDocDatastoreModel();
            const item = <TreatmentDocDatastoreModel>tblData[i];

            treatmentDocDatastoreModel.uuid = item.uuid;
            treatmentDocDatastoreModel.treatment_uuid = item.treatment_uuid;
            treatmentDocDatastoreModel.document_uuid = item.document_uuid;
            treatmentDocDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.TREATMENT_DOC;
            serverDataStoreDataModel.data = treatmentDocDatastoreModel;

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

            // update sync table last synced
            DatabaseHelper.updateSyncStoreLastSynched(data.payload.ackdata.storename, data.payload.ackdata.updatedon);

        }

    }

    public static handlePathologyRecordOrdersResponse(data: CmdModel) {
        // console.log("Pathology record tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const pathologyRecordDatastoreModel = new PathologyRecordDatastoreModel();
            const item = <PathologyRecordDatastoreModel>tblData[i];

            pathologyRecordDatastoreModel.uuid = item.uuid;
            pathologyRecordDatastoreModel.admission_uuid = item.admission_uuid;
            pathologyRecordDatastoreModel.test_performed = item.test_performed;
            pathologyRecordDatastoreModel.test_performed_time = item.test_performed_time;
            pathologyRecordDatastoreModel.test_result = item.test_result;
            pathologyRecordDatastoreModel.comments = item.comments;
            pathologyRecordDatastoreModel.updated_by = item.updated_by;
            pathologyRecordDatastoreModel.updated_on = item.updated_on;
            pathologyRecordDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.PATHOLOGY_RECORD;
            serverDataStoreDataModel.data = pathologyRecordDatastoreModel;

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

            // update sync table last synced
            DatabaseHelper.updateSyncStoreLastSynched(data.payload.ackdata.storename, data.payload.ackdata.updatedon);

        }

    }

    public static handlePathologyRecordDocResponse(data: CmdModel) {
        // console.log("Pathology record doc tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const pathologyRecordDocDatastoreModel = new PathologyRecordDocDatastoreModel();
            const item = <PathologyRecordDocDatastoreModel>tblData[i];

            pathologyRecordDocDatastoreModel.uuid = item.uuid;
            pathologyRecordDocDatastoreModel.pathology_record_uuid = item.pathology_record_uuid;
            pathologyRecordDocDatastoreModel.document_uuid = item.document_uuid;
            pathologyRecordDocDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.PATHOLOGY_RECORD_DOC;
            serverDataStoreDataModel.data = pathologyRecordDocDatastoreModel;

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

            // update sync table last synced
            DatabaseHelper.updateSyncStoreLastSynched(data.payload.ackdata.storename, data.payload.ackdata.updatedon);

        }

    }

    public static handleActionResponse(data: CmdModel) {
        // console.log("action tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const actionDataStoreModel = new ActionDataStoreModel();
            const item = <ActionDataStoreModel>tblData[i];

            actionDataStoreModel.uuid = item.uuid;
            actionDataStoreModel.admission_uuid = item.admission_uuid;
            actionDataStoreModel.conf_type_code = item.conf_type_code;
            actionDataStoreModel.schedule_uuid = item.schedule_uuid;
            actionDataStoreModel.scheduled_time = item.scheduled_time
            actionDataStoreModel.is_deleted = item.is_deleted;
            actionDataStoreModel.updated_by = item.updated_by;
            actionDataStoreModel.updated_on = item.updated_on;
            actionDataStoreModel.sync_pending = SYNC_PENDING.FALSE;

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.ACTION;
            serverDataStoreDataModel.data = actionDataStoreModel;

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

            // update sync table last synced
            DatabaseHelper.updateSyncStoreLastSynched(data.payload.ackdata.storename, data.payload.ackdata.updatedon);

        }

    }

    public static handleUserResponse(data: CmdModel) {
        // console.log("user tbl data", data);

        const tblData = data.payload.ackdata.data

        for (var i = 0; i < tblData.length; i++) {
            const userDatastoreModel = new UserDatastoreModel();
            const item = <UserDatastoreModel>tblData[i];

            userDatastoreModel.usr_id=item.usr_id
            userDatastoreModel.usr_name=item.usr_name
            userDatastoreModel.usr_id=item.usr_id
            userDatastoreModel.urole_name=item.urole_name
            userDatastoreModel.fname=item.fname
            userDatastoreModel.lname=item.lname
            userDatastoreModel.updated_on = item.updated_on;
            userDatastoreModel.sync_pending = SYNC_PENDING.FALSE;

            const serverDataStoreDataModel = new ServerDataStoreDataModel<IDatastoreModel>();
            serverDataStoreDataModel.datastore = SYNC_STORE.USER;
            serverDataStoreDataModel.data = userDatastoreModel;

            new AppMessageDbSyncHandler().handleMessage(serverDataStoreDataModel, ServerHelper.postMessageCallback);

            // update sync table last synced
            DatabaseHelper.updateSyncStoreLastSynched(data.payload.ackdata.storename, data.payload.ackdata.updatedon);

        }

    }

}

