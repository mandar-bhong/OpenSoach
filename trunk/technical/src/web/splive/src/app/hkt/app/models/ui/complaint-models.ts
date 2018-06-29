import { ComplaintFiltrRequest, ComplaintUpdateRequest, ComplaintDetailsResponse } from '../api/complaint-models';
import { SEVERIT_STATE, COMPLAINT_STATE } from '../../../../shared/app-common-constants';


export class ComplaintFilterModel {
    complainttitle: string;
    complaintstate: COMPLAINT_STATE;
    spname: string;
    spid: number;

    copyTo(complaintFiltrRequest: ComplaintFiltrRequest) {
        complaintFiltrRequest.complainttitle = this.complainttitle;
        complaintFiltrRequest.complaintstate = this.complaintstate;
        complaintFiltrRequest.spname = this.spname;
        complaintFiltrRequest.spid = this.spid;

    }
}

export class ComplaintDetailsModel {
    spid: number;
    complainttitle: string;
    description: string;
    complaintby: string;
    employeeid: string;
    severity: SEVERIT_STATE;
    complaintstate: COMPLAINT_STATE;
    complaintid: number;
    remarks: string;
    closedon: Date;
    cpmid: number;
    copyToUpdateRequest(complaintUpdateRequest: ComplaintUpdateRequest) {
        complaintUpdateRequest.complaintid = this.complaintid;
        complaintUpdateRequest.complaintstate = this.complaintstate;
        complaintUpdateRequest.remarks = this.remarks;
        complaintUpdateRequest.closedon = this.closedon;
        complaintUpdateRequest.severity = this.severity;
    }
    copyFrom(complaintDetailsResponse: ComplaintDetailsResponse) {
        this.complaintid = complaintDetailsResponse.complaintid;
        this.cpmid = complaintDetailsResponse.cpmid;
        this.spid = complaintDetailsResponse.spid;
        this.complaintby = complaintDetailsResponse.complaintby;
        this.complainttitle = complaintDetailsResponse.complainttitle;
        this.description = complaintDetailsResponse.description;
        this.complaintstate = complaintDetailsResponse.complaintstate;
        this.severity = complaintDetailsResponse.severity;
        this.remarks = complaintDetailsResponse.remarks;
    }
}
