export class PathlogyReportModel {
    uuid: string;
    admission_uuid: string;
    test_performed: string;
    test_performed_time: string;
    test_result: string;
    comments: string;
    doc_name: string;
    doclist: PathlogyReportDocModel[];
}

export class PathlogyReportDocModel {
    pathology_record_uuid: string;
    document_uuid: string;
    document_name: string;
    doctype: string;
    document_path: string;
    progress:number;
}

export class TeatmentReportModel {
    uuid: string;
    admission_uuid: string;
    treatment_done: string;
    treatment_performed_time: string;
    details: string;
    test_result: string;
    post_observation: string;
    doc_name: string;
    doclist: TreatmentReportDocModel[];
}

export class TreatmentReportDocModel {
    treatment_uuid: string;
    document_uuid: string;
    document_name: string;
    doctype: string;
    document_path: string;
    progress:number;
}