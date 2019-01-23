import "globals";
import { WorkerTasks } from "./worker-tasks.js";

const context: Worker = self as any;

context.onmessage = msg => {
    console.log("Inside TS worker...");
    WorkerTasks.processMessage(msg.data);
};

WorkerTasks.Init(context);

console.log('worker started');