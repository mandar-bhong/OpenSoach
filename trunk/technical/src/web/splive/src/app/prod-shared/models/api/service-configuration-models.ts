import { SERVICEPOINT_STATE } from '../../../shared/app-common-constants';

export class ServiceConfigurationRequest {
    servconfid: number;
    spcid: number;
    conftypecode: string;
    servconfname: string;
    shortdesc: string;
    servconf: string;
}
export class ServiceConfigurationUpdateRequest {
    servconfid: number;
    servconfname: string;
    shortdesc: string;
    servconf: string;
}
export class ServiceConfigurationlistResponse {
    page: number;
    limit: number;
    orderby: number;
    orderdirection: string;

}
export class ServicepointDataListResponse {
    spid: number;
    spname: string;
    spcid: number;
    spcname: string;
    devid: number;
    devname: string;
    servconfid: number;
    spstate: number;
    spstatesince: number;
    createdon: Date;
    updatedon: Date;
}
export class SrevicepointFilterRequest {
    spname: string;
    spcid: number;
    spstate: number;
    devid: number;
}

export class ServicepointAssociateRequest {
    servconfid: number;
    spid: number;
}
export class ServicepointConfigureListResponse {
    servconfid: number;
    servconfname: string;
}
export class ServicepointConfigureTemplateListRequest {
    servconfid: number;
}

export class AssociateServicePointDeviceRequest {
    devid: number;
    spid: number;
}
