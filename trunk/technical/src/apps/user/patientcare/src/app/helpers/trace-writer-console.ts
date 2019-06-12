import * as trace from 'tns-core-modules/trace';

let moment = require("moment");

export class TraceConsoleWriter {
    constructor() { }

    public write(message, category, type) {
        if (!console) {
            return;
        }

        let errorMsg ="";

        console.log(moment().unix());

        console.log(moment.utc().format('MMM DD, YYYY [at] hh:mm:ssA [(UTC)]'));
        console.log(moment.utc());

        if ( (<any>message).constructor.name == "Error"){            
            errorMsg = "Message : " + (message as Error).message +"\r\n" +  "StackTrace : " + (message as Error).stack + "\r\n";
        }else{
            errorMsg = message + "\r\n";
        }

        const traceMessage = new Date().toISOString() + " " + category + ": " + errorMsg;
        

        switch (type) {
            case trace.messageType.info:
                console.info("ConsoleLogger: "+traceMessage);
                break;
            case trace.messageType.warn:
                console.warn("ConsoleLogger: "+traceMessage);
                break;
            case trace.messageType.error:
                console.error("ConsoleLogger: "+traceMessage);
                break;
            default:
                console.log("ConsoleLogger: "+traceMessage);
                break;
        }
    }
}