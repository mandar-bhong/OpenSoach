import { Injectable } from "@angular/core";

// add if building with webpack
import * as ServerDataProcessorWorker from "nativescript-worker-loader!../workers/server-data-processor.worker";

@Injectable()
export class WorkerService {
    public ServerDataProcessorWorker: Worker;
    constructor() {
    }

    initServerDataProcessorWorker()
    {        
        if (global["TNS_WEBPACK"]) {
            console.log('build with web pack');
            this.ServerDataProcessorWorker=new ServerDataProcessorWorker();
        } else {
            console.log('build without web pack');
            this.ServerDataProcessorWorker = new Worker("../workers/server-data-processor.worker");
        }
    }
}