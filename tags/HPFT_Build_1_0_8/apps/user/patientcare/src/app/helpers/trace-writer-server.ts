import * as trace from 'tns-core-modules/trace';
import { HttpHeaders, HttpClient, HttpXhrBackend } from '@angular/common/http';
// import { AppGlobalContext } from '../app-global-context.js';
// import { InfluxDb_Log } from '../app-constants.js';
let moment = require("moment");

export class TraceServerWriter {
    constructor() { }

    public write(message, category, type) {
        const traceMessage = new Date().toISOString() + " " + category + ": " + message;
       
        let msgTypeName = "";

        switch (type) {
            case trace.messageType.info:
                msgTypeName = "Info";
                break;
            case trace.messageType.warn:
                msgTypeName = "Warn";
                break;
            case trace.messageType.log:
                msgTypeName = "Debug";
                break;
            case trace.messageType.error:
                msgTypeName = "Error";
                break;
            default:
                msgTypeName = "Unknown";
                break;
        }

        this.postError(msgTypeName, category, message);
    }

    

    //category message , type
    postError(msgType: string, category: any, errMsg: any ) {
        let httpClient = new HttpClient(new HttpXhrBackend({ build: () => new XMLHttpRequest() }));
        //const filedItems = fields || fields != null ? JSON.stringify(fields) : fields;

        let deviceIdentifier = "123412";//need to updated with actual helper

        let errorMsg ;
        
        let req = `spl.hpft AppComponent="Device:HPFT:${deviceIdentifier}",SubComponent="${category}",LogLevel="${msgType}",`;
    
        if ( (<any>errMsg).constructor.name == "Error"){
            req += `Message="${(errMsg as Error).message}",Error="${(errMsg as Error).stack}",`;                               
        }else{
            req += `Message="${errMsg }",`;
        }

        req += `Time="${moment.utc().format('MMM DD, YYYY [at] hh:mm:ssa [(UTC)]')}"`;

        console.log('request sent to  influxdb', req);
        const headers = new HttpHeaders().set('Content-Transfer-Encoding', 'binary');
        httpClient.post("http://172.105.232.148:8086/write?db=spl", req, { headers: headers, responseType: 'text' }).subscribe((payloadResponse) => {
            console.log('working');
        }, error => {
            console.log(Date.now());
        });
    }
}