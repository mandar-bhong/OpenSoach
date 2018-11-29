import { Component, OnInit, OnDestroy, NgZone } from "@angular/core";
import { DatabaseSchemaService } from "./services/offline-store/database-schema.service";
import { server } from './environments/environment';
var WS = require('nativescript-websockets');

@Component({
    moduleId: module.id,
    selector: "ns-app",
    templateUrl: "app.component.html"
})
export class AppComponent implements OnInit, OnDestroy {

    private socket: any;
    public messages: Array<any>;
    public chatBox: string;
    constructor(private databaseSchemaService: DatabaseSchemaService,
        private zone: NgZone) {
        this.databaseSchemaService.setOfflineDB();
        console.log('server', server);
        this.socket = new WS("ws://echo.websocket.org", []);
        console.log('socket created', this.socket);
        this.messages = [];
        this.chatBox = "";
    }

    ngOnInit() {
        // console.log('socketIO', this.socketIO);
        // this.socketIO.connect();
        this.socket.on('open', socket => {
            this.zone.run(() => {
                this.messages.push("Welcome to the chat!");
                console.log('messages', this.messages);
                this.chatBox = "test message";
                this.send();
            });
        });
        this.socket.on('message', (socket, message) => {
            this.zone.run(() => {
                console.log("on message", message);
                this.messages.push(message);
                console.log('messages', this.messages);
            });
        });
        //this.socket.on('message', function (socket, message) { console.log("Got a message", message); });

        this.socket.on('close', (socket, code, reason) => {
            this.zone.run(() => {
                this.messages.push({ content: "You have been disconnected" });
                console.log('messages', this.messages);
            });
        });
        this.socket.on('error', (socket, error) => {
            console.log("The socket had an error", error);
        });

        this.socket.open();
    }

    ngOnDestroy() {
        //   this.socketIO.disconnect();
        this.socket.close();
    }

    public send() {
        if (this.chatBox) {
            this.socket.send(this.chatBox);
            this.chatBox = "";
        }
    }



}
