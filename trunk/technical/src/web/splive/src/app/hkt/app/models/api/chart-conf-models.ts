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
export class ChartsDetailsResponse {
    servconfid: number;
    spcid: number;
    conftypecode: string;
    servconfname: string;
    shortdesc: string;
    servconf: string;
    createdon: Date;
    updateon: Date;
}
