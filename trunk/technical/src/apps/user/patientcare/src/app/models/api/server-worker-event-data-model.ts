import { SERVER_WORKER_EVENT_MSG_TYPE } from "../../app-constants";

export class ServerWorkerEventDataModel {
	public msgtype: SERVER_WORKER_EVENT_MSG_TYPE;
	public msg: any;
}