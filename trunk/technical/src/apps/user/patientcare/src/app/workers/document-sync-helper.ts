export class DocumentSyncHelper {
    static isDocumentSyncRunning = false;
    static isSyncTriggeredWhileUploading = false;
    static readDocumentsToSync(): Promise<any[]> {
        return new Promise((resolve, reject) => {
            // call sqlite to read the documents from document_tbl
            resolve([]);
        });
    }

    static sync() {
        if (DocumentSyncHelper.isDocumentSyncRunning) {
            DocumentSyncHelper.isSyncTriggeredWhileUploading = true;
            return;
        }

        DocumentSyncHelper.isDocumentSyncRunning = true;

        DocumentSyncHelper.readDocumentsToSync().then(async docList => {
            for (const doc in docList) {
                const isFileUploaded = await DocumentSyncHelper.uploadDoc(doc);
                if (isFileUploaded) {
                    DocumentSyncHelper.deleteDocFromLocalStore(doc);
                }
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
        // delete the document_tbl entry
    }

    static uploadDoc(doc: any): Promise<boolean> {
        return new Promise((resolve, reject) => {
            try {

                // use nativescript-background-http to upload file
                // resolve if file upload is sucessfull else reject
                resolve(true);
            }
            catch (e) {
                console.error('upload error for doc', doc, e);
                reject(e);
            }
        });
    }
}