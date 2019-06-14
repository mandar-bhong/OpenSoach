import { Injectable, Input } from "@angular/core";
import { DatabaseHelper } from "~/app/helpers/database-helper";
// import * as DatabaseWorker from "nativescript-worker-loader!../../workers/database.worker";
import { DB_WORKER_MSG_TYPE } from "../../app-constants";
import { PassDataService } from "../pass-data-service";
import { DBDataProcessorMessageModel } from "~/app/models/db/db-data-processor-message-model";

@Injectable()
export class DatabaseService {

    public static dbWorker: Worker;
    public static currentSeqNo = 0;
    static reqRespMapper: Map<number, any>;

    constructor() {
    }

    initDBWorker() {
        // if (global["TNS_WEBPACK"]) {
        //     console.log('build with web pack');
        //     DatabaseService.dbWorker = new DatabaseWorker();
        // } else {
        //     console.log('build without web pack');
        //     DatabaseService.dbWorker = new Worker("../../workers/database.worker");
        // }

        // DatabaseService.dbWorker.onmessage = DatabaseService.DBWorkerMsgHandle;
        // DatabaseService.reqRespMapper = new Map<number, any>();

        // DatabaseService.dbWorker.onerror = e => {
        //     console.log("database worker error", e);
        // };
    }
    

    public getdbConnection() {
        return DatabaseHelper.getdbConn();
    }

    public closedbConnection() {
        DatabaseHelper.closedbConn();
    }

    public deleteDatabaseInDebugMode() {
        DatabaseHelper.deleteDatabaseInDebugMode();
    }

    public update(key: string, dataList: Array<any>) {
        const response = DatabaseHelper.update(key, dataList);
        return response;
    }

    public selectByID(key: string, paramList: Array<any>): any {
        return DatabaseHelper.selectByID(key,paramList);
        // return DatabaseService.DBWorkerPostMsg(DB_WORKER_MSG_TYPE.SELECT_BY_ID, key, paramList);
    }

    public selectAll(key: string): any {
        return DatabaseHelper.selectAll(key);
        // return DatabaseService.DBWorkerPostMsg(DB_WORKER_MSG_TYPE.SELECT_ALL, key);
    }

    static DBWorkerPostMsg(msgType: DB_WORKER_MSG_TYPE, key: string, paramList?: Array<any>) {
        return new Promise((resolve, reject) => {

            function resolveQueryResponse(queryResp, isSuccess) {
                if (isSuccess) {
                    resolve(queryResp);
                } else {
                    reject(queryResp);
                }
            };

            DatabaseService.currentSeqNo++;
            DatabaseService.reqRespMapper.set(DatabaseService.currentSeqNo, resolveQueryResponse);

            const dbDataProcessorMessageModel = new DBDataProcessorMessageModel
            dbDataProcessorMessageModel.reqid = DatabaseService.currentSeqNo
            dbDataProcessorMessageModel.msgtype = msgType
            dbDataProcessorMessageModel.key = key;            

            if (paramList) {
                dbDataProcessorMessageModel.params = paramList
                DatabaseService.dbWorker.postMessage(dbDataProcessorMessageModel);
            } else {
                DatabaseService.dbWorker.postMessage(dbDataProcessorMessageModel);
            }

        });
    }    

    static DBWorkerMsgHandle(msg) {
        if (DatabaseService.reqRespMapper) {
            if (msg.data.respdata) {
                DatabaseService.reqRespMapper.get(msg.data.reqid)(msg.data.respdata, true);
            } else if (msg.data.error) {
                DatabaseService.reqRespMapper.get(msg.data.reqid)(msg.data.error, false);
            }
            DatabaseService.reqRespMapper.delete(msg.data.reqid);
        } else {
            console.error("request mapper is undefined");
        }

    }

    public closeDBDataProcessorWorker() {
        console.log("closeDBDataProcessorWorker");
        DatabaseService.dbWorker.terminate();
    }

}