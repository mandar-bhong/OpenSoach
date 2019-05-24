export class PathlogyReportModel {
    uuid: string;
    admission_uuid: string;
    test_performed: string;
    test_performed_time: string;
    test_result: string;
    comments: string;
    doc_name: string;
    doclist: string[];
}
export class TeatmentReportModel{
    uuid: string;
    admission_uuid: string;
    treatment_done: string;
    treatment_performed_time: string;
    details: string;
    test_result: string;    
    post_observation: string;
    doc_name: string;
    doclist: string[];
}
