import { SERVER_WORKER_MSG_TYPE, SERVER_WORKER_EVENT_MSG_TYPE, SYNC_STORE, SERVER_SYNC_STATE } from "../app-constants.js";
import { ServerWorkerEventDataModel } from "../models/api/server-worker-event-data-model.js";
import { ServerWorkerContext, SYNC_TYPE } from "./server-worker-context.js";
import { AppMessageUIHandler } from "./app-message-ui-handler.js";
import { ServerHelper } from "./server-helper.js";
import { PlatformHelper } from "../helpers/platform-helper.js";
import { CommandResponseProcessor } from "./command-response-processor.js";
import { DocumentSyncHelper } from "./document-sync-helper.js";
import { TraceCustomCategory } from "../helpers/trace-helper.js";
import * as trace from "tns-core-modules/trace"

var WS = require('nativescript-websockets');

export class WorkerTasks {
    public static socket: any;
    private static isSocketInitialized: boolean;
    private static workerReference: Worker;
    private static retryConnectionTimerStarted = false;
    private static retryConnectionTimer: any;
    
    
    public static Init(worker: Worker) {
        trace.write("Initilizing Worker task",TraceCustomCategory.WORKER,trace.messageType.info);

        WorkerTasks.workerReference = worker;
        ServerHelper.init(WorkerTasks.postMessage);
        ServerHelper.sendToServerCallback = WorkerTasks.sendToServer;
        PlatformHelper.init();
    }

    public static processMessage(msg: any) {
        trace.write("Processing msg in worker" + JSON.stringify(msg) ,TraceCustomCategory.SYNC,trace.messageType.log);
        switch (msg.msgtype) {
            case SERVER_WORKER_MSG_TYPE.INIT_SERVER_INTERFACE:
                // set server worker context
                ServerWorkerContext.serverUrl = msg.data.WebsocketUrl;
                ServerWorkerContext.authToken = msg.data.Token;
                WorkerTasks.initWebSocket();
                break;
            case SERVER_WORKER_MSG_TYPE.CONNECT_SERVER_INTERFACE:
                if (WorkerTasks.isSocketInitialized) {
                    if (WorkerTasks.socket.readyState == 3) // closed
                    {
                        WorkerTasks.socket.open();
                    }
                }
                break;
            case SERVER_WORKER_MSG_TYPE.SEND_MESSAGE:
                msg.data.forEach(element => {
                    const appMessageUIHandler = new AppMessageUIHandler();
                    appMessageUIHandler.handleMessage(element, WorkerTasks.postMessage);
                });
                break;
            case SERVER_WORKER_MSG_TYPE.UPLOAD_DOCUMENT_COMPLETED:
                // call delete document sync
                DocumentSyncHelper.deleteDocFromLocalStore(msg.data)
                break;
        }
    }

    private static initWebSocket() {
        WorkerTasks.socket = new WS(ServerWorkerContext.serverUrl, []);
        console.log('socket created', WorkerTasks.socket);
        trace.write("Socket created" ,TraceCustomCategory.SYNC,trace.messageType.info);

        WorkerTasks.isSocketInitialized = true;
        WorkerTasks.socket.on('open', socket => {
            console.log('messages', "WebSocket opened");
            trace.write("Weboscket opned." ,TraceCustomCategory.SYNC,trace.messageType.info);

            WorkerTasks.raiseSocketConnectionEvent(true);

            //on connect starting sync
            ServerWorkerContext.isSyncInprogress = false;
            ServerWorkerContext.syncType = SYNC_TYPE.NONE
            ServerWorkerContext.syncState = SERVER_SYNC_STATE.NONE;
            ServerHelper.switchSyncState();


        });
        WorkerTasks.socket.on('message', (socket, message) => {
            // console.log("websocket message recieved", message);

            // process resp msg
            CommandResponseProcessor.cmdProcessor(message);

        });

        WorkerTasks.socket.on('close', (socket, code, reason) => {
            console.log('messages', "Websocket disconnected");
            WorkerTasks.isSocketInitialized = true;
            WorkerTasks.raiseSocketConnectionEvent(false);
        });
        WorkerTasks.socket.on('error', (socket, error) => {
            console.log("The socket had an error", error);
            WorkerTasks.isSocketInitialized = true;
            WorkerTasks.raiseSocketConnectionEvent(false);
        });

        WorkerTasks.socket.open();
    }

    private static raiseSocketConnectionEvent(status: boolean) {
        const workerEvent = new ServerWorkerEventDataModel();
        if (status) {
            this.stopRetryConnection();
            workerEvent.msgtype = SERVER_WORKER_EVENT_MSG_TYPE.SERVER_CONNECTED;
        }
        else {
            // delay websocket re-connection until server recieves disconnection event
            setTimeout(() => {
                this.retryConnection();
            }, 60000);
            workerEvent.msgtype = SERVER_WORKER_EVENT_MSG_TYPE.SERVER_DISCONNECTED;
        }

        WorkerTasks.postMessage(workerEvent);
    }

    public static sendToServer(msg: any): void {
        trace.write("",TraceCustomCategory.SYNC,trace.messageType.log);
        WorkerTasks.socket.send(msg);
    }

    public static postMessage(msg: ServerWorkerEventDataModel) {
        WorkerTasks.workerReference.postMessage(msg);
    }

    public static retryConnection() {
        if ((!this.retryConnectionTimer) || (this.retryConnectionTimer==null)) {
            this.retryConnectionTimer = setInterval(() => {
                console.log('connection retrying', new Date());
                this.initWebSocket();
            }, 10 * 1000);
        }
    }

    public static stopRetryConnection() {
        if (this.retryConnectionTimer) {
            console.log('stoping retrying connection', new Date());
            clearInterval(this.retryConnectionTimer);
            this.retryConnectionTimer = null;
        }
    }

}