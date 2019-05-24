import "globals";
import { WorkerTasks } from "./worker-tasks.js";
import * as trace from 'tns-core-modules/trace';
import { TraceCustomCategory, TraceHelper } from "../helpers/trace-helper.js";
const context: Worker = self as any;

TraceHelper.configure();

context.onmessage = msg => {
    console.log("Inside TS worker...");
    WorkerTasks.processMessage(msg.data);
};

WorkerTasks.Init(context);
//trace.write('worker started', TraceCustomCategory.WORKER, trace.messageType.log);