import { DB_WORKER_MSG_TYPE } from "../app-constants.js";
import { DatabaseHelper } from "../helpers/database-helper.js";

export class DBWorkerTasks {

    private static workerReference: Worker;

    public static Init(worker: Worker) {
        console.log("inside db worker tasks..");
        DBWorkerTasks.workerReference = worker;
    }

    public static processMessage(msg: any) {
        console.log(msg);
        switch (msg.msgtype) {
            case DB_WORKER_MSG_TYPE.SELECT_ALL:
                const selectAllResp = DatabaseHelper.selectAll(msg.key);
                selectAllResp.then((result) => {
                    DBWorkerTasks.postMessage({ reqid: msg.reqid, respdata: result });
                }, (err) => {
                    DBWorkerTasks.postMessage({ reqid: msg.reqid, error: err });
                });
                break
            case DB_WORKER_MSG_TYPE.SELECT_BY_ID:
                const selectByIDResp = DatabaseHelper.selectByID(msg.key, msg.params);
                selectByIDResp.then((result) => {
                    DBWorkerTasks.postMessage({ reqid: msg.reqid, respdata: result });
                }, (err) => {
                    DBWorkerTasks.postMessage({ reqid: msg.reqid, error: err });
                });
                break
        }
    }

    public static postMessage(resp: any) {
        DBWorkerTasks.workerReference.postMessage(resp);
    }
}