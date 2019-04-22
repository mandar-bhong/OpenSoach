export class PathologyReportAddRequest {
    uuid: string;
    admissionid: number;
    testperformed: string;
    testresult: string;
    comments: string;
    testperformedtime: string;
    documentuuidlist: string[];
    copyTo() {

    }
}