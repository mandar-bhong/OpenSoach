import { Component, OnInit, OnDestroy } from "@angular/core";
import { WorkerService } from "./services/worker.service";
//TODO: DONT Remove websocket require, if this line is removed. the worker doesnt get access to nativescript-websockets.
// this is a temporary fix and need to be handled through webpack config.
var WS = require('nativescript-websockets');
var bghttp = require("nativescript-background-http");
import * as trace from 'tns-core-modules/trace';
import { TraceCustomCategory } from "./helpers/trace-helper";
import { AppStartupService } from "./services/app-startup.service";

import { displayedEvent, exitEvent, launchEvent, lowMemoryEvent, 
    orientationChangedEvent, resumeEvent, suspendEvent, uncaughtErrorEvent, 
    ApplicationEventData, LaunchEventData, OrientationChangedEventData, UnhandledErrorEventData,
    on as applicationOn, run as applicationRun } from "tns-core-modules/application";

@Component({
    moduleId: module.id,
    selector: "ns-app",
    templateUrl: "app.component.html"
})
export class AppComponent implements OnInit, OnDestroy {

    constructor(private appStartupService: AppStartupService,
        private workerService: WorkerService) {
        trace.write("in app component constructor", TraceCustomCategory.APP_START, trace.messageType.info);


        applicationOn(launchEvent, (args: LaunchEventData) => {
            if (args.android) {
                // For Android applications, args.android is an android.content.Intent class.
                console.log("Launched Android application with the following intent: " + args.android + ".");
            } else if (args.ios !== undefined) {
                // For iOS applications, args.ios is NSDictionary (launchOptions).
                console.log("Launched iOS application with options: " + args.ios);
            }
        });
        
        applicationOn(suspendEvent, (args: ApplicationEventData) => {
            if (args.android) {
                // For Android applications, args.android is an android activity class.
                console.log("Activity: " + args.android);
            } else if (args.ios) {
                // For iOS applications, args.ios is UIApplication.
                console.log("UIApplication: " + args.ios);
            }
        });
        
        applicationOn(resumeEvent, (args: ApplicationEventData) => {
            if (args.android) {
                // For Android applications, args.android is an android activity class.
                console.log("Activity: " + args.android);
            } else if (args.ios) {
                // For iOS applications, args.ios is UIApplication.
                console.log("UIApplication: " + args.ios);
            }
        });
        
        applicationOn(displayedEvent, (args: ApplicationEventData) => {
            console.log("displayedEvent");
        });
        
        applicationOn(orientationChangedEvent, (args: OrientationChangedEventData) => {
            // "portrait", "landscape", "unknown"
            console.log(args.newValue)
        });
        
        applicationOn(exitEvent, (args: ApplicationEventData) => {
            if (args.android) {
                // For Android applications, args.android is an android activity class.
                console.log("Activity: " + args.android);
            } else if (args.ios) {
                // For iOS applications, args.ios is UIApplication.
                console.log("UIApplication: " + args.ios);
            }
        });
        
        applicationOn(lowMemoryEvent, (args: ApplicationEventData) => {
            if (args.android) {
                // For Android applications, args.android is an android activity class.
                console.log("Activity: " + args.android);
            } else if (args.ios) {
                // For iOS applications, args.ios is UIApplication.
                console.log("UIApplication: " + args.ios);
            }
        });
        
        applicationOn(uncaughtErrorEvent, function (args: UnhandledErrorEventData) {
            console.log("Error: " + args.error);
        });



        this.appStartupService.init();
    }

    ngOnInit() {
        trace.write("in app component ngOnInit", TraceCustomCategory.APP_START, trace.messageType.info);
    }

    ngOnDestroy() {
        this.workerService.closeServerDataProcessorWorker();
    }


    
    
    //applicationRun({ moduleName: "app-root" });
}