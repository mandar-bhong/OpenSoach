export class ReportRequest {
    reportid: number;
    spid: number;
    startdate: Date;
    enddate: Date;
}

export class ReportResponse {
    reportheader: any;
    reportdata: string[][];
}