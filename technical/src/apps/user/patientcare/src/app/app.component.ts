import { Component, OnInit, OnDestroy, NgZone } from "@angular/core";
import { DatabaseSchemaService } from "./services/offline-store/database-schema.service";
import { server } from './environments/environment';
import { WorkerService } from "./services/worker.service";
import { Subscription } from "rxjs";
import { ServerDataProcessorMessageModel } from "./models/api/server-data-processor-message-model";
import { SERVER_WORKER_MSG_TYPE } from "~/app/app-constants";
var WS = require('nativescript-websockets');
import * as appSettings from "tns-core-modules/application-settings";
import { APP_MODE } from "./app-constants";
import { AppGlobalContext } from "./app-global-context";
import * as utils from "tns-core-modules/utils/utils";
import { HttpClient } from "@angular/common/http";
import { Router } from "@angular/router";
import * as application from 'application'
import { RouterExtensions } from "nativescript-angular/router";
import { PlatformHelper } from "./helpers/platform-helper";

@Component({
    moduleId: module.id,
    selector: "ns-app",
    templateUrl: "app.component.html"
})
export class AppComponent implements OnInit, OnDestroy {

    private socket: any;
    public messages: Array<any>;
    public chatBox: string;
    private internetConnectionSubscription: Subscription;
    constructor(private databaseSchemaService: DatabaseSchemaService,
        private zone: NgZone,
        private workerService: WorkerService,
        private routerExtensions: RouterExtensions,
        private httpClient: HttpClient) {
        // init PlatformHelper
        PlatformHelper.init();
        this.databaseSchemaService.setOfflineDB();
        // console.log('server', server);
        // this.socket = new WS("ws://echo.websocket.org", []);
        // console.log('socket created', this.socket);
        this.messages = [];
        this.chatBox = "";

        // Initialize the worker here
        this.workerService.initServerDataProcessorWorker();

        // Get APP_MODE

        // const appMode = appSettings.getNumber("APP_MODE", APP_MODE.NONE);
        // AppGlobalContext.AppMode = appMode;

        // if (appMode == APP_MODE.NONE) {
        //     // TODO Dummy code to set the application mode to Shared device
        //     appSettings.setNumber("APP_MODE", APP_MODE.SHARED_DEVICE);
        //     AppGlobalContext.AppMode = APP_MODE.SHARED_DEVICE;
        // }
        // console.log("APP_MODE", appSettings.getNumber("APP_MODE"));

        this.checkIfLoggedIn();
    }

    ngOnInit() {

        // TODO: Dummy code for testing 
        console.log('in app component init');
        // const initModel = new ServerDataProcessorMessageModel();
        // initModel.msgtype = SERVER_WORKER_MSG_TYPE.INIT_SERVER_INTERFACE;
        // this.workerService.ServerDataProcessorWorker.postMessage(initModel);

        // console.log('socketIO', this.socketIO);
        // this.socketIO.connect();
        // this.socket.on('open', socket => {
        //     this.zone.run(() => {
        //         this.messages.push("Welcome to the chat!");
        //         console.log('messages', this.messages);
        //         this.chatBox = "test message";
        //         this.send();
        //     });
        // });
        // this.socket.on('message', (socket, message) => {
        //     this.zone.run(() => {
        //         console.log("on message", message);
        //         this.messages.push(message);
        //         console.log('messages', this.messages);
        //     });
        // });
        //this.socket.on('message', function (socket, message) { console.log("Got a message", message); });

        // this.socket.on('close', (socket, code, reason) => {
        //     this.zone.run(() => {
        //         this.messages.push({ content: "You have been disconnected" });
        //         console.log('messages', this.messages);
        //     });
        // });
        // this.socket.on('error', (socket, error) => {
        //     console.log("The socket had an error", error);
        // });

        // this.socket.open();
    }

    ngOnDestroy() {
        //   this.socketIO.disconnect();
        // this.socket.close();
        // TODO: Send command to worker to disconnect the websocket before terminating server
        this.workerService.closeServerDataProcessorWorker();

        this.internetConnectionSubscription.unsubscribe();
    }

    public send() {
        if (this.chatBox) {
            this.socket.send(this.chatBox);
            this.chatBox = "";
        }
    }

    workerOnMessage(message: MessageEvent) {
        console.log('worker message recieved', message);

    }

    checkIfLoggedIn() {
        // read appSetting to get AUTH_TOKEN && APP_MODE
        // if APP_MODE doesnt exist call checkIfDeviceIsRegistered()
        // if APP_MODE exists && AUTH_TOKEN exists
        // check if token is still valid
        // if valid call initAppStart()
        // else call checkIfDeviceIsRegistered()

        const appMode = appSettings.getNumber("APP_MODE", APP_MODE.NONE);
        const token = appSettings.getString("AUTH_TOKEN");

        console.log("AUTH_TOKEN", token);
        console.log("appMode", appMode);

        if (appMode == APP_MODE.NONE) {
            this.checkIfDeviceIsRegistered();
        } else {
            if (appMode == APP_MODE.SHARED_DEVICE && token != null) {


                // http get method

                this.httpClient.get(
                    this.buildUrl("http://172.105.232.148/api/v1/validateauthtoken",
                        {
                            token: token,
                        })
                )
                    .subscribe((result) => {
                        console.log("result", result);
                        var res = <any>result;
                        if (res.issuccess === true) {
                            console.log("token validate success");
                            this.initAppStart();
                        } else {
                            console.log("token validate fail");
                            this.checkIfDeviceIsRegistered();
                        }
                    }, (error) => {
                        console.log(error);
                    });
            }

        }


    }

    getSerialNumber(): string {
        //TODO: Read the serial number
        // Set the Serial Number in AppGlobalContext

        const serialNumber = "1234567890123456";
        return serialNumber;

    }

    checkIfDeviceIsRegistered() {
        //TODO: Call HTTP API to check if device is registered. getSerialNumber()
        // Handle the response        
        //  Device is registered && is shared device
        // save token in appSetting and AppGlobalContext
        // call initAppStart();
        //Else
        // Navigate to login page


        // const SerialNo = PlatformHelper.API.getSerialNumber();
        const SerialNo = this.getSerialNumber();
        AppGlobalContext.SerialNumber = SerialNo;

        console.log("SerialNo:", SerialNo);

        this.httpClient.post("http://172.105.232.148/api/v1/endpoint/deviceauthorization",
            {
                'serialno': SerialNo,
                'prodcode': 'SPL_HPFT'
            })
            .subscribe(
                res => {
                    console.log("POST Request is successful ", res);
                    this.handleDevAuthResponse(res);
                }, (error) => {
                    console.log(error);
                }
            );

    }

    handleDevAuthResponse(resData) {
        if (resData.issuccess == true) {
            appSettings.setNumber("APP_MODE", APP_MODE.SHARED_DEVICE);
            AppGlobalContext.AppMode = APP_MODE.SHARED_DEVICE;
            appSettings.setString("AUTH_TOKEN", resData.data.token);
            AppGlobalContext.Token = resData.data.token;
            console.log("AppGlobalContext.Token", AppGlobalContext.Token);
            this.initAppStart();
        } else {
            this.routerExtensions.navigate(['login']);
        }
    }

    initAppStart() {
        // post message to worker to connect websocket
        // navigate to patient listing page

        console.log('in initappStart');
        const initModel = new ServerDataProcessorMessageModel();
        initModel.msgtype = SERVER_WORKER_MSG_TYPE.INIT_SERVER_INTERFACE;
        this.workerService.postMessageToServerDataProcessorWorker(initModel);

        this.routerExtensions.navigate(['home']);


    }

    buildUrl(url, parameters) {
        let qs = "";
        for (const key in parameters) {
            if (parameters.hasOwnProperty(key)) {
                const value = parameters[key];
                qs +=
                    "{" + "\"" + key + "\"" + ":" + "\"" + value + "\"" + "&";
            }
        }
        if (qs.length > 0) {
            qs = qs.substring(0, qs.length - 1); //chop off last "&"
            url = url + "?params=" + qs + "}";
        }

        console.log("url", url);

        return url;
    }

}
