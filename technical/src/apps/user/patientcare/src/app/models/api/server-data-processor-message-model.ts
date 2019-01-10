import { SERVER_WORKER_MSG_TYPE } from "../../app-constants";

export class ServerDataProcessorMessageModel {
	public msgtype: SERVER_WORKER_MSG_TYPE;
	public msg: any;
}