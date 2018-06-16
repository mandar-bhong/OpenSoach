export class ComplaintFiltrRequest {
    complainttitle: string;
    complaintstate: number;
}
export class ComplaintDataListResponse {
    complaintid: number;
    spname: string;
    complainttitle: string;
    description: string;
    complaintby: string;
    raisedon: Date;
    complaintstate: number;
    closedon: Date;
}

export class ComplaintUpdateRequest {
    complaintid: number;
    complaintstate: number;
    remarks: string;
    closedon: Date;
    severity: number;
}
export class ComplaintDetailsResponse {
    complaintid: number;
    cpmid: number;
    spid: number;
    complainttitle: string;
    description: string;
    complaintby: string;
    mobileno: string;
    emailid: string;
    employeeid: string;
    severity: number;
    raisedon: Date;
    complaintstate: number;
    closedon: Date;
    remarks: string;
    createdon: Date;
    updatedon: Date;
}
