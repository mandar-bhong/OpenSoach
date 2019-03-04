import { Component, OnInit, OnDestroy, NgZone } from "@angular/core";
import { DatabaseSchemaService } from "./services/offline-store/database-schema.service";
import { WorkerService } from "./services/worker.service";
import { Subscription } from "rxjs";
import { ServerDataProcessorMessageModel } from "./models/api/server-data-processor-message-model";
import { SERVER_WORKER_MSG_TYPE, API_SPL_BASE_URL } from "~/app/app-constants";
import * as appSettings from "tns-core-modules/application-settings";
import { APP_MODE } from "./app-constants";
import { AppGlobalContext } from "./app-global-context";
import { RouterExtensions } from "nativescript-angular/router";
import { PlatformHelper } from "./helpers/platform-helper";
import { PassDataService } from "./services/pass-data-service";
import { ServerApiInterfaceService } from "./services/server-api-interface.service";
//TODO: DONT Remove websocket require, if this line is removed. the worker doesnt get access to nativescript-websockets.
// this is a temporary fix and need to be handled through webpack config.
var WS = require('nativescript-websockets');
import * as trace from 'trace';
import { TraceCustomCategory } from "./helpers/trace-helper";

@Component({
    moduleId: module.id,
    selector: "ns-app",
    templateUrl: "app.component.html"
})
export class AppComponent implements OnInit, OnDestroy {

    public chatBox: string;
    private internetConnectionSubscription: Subscription;
    constructor(private databaseSchemaService: DatabaseSchemaService,
        private zone: NgZone,
        private workerService: WorkerService,
        private routerExtensions: RouterExtensions,
        private passDataService: PassDataService,
        private serverApiInterfaceService: ServerApiInterfaceService) {

        // init PlatformHelper
        PlatformHelper.init();
        this.databaseSchemaService.setOfflineDB().then(() => {
            this.onDbCreate();
        });
    }

    onDbCreate() {
        // Initialize the worker here
        this.workerService.initServerDataProcessorWorker();
        this.checkIfLoggedIn();
    }

    ngOnInit() {
        console.log('in app component init');
    }

    ngOnDestroy() {
        this.workerService.closeServerDataProcessorWorker();

        this.internetConnectionSubscription.unsubscribe();
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
                this.serverApiInterfaceService.get(API_SPL_BASE_URL + "/v1/validateauthtoken",
                    {
                        token: token
                    })
                    .then((result) => {
                        console.log("token validate success");
                        // setting context if token validated
                        AppGlobalContext.Token = token;
                        console.log("AppGlobalContext.Token", AppGlobalContext.Token);

                        // pass device token for any one access
                        AppGlobalContext.WebsocketUrl = appSettings.getString("WEB_SOCKET_URL");
                        trace.write(`AppGlobalContext.WebsocketUrl: ${AppGlobalContext.WebsocketUrl}`, TraceCustomCategory.APP_START, trace.messageType.info);
                        this.initAppStart();
                    }, (error) => {
                        if (!error.handled) {
                            trace.write(`token validate error: ${error}`, trace.categories.Error, trace.messageType.error);
                            this.checkIfDeviceIsRegistered();
                        }
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

        this.serverApiInterfaceService.post(API_SPL_BASE_URL + "/v1/endpoint/deviceauthorization",
            {
                'serialno': SerialNo,
                'prodcode': 'SPL_HPFT'
            })
            .then(
                res => {
                    console.log("POST Request is successful ", res);
                    this.handleDevAuthResponse(res);
                }, (error) => {
                    if (!error.handled) {
                        console.error('deviceauthorization in error', error);
                        this.routerExtensions.navigate(['login']);
                    }
                }
            );

    }

    handleDevAuthResponse(resData) {
        appSettings.setNumber("APP_MODE", APP_MODE.SHARED_DEVICE);
        AppGlobalContext.AppMode = APP_MODE.SHARED_DEVICE;
        appSettings.setString("AUTH_TOKEN", resData.token);
        appSettings.setString("WEB_SOCKET_URL", resData.locationurl);
        AppGlobalContext.Token = resData.token;
        AppGlobalContext.WebsocketUrl = resData.locationurl;
        console.log("AppGlobalContext.Token", AppGlobalContext.Token);
        this.initAppStart();
    }

    initAppStart() {
        // post message to worker to connect websocket
        // navigate to patient listing page

        console.log('in initappStart');
        const initModel = new ServerDataProcessorMessageModel();
        initModel.msgtype = SERVER_WORKER_MSG_TYPE.INIT_SERVER_INTERFACE;
        initModel.data = {};
        initModel.data.WebsocketUrl = AppGlobalContext.WebsocketUrl;
        initModel.data.Token = AppGlobalContext.Token;
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
