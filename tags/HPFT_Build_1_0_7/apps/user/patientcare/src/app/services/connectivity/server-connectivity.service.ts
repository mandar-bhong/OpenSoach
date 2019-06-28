import { Injectable } from "@angular/core";
import { WorkerService } from "~/app/services/worker.service";
import { Subject } from "rxjs";
import { InternetConnectionService } from "~/app/services/connectivity/internet-connection.service";
import { SERVER_WORKER_MSG_TYPE } from "~/app/app-constants";
import { ServerDataProcessorMessageModel } from "~/app/models/api/server-data-processor-message-model";

@Injectable()
export class ServerConnectivityStatusService {
    public ServerConnectionStatus: boolean;
    public InternetStatus: boolean;
    public ServerConnectionSubject = new Subject<boolean>();
    constructor(private workerService: WorkerService,
        private internetConnectionService: InternetConnectionService) {
        this.internetConnectionService.connectionStatus.subscribe(internetStatus => {
            this.InternetStatus = internetStatus;
            if (internetStatus) {
                const connectMessage = new ServerDataProcessorMessageModel();
                connectMessage.msgtype = SERVER_WORKER_MSG_TYPE.CONNECT_SERVER_INTERFACE;
                this.workerService.postMessageToServerDataProcessorWorker(connectMessage);
            }
        });

        this.workerService.ServerConnectionSubject.subscribe(status => {
            this.ServerConnectionStatus = status;
            this.ServerConnectionSubject.next(status);
        });
    }
}