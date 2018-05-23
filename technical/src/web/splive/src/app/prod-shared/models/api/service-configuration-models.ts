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
