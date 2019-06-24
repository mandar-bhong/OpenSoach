import { Component, OnInit, OnDestroy, isDevMode, ViewChild, ElementRef } from "@angular/core";
import { WorkerService } from "./services/worker.service";
//TODO: DONT Remove websocket require, if this line is removed. the worker doesnt get access to nativescript-websockets.
// this is a temporary fix and need to be handled through webpack config.
var WS = require('nativescript-websockets');
var bghttp = require("nativescript-background-http");
import * as trace from 'tns-core-modules/trace';
import { TraceCustomCategory } from "./helpers/trace-helper";
import { AppStartupService } from "./services/app-startup.service";

import { DatabaseService } from "./services/offline-store/database.service";
import { Router, NavigationStart, NavigationEnd } from "@angular/router";
import { filter } from "rxjs/operators";
import { Subscribable, Subscription } from "rxjs";

@Component({
    moduleId: module.id,
    selector: "ns-app",
    styles: [".indicator {color: #FF8912; vertical-align: center;height: 20%;width: 20%;}"],
    templateUrl: "app.component.html"
})
export class AppComponent implements OnInit {
    public isLoading = false;
    navigationStartsSubscription: Subscription;
    navigationEndSubscription: Subscription;
    @ViewChild("act", { static: true }) act: ElementRef;
    constructor(private appStartupService: AppStartupService,
        public router: Router,
        private workerService: WorkerService,
        private databaseService: DatabaseService) {
        trace.write("in app component constructor..", TraceCustomCategory.APP_START, trace.messageType.info);
       this.navigationStartsSubscription = router.events.pipe(
            filter(e => e instanceof NavigationStart)).subscribe((e: NavigationStart) => {             
                this.isLoading = true;
            });
        this.navigationEndSubscription=router.events.pipe(
            filter(e => e instanceof NavigationEnd)).subscribe((e: NavigationEnd) => {                
                setTimeout(() => {
                    this.isLoading = false;
                }, 100);
            });
        this.appStartupService.init();
        
    }

    ngOnInit() {
        trace.write("in app component ngOnInit", TraceCustomCategory.APP_START, trace.messageType.info);
    }

    ngOnDestroy() {
        trace.write("Destroying app component", TraceCustomCategory.APP_START, trace.messageType.info);
        if (this.navigationStartsSubscription) {
            this.navigationStartsSubscription.unsubscribe();
        }
        if(this.navigationEndSubscription){
            this.navigationEndSubscription.unsubscribe();
        }
    }
}



