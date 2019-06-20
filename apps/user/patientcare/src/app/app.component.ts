import { Component, OnInit, OnDestroy, isDevMode } from "@angular/core";
import { WorkerService } from "./services/worker.service";
//TODO: DONT Remove websocket require, if this line is removed. the worker doesnt get access to nativescript-websockets.
// this is a temporary fix and need to be handled through webpack config.
var WS = require('nativescript-websockets');
var bghttp = require("nativescript-background-http");
import * as trace from 'tns-core-modules/trace';
import { TraceCustomCategory } from "./helpers/trace-helper";
import { AppStartupService } from "./services/app-startup.service";

import { DatabaseService } from "./services/offline-store/database.service";

@Component({
    moduleId: module.id,
    selector: "ns-app",
    templateUrl: "app.component.html"
})
export class AppComponent implements OnInit {

    constructor(private appStartupService: AppStartupService,
        private workerService: WorkerService,
        private databaseService: DatabaseService) {
        trace.write("in app component constructor..", TraceCustomCategory.APP_START, trace.messageType.info);

        this.appStartupService.init(); 
    }

    ngOnInit() {
        trace.write("in app component ngOnInit", TraceCustomCategory.APP_START, trace.messageType.info);
    }

    ngOnDestroy() {
        trace.write("Destroying app component", TraceCustomCategory.APP_START, trace.messageType.info);
        
    }
}



