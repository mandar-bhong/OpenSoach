export class ScheduleDataResponse<T> {
    uuid: string;
    patientconfid: number;
    cpmid: number;
    admissionid: number;
    conftypecode: string;
    conf: T;
    enddate: string;
}