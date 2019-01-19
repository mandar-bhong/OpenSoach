import { WorkerTasks } from "./worker-tasks.js";

require('globals');

const context: Worker = self as any;

context.onmessage = msg => {
    console.log("Inside TS worker...");
    WorkerTasks.processMessage(msg.data);
};

WorkerTasks.Init();