 import { ServerDataProcessorMessageModel } from "~/app/models/api/server-data-processor-message-model";
 import { SERVER_WORKER_MSG_TYPE } from "~/app/app-constants.js";
 const timerModule = require("tns-core-modules/timer");
 //require('~/app/app-constants')

require('globals');
var WS = require('nativescript-websockets');

const context: Worker = self as any;

context.onmessage = msg => {
    // setTimeout(() => {
        console.log("Inside TS worker...");
        WorkerTasks.processMessage(msg.data);                                    
    // }, 500)    

};

export class WorkerTasks
{
    private static socket: any;
    public static processMessage(msg:any)
    {
        
        console.log(msg);
        switch(msg.msgtype)
        {
            case SERVER_WORKER_MSG_TYPE.INIT_SERVER_INTERFACE:
            this.initWebSocket();
            break;
            case SERVER_WORKER_MSG_TYPE.SEND_MESSAGE:
            this.socket.send();
            break;
        }
       // this.initWebSocket();
        //(<any>global).postMessage("TS Worker");        
    }

    private static initWebSocket()
    {
        this.socket = new WS("ws://172.105.232.148:8090/ws", []);
        console.log('socket created', this.socket);

        this.socket.on('open', socket => {
                console.log('messages', "WebSocket opened");
                this.socket.send("test message");          
        });
        this.socket.on('message', (socket, message) => {
                console.log("websocket message recieved", message);            
        });
        
        this.socket.on('close', (socket, code, reason) => {
               console.log('messages', "Websocket disconnected");            
        });
        this.socket.on('error', (socket, error) => {
            console.log("The socket had an error", error);
        });

        this.socket.open();

       const timer = timerModule.setInterval(() => {
            const randNumber = Math.floor(Math.random());
            this.socket.send("message"+randNumber);
        }, 10000);
    }
}

