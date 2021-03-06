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
export class ServicepointConfigureListResponse {
    servconfid: number;
    servconfname: string;
}
export class ServicepointConfigureTemplateListRequest {
    servconfid: number;
}

export class ServiceConfigureFilterRequest {
    conftypecode: string;
    servconfname: string;
}
export class ServiceConfigurationResponse {
    servconfid: number;
    spcid: number;
    conftypecode: string;
    servconfname: string;
    shortdesc: string;
    servconf: string;
    createdon: Date;
    updateon: Date;
}

export class ServiceConfigureDataListResponse {
    servconfid: number;
    cpmid: number;
    spcid: number;
    conftypecode: string;
    servconfname: string;
    shortdesc: string;
    createdon: Date;
    updatedon: Date;
}

export class ServicePointWithConfigurationResponse {
    spid: number;
    spname: string;
    servconfid: number;
    spcid: number;
    spcname: string;
}
