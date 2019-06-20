import { DB_WORKER_MSG_TYPE } from "../../app-constants";

export class DBDataProcessorMessageModel {
    public msgtype: DB_WORKER_MSG_TYPE;
    public reqid: number;
    public key: string;
    public params: Array<any>;
}

export class DBDataProcessorMessageResponseModel {
    public reqid: number;
    public respdata: Array<any>;
    public error: string
}