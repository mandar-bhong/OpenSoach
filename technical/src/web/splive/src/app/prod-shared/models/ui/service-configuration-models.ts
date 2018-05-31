import {
    SrevicepointFilterRequest,
    ServicepointAssociateRequest, ServicepointConfigureListResponse, ServicepointConfigureTemplateListRequest
} from '../api/service-configuration-models';

export class ServiceConfigurationModel {
    servconfid: number;
    spcid: number;
    conftypecode: string;
    servconfname: string;
    shortdesc: string;
    servconf: string;
    spid: number;
    copyToAssociateRequest(request: ServicepointAssociateRequest) {
        request.servconfid = this.servconfid;
        request.spid = this.spid;
    }
}
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
export class ConfigureAssociateModel {
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
