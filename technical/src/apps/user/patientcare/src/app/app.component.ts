import { Component, OnInit, OnDestroy, NgZone } from "@angular/core";
import { DatabaseSchemaService } from "./services/offline-store/database-schema.service";
import { server } from './environments/environment';
import { WorkerService } from "./services/worker.service";
import { Subscription } from "rxjs";
import { ServerDataProcessorMessageModel } from "./models/api/server-data-processor-message-model";
import { SERVER_WORKER_MSG_TYPE, API_SPL_BASE_URL } from "~/app/app-constants";
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
import { PassDataService } from "./services/pass-data-service";

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
        private httpClient: HttpClient,
        private passDataService: PassDataService) {

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


                // http get method

                this.httpClient.get(
                    this.buildUrl(API_SPL_BASE_URL + "/v1/validateauthtoken",
                        {
                            token: token,
                        })
                )
                    .subscribe((result) => {
                        console.log("result", result);
                        var res = <any>result;
                        if (res.issuccess === true) {
                            console.log("token validate success");
                            // setting context if token validated
                            AppGlobalContext.Token = token;
                            console.log("AppGlobalContext.Token", AppGlobalContext.Token);

                            // pass device token for any one access
                            this.passDataService.token = AppGlobalContext.Token;
                            AppGlobalContext.WebsocketUrl = appSettings.getString("WEB_SOCKET_URL");
                            console.log("AppGlobalContext.WebsocketUrl", AppGlobalContext.WebsocketUrl);
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

        this.httpClient.post(API_SPL_BASE_URL + "/v1/endpoint/deviceauthorization",
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
            appSettings.setString("WEB_SOCKET_URL", resData.data.locationurl);
            AppGlobalContext.Token = resData.data.token;
            AppGlobalContext.WebsocketUrl = resData.data.locationurl;
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
