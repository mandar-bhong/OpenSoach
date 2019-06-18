import { Injectable } from "@angular/core";
import {environment} from "~/app/environments/environment.tns"

@Injectable({ providedIn: 'root' })
export class AppRepoService {

    public static Instance = new AppRepoService();

    public DBWorkerRequestMapper: Map<number, any> = new Map<number, any>();
    public DBWorkerRequestCounter = 0;

    public API_SPL_BASE_URL = "http://172.105.232.148/api";
    public API_APP_BASE_URL = "http://172.105.232.148:91/api";
    public API_ERROR_LOGGING = 'http://172.105.232.148:8086/write?db=spl';



    constructor() {
        console.log("AppRepo service constructor executing");
        AppRepoService.Instance = this;

        this.API_SPL_BASE_URL = environment.API_SPL_BASE_URL;
        this.API_APP_BASE_URL = environment.API_APP_BASE_URL;
        this.API_ERROR_LOGGING = environment.API_ERROR_LOGGING;
    }

    public PreInit(){
        console.log("AppRepo service PreInit executing");
        this.API_SPL_BASE_URL = environment.API_SPL_BASE_URL;
        this.API_APP_BASE_URL = environment.API_APP_BASE_URL;
        this.API_ERROR_LOGGING = environment.API_ERROR_LOGGING;
    }

}