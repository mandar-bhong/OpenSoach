import { SERVER_WORKER_MSG_TYPE, SERVER_WORKER_EVENT_MSG_TYPE, SYNC_STORE, SERVER_SYNC_STATE } from "../app-constants.js";
import { ServerWorkerEventDataModel } from "../models/api/server-worker-event-data-model.js";
import { ServerWorkerContext } from "./server-worker-context.js";
import { AppMessageUIHandler } from "./app-message-ui-handler.js";
import { ServerHelper } from "./server-helper.js";
import { PlatformHelper } from "../helpers/platform-helper.js";
import { CommandResponseProcessor } from "./command-response-processor.js";

var WS = require('nativescript-websockets');

export class WorkerTasks {
    public static socket: any;
    private static isSocketInitialized: boolean;
    private static workerReference: Worker;
    public static Init(worker: Worker) {
        WorkerTasks.workerReference = worker;
        console.log("in WorkerTasks Init")
        ServerHelper.init(WorkerTasks.postMessage);
        ServerHelper.sendToServerCallback = WorkerTasks.sendToServer;
        PlatformHelper.init();
    }

    public static processMessage(msg: any) {

        console.log(msg);
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
        }
    }

    private static initWebSocket() {
        WorkerTasks.socket = new WS(ServerWorkerContext.serverUrl, []);
        console.log('socket created', WorkerTasks.socket);
        WorkerTasks.isSocketInitialized = true;
        WorkerTasks.socket.on('open', socket => {
            console.log('messages', "WebSocket opened");
            WorkerTasks.raiseSocketConnectionEvent(true);

            //on connect sync data
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
            WorkerTasks.raiseSocketConnectionEvent(false);
        });
        WorkerTasks.socket.on('error', (socket, error) => {
            console.log("The socket had an error", error);
            WorkerTasks.raiseSocketConnectionEvent(false);
        });

        WorkerTasks.socket.open();
    }

    private static raiseSocketConnectionEvent(status: boolean) {
        const workerEvent = new ServerWorkerEventDataModel();
        if (status) {
            workerEvent.msgtype = SERVER_WORKER_EVENT_MSG_TYPE.SERVER_CONNECTED;
        }
        else {
            workerEvent.msgtype = SERVER_WORKER_EVENT_MSG_TYPE.SERVER_DISCONNECTED;
        }

        WorkerTasks.postMessage(workerEvent);
    }

    public static sendToServer(msg: any): void {
        console.log("sendToServer", msg);
        console.log("sendToServer, socket", WorkerTasks.socket);
        WorkerTasks.socket.send(msg);
    }

    public static postMessage(msg: ServerWorkerEventDataModel) {
        WorkerTasks.workerReference.postMessage(msg);
    }
}