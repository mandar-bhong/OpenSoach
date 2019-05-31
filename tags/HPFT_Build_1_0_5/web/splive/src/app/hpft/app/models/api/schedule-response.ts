export class ScheduleDataResponse<T> {
    uuid: string;
    patientconfid: number;
    cpmid: number;
    admissionid: number;
    conftypecode: string;
    conf: T;
    enddate: string;
    startdate: string;
    status: number;
}
export class SchedularConfigData {
    name: string;
    quantity: number;
    foodInst: number;
    frequency: number;
    mornFreqInfo: MornFreqInfo;
    aftrnFreqInfo: AftrnFreqInfo;
    nightFreqInfo: NightFreqInfo
    interval: number;
    startDate: string;
    duration: number;
    startTime: string;
    desc: string;
    endTime: string;
    specificTime: string;
    specificTimes: Array<any>;
    numberofTimes: number;
    endDate: Date
    splinstruction: string; 
    remark: string;  
}
export class MornFreqInfo {
    freqMorn: boolean;
    mornFreqQuantity: number;
}

export class AftrnFreqInfo {
    freqAftrn: boolean;
    aftrnFreqQuantity: number;
}

export class NightFreqInfo {
    freqNight: boolean;
    nightFreqQuantity: number;
}