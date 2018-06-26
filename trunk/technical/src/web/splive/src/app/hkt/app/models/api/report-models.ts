export class ReportRequest {
    reportcode: string;
    lang:string;
    queryparams: any[];
}

export class ReportResponse {
    reportheader: any;
    reportdata: string[][];
}