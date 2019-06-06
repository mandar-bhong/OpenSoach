"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var app_constants_js_1 = require("../app-constants.js");
var server_worker_event_data_model_js_1 = require("../models/api/server-worker-event-data-model.js");
var server_worker_context_js_1 = require("./server-worker-context.js");
var app_message_ui_handler_js_1 = require("./app-message-ui-handler.js");
var server_helper_js_1 = require("./server-helper.js");
var platform_helper_js_1 = require("../helpers/platform-helper.js");
var command_response_processor_js_1 = require("./command-response-processor.js");
var document_sync_helper_js_1 = require("./document-sync-helper.js");
var WS = require('nativescript-websockets');
var WorkerTasks = /** @class */ (function () {
    function WorkerTasks() {
    }
    WorkerTasks.Init = function (worker) {
        WorkerTasks.workerReference = worker;
        console.log("in WorkerTasks Init");
        server_helper_js_1.ServerHelper.init(WorkerTasks.postMessage);
        server_helper_js_1.ServerHelper.sendToServerCallback = WorkerTasks.sendToServer;
        platform_helper_js_1.PlatformHelper.init();
    };
    WorkerTasks.processMessage = function (msg) {
        console.log(msg);
        switch (msg.msgtype) {
            case app_constants_js_1.SERVER_WORKER_MSG_TYPE.INIT_SERVER_INTERFACE:
                // set server worker context
                server_worker_context_js_1.ServerWorkerContext.serverUrl = msg.data.WebsocketUrl;
                server_worker_context_js_1.ServerWorkerContext.authToken = msg.data.Token;
                WorkerTasks.initWebSocket();
                break;
            case app_constants_js_1.SERVER_WORKER_MSG_TYPE.CONNECT_SERVER_INTERFACE:
                if (WorkerTasks.isSocketInitialized) {
                    if (WorkerTasks.socket.readyState == 3) // closed
                     {
                        WorkerTasks.socket.open();
                    }
                }
                break;
            case app_constants_js_1.SERVER_WORKER_MSG_TYPE.SEND_MESSAGE:
                msg.data.forEach(function (element) {
                    var appMessageUIHandler = new app_message_ui_handler_js_1.AppMessageUIHandler();
                    appMessageUIHandler.handleMessage(element, WorkerTasks.postMessage);
                });
                break;
            case app_constants_js_1.SERVER_WORKER_MSG_TYPE.UPLOAD_DOCUMENT_COMPLETED:
                // call delete document sync
                document_sync_helper_js_1.DocumentSyncHelper.deleteDocFromLocalStore(msg.data);
                break;
        }
    };
    WorkerTasks.initWebSocket = function () {
        WorkerTasks.socket = new WS(server_worker_context_js_1.ServerWorkerContext.serverUrl, []);
        console.log('socket created', WorkerTasks.socket);
        WorkerTasks.isSocketInitialized = true;
        WorkerTasks.socket.on('open', function (socket) {
            console.log('messages', "WebSocket opened");
            WorkerTasks.raiseSocketConnectionEvent(true);
            //on connect sync data
            server_worker_context_js_1.ServerWorkerContext.syncState = app_constants_js_1.SERVER_SYNC_STATE.NONE;
            server_helper_js_1.ServerHelper.switchSyncState();
        });
        WorkerTasks.socket.on('message', function (socket, message) {
            // console.log("websocket message recieved", message);
            // process resp msg
            command_response_processor_js_1.CommandResponseProcessor.cmdProcessor(message);
        });
        WorkerTasks.socket.on('close', function (socket, code, reason) {
            console.log('messages', "Websocket disconnected");
            WorkerTasks.raiseSocketConnectionEvent(false);
        });
        WorkerTasks.socket.on('error', function (socket, error) {
            console.log("The socket had an error", error);
            WorkerTasks.raiseSocketConnectionEvent(false);
        });
        WorkerTasks.socket.open();
    };
    WorkerTasks.raiseSocketConnectionEvent = function (status) {
        var workerEvent = new server_worker_event_data_model_js_1.ServerWorkerEventDataModel();
        if (status) {
            workerEvent.msgtype = app_constants_js_1.SERVER_WORKER_EVENT_MSG_TYPE.SERVER_CONNECTED;
        }
        else {
            workerEvent.msgtype = app_constants_js_1.SERVER_WORKER_EVENT_MSG_TYPE.SERVER_DISCONNECTED;
        }
        WorkerTasks.postMessage(workerEvent);
    };
    WorkerTasks.sendToServer = function (msg) {
        console.log("sendToServer", msg);
        WorkerTasks.socket.send(msg);
    };
    WorkerTasks.postMessage = function (msg) {
        WorkerTasks.workerReference.postMessage(msg);
    };
    return WorkerTasks;
}());
exports.WorkerTasks = WorkerTasks;
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoid29ya2VyLXRhc2tzLmpzIiwic291cmNlUm9vdCI6IiIsInNvdXJjZXMiOlsid29ya2VyLXRhc2tzLnRzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiI7O0FBQUEsd0RBQTBIO0FBQzFILHFHQUE2RjtBQUM3Rix1RUFBaUU7QUFDakUseUVBQWtFO0FBQ2xFLHVEQUFrRDtBQUNsRCxvRUFBK0Q7QUFDL0QsaUZBQTJFO0FBQzNFLHFFQUErRDtBQUUvRCxJQUFJLEVBQUUsR0FBRyxPQUFPLENBQUMseUJBQXlCLENBQUMsQ0FBQztBQUU1QztJQUFBO0lBa0dBLENBQUM7SUE5RmlCLGdCQUFJLEdBQWxCLFVBQW1CLE1BQWM7UUFDN0IsV0FBVyxDQUFDLGVBQWUsR0FBRyxNQUFNLENBQUM7UUFDckMsT0FBTyxDQUFDLEdBQUcsQ0FBQyxxQkFBcUIsQ0FBQyxDQUFBO1FBQ2xDLCtCQUFZLENBQUMsSUFBSSxDQUFDLFdBQVcsQ0FBQyxXQUFXLENBQUMsQ0FBQztRQUMzQywrQkFBWSxDQUFDLG9CQUFvQixHQUFHLFdBQVcsQ0FBQyxZQUFZLENBQUM7UUFDN0QsbUNBQWMsQ0FBQyxJQUFJLEVBQUUsQ0FBQztJQUMxQixDQUFDO0lBRWEsMEJBQWMsR0FBNUIsVUFBNkIsR0FBUTtRQUVqQyxPQUFPLENBQUMsR0FBRyxDQUFDLEdBQUcsQ0FBQyxDQUFDO1FBQ2pCLFFBQVEsR0FBRyxDQUFDLE9BQU8sRUFBRTtZQUNqQixLQUFLLHlDQUFzQixDQUFDLHFCQUFxQjtnQkFDN0MsNEJBQTRCO2dCQUM1Qiw4Q0FBbUIsQ0FBQyxTQUFTLEdBQUcsR0FBRyxDQUFDLElBQUksQ0FBQyxZQUFZLENBQUM7Z0JBQ3RELDhDQUFtQixDQUFDLFNBQVMsR0FBRyxHQUFHLENBQUMsSUFBSSxDQUFDLEtBQUssQ0FBQztnQkFDL0MsV0FBVyxDQUFDLGFBQWEsRUFBRSxDQUFDO2dCQUM1QixNQUFNO1lBQ1YsS0FBSyx5Q0FBc0IsQ0FBQyx3QkFBd0I7Z0JBQ2hELElBQUksV0FBVyxDQUFDLG1CQUFtQixF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LEdBQVE7UUFDL0IsT0FBTyxDQUFDLEdBQUcsQ0FBQyxjQUFjLEVBQUUsR0FBRyxDQUFDLENBQUM7UUFDakMsV0FBVyxDQUFDLE1BQU0sQ0FBQyxJQUFJLENBQUMsR0FBRyxDQUFDLENBQUM7SUFDakMsQ0FBQztJQUVhLHVCQUFXLEdBQXpCLFVBQTBCLEdBQStCO1FBQ3JELFdBQVcsQ0FBQyxlQUFlLENBQUMsV0FBVyxDQUFDLEdBQUcsQ0FBQyxDQUFDO0lBQ2pELENBQUM7SUFFTCxrQkFBQztBQUFELENBQUMsQUFsR0QsSUFrR0M7QUFsR1ksa0NBQVciLCJzb3VyY2VzQ29udGVudCI6WyJpbXBvcnQgeyBTRVJWRVJfV09SS0VSX01TR19UWVBFLCBTRVJWRVJfV09SS0VSX0VWRU5UX01TR19UWVBFLCBTWU5DX1NUT1JFLCBTRVJWRVJfU1lOQ19TVEFURSB9IGZyb20gXCIuLi9hcHAtY29uc3RhbnRzLmpzXCI7XHJcbmltcG9ydCB7IFNlcnZlcldvcmtlckV2ZW50RGF0YU1vZGVsIH0gZnJvbSBcIi4uL21vZGVscy9hcGkvc2VydmVyLXdvcmtlci1ldmVudC1kYXRhLW1vZGVsLmpzXCI7XHJcbmltcG9ydCB7IFNlcnZlcldvcmtlckNvbnRleHQgfSBmcm9tIFwiLi9zZXJ2ZXItd29ya2VyLWNvbnRleHQuanNcIjtcclxuaW1wb3J0IHsgQXBwTWVzc2FnZVVJSGFuZGxlciB9IGZyb20gXCIuL2FwcC1tZXNzYWdlLXVpLWhhbmRsZXIuanNcIjtcclxuaW1wb3J0IHsgU2VydmVySGVscGVyIH0gZnJvbSBcIi4vc2VydmVyLWhlbHBlci5qc1wiO1xyXG5pbXBvcnQgeyBQbGF0Zm9ybUhlbHBlciB9IGZyb20gXCIuLi9oZWxwZXJzL3BsYXRmb3JtLWhlbHBlci5qc1wiO1xyXG5pbXBvcnQgeyBDb21tYW5kUmVzcG9uc2VQcm9jZXNzb3IgfSBmcm9tIFwiLi9jb21tYW5kLXJlc3BvbnNlLXByb2Nlc3Nvci5qc1wiO1xyXG5pbXBvcnQgeyBEb2N1bWVudFN5bmNIZWxwZXIgfSBmcm9tIFwiLi9kb2N1bWVudC1zeW5jLWhlbHBlci5qc1wiO1xyXG5cclxudmFyIFdTID0gcmVxdWlyZSgnbmF0aXZlc2NyaXB0LXdlYnNvY2tldHMnKTtcclxuXHJcbmV4cG9ydCBjbGFzcyBXb3JrZXJUYXNrcyB7XHJcbiAgICBwdWJsaWMgc3RhdGljIHNvY2tldDogYW55O1xyXG4gICAgcHJpdmF0ZSBzdGF0aWMgaXNTb2NrZXRJbml0aWFsaXplZDogYm9vbGVhbjtcclxuICAgIHByaXZhdGUgc3RhdGljIHdvcmtlclJlZmVyZW5jZTogV29ya2VyO1xyXG4gICAgcHVibGljIHN0YXRpYyBJbml0KHdvcmtlcjogV29ya2VyKSB7XHJcbiAgICAgICAgV29ya2VyVGFza3Mud29ya2VyUmVmZXJlbmNlID0gd29ya2VyO1xyXG4gICAgICAgIGNvbnNvbGUubG9nKFwiaW4gV29ya2VyVGFza3MgSW5pdFwiKVxyXG4gICAgICAgIFNlcnZlckhlbHBlci5pbml0KFdvcmtlclRhc2tzLnBvc3RNZXNzYWdlKTtcclxuICAgICAgICBTZXJ2ZXJIZWxwZXIuc2VuZFRvU2VydmVyQ2FsbGJhY2sgPSBXb3JrZXJUYXNrcy5zZW5kVG9TZXJ2ZXI7XHJcbiAgICAgICAgUGxhdGZvcm1IZWxwZXIuaW5pdCgpO1xyXG4gICAgfVxyXG5cclxuICAgIHB1YmxpYyBzdGF0aWMgcHJvY2Vzc01lc3NhZ2UobXNnOiBhbnkpIHtcclxuXHJcbiAgICAgICAgY29uc29sZS5sb2cobXNnKTtcclxuICAgICAgICBzd2l0Y2ggKG1zZy5tc2d0eXBlKSB7XHJcbiAgICAgICAgICAgIGNhc2UgU0VSVkVSX1dPUktFUl9NU0dfVFlQRS5JTklUX1NFUlZFUl9JTlRFUkZBQ0U6XHJcbiAgICAgICAgICAgICAgICAvLyBzZXQgc2VydmVyIHdvcmtlciBjb250ZXh0XHJcbiAgICAgICAgICAgICAgICBTZXJ2ZXJXb3JrZXJDb250ZXh0LnNlcnZlclVybCA9IG1zZy5kYXRhLldlYnNvY2tldFVybDtcclxuICAgICAgICAgICAgICAgIFNlcnZlcldvcmtlckNvbnRleHQuYXV0aFRva2VuID0gbXNnLmRhdGEuVG9rZW47XHJcbiAgICAgICAgICAgICAgICBXb3JrZXJUYXNrcy5pbml0V2ViU29ja2V0KCk7XHJcbiAgICAgICAgICAgICAgICBicmVhaztcclxuICAgICAgICAgICAgY2FzZSBTRVJWRVJfV09SS0VSX01TR19UWVBFLkNPTk5FQ1RfU0VSVkVSX0lOVEVSRkFDRTpcclxuICAgICAgICAgICAgICAgIGlmIChXb3JrZXJUYXNrcy5pc1NvY2tldEluaXRpYWxpemVkKSB7XHJcbiAgICAgICAgICAgICAgICAgICAgaWYgKFdvcmtlclRhc2tzLnNvY2tldC5yZWFkeVN0YXRlID09IDMpIC8vIGNsb3NlZFxyXG4gICAgICAgICAgICAgICAgICAgIHtcclxuICAgICAgICAgICAgICAgICAgICAgICAgV29ya2VyVGFza3Muc29ja2V0Lm9wZW4oKTtcclxuICAgICAgICAgICAgICAgICAgICB9XHJcbiAgICAgICAgICAgICAgICB9XHJcbiAgICAgICAgICAgICAgICBicmVhaztcclxuICAgICAgICAgICAgY2FzZSBTRVJWRVJfV09SS0VSX01TR19UWVBFLlNFTkRfTUVTU0FHRTpcclxuICAgICAgICAgICAgICAgIG1zZy5kYXRhLmZvckVhY2goZWxlbWVudCA9PiB7XHJcbiAgICAgICAgICAgICAgICAgICAgY29uc3QgYXBwTWVzc2FnZVVJSGFuZGxlciA9IG5ldyBBcHBNZXNzYWdlVUlIYW5kbGVyKCk7XHJcbiAgICAgICAgICAgICAgICAgICAgYXBwTWVzc2FnZVVJSGFuZGxlci5oYW5kbGVNZXNzYWdlKGVsZW1lbnQsIFdvcmtlclRhc2tzLnBvc3RNZXNzYWdlKTtcclxuICAgICAgICAgICAgICAgIH0pO1xyXG4gICAgICAgICAgICAgICAgYnJlYWs7XHJcbiAgICAgICAgICAgIGNhc2UgU0VSVkVSX1dPUktFUl9NU0dfVFlQRS5VUExPQURfRE9DVU1FTlRfQ09NUExFVEVEOlxyXG4gICAgICAgICAgICAgICAgLy8gY2FsbCBkZWxldGUgZG9jdW1lbnQgc3luY1xyXG4gICAgICAgICAgICAgICAgRG9jdW1lbnRTeW5jSGVscGVyLmRlbGV0ZURvY0Zyb21Mb2NhbFN0b3JlKG1zZy5kYXRhKVxyXG4gICAgICAgICAgICAgICAgYnJlYWs7XHJcbiAgICAgICAgfVxyXG4gICAgfVxyXG5cclxuICAgIHByaXZhdGUgc3RhdGljIGluaXRXZWJTb2NrZXQoKSB7XHJcbiAgICAgICAgV29ya2VyVGFza3Muc29ja2V0ID0gbmV3IFdTKFNlcnZlcldvcmtlckNvbnRleHQuc2VydmVyVXJsLCBbXSk7XHJcbiAgICAgICAgY29uc29sZS5sb2coJ3NvY2tldCBjcmVhdGVkJywgV29ya2VyVGFza3Muc29ja2V0KTtcclxuICAgICAgICBXb3JrZXJUYXNrcy5pc1NvY2tldEluaXRpYWxpemVkID0gdHJ1ZTtcclxuICAgICAgICBXb3JrZXJUYXNrcy5zb2NrZXQub24oJ29wZW4nLCBzb2NrZXQgPT4ge1xyXG4gICAgICAgICAgICBjb25zb2xlLmxvZygnbWVzc2FnZXMnLCBcIldlYlNvY2tldCBvcGVuZWRcIik7XHJcbiAgICAgICAgICAgIFdvcmtlclRhc2tzLnJhaXNlU29ja2V0Q29ubmVjdGlvbkV2ZW50KHRydWUpO1xyXG5cclxuICAgICAgICAgICAgLy9vbiBjb25uZWN0IHN5bmMgZGF0YVxyXG4gICAgICAgICAgICBTZXJ2ZXJXb3JrZXJDb250ZXh0LnN5bmNTdGF0ZSA9IFNFUlZFUl9TWU5DX1NUQVRFLk5PTkU7XHJcbiAgICAgICAgICAgIFNlcnZlckhlbHBlci5zd2l0Y2hTeW5jU3RhdGUoKTtcclxuXHJcblxyXG4gICAgICAgIH0pO1xyXG4gICAgICAgIFdvcmtlclRhc2tzLnNvY2tldC5vbignbWVzc2FnZScsIChzb2NrZXQsIG1lc3NhZ2UpID0+IHtcclxuICAgICAgICAgICAgLy8gY29uc29sZS5sb2coXCJ3ZWJzb2NrZXQgbWVzc2FnZSByZWNpZXZlZFwiLCBtZXNzYWdlKTtcclxuXHJcbiAgICAgICAgICAgIC8vIHByb2Nlc3MgcmVzcCBtc2dcclxuICAgICAgICAgICAgQ29tbWFuZFJlc3BvbnNlUHJvY2Vzc29yLmNtZFByb2Nlc3NvcihtZXNzYWdlKTtcclxuXHJcbiAgICAgICAgfSk7XHJcblxyXG4gICAgICAgIFdvcmtlclRhc2tzLnNvY2tldC5vbignY2xvc2UnLCAoc29ja2V0LCBjb2RlLCByZWFzb24pID0+IHtcclxuICAgICAgICAgICAgY29uc29sZS5sb2coJ21lc3NhZ2VzJywgXCJXZWJzb2NrZXQgZGlzY29ubmVjdGVkXCIpO1xyXG4gICAgICAgICAgICBXb3JrZXJUYXNrcy5yYWlzZVNvY2tldENvbm5lY3Rpb25FdmVudChmYWxzZSk7XHJcbiAgICAgICAgfSk7XHJcbiAgICAgICAgV29ya2VyVGFza3Muc29ja2V0Lm9uKCdlcnJvcicsIChzb2NrZXQsIGVycm9yKSA9PiB7XHJcbiAgICAgICAgICAgIGNvbnNvbGUubG9nKFwiVGhlIHNvY2tldCBoYWQgYW4gZXJyb3JcIiwgZXJyb3IpO1xyXG4gICAgICAgICAgICBXb3JrZXJUYXNrcy5yYWlzZVNvY2tldENvbm5lY3Rpb25FdmVudChmYWxzZSk7XHJcbiAgICAgICAgfSk7XHJcblxyXG4gICAgICAgIFdvcmtlclRhc2tzLnNvY2tldC5vcGVuKCk7XHJcbiAgICB9XHJcblxyXG4gICAgcHJpdmF0ZSBzdGF0aWMgcmFpc2VTb2NrZXRDb25uZWN0aW9uRXZlbnQoc3RhdHVzOiBib29sZWFuKSB7XHJcbiAgICAgICAgY29uc3Qgd29ya2VyRXZlbnQgPSBuZXcgU2VydmVyV29ya2VyRXZlbnREYXRhTW9kZWwoKTtcclxuICAgICAgICBpZiAoc3RhdHVzKSB7XHJcbiAgICAgICAgICAgIHdvcmtlckV2ZW50Lm1zZ3R5cGUgPSBTRVJWRVJfV09SS0VSX0VWRU5UX01TR19UWVBFLlNFUlZFUl9DT05ORUNURUQ7XHJcbiAgICAgICAgfVxyXG4gICAgICAgIGVsc2Uge1xyXG4gICAgICAgICAgICB3b3JrZXJFdmVudC5tc2d0eXBlID0gU0VSVkVSX1dPUktFUl9FVkVOVF9NU0dfVFlQRS5TRVJWRVJfRElTQ09OTkVDVEVEO1xyXG4gICAgICAgIH1cclxuXHJcbiAgICAgICAgV29ya2VyVGFza3MucG9zdE1lc3NhZ2Uod29ya2VyRXZlbnQpO1xyXG4gICAgfVxyXG5cclxuICAgIHB1YmxpYyBzdGF0aWMgc2VuZFRvU2VydmVyKG1zZzogYW55KTogdm9pZCB7XHJcbiAgICAgICAgY29uc29sZS5sb2coXCJzZW5kVG9TZXJ2ZXJcIiwgbXNnKTtcclxuICAgICAgICBXb3JrZXJUYXNrcy5zb2NrZXQuc2VuZChtc2cpO1xyXG4gICAgfVxyXG5cclxuICAgIHB1YmxpYyBzdGF0aWMgcG9zdE1lc3NhZ2UobXNnOiBTZXJ2ZXJXb3JrZXJFdmVudERhdGFNb2RlbCkge1xyXG4gICAgICAgIFdvcmtlclRhc2tzLndvcmtlclJlZmVyZW5jZS5wb3N0TWVzc2FnZShtc2cpO1xyXG4gICAgfVxyXG5cclxufSJdfQ==