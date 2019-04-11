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
export class AssociateServicePointDeviceRequest {
    devid: number;
    spid: number;
}
export class ServicepointAssociateRequest {
    servconfid: number;
    spid: number;
}
export class ServicepointListResponse {
    spid: number;
    spname: string;
}

export class SPCategoriesShortDataResponse {
    spcid: number;
    spcname: string;
}
export class ServicepointDetailsResponse {
    spid: number;
    spcid: number;
    spname: string;
    shortdesc: string;
    spcname: string;
    spstate: number;
    spstatesince: Date;
    createdon: Date;
    updatedon: Date;
}
export class ServicepointDetailsUpdateRequest {
    spid: number;
    spname: string;
    shortdesc: string;
    spstate: number;
    spcid: number;
}
