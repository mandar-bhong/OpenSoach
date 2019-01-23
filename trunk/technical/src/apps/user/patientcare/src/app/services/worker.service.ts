import { Injectable } from "@angular/core";

// add if building with webpack
import * as ServerDataProcessorWorker from "nativescript-worker-loader!../workers/server-data-processor.worker";
import { Subject } from "rxjs";
import { SERVER_WORKER_EVENT_MSG_TYPE } from "~/app/app-constants";
import { ServerDataProcessorMessageModel } from "~/app/models/api/server-data-processor-message-model";
import { ServerWorkerEventDataModel } from "../models/api/server-worker-event-data-model";
import { ServerDataStoreDataModel } from "../models/api/server-data-store-data-model";

@Injectable()
export class WorkerService {
    private ServerDataProcessorWorker: Worker;
    public DataReceivedSubject = new Subject<ServerDataStoreDataModel>();
    public ServerConnectionSubject = new Subject<boolean>();

    constructor() {
    }

    initServerDataProcessorWorker() {
        if (global["TNS_WEBPACK"]) {
            console.log('build with web pack');
            this.ServerDataProcessorWorker = new ServerDataProcessorWorker();
        } else {
            console.log('build without web pack');
            this.ServerDataProcessorWorker = new Worker("../workers/server-data-processor.worker");
        }

        this.ServerDataProcessorWorker.onmessage = m => this.serverWorkerMessageRecieved(m);
    }

    postMessageToServerDataProcessorWorker(message: ServerDataProcessorMessageModel) {
        this.ServerDataProcessorWorker.postMessage(message);
    }

    serverWorkerMessageRecieved(messageEvent: MessageEvent) {
        console.log('worker message recieved', messageEvent);
        const message: ServerWorkerEventDataModel = messageEvent.data;
        switch (message.msgtype) {
            case SERVER_WORKER_EVENT_MSG_TYPE.DATA_RECEIVED:
                message.data.forEach(element => {
                    console.log('subject triggered');
                    this.DataReceivedSubject.next(element);
                });
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
}