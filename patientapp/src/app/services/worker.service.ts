import { Injectable, Input, OnDestroy } from "@angular/core";

// add if building with webpack
import * as ServerDataProcessorWorker from "nativescript-worker-loader!../workers/server-data-processor.worker";
import { Subject } from "rxjs";
import { SERVER_WORKER_EVENT_MSG_TYPE, SYNC_STORE } from "~/app/app-constants";
import { ServerDataProcessorMessageModel } from "~/app/models/api/server-data-processor-message-model";
import { ServerWorkerEventDataModel } from "../models/api/server-worker-event-data-model";
import { ServerDataStoreDataModel } from "../models/api/server-data-store-data-model";
import { IDatastoreModel } from "../models/db/idatastore-model";
import { ScheduleDatastoreModel } from "../models/db/schedule-model";
import { PassDataService } from "./pass-data-service";
import { DocumentHelper } from "../helpers/document_helper";
import { PatientMasterDatastoreModel } from "../models/db/patient-master-model";
import { PatientAdmissionDatastoreModel } from "../models/db/patient-admission-model";
import { ActionDataStoreModel } from "../models/db/action-datastore";
import { ActionTxnDatastoreModel } from "../models/db/action-txn-model";
import { DoctorsOrdersDatastoreModel } from "../models/db/doctors-orders-model";
import { PatientPersonalDetailsDatastoreModel } from "../models/db/patient-personal-details-model";

@Injectable({providedIn:'root'})
export class WorkerService implements OnDestroy{
   
    public ServerDataProcessorWorker: Worker;
    public patientMasterDataReceivedSubject = new Subject<PatientMasterDatastoreModel>();
    public patientAdmissionDataReceivedSubject = new Subject<PatientAdmissionDatastoreModel>();
    public scheduleDataReceivedSubject = new Subject<ScheduleDatastoreModel>();
    public actionDataReceivedSubject = new Subject<ActionDataStoreModel>();
    public actionTxnDataReceivedSubject = new Subject<ActionTxnDatastoreModel>();
    public doctorOrderDataReceivedSubject = new Subject<DoctorsOrdersDatastoreModel>();
    public patientPersonalDetailsDataReceivedSubject = new Subject<PatientPersonalDetailsDatastoreModel>();

    // TODO: remove this, use instead scheduleDataReceivedSubject
    actionsSubject = new Subject<ServerDataStoreDataModel<IDatastoreModel>>();

    //TODO: remove this use instead doctorOrderDataReceivedSubject
    public ServerConnectionSubject = new Subject<boolean>();

    // patientname: string;
    constructor(private passDataService: PassDataService) {
        // this.patientname = this.patientName;
    }
    @Input() patientName: string;
    initServerDataProcessorWorker() {
        if (global["TNS_WEBPACK"]) {
            console.log('build with web pack');
            this.ServerDataProcessorWorker = new ServerDataProcessorWorker();
        } else {
            console.log('build without web pack');
            this.ServerDataProcessorWorker = new Worker("../workers/server-data-processor.worker");
        }

        this.ServerDataProcessorWorker.onmessage = m => this.serverWorkerMessageRecieved(m);
        this.ServerDataProcessorWorker.onerror = e => {
            console.log("worker error", e);
        };
    }

    postMessageToServerDataProcessorWorker(message: ServerDataProcessorMessageModel) {
        this.ServerDataProcessorWorker.postMessage(message);
    }

    serverWorkerMessageRecieved(messageEvent: MessageEvent) {
       // console.log('worker message recieved', messageEvent);
        const message: ServerWorkerEventDataModel = messageEvent.data;
        switch (message.msgtype) {
            case SERVER_WORKER_EVENT_MSG_TYPE.DATA_RECEIVED:
                this.handleDataReceived(message.data);
                break;
            case SERVER_WORKER_EVENT_MSG_TYPE.SERVER_CONNECTED:
                this.ServerConnectionSubject.next(true);
                break;
            case SERVER_WORKER_EVENT_MSG_TYPE.SERVER_DISCONNECTED:
                this.ServerConnectionSubject.next(false);
                break;
            case SERVER_WORKER_EVENT_MSG_TYPE.UPLOAD_DOCUMENT:
                DocumentHelper.uploadDocument(message.data, this);
                break;
        }
    }

    ngOnDestroy(): void {
        console.log("closeServerDataProcessorWorker.");
        this.ServerDataProcessorWorker.postMessage("close");

        this.ServerDataProcessorWorker.terminate();
    }

    handleDataReceived(data: ServerDataStoreDataModel<IDatastoreModel>[]) {
        data.forEach(item => {            
            switch (item.datastore) {
                case SYNC_STORE.PATIENT_MASTER:
                    this.patientMasterDataReceivedSubject.next(<PatientMasterDatastoreModel>item.data);
                    break;
                case SYNC_STORE.PATIENT_ADMISSION:
                    this.patientAdmissionDataReceivedSubject.next(<PatientAdmissionDatastoreModel>item.data);
                    break;
                case SYNC_STORE.SCHEDULE:
                  //  this.actionsSubject.next(item);
                  //  const scheduleDatastoreModel = new ScheduleDatastoreModel();
                  //  Object.assign(scheduleDatastoreModel, item.data);
                 //   this.scheduleDataReceivedSubject.next(scheduleDatastoreModel);                    // TODO: 
                     this.scheduleDataReceivedSubject.next(<ScheduleDatastoreModel>item.data);                    // 
                    // if patient is selected and (<ScheduleDatastoreModel>item.data).admission_uuid equals to selected patient admission_uuid in AppGlobalContext
                    // then notify else do nothing
                   // console.log('get patient data');
                    // const getpatientdata = this.passDataService.getpatientData();
                    // if (getpatientdata.dbmodel.admission_uuid === (<ScheduleDatastoreModel>item.data).admission_uuid) {
                    //     this.scheduleDataReceivedSubject.next(<ScheduleDatastoreModel>item.data);
                    // }
                   // this.scheduleDataReceivedSubject.next(<ScheduleDatastoreModel>item.data);
                 break;
                case SYNC_STORE.ACTION:
                    this.actionDataReceivedSubject.next(<ActionDataStoreModel>item.data);
                    break;
                case SYNC_STORE.DOCTORS_ORDERS:
                    this.doctorOrderDataReceivedSubject.next(<DoctorsOrdersDatastoreModel>item.data);
                    break;
                case SYNC_STORE.ACTION_TXN:
                    this.actionTxnDataReceivedSubject.next(<ActionTxnDatastoreModel>item.data);
                    break;
                case SYNC_STORE.PERSONAL_DETAILS:
                    this.patientPersonalDetailsDataReceivedSubject.next(<PatientPersonalDetailsDatastoreModel>item.data);
                    break;

            }
        });
    }
}