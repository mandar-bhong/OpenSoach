import { Component, OnInit, OnDestroy } from "@angular/core";
import { WorkerService } from "./services/worker.service";
//TODO: DONT Remove websocket require, if this line is removed. the worker doesnt get access to nativescript-websockets.
// this is a temporary fix and need to be handled through webpack config.
var WS = require('nativescript-websockets');
import * as trace from 'trace';
import { TraceCustomCategory } from "./helpers/trace-helper";
import { AppStartupService } from "./services/app-startup.service";

@Component({
    moduleId: module.id,
    selector: "ns-app",
    templateUrl: "app.component.html"
})
export class AppComponent implements OnInit, OnDestroy {

    constructor(private appStartupService: AppStartupService,
        private workerService: WorkerService) {
        trace.write("in app component constructor", TraceCustomCategory.APP_START, trace.messageType.info);
        this.appStartupService.init();
    }

    ngOnInit() {
        trace.write("in app component ngOnInit", TraceCustomCategory.APP_START, trace.messageType.info);
    }

    ngOnDestroy() {
        this.workerService.closeServerDataProcessorWorker();
    }
}