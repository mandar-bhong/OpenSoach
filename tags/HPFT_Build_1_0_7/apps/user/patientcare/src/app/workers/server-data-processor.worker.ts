import "globals";
import { WorkerTasks } from "./worker-tasks.js";
import * as trace from 'tns-core-modules/trace';
import { TraceCustomCategory, TraceHelper } from "../helpers/trace-helper.js";


const context: Worker = self as any;

TraceHelper.configure();

context.onmessage = msg => {
    if (msg.data === "close"){
        trace.write("Closing the worker and deinit worker task ",TraceCustomCategory.WORKER,trace.messageType.info);
        WorkerTasks.DeInit();
    }

    trace.write("Inside TS worker..",TraceCustomCategory.WORKER,trace.messageType.info);
    WorkerTasks.processMessage(msg.data);
};

WorkerTasks.Init(context);
//trace.write('worker started', TraceCustomCategory.WORKER, trace.messageType.log);