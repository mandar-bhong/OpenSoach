export class TaskTemplateResponse {
    taskname: string;
    shortdec: string;
    spcid: number;
}
export class CategoriesShortDataResponse {
    spcid: number;
    spcname: string;
}
export class TaskTemplateRequest {
    spcid: number;
    taskname: string;
    shortdesc: string;
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
