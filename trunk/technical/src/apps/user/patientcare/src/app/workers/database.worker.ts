import "globals";
import { DBWorkerTasks } from "./database-worker-tasks.js";
import * as trace from 'tns-core-modules/trace'
import { TraceCustomCategory } from "../helpers/trace-helper.js";

const context: Worker = self as any;

context.onmessage = msg => {

    trace.write("Inside DB worker",TraceCustomCategory.DB_WORKER,trace.messageType.log);

    DBWorkerTasks.processMessage(msg.data);
};

DBWorkerTasks.Init(context);