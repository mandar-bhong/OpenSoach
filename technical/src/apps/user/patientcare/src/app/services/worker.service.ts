import { Injectable } from "@angular/core";

// add if building with webpack
import * as ServerDataProcessorWorker from "nativescript-worker-loader!../workers/server-data-processor.worker";
import { Subject } from "rxjs";
import { SERVER_WORKER_EVENT_MSG_TYPE } from "~/app/app-constants";

@Injectable()
export class WorkerService {
    public ServerDataProcessorWorker: Worker;
    public ServerNotificationSubject = new Subject<any>();
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

    serverWorkerMessageRecieved(message: MessageEvent) {
        console.log('worker message recieved', message);        
        switch(message.data.msgtype)
        {
            case SERVER_WORKER_EVENT_MSG_TYPE.DATA_RECEIVED:
            console.log('subject triggered');
            this.ServerNotificationSubject.next(message.data);
            break;
        }        
    }
}