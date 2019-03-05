import { DatabaseHelper } from "../helpers/database-helper.js";
import { ServerWorkerEventDataModel } from "../models/api/server-worker-event-data-model.js";
import { SERVER_WORKER_EVENT_MSG_TYPE } from "../app-constants.js";
import { WorkerTasks } from "./worker-tasks.js";

export class DocumentSyncHelper {
    static isDocumentSyncRunning = false;
    static isSyncTriggeredWhileUploading = false;
    static readDocumentsToSync(): Promise<any[]> {
        return new Promise((resolve, reject) => {
            // call sqlite to read the documents from document_tbl

            DatabaseHelper.selectAll("documentlist").then((doclist) => {
                console.log("doclist", doclist);
                resolve(doclist);
            });
        });
    }

    static sync() {
        if (DocumentSyncHelper.isDocumentSyncRunning) {
            DocumentSyncHelper.isSyncTriggeredWhileUploading = true;
            return;
        }

        DocumentSyncHelper.isDocumentSyncRunning = true;

        DocumentSyncHelper.readDocumentsToSync().then(async docList => {
            for (const doc of docList) {
                DocumentSyncHelper.uploadDoc(doc);
            }

            DocumentSyncHelper.isDocumentSyncRunning = false;
            if (DocumentSyncHelper.isSyncTriggeredWhileUploading) {
                DocumentSyncHelper.isSyncTriggeredWhileUploading = false;
                setTimeout(DocumentSyncHelper.sync);
            }
        }, error => {
            DocumentSyncHelper.isDocumentSyncRunning = false;
            return;
        });
    }

    static deleteDocFromLocalStore(doc: any) {

        console.log("in deleteDocFromLocalStore..")

        // delete the document_tbl entry
        if (doc.isFileUploaded) {
            const listData = new Array<any>();
            listData.push(doc.uuid);
            DatabaseHelper.update("document_tbl_delete", listData).then(
                (val) => {
                    console.log("deleted id", val);
                },
                (err) => {
                    console.log("delete err:", err);
                })
        }
    }

    static uploadDoc(doc: any) {

            // use nativescript-background-http to upload file
            // resolve if file upload is sucessfull else reject

            const serverWorkerEventDataModel = new ServerWorkerEventDataModel();
            serverWorkerEventDataModel.msgtype = SERVER_WORKER_EVENT_MSG_TYPE.UPLOAD_DOCUMENT;
            serverWorkerEventDataModel.data = {};
            serverWorkerEventDataModel.data.path = doc.doc_path;
            serverWorkerEventDataModel.data.uuid = doc.uuid;

            WorkerTasks.postMessage(serverWorkerEventDataModel);

        }
        
}