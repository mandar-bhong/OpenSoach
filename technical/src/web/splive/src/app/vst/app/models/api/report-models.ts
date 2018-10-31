export class ReportRequestParams {
    reportreq: ReportRequest[];
    reportfileformat: string;
}

export class ReportRequest {
    reportcode: string;
    lang: string;
    queryparams: any[];
}

export class ReportResponse {
    reportcode: string;
    reportheader: string[];
    reportdata: any[][];
}
