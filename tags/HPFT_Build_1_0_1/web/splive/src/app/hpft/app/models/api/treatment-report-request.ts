export class TreatmentReportRequest {
    uuid: string;
    admissionid: number;
    treatmentdone: string;
    treatmentperformedtime: string;
    details: string;
    postobservation: string;
    documentuuidlist: string[];
}