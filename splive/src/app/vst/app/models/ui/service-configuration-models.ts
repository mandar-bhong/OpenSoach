import { ServiceConfigureFilterRequest } from '../api/service-configuration-models';
import { ServicepointAssociateRequest } from '../api/servicepoint-models';

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
    copyTo(serviceConfigureFilterRequest: ServiceConfigureFilterRequest) {
        serviceConfigureFilterRequest.conftypecode = this.conftypecode;
        serviceConfigureFilterRequest.servconfname = this.servconfname;
    }
}

