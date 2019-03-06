import { AppGlobalContext } from "../app-global-context";
import { ServerDataProcessorMessageModel } from "../models/api/server-data-processor-message-model";
import { SERVER_WORKER_MSG_TYPE } from "../app-constants";
import { WorkerService } from "../services/worker.service";

export class DocumentHelper {

    public static uploadDocument(data: any, workerservice: WorkerService) {
            try {

                // use nativescript-background-http to upload file
                // resolve if file upload is sucessfull else reject

                // file path and url
                var file = data.path;
                var url = "http://172.105.232.148:91/api/v1/document/upload/ep";

                // upload configuration
                var bghttp = require("nativescript-background-http");
                var session = bghttp.session("multipart/form-data");

                var params = [
                    { name: "file", filename: file },
                    { name: "UUID", value: data.uuid }
                ];

                var request = {
                    url: url,
                    method: "POST",
                    headers: {
                        "Content-Type": "multipart/form-data",
                        "Authorization": AppGlobalContext.Token
                    },
                    androidDisplayNotificationProgress: false
                };

                // console.log("params,request:", params, request);

                var task = session.multipartUpload(params, request);

                task.on("error", (e) => {
                    console.log('file upload errror', e);
                });

                task.on("responded", (e) => {
                    // console.log("e", e);
                    if (JSON.parse(e.data).issuccess) {
                        console.log("upload success..");
                        const initdata: any = {};
                        initdata.isFileUploaded = true;
                        initdata.uuid = data.uuid;
                        const initModel = new ServerDataProcessorMessageModel();
                        initModel.msgtype = SERVER_WORKER_MSG_TYPE.UPLOAD_DOCUMENT_COMPLETED;
                        initModel.data = initdata;
                        workerservice.postMessageToServerDataProcessorWorker(initModel);
                    }
                });
            }
            catch (e) {
                console.error('upload error for doc', e);
            }
    }

}