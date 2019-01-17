export class ChartListViewModel {
    dbmodel:any;
}

export class ChartDBModel {
    uuid:string;
    admission_uuid:string;
    conf_type_code:string;
    conf:any;
}

export class MonitorChartModel {
    name: string;
    foodInst: number;
    frequency: number;
    duration: number;
    startDate: string;
    intervalHrs: number;
    startTime: string;
    endTime: string;
    specificTime: string;
    specificTimes: Array<any>;
    desc: string;
}

export class MedChartModel{
    name:string;
    quantity:number;
    foodInst:number;
    frequency:number;
    mornFreqInfo:MornFreqInfo;
    aftrnFreqInfo:AftrnFreqInfo;
    nightFreqInfo:NightFreqInfo    
    intervalHrs: number;
    startDate:string;
    duration:number;    
    startTime:string;
    desc:string;
}

export class MornFreqInfo{
    mornFreq:boolean;
    mornFreqQuantity:number;
}

export class AftrnFreqInfo{
    aftrnFreq:boolean;
    aftrnFreqQuantity:number;
}

export class NightFreqInfo{
    nightFreq:boolean;
    nightFreqQuantity:number;
}

export class IntakeChartModel {
    name: string;
    quantity: string;
    frequency: number;
    duration: number;
    startDate: string;
    intervalHrs: number;
    startTime: string;
    specificTime: string;
    specificTimes: Array<any>;
    desc: string;
}