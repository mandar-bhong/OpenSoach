export class ChartListViewModel {
    dbmodel:any;
}

export class ChartDBModel {
    admissionid:number;
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
    freqMorn:boolean;
    freqAftrn:boolean;  
    freqNight:boolean;      
    intervalHrs: number;
    startDate:string;
    duration:number;    
    startTime:string;
    desc:string;
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