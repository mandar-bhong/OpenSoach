import { Injectable, Input } from "@angular/core";

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

@Injectable()
export class WorkerService {
    public ServerDataProcessorWorker: Worker;
    public DataReceivedSubject = new Subject<ServerDataStoreDataModel<any>>();
    public patientMasterDataReceivedSubject = new Subject<string>();
    public patientAdmissionDataReceivedSubject: Subject<string> = new Subject<string>();
    public scheduleDataReceivedSubject = new Subject<ScheduleDatastoreModel>();

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
        console.log('worker message recieved', messageEvent);
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
        }
    }

    closeServerDataProcessorWorker() {
        this.ServerDataProcessorWorker.terminate();
    }

    handleDataReceived(data: ServerDataStoreDataModel<IDatastoreModel>[]) {
        data.forEach(item => {
            console.log('subject triggered', item);

            switch (item.datastore) {
                case SYNC_STORE.PATIENT_MASTER:
                    this.patientMasterDataReceivedSubject.next(item.data.uuid);
                    break;
                case SYNC_STORE.PATIENT_ADMISSION:
                    this.patientAdmissionDataReceivedSubject.next(item.data.uuid);
                    // console.log('item.data.uuid', item.data.uuid);
                    break;
                case SYNC_STORE.SCHEDULE:
                    // TODO: 
                    // 
                    // if patient is selected and (<ScheduleDatastoreModel>item.data).admission_uuid equals to selected patient admission_uuid in AppGlobalContext
                    // then notify else do nothing
                    const getpatientdata = this.passDataService.getpatientData();
                    console.log('getpatientdata', getpatientdata);
                    if (getpatientdata.dbmodel.admission_uuid === (<ScheduleDatastoreModel>item.data).admission_uuid) {
                        this.scheduleDataReceivedSubject.next(<ScheduleDatastoreModel>item.data);
                        break;
                    }

            }
        });
    }
}