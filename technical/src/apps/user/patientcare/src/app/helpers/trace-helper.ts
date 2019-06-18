import * as trace from 'tns-core-modules/trace';
import { TraceConsoleWriter } from './trace-writer-console.js';
//import Utils from './utils.js';
//import { BuildMode } from '../app-constants.js';
//import { TraceServerWriter } from './trace-writer-server.js';
//HTTPClient is not available in background worker giving compilation error


export enum TraceCustomCategory {
    APP_START = "APP_START",
    SERVICE = "SERVICE",
    APP_EXCEPTION = "APP_EXCEPTION",
    SYNC = "SYNC",
    DATABASE = "DATABASE",
    WORKER = "WORKER",
    DB_WORKER = "DB-WORKER",
    SCHEDULE = 'SCHEDULE'
}

declare var process: any;


export class TraceHelper {

    constructor() {
    }

    static getEnvironmentVars(key: string): string {
        // console.log('process', process);
        if (typeof process !== 'undefined' && process && process.env) {
            return process.env[key];
        } else {
            return "Development"
        }
    }


    static devErrorHandler: trace.ErrorHandler = {
        handlerError(err){
            console.log("Handled in dev error handler");
            trace.write(err,err.name,trace.messageType.error);
        }
    }

    static prodErrorHandler: trace.ErrorHandler = {
        handlerError(err){
            console.log("Handled in prod error handler");
            trace.write(err,err.name,trace.messageType.error);
        }
    }


    static configure() {
        trace.setCategories(trace.categories.concat(            
            trace.categories.Error,
            trace.categories.Debug,
            TraceCustomCategory.APP_START,
            TraceCustomCategory.SYNC,
            TraceCustomCategory.WORKER,
            TraceCustomCategory.DATABASE));

        // uncomment if trace to be enabled for all categories
        // trace.setCategories(trace.categories.concat(
        //     trace.categories.All,
        //     TraceCustomCategory.APP_START,
        //     TraceCustomCategory.SYNC,
        //     TraceCustomCategory.DATABASE));

        trace.clearWriters();
        

        if (false){//Utils.getBuildEnvironment()==BuildMode.PRODUCTION            
            trace.setErrorHandler(TraceHelper.prodErrorHandler);            
        }else{
            trace.addWriter(new TraceConsoleWriter);
            trace.setErrorHandler(TraceHelper.devErrorHandler);
        }
        
        
        trace.enable();
    }
}