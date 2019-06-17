import "globals";
import { DBWorkerTasks } from "./database-worker-tasks.js";
const context: Worker = self as any;

context.onmessage = msg => {
    console.log("Inside DB TS worker....");
    DBWorkerTasks.processMessage(msg.data);
};

DBWorkerTasks.Init(context);