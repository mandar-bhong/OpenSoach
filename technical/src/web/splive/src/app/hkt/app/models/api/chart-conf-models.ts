export class TaskTemplateResponse {
    taskname: string;
    shortdec: string;
    spcid: number;
}

export class TaskTemplateRequest {
    spcid: number;
    taskname: string;
    shortdesc: string;
}
