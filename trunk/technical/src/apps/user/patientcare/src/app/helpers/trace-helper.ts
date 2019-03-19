import * as trace from 'trace';
import { TraceConsoleWriter } from './trace-writer-console.js';

export enum TraceCustomCategory {
    APP_START = "APP_START",
    SYNC = "SYNC",
    DATABASE = "DATABASE",
    WORKER="WORKER",
    SCHEDULE='SCHEDULE'
}
export class TraceHelper {
    constructor(){
        console.log('in TraceHelper constructor');
    }
    static configure() {
        trace.setCategories(trace.categories.concat(
            trace.categories.Error,
            trace.categories.Debug,
            TraceCustomCategory.APP_START,
            TraceCustomCategory.SYNC,
            TraceCustomCategory.DATABASE));

        // uncomment if trace to be enabled for all categories
        // trace.setCategories(trace.categories.concat(
        //     trace.categories.All,
        //     TraceCustomCategory.APP_START,
        //     TraceCustomCategory.SYNC,
        //     TraceCustomCategory.DATABASE));

        trace.clearWriters();
        trace.addWriter(new TraceConsoleWriter);
        trace.enable();
    }
}