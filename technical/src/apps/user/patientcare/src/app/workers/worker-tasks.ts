import { Dummy } from "./dummy.js";
import { SERVER_WORKER_MSG_TYPE, SERVER_WORKER_EVENT_MSG_TYPE } from "../app-constants.js";
import { ServerWorkerEventDataModel } from "../models/api/server-worker-event-data-model.js";
import { ServerWorkerContext } from "./server-worker-context.js";

var WS = require('nativescript-websockets');

export class WorkerTasks {
    public static socket: any;
    private static isSocketInitialized: boolean;
    public static Init()
    {
        console.log("in WorkerTasks Init")
        Dummy.sendToServerCallback = WorkerTasks.sendToServer;
        ServerWorkerContext.ContextVar1="Worker Initialized";
        console.log('ServerWorkerContext.ContextVar1', ServerWorkerContext.ContextVar1);
    }

    public static processMessage(msg: any) {

        console.log(msg);
        switch (msg.msgtype) {
            case SERVER_WORKER_MSG_TYPE.INIT_SERVER_INTERFACE:
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
                WorkerTasks.socket.send();
                break;
        }               
    }

    private static initWebSocket() {
        WorkerTasks.socket = new WS("ws://172.105.232.148:8090/ws", []);
        console.log('socket created', WorkerTasks.socket);
        WorkerTasks.isSocketInitialized = true;
        WorkerTasks.socket.on('open', socket => {
            console.log('messages', "WebSocket opened");
            WorkerTasks.raiseSocketConnectionEvent(true);

            // TODO: Dummy Code to trigger data send
            Dummy.DummyMethod();
        });
        WorkerTasks.socket.on('message', (socket, message) => {
            console.log("websocket message recieved", message);
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

        (<any>global).postMessage(workerEvent);
    }

   public static sendToServer(msg:any):void
    {
        console.log("sendToServer",msg);
        console.log("sendToServer, socket",WorkerTasks.socket);
        WorkerTasks.socket.send(msg);
        console.log('ServerWorkerContext.ContextVar1', ServerWorkerContext.ContextVar1);
    }
}