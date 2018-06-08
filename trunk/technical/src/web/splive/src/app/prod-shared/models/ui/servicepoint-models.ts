import {
    ServicepointConfigureListResponse,
    ServicepointConfigureTemplateListRequest,
} from '../api/service-configuration-models';
import {
    ServicepointAssociateRequest, SrevicepointFilterRequest,
    ServicepointDetailsResponse, ServicepointDetailsUpdateRequest
} from '../api/servicepoint-models';
import { SERVICEPOINT_STATE } from '../../../shared/app-common-constants';

export class ServicepointFilterModel {
    spname: string;
    spcid: number;
    spstate: number;
    devid: number;
    servconfid: number;

    copyTo(srevicepointFilterRequest: SrevicepointFilterRequest) {
        srevicepointFilterRequest.spcid = this.spcid;
        srevicepointFilterRequest.spname = this.spname;
        srevicepointFilterRequest.spstate = this.spstate;
        srevicepointFilterRequest.devid = this.devid;
    }
}

export class ServicePointServiceConfigureAssociateModel {
    servconfid: number;
    spid: number;
    servconfname: string;

    copyToAssociateRequest(request: ServicepointAssociateRequest) {
        request.servconfid = this.servconfid;
        request.spid = this.spid;
    }
    copyFrom(details: ServicepointConfigureListResponse) {
        this.servconfid = details.servconfid;
        this.servconfname = details.servconfname;
    }
    copyTo(request: ServicepointConfigureTemplateListRequest) {
        request.servconfid = this.servconfid;
    }
}
export class ServicePointDetailsModel {
    spid: number;
    spcid: number;
    spname: string;
    // spcname: string;
    spstate: SERVICEPOINT_STATE;
    spstatesince: Date;
    copyFrom(servicepointDetailsResponse: ServicepointDetailsResponse) {
        this.spid = servicepointDetailsResponse.spid;
        this.spcid = servicepointDetailsResponse.spcid;
        this.spname = servicepointDetailsResponse.spname;
        this.spstate = servicepointDetailsResponse.spstate;
        // this.spcname = servicepointDetailsResponse.spcname;
    }
    copyTo(servicepointDetailsUpdateRequest: ServicepointDetailsUpdateRequest) {
        servicepointDetailsUpdateRequest.spid = this.spid;
        servicepointDetailsUpdateRequest.spname = this.spname;
        servicepointDetailsUpdateRequest.spstate = this.spstate;
        servicepointDetailsUpdateRequest.spcid = this.spcid;
    }
}
