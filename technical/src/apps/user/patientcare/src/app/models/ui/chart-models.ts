export class ChartListViewModel {
    dbmodel: any;
}

export class ChartDBModel {   
    conf_type_code: string;
    conf: any;
    uuid:string;
    admission_uuid:string;
  
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

export class MedChartModel {
    name: string;
    quantity: number;
    foodInst: number;
    frequency: number;
    mornFreqInfo:MornFreqInfo;
    aftrnFreqInfo:AftrnFreqInfo;
    nightFreqInfo:NightFreqInfo  
    intervalHrs: number;
    startDate: string;
    duration: number;
    startTime: string;
    desc: string;
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

export class Schedulardata {
    uuid:string;
    admission_uuid:string;
    conf_type_code: string;
    conf: MedChartModel;
}

export class ConfigData {
    name: string;
    desc: string;
    freqMorn: boolean;
    freqAftrn: boolean;
    freqNight: boolean;
    quantity: number;
    startDate: string;
    frequency: number;
    foodInst: number;
    duration: string;
    intervalHrs: string;
    startTime: string;
    
}
export class ActionListModel {
    uuid:string;
    admission_id:number;
    chart_conf_id:number;
    exec_time:string;
}

export class MornFreqInfo{
    freqMorn:boolean;
    mornFreqQuantity:number;
}

export class AftrnFreqInfo{
    freqAftrn:boolean;
    aftrnFreqQuantity:number;
}

export class NightFreqInfo{
    freqNight:boolean;
    nightFreqQuantity:number;
}